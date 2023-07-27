package main

func ignoreTemplate() string {
	t := `
{{$name := index . "name"}}
{{$lower := index . "lower"}}
go.sum
node_modules
package-lock.json
package.json
assets/css/tail.min.css
.DS_Store
{{$lower}}
`
	return t
}
