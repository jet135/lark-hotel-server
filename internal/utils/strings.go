package utils

import (
	"regexp"
	"strings"
)

var (
	StrIsUndefined = 0
	StrIsPhone     = 1
	StrIsName      = 2
)

func CutPrefix(s, prefix string) (string, bool) {
	if strings.HasPrefix(s, prefix) {
		return strings.TrimPrefix(s, prefix), true
	}
	return s, false
}

func EitherCutPrefix(s string, prefix ...string) (string, bool) {
	// 任一前缀匹配则返回剩余部分
	for _, p := range prefix {
		if strings.HasPrefix(s, p) {
			return strings.TrimPrefix(s, p), true
		}
	}
	return s, false
}

// trim space and equal
func TrimEqual(s, prefix string) (string, bool) {
	if strings.TrimSpace(s) == prefix {
		return "", true
	}
	return s, false
}

func EitherTrimEqual(s string, prefix ...string) (string, bool) {
	// 任一前缀匹配则返回剩余部分
	for _, p := range prefix {
		if strings.TrimSpace(s) == p {
			return "", true
		}
	}
	return s, false
}

// StrDefine 判断输入的字符串含义
func StrDefine(input string) int {
	// 正则表达式匹配中文姓名（假设姓名只包含汉字）
	chineseNameRegex := regexp.MustCompile(`^[\p{Han}]{2,}$`)

	// 正则表达式匹配中国大陆手机号
	phoneNumberRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)

	if chineseNameRegex.MatchString(input) {
		return StrIsName
	} else if phoneNumberRegex.MatchString(input) {
		return StrIsPhone
	} else {
		return StrIsUndefined
	}
}
