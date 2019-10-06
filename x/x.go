package x

import (
	"bytes"
	"fmt"
	"io"

	"upm/x/tar"
	"upm/x/gz"
	"upm/x/xz"
)

func Extract(src io.Reader, to string, args []string) error {
	var readBuf, writeBuf bytes.Buffer

	compRouter := map[string] func(io.Reader, io.Writer) error {
		"GZ": gz.Decompress,
		"XZ": xz.Decompress,
	}

	archRouter := map[string] func(io.Reader, string) error {
		"TAR": tar.Dearchive,
	}

	if _, err := writeBuf.ReadFrom(src); err != nil {
		return err
	}

	// Reversing args as we extract the file from end to start
	for left, right := 0, len(args) - 1; left < right; left, right = left + 1, right - 1 {
		args[left], args[right] = args[right], args[left]
	}

	for _, t := range args {
		comp, arch := true, true

		if _, err := readBuf.ReadFrom(&writeBuf); err != nil {
			return err
		}
		writeBuf.Reset()

		if _, exists := compRouter[t]; !exists {
			comp = false
		}

		if _, exists := archRouter[t]; !exists {
			arch = false
		}

		if !comp && !arch {
			return fmt.Errorf("unknown file type - %s", t)
		} else if comp {
			if err := compRouter[t](&readBuf, &writeBuf); err != nil {
				return err
			}
		} else if arch {
			if err := archRouter[t](&readBuf, to); err != nil {
				return err
			}
			break
		}
	}
	return nil
}

