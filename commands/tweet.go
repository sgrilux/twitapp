package commands

import (
	"github.com/sgrilux/twitapp/pkg/apiclient"
	"github.com/sgrilux/twitapp/twitapp"
)

type TweetCommand struct{}

func (command *TweetCommand) Execute(args []string) error {
	creds := twitapp.Credentials{
		ConsumerKey:    TwitApp.ConsumerKey,
		ConsumerSecret: TwitApp.ConsumerSecret,
		AccessToken:    TwitApp.AccessToken,
		AccessSecret:   TwitApp.AccessSecret,
	}

	message := args[0]

	client, err := apiclient.NewTwitterClient(creds)
	if err != nil {
		return err
	}

	if _, _, err := client.Tweet(message); err != nil {
		return err
	}
	return nil
}
