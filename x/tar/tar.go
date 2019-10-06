package tar

import (
	"fmt"
	"os"
	"io"
	"archive/tar"
	"path/filepath"

	"upm/logger"
)

func Dearchive(src io.Reader, dest string) error {
	logger.SetPrefix("tar: ")

	tarReader := tar.NewReader(src)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			newDirPath := filepath.Join(dest, header.Name)

			if err := os.MkdirAll(newDirPath, 0755); err != nil {
				return err
			}
			logger.Infof("Created directory in %s", newDirPath)

		case tar.TypeReg:
			newFilePath := filepath.Join(dest, header.Name)

			outFile, err := os.Create(newFilePath)
			if err != nil {
				return err
			}
			logger.Infof("Created file in %s", newFilePath)

			defer outFile.Close()

			if count, err := io.Copy(outFile, tarReader); err != nil {
				return err
			} else {
				logger.Debugf("Wrote %d bytes to %s", count, newFilePath)
			}

			if err := outFile.Close(); err != nil {
				return err
			}

		default:
			return fmt.Errorf(
				"unknown type: %s in %s",
				header.Typeflag,
				header.Name)
		}
	}
	return nil
}

