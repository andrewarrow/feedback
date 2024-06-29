#!/usr/bin/env /Users/aa/rick/foo/bin/python3

import sys
import os
import shutil
from gomain import gomain
from placeit import placeit
from feedback import feedback

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
./{{name}} render
go build
echo 3
./{{name}} run 3000
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
    placeit("views/text.html", {}, "")
    placeit("views/welcome.html", {}, "hello")

path = sys.argv[1]
name = sys.argv[2]

def main():
    try:
      os.makedirs(path+"/"+name)
      os.makedirs(path+"/"+name+"/"+"views")
      os.makedirs(path+"/"+name+"/"+"app")
      os.makedirs(path+"/"+name+"/"+"browser")
      markup = path+"/"+name+"/"+"markup"
      os.makedirs(markup)
      js = path+"/"+name+"/"+"assets/javascript"
      os.makedirs(js)
      shutil.copyfile("../assets/javascript/wasm_exec.js", js+"/wasm_exec.js")
      shutil.copyfile("../markup/application_layout.mu", markup+"/application_layout.mu")
    except OSError as e:
      pass

    gomod()
    gomain()
    run()
    ignore()
    feedback()

if __name__ == "__main__":
    main()
