package pkg

import (
	"path/filepath"
	//"upm/logger"
)

//var Log logger.UPMLogger

/*
	Definition of a package 
*/

type URL string

type Hash string

type Person struct {
	Name string
	Email string
}

type PkgHead struct {
	Name string
	Description string
	Version string
	Section string
	Architecture string
	HashSum Hash
}

type PkgBody struct {
	VCS URL
	Developers []Person
	Maintainers []Person
	BuildDepends []PkgHead
	Depends []PkgHead
}

type PkgAPI interface {
	Build()
	Install()
	Remove()
	Purge()
}

type Pkg struct {
	Head PkgHead
	Body PkgBody
	API PkgAPI
}

func Unpack(from string) Pkg {
	router := map[string] func(string) Pkg {
		".deb": DEBUnpack,
		".upm": UPMUnpack,
	}
	ext := filepath.Ext(from)
	if unpack, exists := router[ext]; !exists {
		Log.Error("Unpack: Can't unpack file with extension %s\n", ext)
	} else {
		return unpack(from)
	}
	return Pkg{}
}

