package filedump

import (
	"io"

	"github.com/blbgo/general"
)

type factory struct {
	Config
}

// NewFactory provides a general.DumperFactory interface that creates fileDump (returned as
// general.Dumper)
func NewFactory(config Config) general.DumperFactory {
	return &factory{Config: config}
}

// New creates a new fileLog and returns it as a general.Logger
func (r *factory) New(name string) (general.Dumper, error) {
	if name == "" {
		name = "dump"
	}
	return new(r.Config, name)
}

// Dump is a helper method that creates a Dumper, writes once and closes
func (r *factory) Dump(name string, data []byte) error {
	dumper, err := r.New(name)
	if err != nil {
		return err
	}
	closer, ok := dumper.(io.Closer)
	if ok {
		defer closer.Close()
	}

	return dumper.Dump(data)
}

// DumpObj is a helper method that creates a Dumper, writes once and closes
func (r *factory) DumpObj(name string, obj interface{}) error {
	dumper, err := r.New(name)
	if err != nil {
		return err
	}
	closer, ok := dumper.(io.Closer)
	if ok {
		defer closer.Close()
	}

	return dumper.DumpObj(obj)
}
