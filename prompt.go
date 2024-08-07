package go_llama_agentic_system

// Much of the code comments below were copied or derived from
// https://llama.meta.com/docs/model-cards-and-prompt-formats/llama3_1/. Go to the model card for more info

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func CreatePrompt(msgs []Message) string {
	var p strings.Builder
	p.WriteString(BeginningOfText)
	for _, msg := range msgs {
		p.WriteString(StartHeader)
		p.WriteString(msg.Role)
		p.WriteString(EndHeader)
		p.WriteString("\n\n")

		p.WriteString(msg.Content)
		if msg.IPython() || msg.CustomFunctionCall() {
			p.WriteString(EndOfMessage)
		}
	}
	return p.String()
}

type SystemPromptParams struct {
	IPythonEnabled bool
	BuiltInTools   []BuiltinTool
	Functions      []Function
	Instruction    string
}

func SystemPrompt(sp SystemPromptParams) string {
	var s strings.Builder
	// code interpreter
	if sp.IPythonEnabled {
		s.WriteString("Environment: ipython\n")
	}

	// built in tools
	if len(sp.BuiltInTools) > 0 {
		s.WriteString("Tools: ")
		s.WriteString(strings.Join(sp.BuiltInTools, ", "))
		s.WriteRune('\n')
	}

	// Knowledge cut off
	s.WriteString("\nCutting Knowledge Date: December 2023\nToday Date: ")
	s.WriteString(time.Now().Format("02 Jan 2006"))
	s.WriteString("\n\n")

	// Custom functions
	if len(sp.Functions) > 0 {
		s.WriteString("You have access to the following functions:\n\n")
		for _, f := range sp.Functions {
			s.WriteString(fmt.Sprintf("Use the function '%s' to '%s'\n", f.Name, f.Description))
			jsonDef, _ := json.Marshal(f.Definition)
			s.WriteString(string(jsonDef))
			s.WriteString("\n\n")
		}
		s.WriteString(
			`Think very carefully before calling functions.
If you choose to call a function ONLY reply in the following format with no prefix or suffix:

<function=example_function_name>{{"example_name": "example_value"}}</function>

Reminder:
- If looking for real time information use relevant functions before falling back to brave_search
- Function calls MUST follow the specified format, start with <function= and end with </function>
- Required parameters MUST be specified
- Only call one function at a time
- Put the entire function call reply on one line

`,
		)
	}

	// Custom agent instructions
	if sp.Instruction == "" {
		sp.Instruction = "You are a helpful Assistant."
	}
	s.WriteString(sp.Instruction)
	s.WriteRune('\n')
	return s.String()
}
