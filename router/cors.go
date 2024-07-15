package router

func setCorsOptions(c *Context) {
	c.Writer.Header().Set("Allow", "GET,POST,PUT,PATCH,DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline' http://localhost")
	c.Writer.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubdomains")

}

func setCors(c *Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
}
