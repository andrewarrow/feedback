package main

func modTemplate() string {
	t := `{{$package := index . "package"}} module {{$package}}

replace github.com/andrewarrow/feedback => /Users/aa/os/feedback

go 1.19

require (
	github.com/andrewarrow/feedback v0.0.0-20230629214121-08868362ccbe
	golang.org/x/crypto v0.8.0
)`

	return t
}
