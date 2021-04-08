package commands

import (
	"fmt"

	"github.com/sgrilux/twitapp/pkg/apiclient"
	"github.com/sgrilux/twitapp/twitapp"
)

type StatsCommand struct{}

func (command *StatsCommand) Execute(args []string) error {
	creds := twitapp.Credentials{
		ConsumerKey:    TwitApp.ConsumerKey,
		ConsumerSecret: TwitApp.ConsumerSecret,
		AccessToken:    TwitApp.AccessToken,
		AccessSecret:   TwitApp.AccessSecret,
	}

	client, err := apiclient.NewTwitterClient(creds)
	if err != nil {
		return err
	}

	user, err := client.GetUser("")
	if err != nil {
		return err
	}

	followerList, err := client.GetFollowerList(user)
	if err != nil {
		return err
	}

	followingList, err := client.GetFollowingList(user)
	if err != nil {
		return err
	}

	fmt.Println("Statistics")
	fmt.Printf("Followers: %d\n", len(followerList))
	fmt.Printf("Following: %d\n", len(followingList))

	return nil
}
