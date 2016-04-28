package lantern

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestFeedProvider struct{}

type TestFeedRetriever struct{}

func (provider *TestFeedProvider) AddSource(source string) {

}

func (provider *TestFeedProvider) DisplayError(errMsg string) {

}

func (retriever *TestFeedRetriever) AddFeed(title, description,
	image, link string) {

}

func getFeed(t *testing.T, locale string, proxyAddr string) {
	provider := &TestFeedProvider{}
	GetFeed(locale, proxyAddr, provider)

	feed := CurrentFeed()
	if !assert.NotNil(t, feed) {
		return
	}
	if !assert.NotEqual(t, 0, NumFeedEntries(),
		"No feed entries after processing") {
		return
	}

	assert.Equal(t, NumFeedEntries(), len(feed.Items["all"]),
		"All feed items should be equal to total entries")
}

func TestGetFeed(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "testconfig")
	if !assert.NoError(t, err, "Unable to create temp configDir") {
		return
	}
	defer os.RemoveAll(tmpDir)
	result, err := Start(tmpDir, 5000)
	if !assert.NoError(t, err, "Should have been able to start lantern") {
		return
	}
	locales := []string{"en_US", "fa_IR", "invalid"}
	for _, l := range locales {
		getFeed(t, l, result.HTTPAddr)
	}
}
