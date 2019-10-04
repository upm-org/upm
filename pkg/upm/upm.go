package upm

import (
	"encoding/binary"
	"io"
	"os"

	"upm/logger"
	"upm/pkg"
)

/*
	UPM Package specification 

	.upm file extension,

	File consists of chunks seperated by newline

	Chunk headers are 5 byte length (4-name, 1-newline)
	Chunk Fields are 5 byte length as well (4-Name, 1-space)
	After field content newline character is expected

	File structure:
		META:
			4 bytes - UPM signature

		HEAD:
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

		BODY:
			1 byte - data length
			Archive type (native support: tar.xz, tar.gz) (TODO: archive mapper plugin)
			
			~DATA~
*/

func init() {
	logger.SetPrefix("upm: ")
}

func bytesToInt(s []byte) uint64 {
	var b [8]byte
	copy(b[8-len(s):], s)
	return binary.BigEndian.Uint64(b[:])
}

func Pack(info pkg.PKGInfo, dir, to string) {
	//var file *os.File
}

var unpackPath *string

func Unpack(from, to string) (*pkg.PKG, error) {
	var file *os.File
	var res pkg.PKG

	unpackPath = &to

	file, err := os.Open(from)
	if err != nil {
		logger.Fatal(err)
	}

	for _, chunk := range chunkHandler {
		logger.Debug("Reading chunk", chunk.Name)

		for _, field := range chunk.Fields {
			logger.Debug("Reading field", field.Name)

			lengthBuf := make([]byte, field.LengthSize)

			count, err := file.Read(lengthBuf)

			if err != nil {
				return nil, err
			}

			length := bytesToInt(lengthBuf)
			logger.Debugf("upm: Read field %s-LENGTH %d bytes - %d", field.Name, count, length)

			data := make([]byte, length)

			count, err = file.Read(data)

			if err != nil {
				return nil, err
			}

			logger.Debugf("upm: Read field %s %d bytes - %s", field.Name, count, string(data))

			if err := field.Parse(&res, data); err != nil {
				return nil, err
			}
		}
	}

	if offset, err := file.Seek(0, io.SeekCurrent); err != nil {
		return nil, err
	} else {
		logger.Debugf("Finished reading file - %d bytes", offset)
	}

	return nil, nil
}

