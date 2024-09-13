package license_test

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/caffeine-addictt/waku/cmd/license"
	"github.com/caffeine-addictt/waku/cmd/utils"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func TestLicenseJson(t *testing.T) {
	fi, err := os.OpenFile(filepath.Join("..", "..", "licenses", "license.json"), os.O_RDONLY, utils.FilePerms)
	assert.NoError(t, err)

	buf, err := io.ReadAll(fi)
	assert.NoError(t, err)

	var l license.License
	err = json.Unmarshal(buf, &l)
	assert.NoError(t, err)
}
