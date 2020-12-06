package racer

import "testing"

func TestRacer(t *testing.T) {
	urlA := "www.baidu.com"
	urlB := "www.json.cn"
	winner := Racer(urlA, urlB)
	want := urlB
	if winner != want {
		t.Errorf("got %v but want %v", winner, want)
	}
}
