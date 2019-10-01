package manager

import (
	"upm/pkg"
	"upm/pkg/upm"
	//"upm/pkg/rpm"
	"upm/pkg/deb"
	"upm/logger"
	"path/filepath"
)

var Log logger.UPMLogger

func Unpack(from string) pkg.Pkg {
	router := map[string] func(string) pkg.Pkg {
		".deb": deb.Unpack,
		".upm": upm.Unpack,
	}
	ext := filepath.Ext(from)
	if unpack, exists := router[ext]; !exists {
		Log.Error("Unpack: Can't unpack file with extension %s\n", ext)
	} else {
		return unpack(from)
	}
	return pkg.Pkg{}
}

