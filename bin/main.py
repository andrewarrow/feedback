#!/usr/bin/env /Users/aa/rick/foo/bin/python3

import sys
import os
from helper import foo

def replace_template_vars(template, replacements):
    for key, value in replacements.items():
        template = template.replace(f"{{{{{key}}}}}", value)
    return template

def gomod():
    template = """\
module {{name}}

replace github.com/andrewarrow/feedback => /Users/aa/os/feedback

go 1.21.0

require github.com/andrewarrow/feedback v0.0.0-20240617025030-9eb1fcd3b846
    """

    replacements = {
        "name": name,
    }
    placeIt("go.mod", replacements, template)

def placeIt(filename, replacements, template):

    result = replace_template_vars(template, replacements)

    output_filename = path+"/"+name+"/"+filename
    with open(output_filename, 'w') as file:
        file.write(result)
    return output_filename

def gomain():
    template = """\
package main

import "fmt"

//go:embed app/feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

var buildTag string

func main() {
  fmt.Println("wfwe")
}
    """

    placeIt("main.go", {}, template)

def run():
    template = """\
go mod tidy
go build
./{{name}}
    """

    rpath = placeIt("run", {"name": name}, template)
    st = os.stat(rpath)
    os.chmod(rpath, st.st_mode | 0o111)

def ignore():
    template = """\
go.sum
{{name}}
node_modules
package*.json
json.wasm.gz
.DS_Store
views
tail.min.css
    """

    placeIt(".gitignore", {"name": name}, template)

def ignore():
    template = """\
    """

    placeIt("views/text.html", {}, template)


path = sys.argv[1]
name = sys.argv[2]

def main():
    os.makedirs(path+"/"+name)
    os.makedirs(path+"/"+name+"/"+"views")
    gomod()
    gomain()
    run()
    ignore()
    foo()

if __name__ == "__main__":
    main()
