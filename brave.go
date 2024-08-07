package go_llama_agentic_system

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const BraveSearchUrl = "https://api.search.brave.com/res/v1/web/search"

type BraveSearch struct {
	ApiKey  string
	BaseUrl string
}

func NewBraveSearch(apiKey string) BraveSearch {
	return BraveSearch{ApiKey: apiKey, BaseUrl: BraveSearchUrl}
}

func (b BraveSearch) Query(q string, topK uint) ([]SearchResult, error) {
	reqUrl, _ := url.Parse(b.BaseUrl)
	payload := url.Values{}
	payload.Set("q", q)
	reqUrl.RawQuery = payload.Encode()
	req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	headers := make(http.Header)
	headers.Set("X-Subscription-Token", b.ApiKey)
	headers.Set("Accept-Encoding", "gzip")
	headers.Set("Accept", "application/json")
	req.Header = headers
	resp, respErr := http.DefaultClient.Do(req)
	if respErr != nil {
		return nil, respErr
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var reader io.Reader
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		gzipReader, gzipErr := gzip.NewReader(resp.Body)
		if gzipErr != nil {
			return nil, gzipErr
		}
		defer gzipReader.Close()
		reader = gzipReader
	default:
		reader = resp.Body
	}

	return b.cleanBraveResponse(reader, topK)
}

func (b BraveSearch) cleanBraveResponse(data io.Reader, k uint) ([]SearchResult, error) {
	content, _ := io.ReadAll(data)
	dec := json.NewDecoder(bytes.NewReader(content))
	results := make([]map[string]interface{}, 0)
	if decErr := dec.Decode(&results); decErr != nil {
		return nil, decErr
	}

	return nil, nil
}
