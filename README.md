# go-stock-bot
go-stock-bot is a slack chat bot based on [mybot](https://github.com/rapidloop/mybot)


## mybot (original README)

`mybot` is an working Slack bot written in Go. Fork it and use it to build
your very own cool Slack bot!

Check the [blog post](https://www.opsdash.com/blog/slack-bot-in-golang.html)
for a description of mybot internals.

Follow us on Twitter today! [@therapidloop](https://twitter.com/therapidloop)

## go-stock-bot modifications
* no triggering command, just use '$' or '￥' to prefix your stock symbol to trigger bot action

	```
you> Oh, $TEAM is up!
bot> Atlassian Corporation Plc (TEAM) is trading at $23.85 (change +2.14%)
	```
* displaying change percentages along with current prices
* support multiple symbols in one message
* added unit testing
