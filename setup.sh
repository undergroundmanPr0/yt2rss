#!/bin/sh
echo "This is for setting up your tool!"
[[ ! -f go.mod ]] && go mod init yt2rss
go get google.golang.org/api/youtube/v3
go get google.golang.org/api/googleapi/transport
go mod tidy
go build && echo "Build complete!"
cp ./yt2rss /usr/bin/yt2rss
cp config.json /etc/config.json
