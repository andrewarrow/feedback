package router

func handleAdmin(c *Context, second, third string) {
	c.Layout = "admin_layout.html"
	if second == "" && third == "" && c.Method == "GET" {
		handleAdminIndex(c)
		return
	}
	if second == "users" && third == "" && c.Method == "GET" {
		handleAdminUsersIndex(c)
		return
	}
	c.NotFound = true
}

func handleAdminIndex(c *Context) {
	c.SendContentInLayout("admin_index.html", nil, 200)
}

func handleAdminUsersIndex(c *Context) {
	c.SendContentInLayout("admin_users_index.html", nil, 200)
}
