<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [go爬虫框架-colly](#go爬虫框架-colly)   
   - [简单例子](#简单例子)   
   - [type Collector](#type-collector)   
      - [func NewCollector(options ...CollectorOption) *Collector](#func-newcollectoroptions-collectoroption-collector)   
   - [type HTMLElement](#type-htmlelement)   
      - [func (h *HTMLElement) Attr(k string) string](#func-h-htmlelement-attrk-string-string)   
      - [func (h *HTMLElement) ForEach(goquerySelector string, callback func(int, *HTMLElement))](#func-h-htmlelement-foreachgoqueryselector-string-callback-funcint-htmlelement)   
   - [参考](#参考)   

<!-- /MDTOC -->
# go爬虫框架-colly


## 简单例子

```
package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}
```

* 初始化采集器，然后注册一系列回调函数，最后使用Visit爬取站点
* 关键是各种注册方法的调用
* 使用简单，方法众多，基本是每个阶段都可以回调
* 官方入门实例
http://go-colly.org/docs/introduction/start/

```
c.OnRequest(func(r *colly.Request) {
    fmt.Println("Visiting", r.URL)
})

c.OnError(func(_ *colly.Response, err error) {
    log.Println("Something went wrong:", err)
})

c.OnResponseHeaders(func(r *colly.Response) {
    fmt.Println("Visited", r.Request.URL)
})

c.OnResponse(func(r *colly.Response) {
    fmt.Println("Visited", r.Request.URL)
})

c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    e.Request.Visit(e.Attr("href"))
})

c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
    fmt.Println("First column of a table row:", e.Text)
})

c.OnXML("//h1", func(e *colly.XMLElement) {
    fmt.Println(e.Text)
})

c.OnScraped(func(r *colly.Response) {
    fmt.Println("Finished", r.Request.URL)
})
```




## type Collector

```
type Collector struct {
	// UserAgent is the User-Agent string used by HTTP requests
	UserAgent string
	// MaxDepth limits the recursion depth of visited URLs.
	// Set it to 0 for infinite recursion (default).
	MaxDepth int
	// AllowedDomains is a domain whitelist.
	// Leave it blank to allow any domains to be visited
	AllowedDomains []string
	// DisallowedDomains is a domain blacklist.
	DisallowedDomains []string
	// DisallowedURLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request will be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters
	// Leave it blank to allow any URLs to be visited
	DisallowedURLFilters []*regexp.Regexp

	// Leave it blank to allow any URLs to be visited
	URLFilters []*regexp.Regexp

	// AllowURLRevisit allows multiple downloads of the same URL
	AllowURLRevisit bool
	// MaxBodySize is the limit of the retrieved response body in bytes.
	// 0 means unlimited.
	// The default value for MaxBodySize is 10MB (10 * 1024 * 1024 bytes).
	MaxBodySize int
	// CacheDir specifies a location where GET requests are cached as files.
	// When it's not defined, caching is disabled.
	CacheDir string
	// IgnoreRobotsTxt allows the Collector to ignore any restrictions set by
	// the target host's robots.txt file.  See http://www.robotstxt.org/ for more
	// information.
	IgnoreRobotsTxt bool
	// Async turns on asynchronous network communication. Use Collector.Wait() to
	// be sure all requests have been finished.
	Async bool
	// ParseHTTPErrorResponse allows parsing HTTP responses with non 2xx status codes.
	// By default, Colly parses only successful HTTP responses. Set ParseHTTPErrorResponse
	// to true to enable it.
	ParseHTTPErrorResponse bool
	// ID is the unique identifier of a collector
	ID uint32
	// DetectCharset can enable character encoding detection for non-utf8 response bodies
	// without explicit charset declaration. This feature uses https://github.com/saintfish/chardet
	DetectCharset bool

	// CheckHead performs a HEAD request before every GET to pre-validate the response
	CheckHead bool
	// TraceHTTP enables capturing and reporting request performance for crawler tuning.
	// When set to true, the Response.Trace will be filled in with an HTTPTrace object.
	TraceHTTP bool
	// contains filtered or unexported fields
}
```

Collector provides the scraper instance for a scraping job

* 关键类型，采集器对象


### func NewCollector(options ...CollectorOption) *Collector

```
NewCollector creates a new Collector instance with default configuration
```

* 创建一个采集器，若没有提供options则采用默认配置
* 实例化时候要进行配置，配置包括如下：

```
COLLY_ALLOWED_DOMAINS (comma separated list of domains)
COLLY_CACHE_DIR (string)
COLLY_DETECT_CHARSET (y/n)
COLLY_DISABLE_COOKIES (y/n)
COLLY_DISALLOWED_DOMAINS (comma separated list of domains)
COLLY_IGNORE_ROBOTSTXT (y/n)
COLLY_FOLLOW_REDIRECTS (y/n)
COLLY_MAX_BODY_SIZE (int)
COLLY_MAX_DEPTH (int - 0 means infinite)
COLLY_PARSE_HTTP_ERROR_RESPONSE (y/n)
COLLY_USER_AGENT (string)
```


## type HTMLElement

* 用于表示HTML元素

```
type HTMLElement struct {
	// Name is the name of the tag
	Name string
	Text string

	// Request is the request object of the element's HTML document
	Request *Request
	// Response is the Response object of the element's HTML document
	Response *Response
	// DOM is the goquery parsed DOM object of the page. DOM is relative
	// to the current HTMLElement
	DOM *goquery.Selection
	// Index stores the position of the current element within all the elements matched by an OnHTML callback
	Index int
	// contains filtered or unexported fields
}
```

### func (h *HTMLElement) Attr(k string) string

```
Attr returns the selected attribute of a HTMLElement or empty string if no attribute found
```

* 用于获取引用为h的html元素中对应属性k的值，必输返回字符串
* 以 '<>' 表示html元素，所有'<>' 内属性均为键值对

### func (h *HTMLElement) ForEach(goquerySelector string, callback func(int, *HTMLElement))

```
ForEach iterates over the elements matched by the first argument and calls the callback function on every HTMLElement match.
```


## 参考

* <https://pkg.go.dev/github.com/gocolly/colly/v2#Collector>



---
