package license

import (
	"fmt"
	"io"
	"net/http"

	"github.com/caffeine-addictt/waku/cmd/options"
	"github.com/caffeine-addictt/waku/pkg/log"
	"github.com/caffeine-addictt/waku/pkg/version"
	"github.com/goccy/go-json"
)

const (
	LICENSE_LIST = "license.json"
	BASE_URL     = "https://raw.githubusercontent.com/caffeine-addictt/waku/%s/licenses/"
)

// The global "cache" per say so we only
// need to hit the endpoint once per session.
var Licenses *[]License

func GetLicenseFetchUrl() string {
	var url string
	if options.NewOpts.Branch.Value() != "" {
		url = fmt.Sprintf(BASE_URL, options.NewOpts.Branch.Value())
	} else {
		url = fmt.Sprintf(BASE_URL, "v"+version.Version)
	}
	return url
}

// GetLicenses returns the list of licenses from the GitHub API
// or returns the cached list if it exists.
func GetLicenses() (*[]License, error) {
	if Licenses != nil {
		return Licenses, nil
	}

	url := GetLicenseFetchUrl() + LICENSE_LIST
	log.Infof("fetching licenses from %s...\n", url)
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/plain")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	log.Debugln("reading http stream")
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var l struct {
		Licenses []License `json:"licenses"`
	}

	log.Debugln("unmarshalling license json")
	if err := json.Unmarshal(body, &l); err != nil {
		return nil, err
	}

	Licenses = &l.Licenses
	return Licenses, nil
}
