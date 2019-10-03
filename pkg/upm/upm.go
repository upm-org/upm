package upm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"upm/logger"
	"upm/pkg"
	"upm/x"
)

/*
	UPM Package specification 

	.upm file extension,

	File consists of chunks seperated by newline

	Chunk headers are 5 byte length (4-name, 1-newline)
	Chunk fields are 5 byte length as well (4-name, 1-space)
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
			Archive type (native support: tar.xz, tar.gz) (TODO: archieve mapper plugin)
			
			~DATA~
*/

func bytesToInt(s []byte) uint64 {
	var b [8]byte
	copy(b[8-len(s):], s)
	return binary.BigEndian.Uint64(b[:])
}

func Pack(info pkg.PkgInfo, dir, to string) {
	//var file *os.File
}

func Unpack(from, to string) (*pkg.Pkg, error) {
	var file *os.File
	var res pkg.Pkg
	Log := logger.Log
	Log.Prefix = "upm: "

	type ChunkField struct {
		name       string
		lengthSize int
		parse      func([]byte) error
	}

	type ChunkFields []ChunkField

	type Chunk struct {
		name   string
		fields ChunkFields
	}

	type Chunks []Chunk

	// distanceMap shows us how many bytes are reserved for field length
	var distanceMap = Chunks{
		Chunk{
			"META",
			ChunkFields{
				ChunkField{
					"SIGN",
					1,
					func (data []byte) error {
						signature := string(data)
						if signature != "UPM" {
							return fmt.Errorf("Wrong signature! Expected UPM, got %s", signature)
						}

						return nil
					},
				},
			},
		},
		Chunk{
			"HEAD",
			ChunkFields{
				ChunkField{
					"NAME",
					1,
					func (data []byte) error {
						res.Head.Name = string(data)

						return nil
					},
				},
				ChunkField{
					"DESC",
					2,
					func (data []byte) error {
						res.Head.Description = string(data)

						return nil
					},
				},
				ChunkField{
					"VERS",
					1,
					func (data []byte) error {
						res.Head.Version = string(data)

						return nil
					},
				},
				ChunkField{
					"SECT",
					1,
					func (data []byte) error {
						res.Head.Section = string(data)

						return nil
					},
				},
				ChunkField{
					"ARCH",
					1,
					func (data []byte) error {
						res.Head.Architecture = string(data)

						return nil
					},
				},
			},
		},
		Chunk{
			"BODY",
			ChunkFields{
				ChunkField{
					"COMP",
					1,
					func (data []byte) error {
						var buf bytes.Buffer

						fileTypes := strings.Split(string(data), ".")
						compressed, err := ioutil.ReadAll(file)
						if err != nil {
							return err
						}

						offset, _ := file.Seek(0, 1)

						Log.Infof("Finished reading the file, %d bytes", offset)

						if count, err := buf.Write(compressed); err != nil {
							return err
						} else {
							Log.Debugf("%d bytes copied to buffer", count)
						}
						if err = x.Extract(&buf, to, fileTypes); err != nil {
							return err
						}

						return nil
					},
				},
			},
		},
	}


	file, err := os.Open(from)
	if err != nil {
		Log.Fatalf("%s", err)
	}

	for _, chunk := range distanceMap {
		for _, field := range chunk.fields {
			lengthBuf := make([]byte, field.lengthSize)

			count, err := file.Read(lengthBuf)
			if err != nil {
				return nil, err
			}
			length := bytesToInt(lengthBuf)
			Log.Debugf("Read field %s-LENGTH %d bytes - %d", field.name, count, length)

			data := make([]byte, length)

			count, err = file.Read(data)
			if err != nil {
				return nil, err
			}
			Log.Debugf("Read field %s %d bytes - %s", field.name, count, string(data))

			// Skipping newline
			_, err = file.Seek(1, 1)

			if err := field.parse(data); err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}

