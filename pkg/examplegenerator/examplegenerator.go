package examplegenerator

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/flynn/flynn/pkg/httprecorder"
)

type Example struct {
	Name   string
	Runner func()
}

func WriteOutput(r *httprecorder.Recorder, examples []Example, out *os.File) error {
	res := make(map[string]*httprecorder.CompiledRequest)
	for _, ex := range examples {
		ex.Runner()
		res[ex.Name] = r.GetRequests()[0]
	}

	var err error
	data, err := json.MarshalIndent(res, "", "\t")
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "len(data) = %d\n", len(data))
	n, err := out.Write(data)
	out.Sync()
	fmt.Fprintf(out, "written = %d\n", n)
	time.Sleep(5 * time.Second)
	return err
}
