package seeds


var seeds = [...]seed{
	Subreddit{},
}

func Execute() {
	for _, x := range seeds {
		x.Run()
	}
}