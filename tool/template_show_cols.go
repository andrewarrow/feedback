package main

func showColsTemplate() string {
	t := `
{{$name := index . "name"}}
{{$lower := index . "lower"}}
{{ "{{" }} define "_{{$lower}}_show_col1" {{ "}}" }}
{{ "{{" }} $row := index . "row" {{ "}}" }}
{{ "{{" }} $row {{ "}}" }}
{{ "{{" }} end {{ "}}" }}

{{ "{{" }} define "_{{$lower}}_show_col2" {{ "}}" }}

{{ "{{" }} template "_editable_fields" . {{ "}}" }}

{{ "{{" }} end {{ "}}" }}
`
	return t
}
