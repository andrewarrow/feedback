#!/usr/bin/env python3

import sys
import os
import shutil
from gomain import gomain
from placeit import placeit
from feedback import feedback
from browser import browser
from wasm import wasm

def gomod():
    template = """\
module {{name}}

//replace github.com/andrewarrow/feedback => /Users/aa/os/feedback

go 1.21.0

require github.com/andrewarrow/feedback v0.0.0-20240630205923-a0616bffee08
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
cp main.go save_main
cp wasm/main.go .
GOOS=js GOARCH=wasm go build -ldflags="-s -w -X main.useLive=true" -o assets/other/json.wasm 
cd assets/other
rm json.wasm.gz
gzip json.wasm
cd ../..
if [ $? -eq 0 ]; then
    tailwindcss -i assets/css/tail.components.css -o assets/css/tail.min.css --minify
    uuid=$(uuidgen); go build -ldflags="-X main.buildTag=$uuid"
    echo 3
    ./{{name}} run 3000
else
  mv save_main main.go
fi
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

    tailwind = """\
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer components {

}
    """

    tailwindconfig = """\
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['views/*.html',],
  theme: {
    extend: {
      colors: {
        'cream': '#EFDECD',
        'lime': '#8FBC8F',
        'a-blue': '#4A88EE',
        'a-dark': '#00364d',
        'a-good': '#00364d'
      },
      fontFamily: {
        pragmatica: ['Pragmatica'],
        familjen: ['Familjen Grotesk'],
      },
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["light", "dark", "luxury", "sunset"],
  },
}
    """

    placeit(".gitignore", {"name": name}, template)
    placeit("views/text.html", {}, "")
    placeit("tailwind.config.js", {}, tailwindconfig)
    placeit("assets/css/tail.components.css", {}, tailwind)

path = sys.argv[1]
name = sys.argv[2]

def main():
    try:
      os.makedirs(path+"/"+name)
      os.makedirs(path+"/"+name+"/"+"views")
      os.makedirs(path+"/"+name+"/"+"app")
      os.makedirs(path+"/"+name+"/"+"browser")
      os.makedirs(path+"/"+name+"/"+"wasm")
      os.makedirs(path+"/"+name+"/"+"assets/css")
      os.makedirs(path+"/"+name+"/"+"assets/images")
      os.makedirs(path+"/"+name+"/"+"assets/other")
      markup = path+"/"+name+"/"+"markup"
      os.makedirs(markup)
      js = path+"/"+name+"/"+"assets/javascript"
      os.makedirs(js)
      shutil.copyfile("../assets/javascript/wasm_exec.js", js+"/wasm_exec.js")
      shutil.copyfile("../markup/application_layout.mu", markup+"/application_layout.mu")
      shutil.copyfile("../markup/navbar.mu", markup+"/navbar.mu")
      shutil.copyfile("../markup/404.mu", markup+"/404.mu")
      shutil.copyfile("../markup/welcome.mu", markup+"/welcome.mu")
    except OSError as e:
      pass

    gomod()
    gomain()
    run()
    ignore()
    feedback()
    browser()
    wasm()

if __name__ == "__main__":
    main()
