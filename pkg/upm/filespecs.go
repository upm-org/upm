package upm

import (
	"bytes"
	"errors"
	"strings"

	"upm/logger"
	"upm/pkg"
	"upm/x"
)

var ErrWrongSignature = errors.New("wrong signature, should be upm")

type ChunkField struct {
	Name       string
	LengthSize int
	Parse      func(*pkg.PKG, []byte) error
}

type ChunkFields []ChunkField

type Chunk struct {
	Name   string
	Fields ChunkFields
}

type Chunks []Chunk

var chunkHandler = Chunks{
	Chunk{
		"META",
		ChunkFields{
			ChunkField{
				"SIGN",
				1,
				func (res *pkg.PKG, data []byte) error {
					signature := string(data)
					if signature != "UPM" {
						return ErrWrongSignature
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
				func (res *pkg.PKG, data []byte) error {
					res.Head.Name = string(data)

					return nil
				},
			},
			ChunkField{
				"DESC",
				2,
				func (res *pkg.PKG, data []byte) error {
					res.Head.Description = string(data)

					return nil
				},
			},
			ChunkField{
				"VERS",
				1,
				func (res *pkg.PKG, data []byte) error {
					res.Head.Version = string(data)

					return nil
				},
			},
			ChunkField{
				"SECT",
				1,
				func (res *pkg.PKG, data []byte) error {
					res.Head.Section = string(data)

					return nil
				},
			},
			ChunkField{
				"ARCH",
				1,
				func (res *pkg.PKG, data []byte) error {
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
				func (res *pkg.PKG, data []byte) error {
					res.Body = pkg.PKGBody{Pack: strings.Split(string(data), "."), Data: new(bytes.Buffer)}

					return nil
				},
			},
			ChunkField{
				"DATA",
				8,
				func (res *pkg.PKG, data []byte) error {
					if count, err := res.Body.Write(data); err != nil {
						return err
					} else {
						logger.Debugf("%d bytes copied to buffer", count)
					}

					if err := x.Extract(&res.Body, *unpackPath, res.Body.Pack); err != nil {
						return err
					}

					return nil
				},
			},
		},
	},
}
