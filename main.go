package main

import (
	"bufio"
	"educationalsp/lsp"
	"educationalsp/rpc"
	"encoding/json"
	"log"
	"os"
)

func main() {
	logger := getLogger("/home/aditya/Projects/LSP/log.txt")
	logger.Println("Started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error %s", err)
			continue
		}
		handleMsg(logger, method, content)
	}

}

func handleMsg(logger *log.Logger, method string, content []byte) {
	logger.Printf("We received message with Methods: %s", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("We have an error in unmarshalling the request: %s", err)
		}
		logger.Printf("Connected to: %s  %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version,
		)
		//Lets reply
		msg := lsp.NewInitializeResponse(request.Id)
		// logger.Println("The msg got from NewInitializeResponse")
		// logger.Println(msg)
		reply := rpc.EncodeMessage(msg)
		// logger.Println("The reply got from Encode Message")
		// logger.Println(reply)

		write := os.Stdout
		write.Write([]byte(reply))

		logger.Printf("Sent the reply!")

	}

}

func getLogger(fileName string) *log.Logger {
	logfile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Didn't get a correct file")
	}
	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
