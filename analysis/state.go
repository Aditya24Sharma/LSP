package analysis

import (
	"educationalsp/lsp"
	"fmt"
	"strings"
)

type State struct {
	Document map[string]string
}

func NewState() State {
	return State{Document: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Document[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Document[uri] = text
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	document := s.Document[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			Id:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Characters: %d, Position: %v", uri, len(document), position),
		},
	}
}

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			Id:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 2,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}

func (s *State) CodeAction(id int, uri string) lsp.CodeActionsResponse {

	text := s.Document[uri]

	action := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {

		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range: lsp.Range{
						Start: lsp.Position{
							Line:      row,
							Character: idx,
						},
						End: lsp.Position{
							Line:      row,
							Character: idx + len("VS Code"),
						},
					},
					NewText: "NeoVim",
				}}
			action = append(action, lsp.CodeAction{
				Title: "Replace VS Code with a superior editor",
				Edit: &lsp.WorkSpaceEdit{
					Changes: replaceChange,
				},
			})
			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{
					Range: lsp.Range{
						Start: lsp.Position{
							Line:      row,
							Character: idx,
						},
						End: lsp.Position{
							Line:      row,
							Character: idx + len("VS Code"),
						},
					},
					NewText: "VS C*de",
				}}
			action = append(action, lsp.CodeAction{
				Title: "Censor to VS C*de",
				Edit: &lsp.WorkSpaceEdit{
					Changes: censorChange,
				},
			})
		}
	}
	response := lsp.CodeActionsResponse{
		Response: lsp.Response{
			RPC: "2.0",
			Id:  &id,
		},
		Result: action,
	}
	return response
}
