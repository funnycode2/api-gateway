package util

import (
	"regexp"
	"strings"
	"errors"
)

var (
	//主机名(含端口)正则表达式 由于写不出一个, 只能拆成多个
	//最简单的主机名,如: localhost
	HostRegxp0 = regexp.MustCompile("^\\w+$")
	//次简单的主机名,如: localhost:8080
	HostRegxp1 = regexp.MustCompile("^\\w+:\\d+$")
	//没有端口号的一般主机名,如: jandan.net
	HostRegxp2 = regexp.MustCompile("^\\w+(\\.\\w+)+$")
	//一般的主机名,如: jandan.net:80
	HostRegxp3 = regexp.MustCompile("^\\w+(\\.\\w+)+:\\d+$")
)

//匹配主机名+端口
func MatchHost(host string) bool {
	return !HostRegxp0.MatchString(host) &&
		!HostRegxp1.MatchString(host) &&
		!HostRegxp2.MatchString(host) &&
		!HostRegxp3.MatchString(host)
}

var emptyUrlError = errors.New("empty url not allowed")

//将url正则化如: url := "\\aaa\\\\bb/\\" 转化成  /aaa/bb
func NormalizeUri(url string) (string, error) {
	url = strings.Replace(url, "\\", "/", -1)
	if len(url) == 0 {
		return "", emptyUrlError
	}
	splits := strings.Split(url, "/")
	var normalized string
	for _, split := range splits {
		if trimmed := strings.TrimSpace(split); len(trimmed) > 0 {
			normalized = normalized + "/" + trimmed
		}
	}
	if len(normalized) == 0 {
		return "", emptyUrlError
	}
	return normalized, nil
}
