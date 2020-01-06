package potato_test

import "os"

func getAPIURL() string {
	apiURL := os.Getenv("POC_GO_API_URL")
	if apiURL != "" {
		return apiURL
	}

	return "https://potato.potatocoin.com"
}
