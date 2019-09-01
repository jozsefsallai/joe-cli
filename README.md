# Joe CLI

My own CLI for everyday tasks. (WIP)

While in development, installation looks like this:

```sh
git clone git@github.com:jozsefsallai/joe-cli.git
cd joe-cli
go get -u github.com/golang/dep/cmd/dep
dep ensure
go build -o -i joe github.com/jozsefsallai/joe-cli
mv ./joe /somewhere/in/your/PATH
```

## License

MIT.
