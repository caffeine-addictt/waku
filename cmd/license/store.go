package license

import (
	"io"
	"net/http"

	"github.com/goccy/go-json"
)

const API_VERSION = "2022-11-28"

// The global "cache" per say so we only
// need to hit the endpoint once per session.
var Licenses *[]License

// GetLicenses returns the list of licenses from the GitHub API
// or returns the cached list if it exists.
func GetLicenses() (*[]License, error) {
	if Licenses != nil {
		return Licenses, nil
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/licenses", http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", API_VERSION)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var l []License
	if err := json.Unmarshal(body, &l); err != nil {
		return nil, err
	}

	Licenses = &l
	return Licenses, nil
}
