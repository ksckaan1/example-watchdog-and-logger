package main

import (
	"fmt"
	"log"
	"ps-checker/utils"
	"time"
)

var (
	psDestination = "face-reg"

	stopCommands  = []string{"bash", "/var/Face-Recognition/stop.sh"}
	startCommands = []string{"bash", "/var/Face-Recognition/start.sh"}
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}

func main() {

	hasSudo := true

	isInfinite := true

	for isInfinite {

		fmt.Println("checking ps...")
		_, err := utils.CheckPS(psDestination)

		if err != nil {
			//if error exists, that means process not found

			utils.PrintLog(psDestination + " not found. restarting...")

			stopOP, err := utils.ProcessStop(hasSudo, stopCommands)
			if err != nil {
				utils.PrintLog("Error when playing stop.sh command / err: " + err.Error())

			}

			fmt.Println(stopOP)
			startOP, err := utils.ProcessStart(hasSudo, startCommands)
			if err != nil {
				utils.PrintLog("Error when playing start.sh command / err: " + err.Error())
			}

			fmt.Println(startOP)
		} else {
			fmt.Printf("%s is working", psDestination)
		}

		time.Sleep(time.Second * 5)
	}

}
