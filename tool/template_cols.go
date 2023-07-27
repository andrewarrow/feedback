package main

func colsTemplate() string {
	t := `
{{ "{{" }} define "_client_col1" }}
{{ "{{" }} $row := index . "row" }}
{{ "{{" }} $guid := index $row "guid" }}
<a href="/sd/clients/{{ "{{" }}$guid{{ "}}" }}" class="underline">
{{ "{{" }} index $row "name" }}
</a>
{{ "{{" }} end }}

{{ "{{" }} define "_client_col2" }}
{{ "{{" }} $row := index . "row" }}
{{ "{{" }} index $row "street1" }}
<br/>
{{ "{{" }} index $row "street2" }}
{{ "{{" }} end }}

{{ "{{" }} define "_client_col3" }}
{{ "{{" }} $row := index . "row" }}
{{ "{{" }} index $row "city" }},
{{ "{{" }} index $row "state" }}
{{ "{{" }} index $row "zip" }}
{{ "{{" }} index $row "country" }}
{{ "{{" }} end }}

{{ "{{" }} define "_client_col4" }}
{{ "{{" }} $row := index . "row" }}
{{ "{{" }} index $row "created_at_human" }}
<br/>
{{ "{{" }} index $row "created_at_ago" }}
{{ "{{" }} end }}
`
	return t
}
