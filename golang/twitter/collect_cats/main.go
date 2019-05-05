package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_API_TOKEN"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_SECRET_API_TOKEN"))
	return anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_SECRET_ACCESS_TOKEN"))
}

func main() {
	loadEnv()

	api := getTwitterApi()

	v := url.Values{}
	v.Set("count", "30")

	searchResult, err := api.GetSearch("golang", v)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}
}
