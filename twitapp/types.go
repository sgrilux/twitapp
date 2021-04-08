package twitapp

import "github.com/dghubble/go-twitter/twitter"

type Credentials struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

type FollowerList []twitter.User
type FollowingList []twitter.User
