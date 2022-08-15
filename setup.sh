echo "this is for setting up your api"
[[ ! -f go.mod ]] && go mod init yt2rss
go get google.golang.org/api/youtube/v3
go get google.golang.org/api/googleapi/transport
go mod tidy
go build && echo "Build complete! Now use test_script.sh"
cp ./yt2rss /usr/bin/yt2rss
