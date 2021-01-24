# mcsi-cli

mcsi-cli is a command line tool for getting information about minecraft servers via [this api](http://mcsrvstat.us/)

## Usage

`mcsi-cli` - shows the help page
`mcsi-cli <url>` - gives you the basic information about the minecraft server at the url
`mcsi-cli -a <url>` - gives you all of the information except for the icon

## Install

I'm not yet sure how this works, but i think that this will work:

First you need `go` to install this, check on [here](https://golang.org/dl/) for that
Then run this in your favorite posix shell. i don't know how to do this on windows, if you know, please submit a pull request
```bash
go get github.com/w1kee/mcsi-cli
cd $GOPATH/src/github.com/w1kee/mcsi-cli
go install
go build
sudo cp mcsi-cli /usr/bin
```
hope that works, i didn't check, but don't see a reason why it won't

## Contribute

if you want to contribute to this project or fix my noob go code, be my guest. i started learning go only a few days ago, so i'm probably doing some beginner errors

## Credits
- [Anders G. Jørgensen](https://spirit55555.dk/) for making [mcsrvstat.us](https://mcsrvstat.us/) which provides the api
- [w1ke](https://w1ke.cz/)
 for making the cli

## License

this program is licensed under the mit license, so it's open source
