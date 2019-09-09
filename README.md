# Joe CLI

My own CLI for everyday tasks. (WIP)

While in development, installation looks like this:

```sh
git clone git@github.com:jozsefsallai/joe-cli.git
cd joe-cli
go get -u github.com/golang/dep/cmd/dep
dep ensure
go build -i -o joe github.com/jozsefsallai/joe-cli
mv ./joe /somewhere/in/your/PATH
```

## Command List

  * joe upload [file] - requires an AWS IAM key+secret pair in ~/.joerc.json
  * joe weather now [lat?] [long?] - requires a [DarkSky](https://darksky.net/) API key in ~/.joerc.json
  * joe weather tomorrow [lat?] [long?] - requires a [DarkSky](https://darksky.net/) API key in ~/.joerc.json
  * joe ip - get your IPv4 and IPv6 addresses
  * joe ip v4 - get your IPv4 address
  * joe ip v6 - get your IPv6 address

## License

MIT.
