package main

import (
	"fmt"
	"streamer"
)

func main() {
	const numJobs = 4
	const numWorkers = 2

	notifyChannel := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChannel)

	videoQueue := make(chan  streamer.VideoProcessingJob, numJobs))
	defer close(videoQueue)

	// Get a woker pool
	wp := streamer.New(videoQueue, numWorkers)
	fmt.Println("wp:", wp)

	// start a worker pool
	

	// Create 4 videos to sent to worker pool

	// send the videos to wokere pool
}
