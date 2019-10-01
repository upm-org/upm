package extar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"

	"upm/config"
	"upm/logger"
)

func ExtractTarGz(gzipStream io.Reader, target string) {
	Log := logger.Log
	Log.SetPrefix("extar: ")

	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		Log.Fatal("NewReader failed: %s", err)
	}

	tarReader := tar.NewReader(uncompressedStream)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			Log.Fatal("Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(filepath.Join(config.Config.Cache.Dir, header.Name), 0755); err != nil {
				Log.Fatal("Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			outFile, err := os.Create(filepath.Join(config.Config.Cache.Dir, header.Name))
			if err != nil {
				Log.Fatal("Create() failed: %s", err.Error())
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tarReader); err != nil {
				Log.Fatal("Copy() failed: %s", err.Error())
			}
		default:
			Log.Fatal(
				"Uknown type: %s in %s",
				header.Typeflag,
				header.Name)
		}
	}
}
