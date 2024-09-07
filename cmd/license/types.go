package license

import (
	"io"
	"net/http"
	"regexp"

	"github.com/goccy/go-json"
)

// Returned from the "api.github.com/licenses" endpoint.
//
// This is not the full struct, but just what we care about.
//
// Reference: https://docs.github.com/en/rest/licenses/licenses?apiVersion=2022-11-28#get-all-commonly-used-licenses
type License struct {
	// This is the full name of the License.
	// Example: "MIT License"
	Name string `json:"name"`

	// This is the SPDX ID of the License.
	// Example: "MIT"
	Spdx string `json:"spdx_id"`

	// This is the URL to the license text.
	Url string `json:"url"`
}

func (license *License) GetLicenseText() (*LicenseText, error) {
	req, err := http.NewRequest(http.MethodGet, license.Url, http.NoBody)
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

	var l LicenseText
	if err := json.Unmarshal(body, &l); err != nil {
		return nil, err
	}

	return &l, nil
}

// Returned from the "api.github.com/licenses/<license>" endpoint.
//
// This is not the full struct, but just what we care about.
//
// Reference: https://docs.github.com/en/rest/licenses/licenses?apiVersion=2022-11-28#get-a-license
type LicenseText struct {
	// This is the Implementation text for the license.
	//
	// Not confirmed, but the syntax should contain "[year]" etc. we coould
	// grep for to do templating.
	Implementation string `json:"implementation"`

	// This is the full text of the license.
	Body string `json:"body"`
}

// GetWants will fetch the "wants" from
// the Body of the LicenseText struct.
func (l *LicenseText) GetWants() []string {
	re := regexp.MustCompile(`\[[^\]]*\]`)
	return re.FindAllString(l.Body, -1)
}
