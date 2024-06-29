#!/usr/bin/env /Users/aa/rick/foo/bin/python3

import sys
import os
from gomain import gomain
from placeit import placeit

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
    placeit("go.mod", replacements, template)

def run():
    template = """\
go mod tidy
go build
./{{name}}
    """

    rpath = placeit("run", {"name": name}, template)
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

    placeit(".gitignore", {"name": name}, template)

def ignore():
    template = """\
    """

    placeit("views/text.html", {}, template)


path = sys.argv[1]
name = sys.argv[2]

def main():
    try:
      os.makedirs(path+"/"+name)
    except OSError as e:
      pass

    try:
      os.makedirs(path+"/"+name+"/"+"views")
    except OSError as e:
      pass

    gomod()
    gomain()
    run()
    ignore()

if __name__ == "__main__":
    main()
