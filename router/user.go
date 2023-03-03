package router

func (r *Router) IsUserRequired(path string, method string) bool {
	//fmt.Println(path, method)
	if path == "/sessions/new/" {
		return false
	}
	if path == "/sessions/" {
		return false
	}
	return true
}
