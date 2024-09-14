package license_test

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/caffeine-addictt/waku/cmd/license"
	"github.com/caffeine-addictt/waku/cmd/utils"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func TestLicenseJson(t *testing.T) {
	licenseRoot := filepath.Join("..", "..", "licenses")
	fi, err := os.OpenFile(filepath.Join(licenseRoot, "license.json"), os.O_RDONLY, utils.FilePerms)
	assert.NoError(t, err)
	defer fi.Close()

	buf, err := io.ReadAll(fi)
	assert.NoError(t, err)

	var l struct {
		Licenses []license.License `json:"licenses"`
	}
	err = json.Unmarshal(buf, &l)
	assert.NoError(t, err)

	for i, tc := range l.Licenses {
		t.Run(fmt.Sprintf("%d - %s", i, tc.Spdx), func(t *testing.T) {
			assert.NotEmpty(t, tc.Filename, "filename is empty")
			assert.NotEmpty(t, tc.Name, "name is empty")
			assert.NotEmpty(t, tc.Spdx, "spdx is empty")

			// check formatting
			assert.Equal(t, strings.TrimSpace(tc.Filename), tc.Filename, "filename poorly formatted")
			assert.Equal(t, strings.TrimSpace(tc.Name), tc.Name, "name poorly formatted")
			assert.Equal(t, strings.TrimSpace(tc.Spdx), tc.Spdx, "spdx poorly formatted")

			for _, v := range tc.Wants {
				s := strings.TrimSpace(v)
				assert.NotEmpty(t, s, "wants empty")
				assert.Equal(t, s, v, "wants poorly formatted")

				assert.False(t, strings.HasPrefix(s, "["), fmt.Sprintf("want %s should not have leading [", s))
				assert.False(t, strings.HasSuffix(s, "]"), fmt.Sprintf("want %s should not have trailing ]", s))
			}

			fi, err := os.OpenFile(filepath.Join(licenseRoot, tc.Filename), os.O_RDONLY, utils.FilePerms)
			assert.NoError(t, err)
			defer fi.Close()

			buf, err := io.ReadAll(fi)
			assert.NoError(t, err)
			assert.NotEmpty(t, buf)
		})
	}
}
