package main

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	startTime := time.Now()
	imgCounter := 0

	// Create screenshots folder if it doesn't exist
	if _, err := os.Stat("screenshots"); os.IsNotExist(err) {
		err := os.Mkdir("screenshots", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for {
		// Take screenshot
		bounds := screenshot.GetDisplayBounds(0)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Create new file with incremented number in filename
		imgCounter++
		imgName := fmt.Sprintf("image_%d.png", imgCounter)
		imgPath := filepath.Join("screenshots", imgName)
		imgFile, err := os.Create(imgPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer imgFile.Close()

		// Encode image and save to file
		err = png.Encode(imgFile, img)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Calculate elapsed time and sleep for 10 seconds
		elapsedTime := time.Since(startTime)
		fmt.Printf("Took screenshot %s, elapsed time: %v\n", imgPath, elapsedTime)
		time.Sleep(10 * time.Second)
	}
}
