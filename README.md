# char-dump
A pastebin web-app written in go using fiber library.

## build dependencies

- [fiber](https://gofiber.io/)

## build instructions

```bash
#install dependencies
go get -d ./...

#compile program
go build

#alternatively, run this for an optimized release version
go build -ldflags "-s -w"

#run server
./char-dump

#to clean database.json, run
printf {}> database.json
```