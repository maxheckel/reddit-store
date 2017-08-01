package fetcher

import (
	"databaseConn"
	"models"
	"net/http"
	"io/ioutil"
	log "log"
	"encoding/json"
	"regexp"
	"sync"
)

var threshold int = 5

type Posts struct {
	Data struct{
		Children []struct{
			Kind string
			Data models.Post
		}
	}
}
var wg sync.WaitGroup

func Execute() {
	db := databaseConn.DB{}.GetDB()
	subreddits := []models.Subreddit{}
	db.Find(&subreddits)
	for _, subreddit := range subreddits {
		subredditUri := subreddit.Subreddit
		url := "https://www.reddit.com/r/" + subredditUri + "/top.json?sort=top&t=day"
		posts := fetchUrl(url, &Posts{})
		var subreddit = models.Subreddit{}
		db.Where("subreddit = ?", subredditUri).First(&subreddit)
		for x := 0; x < threshold; x++ {
			post := posts.Data.Children[x].Data
			post.SubredditID = subreddit.ID
			var Post = new(models.Post)
			db.FirstOrInit(Post, post)
			db.Create(&Post)
			matches, _ := regexp.MatchString("jpg|gifv|gif|png", Post.Url)
			if matches {
				wg.Add(1)
				go fetchImage(*Post, &wg)
			}
		}
	}
	wg.Wait()
}


func fetchUrl(url string, posts *Posts) *Posts{
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal([]byte(body), &posts); err != nil {
		panic(err)
	}
	return posts
}
