package router

type Controller interface {
	Index(*Context)
	New(*Context)
	Create(*Context, string)
	CreateWithJson(*Context, string)
	Show(*Context, string)
	Destroy(*Context)
}
