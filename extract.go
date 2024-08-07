package go_llama_agentic_system

import (
	"encoding/json"
	"regexp"
	"strings"
)

var BuiltinToolPattern = regexp.MustCompile(`\b(?P<tool_name>\w+)\.call\(query="(?P<query>[^"]*)"\)`)
var CustomToolCallPattern = regexp.MustCompile(`<function=(?P<function_name>[^}]+)>(?P<args>{.*?})`)

func ExtractBuiltinTool(s string) (BuiltinTool, string) {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, IPython) {
		s = strings.TrimPrefix(s, IPython)
	}
	if strings.HasSuffix(s, EndOfMessage) {
		s = strings.TrimSuffix(s, EndOfMessage)
	}
	// Find matches
	matches := BuiltinToolPattern.FindStringSubmatch(s)

	if len(matches) < 2 {
		return "", ""
	}
	// Get the index of named capturing groups
	toolNameIndex := BuiltinToolPattern.SubexpIndex("tool_name")
	queryIndex := BuiltinToolPattern.SubexpIndex("query")

	// Extract the values
	toolName := matches[toolNameIndex]
	query := matches[queryIndex]
	return toolName, query
}

func ExtractCustomTool(s string) (string, json.RawMessage) {
	s = strings.TrimSpace(s)
	if strings.HasSuffix(s, EndOfMessage) {
		s = strings.TrimSuffix(s, EndOfMessage)
	}
	// Find matches
	matches := CustomToolCallPattern.FindStringSubmatch(s)

	if len(matches) < 2 {
		return "", nil
	}
	// Get the index of named capturing groups
	functionNameIndex := BuiltinToolPattern.SubexpIndex("function_name")
	argsIndex := BuiltinToolPattern.SubexpIndex("args")

	// Extract the values
	functionName := matches[functionNameIndex]
	argsJson := matches[argsIndex]

	return functionName, json.RawMessage(argsJson)
}
