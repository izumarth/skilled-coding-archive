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

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	// Get a woker pool
	wp := streamer.New(videoQueue, numWorkers)

	// start a worker pool
	wp.Run()
	fmt.Println("Worker pool started. Press enter to continue")
	_, _ = fmt.Scanln()

	// Create 4 videos to sent to worker pool
	video1 := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "mp4", notifyChannel, nil)
	video2 := wp.NewVideo(2, "./input/puppy1.mp4", "./output", "mp4", notifyChannel, nil)
	video3 := wp.NewVideo(3, "./input/puppy1.mp4", "./output", "mp4", notifyChannel, nil)
	video4 := wp.NewVideo(4, "./input/puppy1.mp4", "./output", "mp4", notifyChannel, nil)

	// send the videos to wokere pool
	videoQueue <- streamer.VideoProcessingJob{Video: video1}
	videoQueue <- streamer.VideoProcessingJob{Video: video2}
	videoQueue <- streamer.VideoProcessingJob{Video: video3}
	videoQueue <- streamer.VideoProcessingJob{Video: video4}

	for i := 1; i <= numJobs; i++ {
		msg := <-notifyChannel
		fmt.Println("i:", i, "msg:", msg)
	}

	fmt.Println("Done.")
}
