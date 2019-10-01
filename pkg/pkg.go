package pkg

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

