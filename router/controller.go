package router

type Controller interface {
	Index(*Context)
	New(*Context)
	Create(*Context, string)
	CreateWithId(*Context, string, string)
	CreateWithJson(*Context, string)
	Show(*Context, string)
	Destroy(*Context)
}
