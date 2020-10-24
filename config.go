package filedump

import (
	"github.com/blbgo/general"
)

// Config provides config values for fileDump implementations
type Config interface {
	DumpPath() string
}

type config struct {
	DumpPathValue string
}

// NewConfig provides a filedump.Config
func NewConfig(c general.Config) (Config, error) {
	r := &config{}
	var err error

	r.DumpPathValue, err = c.Value("FileDump", "DumpPath")
	if err != nil {
		return nil, err
	}

	return r, nil
}

// DumpPath method of filedump.Config, returns the path where dump files should be created
func (r *config) DumpPath() string {
	return r.DumpPathValue
}
