package main

import (

	"fmt"
	"google.golang.org/api/youtube/v3"
	"net/http"
	"log"
        "google.golang.org/api/googleapi/transport"
	"os"
	"strings"
)



func getResponse(service *youtube.Service, part string, channelID string) *youtube.ChannelListResponse {
	request := service.Channels.List([]string{part}).Id(channelID)
	response, err := request.Do()
	if err != nil {
		fmt.Println("Response not found enter correct channel ID")
		os.Exit(1)
	}
	return response
}

func playlistItemsList(service *youtube.Service, part string, playlistId string, pageToken string) *youtube.PlaylistItemListResponse {
	call := service.PlaylistItems.List([]string{part}).MaxResults(50)
	call = call.PlaylistId(playlistId)
	if pageToken != "" {
		call = call.PageToken(pageToken)
	}
	response, err := call.Do()
	if err != nil {
		fmt.Println("Error fetching videos")
		os.Exit(1)
	}
	return response
}

func removeForbiddenChar(content string) string {
	content = strings.Replace(content,  " & ", "and", -2)
	content = strings.Replace(content, "<", "less than", -2)
	return content
} 


func createRSSFeed(channelID string, apiID string, filename string) {

	client := &http.Client{
		Transport: &transport.APIKey{Key: apiID},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
		os.Exit(1)
	}

	response := getResponse(service, "snippet,contentDetails", channelID)
	
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		fmt.Println("Error creating/opening the file.")
		os.Exit(1)
	}

	count := 0
	writable := "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n<rss version=\"2.0\">\n"

	for _, channel := range response.Items {
		writable = writable + "<channel>\n"
		user := channel.Snippet.Title
		writable = writable + "<title>"+ user + "(yt2rss api)</title>\n"

		playlistId := channel.ContentDetails.RelatedPlaylists.Uploads
		writable = writable +"<link>" + "https://www.youtube.com/channel/"+playlistId+ "/</link>\n"

		nextPageToken := ""
		for {
			playlistResponse := playlistItemsList(service, "snippet", playlistId, nextPageToken)
			
			for _, playlistItem := range playlistResponse.Items {
				title := playlistItem.Snippet.Title
				title = removeForbiddenChar(title)
				description := playlistItem.Snippet.Description
				description = removeForbiddenChar(description)
				videoId := playlistItem.Snippet.ResourceId.VideoId
				pubDate := playlistItem.Snippet.PublishedAt
				count = count + 1
				writable = writable + "<item>\n" +"<title>"+title+"</title>\n" + "<description>"+description +"</description>\n" + "<link>https://www.youtube.com/watch?v="+videoId+"</link>\n" + "<pubDate>"+pubDate+"</pubDate>\n" + "</item>\n"
			}
			nextPageToken = playlistResponse.NextPageToken
			if nextPageToken == "" {
				break
			}
		}

	}

	writable = writable + "</channel>\n</rss>"

	length, err := file.WriteString(writable)

	if err != nil {
		fmt.Println("Error writing into the file", length)
		os.Exit(1)
	}
	fmt.Println("Videos fetched ", count)
	defer file.Close()

	fmt.Println("Link to RSS feed: file://"+filename)
	return
}
