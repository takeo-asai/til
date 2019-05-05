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

func getTwitterAPI() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_API_TOKEN"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_SECRET_API_TOKEN"))
	return anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_SECRET_ACCESS_TOKEN"))
}

func main() {
	loadEnv()

	api := getTwitterAPI()
	v := url.Values{}
	v.Set("track", "猫,cat") // , で or 検索
	s := api.PublicStreamFilter(v)

	for t := range s.C {
		switch v := t.(type) {
		case anaconda.Tweet:
			for _, media := range v.ExtendedEntities.Media {
				url := media.Media_url_https
				videos := media.VideoInfo.Variants
				vURL := ""
				for _, video := range videos {
					if video.ContentType == "video/mp4" {
						vURL = video.Url
						break
					}
				}
				fmt.Printf("%-15s: [%v] %v, video: %v\n", v.User.ScreenName, v.User.Lang, url, vURL)
			}
		}
	}
}
