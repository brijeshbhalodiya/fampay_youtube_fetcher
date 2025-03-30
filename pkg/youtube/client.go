package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func (c *YoutubeClient) FetchVideos(query string, maxResults int, fetchVideoAfterUTC string) (*SearchResponse, error) {
	parsedUrl, err := url.Parse(BASE_URL + "/search")
	if err != nil {
		fmt.Println(err)
	}

	queryParams := parsedUrl.Query()
	queryParams.Set("publishedAfter", fetchVideoAfterUTC)
	queryParams.Set("part", "snippet")
	queryParams.Set("order", "date")
	queryParams.Set("q", query)
	queryParams.Set("maxResults", "5")
	queryParams.Set("key", c.apiKey)
	parsedUrl.RawQuery = queryParams.Encode()

	req, err := http.NewRequest("GET", parsedUrl.String(), nil)
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
