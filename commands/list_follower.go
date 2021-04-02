package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/sgrilux/twitapp/pkg/apiclient"
)

type ListFollowerCommands struct {
	Output string `short:"o" long:"output" description:"Print output to a file!" positional-args:"yes"`
}

func (command *ListFollowerCommands) Execute(args []string) error {
	creds := apiclient.Credentials{
		ConsumerKey:    TwitApp.ConsumerKey,
		ConsumerSecret: TwitApp.ConsumerSecret,
		AccessToken:    TwitApp.AccessToken,
		AccessSecret:   TwitApp.AccessSecret,
	}

	client, err := apiclient.NewTwitterClient(creds)
	if err != nil {
		return err
	}

	var userName string
	if len(args) > 0 {
		userName = args[0]
	}

	user, err := client.GetUser(userName)
	if err != nil {
		return err
	}

	var output *os.File
	if command.Output != "" {
		output, err = os.Create(command.Output)
		if err != nil {
			log.Fatal(err)
		}

		defer output.Close()
	}
	// Print followers
	var cursor int64

	cursor = -1
	for {
		followers, _, err := client.GetFollowerList(client.GetFollowerListParams(user, cursor))
		if err != nil {
			return err
		}

		for _, f := range followers.Users {
			if err := printUser(output, f.ScreenName); err != nil {
				return err
			}
		}

		cursor = followers.NextCursor
		if cursor == 0 {
			break
		}
		// time.Sleep(1 * time.Second)
	}
	return nil
}

func printUser(file *os.File, user string) error {
	if file != nil {
		_, err := file.WriteString(user + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file")
		}
	} else {
		fmt.Printf("%s\n", user)
	}
	return nil
}
