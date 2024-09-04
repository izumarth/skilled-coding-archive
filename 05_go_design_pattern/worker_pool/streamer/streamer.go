package streamer

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/tsawler/toolbox"
)

type ProcessingMessage struct {
	ID         int
	Successful bool
	Message    string
	OutputFile string
}

type VideoProcessingJob struct {
	Video Video
}

type Processor struct {
	Engine Encoder
}

type Video struct {
	ID           int
	InputFile    string
	OutputDir    string
	EncodingType string
	NotifyChan   chan ProcessingMessage
	Options      *VideoOptions
	Encoder      Processor
}

type VideoOptions struct {
	RenameOutput    bool
	SegmentDuration int
	MaxRate1080p    string
	MaxRate720p     string
	MaxRate480p     string
}

func (vd *VideoDispatcher) NewVideo(id int, input, output, encTYpe string, notifyChan chan ProcessingMessage, ops *VideoOptions) Video {
	if ops == nil {
		ops = &VideoOptions{}
	}

	fmt.Println("NewVideo: New Video created:", id, input)

	return Video{
		ID:           id,
		InputFile:    input,
		OutputDir:    output,
		EncodingType: encTYpe,
		NotifyChan:   notifyChan,
		Encoder:      vd.Processor,
		Options:      ops,
	}
}

func (v *Video) encode() {
	var fileName string

	switch v.EncodingType {
	case "mp4":
		// encode the video
		fmt.Println("v.endode(): About to encode to mp4", v.ID)
		name, err := v.encodeToMP4()
		if err != nil {
			// send information to the notifyChan
			v.sendToNotifyChan(false, "", fmt.Sprintf("encode failed for %d: %s", v.ID, err.Error()))
			return
		}
		fileName = fmt.Sprintf("%s.mp4", name)

	default:
		fmt.Println("v.endode(): error trying to encode video", v.ID)
		v.sendToNotifyChan(false, "", fmt.Sprintf("error proccesing for %d", v.ID))
		return
	}
	fmt.Println("v.endode(): sending success message ", v.ID)
	v.sendToNotifyChan(true, fileName, fmt.Sprintf("video id %d processed and saved as %s", v.ID, v.OutputDir))
}

func (v *Video) encodeToMP4() (string, error) {
	baseFileName := ""
	fmt.Println("v.encodeToMP4():about to try to encode video id", v.ID)
	if !v.Options.RenameOutput {
		// Get the basefile name
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, filepath.Ext(b))
	} else {
		var t toolbox.Tools
		baseFileName = t.RandomString(10)
	}

	err := v.Encoder.Engine.EncodeToMP4(v, baseFileName)
	if err != nil {
		return "", err
	}
	fmt.Println("v.encodeToMP4(): succesesfullluy encoded video id", v.ID)

	return baseFileName, nil
}

func (v *Video) sendToNotifyChan(successful bool, fileName, message string) {
	fmt.Println("v.sendToNotifyChan(): sending message to nofitfyc Chan for video id", v.ID)
	v.NotifyChan <- ProcessingMessage{
		ID:         v.ID,
		Successful: successful,
		Message:    message,
		OutputFile: fileName,
	}
}

func New(jobQueue chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {
	fmt.Println("v.New(): creating worker pool")
	workerPool := make(chan chan VideoProcessingJob, maxWorkers)

	var e VideoEncoder
	p := Processor{
		Engine: &e,
	}

	return &VideoDispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		WorkerPool: workerPool,
		Processor:  p,
	}
}
