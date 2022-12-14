# yt2rss

A simple command line tool to fetch the video links and details of your favorite YouTube channel into an RSS(Really Simple Syndication) feed on your local machine which you can use with different RSS readers like [Newsboat](https://newsboat.org/), etc.

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/dwyl/auth_plug/Elixir%20CI?label=build&style=flat-square)

## Problem

As a newsboat user on Linux(running Ubuntu) I faced this problem with using the RSS feed that I fetched from YouTube using https://www.youtube.com/feeds/videos.xml?channel_id=UCH8JwgaHCkhdfERVkGbLl2g . 

When I added this link to my newsboat urls list, it rendered only the 15 latest videos on this channel, while the videos on the channel count to about 400. 

I tried finding the solution to get all videos uploaded on a YouTube channel. 

This took me some time to look around to finally come across to use the youtube api in order to fetch all the videos on a channel. 

I took the reference of this answer on Stack Overflow https://stackoverflow.com/questions/18953499/youtube-api-to-fetch-all-videos-on-a-channel I liked the answer with the PHP code but since I have the interest in learning Golang, I chose to create my own tool to convert the response fetched using the YouTube API into an RSS xml file.

## Prerequisites and Installation

Before using this tool you need to have some prerequisites as listed below:

1. First get the repository cloned or downloaded on your machine using `git clone https://github.com/undergroundmanPr0/yt2rss.git` from the `main` branch in order to build it from source. To do so, first download and install git using `sudo apt install git-all`. After cloning the repository, change the directory using `cd yt2rss`.
2. You will need a YouTube Data API v3 key which you can get by following this video https://www.youtube.com/watch?v=TE66McLMMEw watching upto 3:21.
3. After you get your API key make sure that you update the `./src/config.json` replacing `[API_KEY]` with your own API key and changing the `[LOCATION]` to a preferred one. Keeping the location blank will create the RSS xml file in the same directory where you will run this command line.
4. Make sure that you have Golang already installed on your machine. If not just run the command on your terminal `sudo apt install golang-go` The Golang will be installed on your machine. The version of Golang used to create this tool is `go1.18.1 linux/amd64`.
5.  This step is also important, that is getting the channel id of your favorite YouTube channel. Go to your favorite YouTube creator's channel page and copy the channel URL and paste in https://commentpicker.com/youtube-channel-id.php to find the channel id. Or you can also search by right clicking on the channel page and selecting view page source and finding channel_id in there.
6. After having followed the above steps run `sudo bash ./src/setup.sh` on your terminal.

Now, you're all set to run this tool on your local machine and get the RSS feed of any YouTube channel page.

## Usage

In order to use the tool run the following command replacing `[CHANNEL_ID]` and `[FILE_NAME]` with the channel id you fetched above and the file name where of your choice where you want to save the RSS feed and hit `Enter`.

	yt2rss get -c [CHANNEL_ID] -f [FILE_NAME]

You will receive a file link as `file://[LOCATION]/[FILE_NAME].xml` which you can copy and paste in your RSS feed reader url list.

In order to update your RSS feed file for a particular [CHANNEL_ID], just run the above command with the same [CHANNEL_ID] and [FILE_NAME].

## For updating the project

In order to update the project and synchronise your cloned repository with the latest version of the project, inside `yt2rss` folder, run `git pull`. 

You will get the latest updates of this project from time to time.

## For contributing

If you find any sort of issues while using yt2rss raise one here [Issues](https://github.com/undergroundmanPr0/yt2rss/issues)

If you want to work on some features or enhancement on yt2rss then consider reading the [CONTRIBUTING.md](https://github.com/undergroundmanPr0/yt2rss/blob/main/docs/CONTRIBUTING.md)
