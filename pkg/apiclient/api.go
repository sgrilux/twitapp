package apiclient

import (
	"fmt"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/sgrilux/twitapp/twitapp"
)

type TwitterClient struct {
	twitter twitter.Client
	user    twitter.User
}

func NewTwitterClient(creds twitapp.Credentials) (*TwitterClient, error) {

	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessSecret)

	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	return &TwitterClient{
		twitter: *client,
		user:    *user,
	}, nil
}

func (client *TwitterClient) Tweet(msg string) (*twitter.Tweet, *http.Response, error) {
	if msg == "" {
		return nil, nil, fmt.Errorf("msg is empy")
	}
	return client.twitter.Statuses.Update(msg, nil)
}

func (client *TwitterClient) GetFollowerList(user *twitter.User) (followerList twitapp.FollowerList, err error) {
	var cursor int64

	cursor = -1
	for {
		followers, _, err := client.twitter.Followers.List(getFollowerListParams(user, cursor))
		if err != nil {
			return nil, err
		}

		followerList = append(followerList, followers.Users...)

		cursor = followers.NextCursor
		if cursor == 0 {
			break
		}
	}

	return followerList, nil
}

func (client *TwitterClient) GetFollowingList(user *twitter.User) (followingList twitapp.FollowingList, err error) {
	var cursor int64

	cursor = -1
	for {
		followers, _, err := client.twitter.Friends.List(getFriendListParams(user, cursor))
		if err != nil {
			return nil, err
		}

		followingList = append(followingList, followers.Users...)

		cursor = followers.NextCursor
		if cursor == 0 {
			break
		}
	}

	return followingList, nil
}

func (client *TwitterClient) GetUser(user string) (*twitter.User, error) {
	if user == "" {
		return &client.user, nil
	}

	params := twitter.UserShowParams{
		ScreenName: user,
	}

	twitterUser, _, err := client.twitter.Users.Show(&params)
	if err != nil {
		return nil, err
	}

	return twitterUser, nil
}

func getFollowerListParams(user *twitter.User, cursor int64) *twitter.FollowerListParams {
	return &twitter.FollowerListParams{
		UserID:     user.ID,
		ScreenName: user.ScreenName,
		Cursor:     cursor,
		Count:      200,
	}
}

func getFriendListParams(user *twitter.User, cursor int64) *twitter.FriendListParams {
	return &twitter.FriendListParams{
		UserID:     user.ID,
		ScreenName: user.ScreenName,
		Cursor:     cursor,
		Count:      200,
	}
}
