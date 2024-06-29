#!/usr/bin/env /Users/aa/rick/foo/bin/python3

def replace_template_vars(template, replacements):
    for key, value in replacements.items():
        template = template.replace(f"{{{{{key}}}}}", value)
    return template

def foo():
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

    output_filename = "test.txt"
    with open(output_filename, 'w') as file:
        file.write(result)

foo()
