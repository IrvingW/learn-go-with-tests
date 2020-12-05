package concurrency

import (
	"testing"
	"time"
)

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Microsecond)
	return false
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a demo url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
