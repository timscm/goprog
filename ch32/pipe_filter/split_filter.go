package pipefilter

import (
	"errors"
	"strings"
)

var ErrSplitFilterWrongFormat = errors.New("input data should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	// 检查数据格式/类型，是否可以处理
	str, ok := data.(string)
	if !ok {
		return nil, ErrSplitFilterWrongFormat
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
