package filedump

import (
	"io"
	"os"
	"testing"

	"github.com/blbgo/general"
)

func TestFileDumper(t *testing.T) {
	r := general.Dumper(&fileDump{})

	// test that fileDump as a general.Dumper can be asserted to an io.Closer
	closer, ok := r.(io.Closer)
	if !ok {
		t.Error("fileDump Dumper is not an io.Closer")
	}

	// test that calling close calls os.File with nil causing a os.ErrInvalid error
	err := closer.Close()
	if err != os.ErrInvalid {
		t.Errorf("Error calling Close is not os.ErrInvalid: %v", err)
	}
}
