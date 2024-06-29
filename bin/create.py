#!/usr/bin/env /Users/aa/rick/foo/bin/python3

import sys

def replace_template_vars(template, replacements):
    for key, value in replacements.items():
        template = template.replace(f"{{{{{key}}}}}", value)
    return template

def foo():
    path = sys.argv[1]
    name = sys.argv[2]
    template = """\
    Hello, {{name}}!
    Welcome to {{place}}.
    We hope you enjoy your stay.
    """

    replacements = {
        "name": "Alice",
        "place": "Wonderland"
    }

    result = replace_template_vars(template, replacements)

    output_filename = path+"/"+name+"/"+"go.mod"
    with open(output_filename, 'w') as file:
        file.write(result)

foo()
