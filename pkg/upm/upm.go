package upm

import (
	"encoding/binary"
	"log"
	"os"

	"upm/config"
	"upm/extar"
	"upm/logger"
	"upm/pkg"
)

/*
	UPM Package specification 

	.upm file extension,

	Chunks of data:
		SIGN:
			4 bytes - UPM signature

		NEWLINE

		META:
			1 byte - data length
			Name of the package

			2 byte - data length
			Description of the package

			1 byte - data length
			Version of the package

			1 byte - data length
			Section of the package

			1 byte - data length
			Architecture of the package

		NEWLINE

		DATA:
			Packed tar.{gz, xz} archive
*/

func Unpack(from string) pkg.Pkg {
	/*
		Parsing file
	*/
	var res pkg.Pkg
	var file *os.File

	Log := logger.Log
	Log.SetPrefix(".upm/Unpack: ")

	readNext := func(n uint64) []byte {
		data := make([]byte, n)

		count, err := file.Read(data)
		if err != nil {
			Log.Fatal("%s", err)
		}
		Log.Debug("Read %d bytes - %s", count, string(data))
		return data
	}

	bytesToInt := func(s []byte) uint64 {
		var b [8]byte
		copy(b[8-len(s):], s)
		return binary.BigEndian.Uint64(b[:])
	}

	Log.Info("Parsing %s package", from)
	Log.Info("Reading Head")

	file, err := os.Open(from)
	if err != nil {
		Log.Fatal("%s", err)
	}

	// Parsing signature
	if signature := string(readNext(4)); signature != "UPM\n" {
		Log.Error("Bad signature\nExpected: UPM\nGot: %s", string(signature))
	} else {
		Log.Debug("signature - %s", signature)
	}

	// Parsing name
	nameLength := bytesToInt(readNext(1))
	res.Head.Name = string(readNext(nameLength))

	// Parsing description
	descLength := bytesToInt(readNext(2))
	res.Head.Description = string(readNext(descLength))

	// Reading version
	versLength := bytesToInt(readNext(1))
	res.Head.Version = string(readNext(versLength))

	// Reading section
	sectLength := bytesToInt(readNext(1))
	res.Head.Section = string(readNext(sectLength))

    // Reading architecture
	archLength := bytesToInt(readNext(1))
	res.Head.Architecture = string(readNext(archLength))

	// Reading newline
	if newline := string(readNext(1)); newline != "\n" {
		log.Fatal("Expected newline, got %s", newline)
	}

	Log.Info("Successfully read package Head")

	Log.Info("Extracting package Body to %s", config.Config.Cache.Dir)

	extar.ExtractTarGz(file, config.Config.Cache.Dir)

	Log.Info("Successfully extracted package Body")

	return res
}
