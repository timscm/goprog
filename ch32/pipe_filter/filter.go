package pipefilter

// input
type Request interface{}

// output
type Response interface{}

type Filter interface {
	Process(data Request) (Response, error)
}
