package pipefilter

import "errors"

var ErrSumFilterWrongFormat = errors.New("input should be []int")

type SumFilter struct{}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	elems, ok := data.([]int)
	if !ok {
		return nil, ErrSumFilterWrongFormat
	}
	ret := 0
	for _, elem := range elems {
		ret += elem
	}
	return ret, nil
}
