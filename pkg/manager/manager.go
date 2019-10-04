package manager

import (
	"fmt"

	"upm/pkg"
	"upm/pkg/upm"
	//"upm/pkg/rpm"
	//"upm/pkg/deb"
	"path/filepath"
)


func Unpack(from, to string) /*(*pkg.PKG,*/ error/*)*/ {
	router := map[string] func(string, string) (*pkg.PKG, error) {
		//".deb": deb.Unpack,
		".upm": upm.Unpack,
	}
	ext := filepath.Ext(from)

	routedUnpack, exists := router[ext]
	if !exists {
		return /*nil,*/ fmt.Errorf("Unpack: Can't unpack file with extension %s\n", ext)
	}
	/*res*/_, err := routedUnpack(from, to)
	if err != nil {
		return /*nil,*/ err
	}
	return /*res,*/ nil
}

