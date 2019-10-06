package gz

import (
	"compress/gzip"
	"io"

	"upm/logger"
)

func Decompress(src io.Reader, dest io.Writer) error {
	logger.SetPrefix("gz: ")

	decompressed, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	if count, err := io.Copy(dest, decompressed); err != nil {
		return err
	} else {
		logger.Debugf("Decompressed %d bytes", count)
	}

	return nil
}

