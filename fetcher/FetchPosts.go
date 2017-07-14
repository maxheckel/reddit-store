package fetcher

import (
	"databaseConn"
	"models"
	//"fmt"
	"net/http"
	"io/ioutil"

	log "log"
)

func Execute() {
	db := databaseConn.DB{}.GetDB()
	subreddits := []models.Subreddit{}
	db.Find(&subreddits)
	for _, subreddit := range subreddits {
		client := &http.Client{}

		req, err := http.NewRequest("GET", "https://www.reddit.com/r/"+subreddit.Subreddit+".json", nil)
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

		log.Println(string(body))
	}
}
