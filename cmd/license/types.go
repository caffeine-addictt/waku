package license

import (
	"io"
	"net/http"
	"strings"

	"github.com/goccy/go-json"
)

// Returned from the licenses/ directory.
type License struct {
	// This is the full name of the License.
	// Example: "MIT License"
	Name string `json:"name"`

	// This is the SPDX ID of the License.
	// Example: "MIT"
	Spdx string `json:"spdx"`

	// The filename in 'licenses/'
	Filename string `json:"filename"`

	// That values the licens wants
	Wants LicenseWants `json:"wants"`
}

func (license *License) GetLicenseText() (string, error) {
	req, err := http.NewRequest(http.MethodGet, BASE_URL+license.Filename, http.NoBody)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "text/plain")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

type LicenseWants []string

// Will clean up license wants
func (l *LicenseWants) UnmarshalJSON(data []byte) error {
	var license []string

	if err := json.Unmarshal(data, &license); err != nil {
		return err
	}

	for i := range license {
		s := strings.TrimSpace(license[i])
		s = strings.TrimPrefix(strings.TrimSuffix(s, "]"), "[")
		license[i] = strings.ToLower(s)
	}

	*l = license
	return nil
}
