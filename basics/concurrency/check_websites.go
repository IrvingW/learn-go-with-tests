package concurrency

// WebsiteChecker 接口接收一个字符串，返回他是否正确
type WebsiteChecker func(string) bool

type result struct {
	url   string
	valid bool
}

// CheckWebsites 函数接受两个参数，第一个参数是一个检查网址的函数，另外一个参数是需要检查的网址列表
// 返回一个map，key是网址，value是是否正确
func CheckWebsites(wc WebsiteChecker, urls []string) (results map[string]bool) {
	results = make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.url] = result.valid
	}
	return
}
