package pkg

import (
	"os"
	"upm/logger"
)

var Log logger.UPMLogger

/*
	UPM Package specification 

	.upm file extension,

	Chunks of data:
		SIGN:
			4 bytes - UPM\n signature
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
		DATA:
			Packed tar.{gz, xz} archive
*/

func UPMUnpack(from string) Pkg {
	/*
		Parsing file
	*/
	Log.Info("Parsing %s package", from)

	file, err := os.Open(from)
	if err != nil {
		Log.Fatal(err)
	}

	signature := make([]byte, 4)
	count, err := file.Read(signature)
	if err != nil {
		Log.Fatal(err)
	}
	Log.Debug("read %d bytes: %q\n", count, signature[:count])

	if string(signature) != "UPM\n" {
		Log.Error("Bad signature\nExpected: UPM\nGot: %s", string(signature))
	}
	return Pkg{}
}
