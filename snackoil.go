package main
/*
This bot is using Dalton Hubble's go-twitter library.
it's a Go client library for the Twitter API.
https://github.com/dghubble/go-twitter <- Check it out, he has cool stuff! :)
*/
import (
	"fmt"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Get credentials via env vars.
var (
	ConsumerKey       = os.Getenv("CONSUMER_KEY")
	ConsumerKeySecret = os.Getenv("CONSUMER_SECRET")
	AccessToken       = os.Getenv("ACCESS_TOKEN")
	AccessSecret      = os.Getenv("ACCESS_TOKEN_SECRET")
)

func retweet(){
	config := oauth1.NewConfig(ConsumerKey, ConsumerKeySecret)
	token := oauth1.NewToken(AccessToken, AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Client
	client := twitter.NewClient(httpClient)
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("Username:%+v\n", user.ScreenName)

	// Search the last 20 #snackoil to retweet
	searchParams := &twitter.SearchTweetParams{
		Query: "%23" + "snackoil",
		Count: 20,
	}
	searchResult, _, _ := client.Search.Tweets(searchParams)
	if len(searchResult.Statuses) == 0 {
		os.Exit(0)
	}
	for _, tweet := range searchResult.Statuses {
		client.Statuses.Retweet(tweet.ID, &twitter.StatusRetweetParams{})
		fmt.Printf("RETWEETED: %+v\n", tweet.Text)
	}
}

func main() {
	for {
		go retweet()
		time.Sleep(10 * time.Second)
	}
}
