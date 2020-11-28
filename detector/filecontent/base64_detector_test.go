package filecontent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64DetectorShouldNotDetectSafeText(t *testing.T) {
	s := "pretty safe"
	bd := Base64Detector{}
	bd.initBase64Map()

	res := bd.CheckBase64Encoding(s)
	assert.Equal(t, "", res)
}
