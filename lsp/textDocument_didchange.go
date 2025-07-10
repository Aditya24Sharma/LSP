package lsp

type DidChangeTextDocumentNotification struct {
	Notification
	Params `json:"params"`
}

type Params struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}
