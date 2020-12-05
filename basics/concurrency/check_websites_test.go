package concurrency

import (
	"reflect"
	"testing"
)

func MockWebsiteChecker(url string) bool {
	if url == "http://google.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://baidu.com",
		"http://json.cn",
	}
	want := map[string]bool{
		"http://google.com": false,
		"http://baidu.com":  true,
		"http://json.cn":    true,
	}
	got := CheckWebsites(MockWebsiteChecker, websites)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}
