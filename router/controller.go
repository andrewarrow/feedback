package router

type Controller interface {
	Index(*Context)
	New(*Context)
	Create(*Context)
	CreateWithId(*Context, string)
	CreateWithJson(*Context, string)
	Show(*Context, string)
	Destroy(*Context)
}
