# Feedback

This is a [rails](https://rubyonrails.org) inspired [golang](https://go.dev/) framework that uses [Hacker News](https://news.ycombinator.com/) style article submission, comments, and voting.
 
This repo is also a framework you can fork and turn into your own app.
                
Our sample app is live at [RemoteRenters](https://remoterenters.com/)

```
func handleContext(c *Context) {
  tokens := c.tokens

  if len(tokens) == 3 { //          /foo/
    handlePathContext(c, tokens[1], "", "")
  } else if len(tokens) == 4 { //   /foo/bar/
    handlePathContext(c, tokens[1], tokens[2], "")
  } else if len(tokens) == 5 { //   /foo/bar/more/
    handlePathContext(c, tokens[1], tokens[2], tokens[3])
  } else {
    c.notFound = true
  }
}
```

This is the heart of the routing code.  You can have three levels of:

/foo/

/foo/bar/

/foo/bar/more/

```
func handlePathContext(c *Context, first, second, third string) {
  if first == "models" {
    handleModels(c, second, third)
  } else if first == "sessions" {
    handleSessions(c, second, third)
  } else if first == "stories" {
    handleStories(c, second, third)
  } else if first == "users" {
    handleUsers(c, second, third)
  } else if first == "comments" {
    handleComments(c, second, third)
  } else if first == "about" {
    handleAbout(c, second, third)
  } else if first == "fresh" {
    handleFresh(c, second, third)
  } else {
    c.notFound = true
  }
}
```

Each main section calls it's controller's handle method with second and third.

Second might be "" if it's just /foo/
Third will also be ""
Second will be "something" if it's /foo/something/
And third will only not be "" if you have all three parts.

