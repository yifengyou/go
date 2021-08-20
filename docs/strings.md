# strings

## func Contains(s, substr string) bool

s字符串中是否包含substr字符串，返回布尔值

1. substr若为字面量，规范用""，若是一个字符也""，别杠

```
fmt.Println(strings.Contains(s2, 's'))
// fmt.Println(strings.Contains(s2, 's'))
fmt.Println(strings.Contains(s2, ""))
// true   任何字符串都包含空串
```

## func ContainsAny(s, chars string) bool

s字符串中是否包含chars字符串中任意一个字符串

```
fmt.Println(strings.ContainsAny(s2, ""))
// false   空串默认排除在外
```



## func HasPrefix(s, prefix string) bool

判断s是否有前缀字符串prefix




## func HasSuffix(s, suffix string) bool

判断s是否有后缀字符串suffix








































---
