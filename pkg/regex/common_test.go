package commonregex

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonRegex(t *testing.T) {
	tests := []struct {
		name     string
		expected []byte
		value    []byte
		regex    string
	}{
		{
			name:     "domain: good domain should return domain",
			expected: []byte("google.com"),
			value:    []byte("google.com"),
			regex:    Domain,
		},
		{
			name:     "domain: bad domain should not return domain",
			expected: []byte(nil),
			value:    []byte("notadomain"),
			regex:    Domain,
		},
		{
			name:     "ip: good ip should return ip",
			expected: []byte("8.8.8.8"),
			value:    []byte("8.8.8.8"),
			regex:    Ip,
		},
		{
			name:     "ip: bad ip should not return ip",
			expected: []byte(nil),
			value:    []byte("256.256.256.256"),
			regex:    Ip,
		},
		{
			name:     "sha1: good sha1 should return sha1",
			expected: []byte("f572d396fae9206628714fb2ce00f72e94f2258f"),
			value:    []byte("f572d396fae9206628714fb2ce00f72e94f2258f"),
			regex:    Sha1,
		},
		{
			name:     "sha1: bad (shorter that 40 chars) sha1 should not return sha1",
			expected: []byte(nil),
			value:    []byte("adf"),
			regex:    Sha1,
		},
		{
			name:     "sha1: bad (longer that 40 chars) sha1 should not return sha1",
			expected: []byte(nil),
			value:    []byte("f572d396fae9206628714fb2ce00f72e94f2258fadf"),
			regex:    Sha1,
		},
		{
			name:     "md5: good md5 should return md5",
			expected: []byte("b1946ac92492d2347c6235b4d2611184"),
			value:    []byte("b1946ac92492d2347c6235b4d2611184"),
			regex:    Md5,
		},
		{
			name:     "md5: bad (shorter that 32 chars) md5 should not return md5",
			expected: []byte(nil),
			value:    []byte("adsf"),
			regex:    Md5,
		},
		{
			name:     "md5: bad (longer that 32 chars) md5 should not return md5",
			expected: []byte(nil),
			value:    []byte("b1946ac92492d2347c6235b4d2611184adf"),
			regex:    Md5,
		},
		{
			name:     "sha256: good sha256 should return sha256",
			expected: []byte("5891b5b522d5df086d0ff0b110fbd9d21bb4fc7163af34d08286a2e846f6be03"),
			value:    []byte("5891b5b522d5df086d0ff0b110fbd9d21bb4fc7163af34d08286a2e846f6be03"),
			regex:    Sha256,
		},
		{
			name:     "sha256: bad (shorter that 32 chars) sha256 should not return sha256",
			expected: []byte(nil),
			value:    []byte("adsf"),
			regex:    Sha256,
		},
		{
			name:     "sha256: bad (longer that 65 chars) sha256 should not return sha256",
			expected: []byte(nil),
			value:    []byte("5891b5b522d5df086d0ff0b110fbd9d21bb4fc7163af34d08286a2e846f6be03adf"),
			regex:    Sha256,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			regex, err := regexp.Compile(test.regex)
			if err != nil {
				panic(err)
			}
			actual := regex.Find(test.value)
			assert.Equal(t, test.expected, actual)
		})
	}
}
