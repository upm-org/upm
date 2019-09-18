package pkg

/*
	Definition of a package as an interface.
*/

type URL string

type Hash string

type Person struct {
	Name string
	Email string
}

type PkgHead interface {
	Name string
	Version string
	Section string
	Architecture string
	HashSum Hash
}

type PkgBody interface {
	VCS URL
	Developers []Person
	Maintainers []Person
	BuildDepends []PkhHead
	Depends []PkgHead
}

type PkgAPI interface {
	Unpack()
	Build()
	Install()
	Remove()
	Purge()
}

type Pkg interface {
	Head PkgHead
	Body PkgBody
	API PkgAPI
}
