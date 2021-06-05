package main

import (
	"log"
	"os/exec"
	"path"
	"path/filepath"
)

func convertVideo(pattern string) {
	videos, err := filepath.Glob(pattern)

	if err != nil {
		log.Fatal(err)
		return
	}

	if videos == nil {
		return
	}

	ExecParWait(videos, func(video string) {
		ext := path.Ext(video)
		oggfile := video[0:len(video)-len(ext)] + ".ogg"

		cmd := exec.Command("/usr/bin/ffmpeg",
			"-i", video,
			"-ab", "192k",
			"-ar", "48000",
			"-ac", "2",
			"-vn",
			"-codec:a", "libopus",
			oggfile)

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	})
}

func main() {
	patterns := [3]string{"*.mp4", "*.webm", "*.mkv"}

	for i := 0; i < len(patterns); i++ {
		convertVideo(patterns[i])
	}
}
