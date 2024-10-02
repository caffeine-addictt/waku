package license

import (
	"io"
	"net/http"

	"github.com/caffeine-addictt/waku/pkg/version"
	"github.com/goccy/go-json"
)

const (
	LICENSE_LIST = "license.json"
	BASE_URL     = "https://raw.githubusercontent.com/caffeine-addictt/waku/v" + version.Version + "/licenses/"
)

// The global "cache" per say so we only
// need to hit the endpoint once per session.
var Licenses *[]License

// GetLicenses returns the list of licenses from the GitHub API
// or returns the cached list if it exists.
func GetLicenses() (*[]License, error) {
	if Licenses != nil {
		return Licenses, nil
	}

	req, err := http.NewRequest(http.MethodGet, BASE_URL+LICENSE_LIST, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/plain")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var l struct {
		Licenses []License `json:"licenses"`
	}
	if err := json.Unmarshal(body, &l); err != nil {
		return nil, err
	}

	Licenses = &l.Licenses
	return Licenses, nil
}
