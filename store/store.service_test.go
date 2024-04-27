package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func Init() *StorageService {
	return testStoreService.Init()
}
func TestStoreInit(t *testing.T) {
	testStoreService = Init()
	assert.True(t, testStoreService.redisClient != nil)
}
func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortURL := "Jsz4k57oAX"

	// Persist data mapping
	err := testStoreService.Save(shortURL, initialLink)
	if err != nil {
		t.Errorf("Error saving URL: %v", err)
	}

	// Retrieve initial URL

	retrievedUrl, err := testStoreService.Get(shortURL)
	if err != nil {
		t.Errorf("Error retrieving URL: %v", err)
	}

	assert.Equal(t, initialLink, retrievedUrl)
}
