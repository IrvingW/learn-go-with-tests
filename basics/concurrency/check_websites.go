package concurrency

// WebsiteChecker 接口接收一个字符串，返回他是否正确
type WebsiteChecker func(string) bool

// CheckWebsites 函数接受两个参数，第一个参数是一个检查网址的函数，另外一个参数是需要检查的网址列表
// 返回一个map，key是网址，value是是否正确
func CheckWebsites(wc WebsiteChecker, urls []string) (results map[string]bool) {
	results = make(map[string]bool)
	for _, url := range urls {
		results[url] = wc(url)
	}
	return
}
