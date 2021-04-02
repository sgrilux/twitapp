# twitapp

This tool allows you to interact with Twitter. Here you can find the list of supported funcionalities:

- Send tweets
- Print to STDOUT or to a file yours or someone else's followers 

## Prerequisites

To allow the tool to interact with twitter you need to apply for a developer account and generate Consumer Key and Secret as well as Access Token and Secret.
For more details please follow Twitter developer [documentation](https://developer.twitter.com/en/docs/developer-portal/overview)

## Run

### Docker

```bash
docker run -it --rm \
    -e CONSUMER_KEY=<twitter_consumer_key> \
    -e CONSUMER_SECRET=<twitter_consumer_secret> \
    -e ACCESS_TOKEN=<twitter_access_token> \
    -e ACCESS_SECRET=<twitter_access_secretZ \
    sgrilux/twitapp <command>
```

For example:

This command print a list of followers into the output folder.
```
docker run -it --rm \
    -e CONSUMER_KEY=<twitter_consumer_key> \
    -e CONSUMER_SECRET=<twitter_consumer_secret> \
    -e ACCESS_TOKEN=<twitter_access_token> \
    -e ACCESS_SECRET=<twitter_access_secretZ \
    -v $(pwd)/output:/tmp \ 
    twitapp list-follower -o /tmp/list-followers-$(date +%Y%m%d).txt
```

### From Source

```bash
git clone git@github.com:sgrilux/twitapp.git

cd twitapp
make build_mac 
# make build_linux

bin/twitapp <command>
```

## Usage

```bash
Usage:
  twitapp [OPTIONS] <help | list-follower | tweet>

Application Options:
      --consumer-key=ConsumerKey          Twitter Consumer Key [$CONSUMER_KEY]
      --consumer-secret=ConsumerSecret    Twitter Consumer Secret
                                          [$CONSUMER_SECRET]
      --access-token=AccessToken          Twitter Access Token [$ACCESS_TOKEN]
      --access-secret=AccessSecret        Twitter Access Secret [$ACCESS_SECRET]
  -v, --version                           Print the version of Fly and exit

Help Options:
  -h, --help                              Show this help message

Available commands:
  help           Print this help message
  list-follower  Show a list of users following the specified user. If none is given it will show the list of followers of the current user
  tweet          Tweet a message
```