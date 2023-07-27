package main

func colsTemplate() string {
	t := `
{{$name := index . "name"}}
{{$lower := index . "lower"}}
{{ "{{" }} define "_{{$lower}}_col1" {{ "}}" }}
{{ "{{" }} $row := index . "row" {{ "}}" }}
{{ "{{" }} $guid := index $row "guid" {{ "}}" }}
<a href="/sd/clients/{{ "{{" }}$guid{{ "}}" }}" class="underline">
{{ "{{" }} index $row "name" {{ "}}" }}
</a>
{{ "{{" }} end {{ "}}" }}

{{ "{{" }} define "_{{$lower}}_col2" {{ "}}" }}
{{ "{{" }} $row := index . "row" {{ "}}" }}
{{ "{{" }} index $row "street1" {{ "}}" }}
<br/>
{{ "{{" }} index $row "street2" {{ "}}" }}
{{ "{{" }} end {{ "}}" }}

{{ "{{" }} define "_{{$lower}}_col3" {{ "}}" }}
{{ "{{" }} $row := index . "row" {{ "}}" }}
{{ "{{" }} index $row "city" {{ "}}" }},
{{ "{{" }} index $row "state" {{ "}}" }}
{{ "{{" }} index $row "zip" {{ "}}" }}
{{ "{{" }} index $row "country" {{ "}}" }}
{{ "{{" }} end {{ "}}" }}

{{ "{{" }} define "_{{$lower}}_col4" {{ "}}" }}
{{ "{{" }} $row := index . "row" {{ "}}" }}
{{ "{{" }} index $row "created_at_human" {{ "}}" }}
<br/>
{{ "{{" }} index $row "created_at_ago" {{ "}}" }}
{{ "{{" }} end {{ "}}" }}
`
	return t
}
