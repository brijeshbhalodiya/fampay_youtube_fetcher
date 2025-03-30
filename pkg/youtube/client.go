package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BASE_URL = "https://www.googleapis.com/youtube/v3"
)

type YoutubeClient struct {
	apiKey     string
	httpClient *http.Client
}

func NewYoutubeClient(apiKey string) *YoutubeClient {
	return &YoutubeClient{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

func (c *YoutubeClient) FetchVideos(ctx context.Context, query string, maxResults int) (*SearchResponse, error) {
	url := fmt.Sprintf("%s/search?part=snippet&q=%s&maxResults=%d&type=video&key=%s",
		BASE_URL, query, maxResults, c.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var searchResp SearchResponse
	if err := json.Unmarshal(body, &searchResp); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &searchResp, nil
}
