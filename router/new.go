package router

type Router struct {
	Paths map[string]string
}

func NewRouter() *Router {
	r := Router{}
	r.Paths = map[string]string{}
	r.Paths["/admin/users"] = "GET"
	return &r
}
