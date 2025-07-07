package main

import (
	"bufio"
	"educationalsp/rpc"
	"log"
	"os"
)

func main() {
	logger := getLogger("/home/aditya/Projects/LSP/log.txt")
	logger.Println("Started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		mssg := scanner.Text()
		handleMssg(logger, mssg)
	}

}

func handleMssg(logger *log.Logger, mssg any) {
	logger.Println(mssg)

}

func getLogger(fileName string) *log.Logger {
	logfile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Didn't get a correct file")
	}
	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
