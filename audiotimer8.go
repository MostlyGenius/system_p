package main

import (
	"log"
	"os"
	"time"

	"golang.org/x/sys/windows/svc"
)

type myService struct{}

func (m *myService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (svcSpecificEC bool, exitCode uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown

	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

mainLoop:
	for {
		select {
		case req := <-r:
			switch req.Cmd {
			case svc.Stop, svc.Shutdown:
				changes <- svc.Status{State: svc.StopPending}
				break mainLoop
			default:
				log.Printf("unexpected control request #%d", req)
			}
		default:
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			_, err := file.WriteString(currentTime + "\n")
			if err != nil {
				log.Println("Error writing to the file:", err)
			}
			time.Sleep(1 * time.Minute)
		}
	}

	return
}

func main() {
	isInteractive, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Fatalf("failed to determine if we are running in an interactive session: %v", err)
	}
	if isInteractive {
		log.Println("This program is designed to run as a Windows Service. Please install and start the service.")
		return
	}

	runService("MyTimeWriterSvc", &myService{})
}

func runService(name string, s svc.Handler) {
	err := svc.Run(name, s)
	if err != nil {
		log.Printf("Service %s failed: %v", name, err)
	}
}
