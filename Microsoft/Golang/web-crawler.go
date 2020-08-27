/**
 * // This is HtmlParser's API interface.
 * // You should not implement it, or speculate about its implementation
 * type HtmlParser struct {
 *     func GetUrls(url string) []string {}
 * }
 */

func crawl(startUrl string, htmlParser HtmlParser) []string {
    urls := make([]string, 0)
    queue := make([]string, 0)
    visited := make(map[string]bool)
    
    u, _ := url.Parse(startUrl)
    originHost := u.Host
    
    urls = append(urls, startUrl)
    visited[startUrl] = true
    
    for _, url := range htmlParser.GetUrls(startUrl) {
        if _, ok := visited[url]; !ok && hasHostname(url, originHost) {
            queue = append(queue, url)
        }
    }
    
    for len(queue) > 0 {
        url := queue[0]
        queue = queue[1:]
        
        if _, ok := visited[url]; ok {
            continue
        }
        urls = append(urls, url)
        visited[url] = true
        
        for _, url := range htmlParser.GetUrls(url) {
            if hasHostname(url, originHost) {
                queue = append(queue, url)
            }
        }
    }
    
    return urls
}

func hasHostname(URL, hostname string) bool {
    u, _ := url.Parse(URL)
    host := u.Host
    return hostname == host
}
