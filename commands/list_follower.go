package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/sgrilux/twitapp/pkg/apiclient"
	"github.com/sgrilux/twitapp/twitapp"
)

type ListFollowerCommands struct {
	Output string `short:"o" long:"output" description:"Print output to a file!" positional-args:"yes"`
}

func (command *ListFollowerCommands) Execute(args []string) error {
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

	followerList, err := client.GetFollowerList(user)
	if err != nil {
		return err
	}

	for _, f := range followerList {
		if err := printUser(output, f.ScreenName); err != nil {
			return err
		}
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
