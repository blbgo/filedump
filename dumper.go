package filedump

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/blbgo/general"
)

type fileDump struct {
	*os.File
}

// New provides a general.Dumper interface that uses a file as storage
func New(config Config) (general.Dumper, error) {
	return new(config, "dump")
}

func new(config Config, name string) (general.Dumper, error) {
	file, err := os.Create(
		filepath.Join(
			config.DumpPath(),
			fmt.Sprintf("%v%v.txt", name, time.Now().Format("2006-01-02T15-04-05_000")),
		),
	)
	if err != nil {
		return nil, err
	}
	return &fileDump{File: file}, nil
}

func (r *fileDump) Dump(data []byte) error {
	_, err := r.Write(data)
	return err
}

func (r *fileDump) DumpObj(obj interface{}) error {
	data, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return err
	}
	_, err = r.Write(data)
	return err
}

func (r *fileDump) Close() error {
	return r.File.Close()
}
