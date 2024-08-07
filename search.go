package go_llama_agentic_system

type SearchResult struct {
	Type            string                 `json:"type"`
	Title           string                 `json:"title"`
	URL             string                 `json:"URL"`
	Description     string                 `json:"description"`
	ExtraAttributes map[string]interface{} `json:",inline"`
}
