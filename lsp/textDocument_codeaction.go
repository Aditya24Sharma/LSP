package lsp

type CodeActionsRequest struct {
	Request
	Params CodeActionsParams `json:"params"`
}

type CodeActionsParams struct {
	TextDocumentIdentifier TextDocumentIdentifier `json:"textDocument"`
	Range                  Range                  `json:"range"`
	Context                CodeActionContext      `json:"context"`
}

type CodeActionContext struct {
}

type CodeActionsResponse struct {
	Response
	Result []CodeAction `json:"result"`
}

type CodeAction struct {
	Title string         `json:"title"`
	Edit  *WorkSpaceEdit `json:"edit"`
}

type CodeActionsResult struct {
	Contents string `json:"contents"`
}
