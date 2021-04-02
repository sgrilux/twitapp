package commands

type TwitAppCommand struct {
	Help HelpCommand `command:"help" description:"Print this help message"`

	ConsumerKey    string `long:"consumer-key" value-name:"ConsumerKey" env:"CONSUMER_KEY" description:"Twitter Consumer Key" required:"true"`
	ConsumerSecret string `long:"consumer-secret" value-name:"ConsumerSecret" env:"CONSUMER_SECRET" description:"Twitter Consumer Secret" required:"true"`
	AccessToken    string `long:"access-token" value-name:"AccessToken" env:"ACCESS_TOKEN" description:"Twitter Access Token" required:"true"`
	AccessSecret   string `long:"access-secret" value-name:"AccessSecret" env:"ACCESS_SECRET" description:"Twitter Access Secret" required:"true"`

	Tweet        TweetCommand         `command:"tweet" description:"Tweet a message"`
	ListFollower ListFollowerCommands `command:"list-follower" description:"Show a list of users following the specified user. If none is given it will show the list of followers of the current user"`

	Version func() `short:"v" long:"version" description:"Print the version of Fly and exit"`
}

var TwitApp TwitAppCommand
