package apiclient

import (
	"fmt"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TwitterClient struct {
	twitter twitter.Client
	user    twitter.User
}

type Credentials struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func NewTwitterClient(creds Credentials) (*TwitterClient, error) {

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

func (client *TwitterClient) GetFollowerList(params twitter.FollowerListParams) (*twitter.Followers, *http.Response, error) {
	return client.twitter.Followers.List(&params)
}

func (client *TwitterClient) GetUser(user string) (*twitter.User, error) {
	if user == "" {
		return client.getAccount(), nil
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

func (client *TwitterClient) GetFollowerListParams(user *twitter.User, cursor int64) twitter.FollowerListParams {
	return twitter.FollowerListParams{
		UserID:     user.ID,
		ScreenName: user.ScreenName,
		Cursor:     cursor,
		Count:      200,
	}
}

func (client *TwitterClient) getAccount() *twitter.User {
	return &client.user
}
