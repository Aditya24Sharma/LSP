package main

import (
	"bufio"
	"educationalsp/analysis"
	"educationalsp/lsp"
	"educationalsp/rpc"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	logger := getLogger("/home/aditya/Projects/LSP/log.txt")
	logger.Println("Started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error %s", err)
			continue
		}
		handleMsg(logger, writer, state, method, content)
	}

}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func handleMsg(logger *log.Logger, writer io.Writer, state analysis.State, method string, content []byte) {
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

		writeResponse(writer, msg)
		logger.Printf("Sent the reply!")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("textDocument/didOpen error %s", err)
			return
		}

		logger.Printf("Opened: %s", request.Params.TextDocument.URI)
		diagnostics := state.OpenDocument(
			request.Params.TextDocument.URI,
			request.Params.TextDocument.Text,
		)
		writeResponse(writer, lsp.PublishDiagnosticsNotification{
			Notification: lsp.Notification{
				RPC:    "2.0",
				Method: "textDocument/publishDiagnostics",
			},
			Params: lsp.PublishDiagnosticsParams{
				URI:         request.Params.TextDocument.URI,
				Diagnostics: diagnostics,
			},
		})
	case "textDocument/didChange":
		var request lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("textDocument/didChange error %s", err)
			return
		}

		logger.Printf("Changed : %s", request.Params.TextDocument.URI)

		for _, change := range request.ContentChanges {
			diagnostics := state.UpdateDocument(
				request.Params.TextDocument.URI,
				change.Text,
			)
			writeResponse(writer, lsp.PublishDiagnosticsNotification{
				Notification: lsp.Notification{
					RPC:    "2.0",
					Method: "textDocument/publishDiagnostics",
				},
				Params: lsp.PublishDiagnosticsParams{
					URI:         request.TextDocument.URI,
					Diagnostics: diagnostics,
				},
			},
			)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("textDocument/hover error %s", err)
			return
		}
		response := state.Hover(request.Id, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response)

	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("textDocument/definition error %s", err)
			return
		}
		response := state.Definition(request.Id, request.Params.TextDocument.URI, request.Params.Position)
		writeResponse(writer, response)

	case "textDocument/codeAction":
		var request lsp.CodeActionsRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("textDocument/codeAction error %s", err)
			return
		}
		response := state.CodeAction(request.Id, request.Params.TextDocument.URI)
		writeResponse(writer, response)

	case "textDocument/completion":
		var request lsp.CompletionRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("textDocument/completion error %s", err)
			return
		}
		response := state.Completion(request.Id, request.Params.TextDocument.URI)
		writeResponse(writer, response)
	}

}
func getLogger(fileName string) *log.Logger {
	logfile, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Didn't get a correct file")
	}
	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
