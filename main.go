package main

import (
	"log"
	"ps-checker/utils"
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

	utils.PrintLog("nabersin")

	//hasSudo := true
	//
	//isInfinite := true
	//
	//
	//for isInfinite {
	//
	//	fmt.Println("checking ps...")
	//	_, err := utils.CheckPS(psDestination)
	//
	//	if err != nil {
	//		//if error exists, that means process not found
	//
	//		stopOP, err := utils.ProcessStop(hasSudo, stopCommands)
	//		if err != nil {
	//			fmt.Println(" stop error verdi", err.Error())
	//		}
	//
	//		fmt.Println(stopOP)
	//		startOP, err := utils.ProcessStart(hasSudo, startCommands)
	//		if err != nil {
	//			fmt.Println("start error verdi", err.Error())
	//		}
	//
	//		fmt.Println(startOP)
	//	} else {
	//		fmt.Printf("%s is working", psDestination)
	//	}
	//
	//	time.Sleep(time.Second * 5)
	//}

}
