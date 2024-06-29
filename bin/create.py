#!/usr/bin/env /Users/aa/rick/foo/bin/python3

import sys

path = sys.argv[1]
name = sys.argv[2]

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

    result = replace_template_vars(template, replacements)

    output_filename = path+"/"+name+"/"+"go.mod"
    with open(output_filename, 'w') as file:
        file.write(result)

gomod()
