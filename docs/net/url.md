# url

```
import "net/url"
```

* url包解析URL并实现了查询



## type URL

```
// A URL represents a parsed URL (technically, a URI reference).
//
// The general form represented is:
//
//	[scheme:][//[userinfo@]host][/]path[?query][#fragment]
//
// URLs that do not start with a slash after the scheme are interpreted as:
//
//	scheme:opaque[?query][#fragment]
//
// Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/.
// A consequence is that it is impossible to tell which slashes in the Path were
// slashes in the raw URL and which were %2f. This distinction is rarely important,
// but when it is, the code should use RawPath, an optional field which only gets
// set if the default encoding is different from Path.
//
// URL's String method uses the EscapedPath method to obtain the path. See the
// EscapedPath method for more details.
type URL struct {
	Scheme      string
	Opaque      string    // encoded opaque data
	User        *Userinfo // username and password information
	Host        string    // host or host:port
	Path        string    // path (relative paths may omit leading slash)
	RawPath     string    // encoded path hint (see EscapedPath method)
	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	RawQuery    string    // encoded query values, without '?'
	Fragment    string    // fragment for references, without '#'
	RawFragment string    // encoded fragment hint (see EscapedFragment method)
}
```

### func (u *URL) Query() Values

* Query方法解析RawQuery字段并返回其表示的Values类型键值对。




```
func Parse(rawurl string) (url *URL, err error)
func ParseRequestURI(rawurl string) (url *URL, err error)
func (u *URL) IsAbs() bool
func (u *URL) RequestURI() string
func (u *URL) String() string
func (u *URL) Parse(ref string) (*URL, error)
func (u *URL) ResolveReference(ref *URL) *URL
```









---
