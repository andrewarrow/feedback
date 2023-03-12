# Feedback

This is a [rails](https://rubyonrails.org) inspired [golang](https://go.dev/) framework that uses [Hacker News](https://news.ycombinator.com/) style article submission, comments, and voting in its demo app [RemoteRenters](https://github.com/andrewarrow/remoterenters).
 
Look at [main.go](https://github.com/andrewarrow/remoterenters/blob/main/main.go) to see how to use feedback.

Our sample app is live at [remoterenters.com](https://remoterenters.com/)

```
  r.Paths = map[string]func(*Context, string, string){}
  r.Paths["models"] = handleModels
  r.Paths["sessions"] = handleSessions
  r.Paths["users"] = handleUsers
  r.Paths["about"] = handleAbout
```

This is the heart of the routing code.  You can have three levels of:

/foo/

/foo/bar/

/foo/bar/more/

That's why each top level path takes a func with two strings.
Some of the paths you get built in to feedback like `sessions` and `users` since every app will need that logic.

But notice in [main.go](https://github.com/andrewarrow/remoterenters/blob/main/main.go) how this app adds more routes.

```
func HandleSomething(c *router.Context, second, third string) {
  if second == "" {
    c.SendContentInLayout("something_index.html", nil, 200)
  } else if third != "" {
    c.NotFound = true
  } else {
    c.NotFound = true
  }
}
```

Each controller has a HandleSomething func like this.
