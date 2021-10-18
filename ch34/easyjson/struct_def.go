// go install github.com/mailru/easyjson/...
// go/bin/easyjson -all struct_def.go
// --> struct_def_easyjson.go
package structdef

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}
