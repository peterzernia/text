# Text

A simple command-line interface to send a Telegram message

## Usage

```
text me 'hello world'
```

```
docker build . && text me 'build complete'
```

Text can be used in combination with [wait-on](https://github.com/jeffbski/wait-on) to wait for internet to come back:

```
wait-on https://news.ycombinator.com/ && text me 'internet is back'
```

## Setup

- Create a new [telegram bot](https://core.telegram.org/bots#6-botfather)
- Start a chat with your new bot
- Get the [chat id](https://stackoverflow.com/a/32572159)
- Set TEXT_BOT_TOKEN=<your bot token> && TEXT_ME=<your chat id>
- run `make build` and `make copy` to build and move the binary into /usr/local/bin

Note: you can text other people/groups with the chat id specific to that user/group. Simply set the
env variable TEXT_PERSON=<other chat id> and run `text person 'hey'` where person is the username/group name
