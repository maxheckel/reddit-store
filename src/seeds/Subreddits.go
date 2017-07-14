package seeds

import "models"
import (
	"databaseConn"
	//"fmt"
)

type Subreddit struct {
	Seed
}

func (seed Seed) Run() {
	var data = make([]map[string]string, 2)
	data[0] = map[string]string{
		"name" : "Ask Reddit",
		"subreddit": "askreddit",
	}
	data[1] = map[string]string{
		"name" : "Gaming",
		"subreddit": "gaming",
	}
	var db = databaseConn.DB{}.GetDB()

	for _, row := range data {
		var Subreddit = new(models.Subreddit)
		db.FirstOrInit(Subreddit, models.Subreddit{
			Name: row["name"],
			Subreddit: row["subreddit"],
		})
		db.Create(Subreddit)
	}

}

