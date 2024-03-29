package main

func runTemplate() string {

	t := `{{$package := index . "package"}}# https://tailwindcss.com/blog/standalone-cli
tailwindcss -i assets/css/tail.components.css -o assets/css/tail.min.css --minify
uuid=$(uuidgen); go build -ldflags="-X main.buildTag=$uuid"
./{{$package}} run 3000`
	return t
}
