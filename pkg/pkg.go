package pkg

import "bytes"

/*
	Definition of a package 
*/

type URL string

type Hash string

type Person struct {
	Name string
	Email string
}

type PKGHead struct {
	Name string
	Description string
	Version string
	Section string
	Architecture string
}

type PKGBody struct {
	Pack []string
	Data *bytes.Buffer
}

func (b *PKGBody) Read(p []byte) (n int, err error) {
	return b.Data.Read(p)
}

func (b *PKGBody) Write(p []byte) (n int, err error) {
	return b.Data.Write(p)
}

type PKGInfo struct {
	Vcs URL
	Developers []Person
	Maintainers []Person
	BuildDepends []PKGHead
	Depends []PKGHead
}

type PKGAPI interface {
	Build()
	Install()
	Remove()
	Purge()
}

type PKG struct {
	Head PKGHead
	Body PKGBody
	Info PKGInfo
	API  PKGAPI
}
