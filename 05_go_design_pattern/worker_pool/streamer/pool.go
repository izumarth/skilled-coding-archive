package streamer

import "fmt"

type VideoDispatcher struct {
	WorkerPool chan chan VideoProcessingJob
	maxWorkers int
	jobQueue   chan VideoProcessingJob
	Processor  Processor
}

// new VideoWorker
type videoWorker struct {
	id         int
	jobQueue   chan VideoProcessingJob
	workerPool chan chan VideoProcessingJob
}

// type videoWorker
func newVideoWorker(id int, workerPool chan chan VideoProcessingJob) videoWorker {
	fmt.Println("newVideoWorker(): creating video worker id", id)
	return videoWorker{
		id:         id,
		jobQueue:   make(chan VideoProcessingJob),
		workerPool: workerPool,
	}
}

// start()
func (w videoWorker) start() {
	fmt.Println("w.start(): starting worker ", w.id)
	go func() {
		for {
			// Add jobQueue to the worker pool.
			w.workerPool <- w.jobQueue

			// Wait for  a job come back.
			job := <-w.jobQueue

			// Process the job
			w.processVideoJob(job.Video)
		}
	}()
}

// Run()
func (vd *VideoDispatcher) Run() {
	fmt.Println("vd.Run(): starting worker pool by running workers")
	for i := 0; i < vd.maxWorkers; i++ {
		fmt.Println("vd.Run(): starting worker id", i+1)
		worker := newVideoWorker(i+1, vd.WorkerPool)
		worker.start()
	}

	go vd.dispatch()
}

// dispatch
func (vd *VideoDispatcher) dispatch() {
	for {
		// wait for a job to come in
		job := <-vd.jobQueue
		fmt.Println("vd.dispatch(): sending job", job.Video.ID, "to wokere job queue")

		go func() {
			workerJobQueue := <-vd.WorkerPool
			workerJobQueue <- job
		}()
	}
}

// processVideoJob
func (w videoWorker) processVideoJob(video Video) {
	fmt.Println("w.processVideoJob(): starting encode on ", video.ID)
	video.encode()
}
