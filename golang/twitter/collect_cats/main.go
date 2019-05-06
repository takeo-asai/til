package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"net/http"
	"io"
	"path/filepath"

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

func saveImage(url string) {
	response, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()

	name := filepath.Base(url)
    file, err := os.Create(filepath.Join("images/", name))
    if err != nil {
        panic(err)
    }
    defer file.Close()

    io.Copy(file, response.Body)
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
				if url != "" {// まずは画像だけ
					saveImage(url)
				}
				fmt.Printf("%-15s: [%v] %v, video: %v\n", v.User.ScreenName, v.User.Lang, url, vURL)
			}
		}
	}
}
