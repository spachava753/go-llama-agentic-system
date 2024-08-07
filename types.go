package go_llama_agentic_system

import (
	"encoding/json"
	"strings"
)

type ControlToken = string

const (
	// BeginningOfText specifies the start of the prompt
	BeginningOfText ControlToken = "<|begin_of_text|>"
	// EndOfText means model will cease to generate more tokens. This token is generated only by the base models.
	EndOfText ControlToken = "<|end_of_text|>"
	// FinetuneRightPad token is used for padding text sequences to the same length in a batch
	FinetuneRightPad ControlToken = "<|finetune_right_pad_id|>"
	// StartHeader token starts header to define the role for a particular message.
	// The possible roles are: [system, user, assistant and ipython]
	StartHeader ControlToken = "<|start_header_id|>"
	// EndHeader token ends header to define the role for a particular message.
	// The possible roles are: [system, user, assistant and ipython]
	EndHeader ControlToken = "<|end_header_id|>"
	// EndOfMessage means "end of message". A message represents a possible stopping point for execution where the model
	// can inform the executor that a tool call needs to be made. This is used for multi-step interactions between the
	// model and any available tools. This token is emitted by the model when the Environment: ipython instruction is
	// used in the system prompt, or if the model calls for a built-in tool.
	EndOfMessage ControlToken = "<|eom_id|>"
	// EndOfTurn means "end of turn". Represents when the model has determined that it has finished interacting with the
	// user message that initiated its response. This is used in two scenarios:
	// - at the end of a direct interaction between the model and the user
	// - at the end of multiple interactions between the model and any available tools
	//
	// This token signals to the executor that the model has finished generating a response
	EndOfTurn ControlToken = "<|eot_id|>"
	// IPython is a special tag used in the model’s response to signify a built in tool call or use the code interpreter.
	IPython ControlToken = "<|python_tag|>"
)

type Role = string

const (
	SystemRole    Role = "system"
	AssistantRole Role = "assistant"
	UserRole      Role = "user"
	IPythonRole   Role = "ipython"
)

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

func (m Message) IPython() bool {
	return strings.HasPrefix(m.Content, IPython)
}

func (m Message) CustomFunctionCall() bool {
	funcName, _ := ExtractCustomTool(m.Content)
	// might not have any arguments, but will definitely have a function name
	return len(funcName) > 0
}

type BuiltinTool = string

const (
	BraveSearch  BuiltinTool = "brave_search"
	WolframAlpha BuiltinTool = "wolfram_alpha"
)

type Function struct {
	Name, Description string
	Definition        json.RawMessage
}
