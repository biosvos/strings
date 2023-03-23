package strings

import (
	"bytes"
	"regexp"
)

type Strings string

func (s Strings) Regex(regex string) ([]*MatchResult, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	matches := re.FindAllStringSubmatch(string(s), -1)
	var ret []*MatchResult
	for _, match := range matches {
		ret = append(ret, &MatchResult{
			Result: match[0],
			Groups: match[1:],
		})
	}
	return ret, nil
}

type MatchResult struct {
	Result string
	Groups []string
}

func (s Strings) Replace(regex string, to string) (string, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	ret := re.ReplaceAllString(string(s), to)
	return ret, nil
}

func Join(strings []string, separator string) string {
	if len(strings) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	buffer.WriteString(strings[0])
	for _, str := range strings[1:] {
		buffer.WriteString(separator)
		buffer.WriteString(str)
	}
	return buffer.String()
}
