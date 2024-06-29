
import sys
import os

path = sys.argv[1]
name = sys.argv[2]

def replace_template_vars(template, replacements):
    for key, value in replacements.items():
        template = template.replace(f"{{{{{key}}}}}", value)
    return template

def placeit(filename, replacements, template):

    result = replace_template_vars(template, replacements)

    output_filename = path+"/"+name+"/"+filename
    with open(output_filename, 'w') as file:
        file.write(result)
    return output_filename
