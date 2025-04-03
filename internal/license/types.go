package license

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/caffeine-addictt/waku/pkg/log"
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

	// That values the license wants
	Wants LicenseWants `json:"wants"`
}

func (license *License) GetLicenseText() (string, error) {
	url := GetLicenseFetchUrl() + license.Filename
	log.Infof("fetching license text from %s...\n", url)

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "text/plain")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	log.Debugln("reading http stream")
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	txt := string(body)
	log.Debugln("replacing [year] with current year")
	for i := range license.Wants {
		if license.Wants[i] == "year" {
			txt = strings.ReplaceAll(txt, "[year]", fmt.Sprintf("%d", time.Now().UTC().Year()))
			license.Wants = append(license.Wants[:i], license.Wants[i+1:]...)
			break
		}
	}

	return txt, nil
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
