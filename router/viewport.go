package router

import "html/template"

// viewport := r.Header.Get("Viewport")

var viewport = template.HTML(
	`<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link href="https://fonts.bunny.net" rel="preconnect"/>
<link href="https://fonts.bunny.net/css?family=poppins:400,500,700,800" rel="stylesheet"/>
<link href="https://fonts.bunny.net/css?family=montserrat:400,500,700,800" rel="stylesheet"/>
<link href="https://fonts.bunny.net/css?family=oxygen-mono:400" rel="stylesheet" />
<link href="https://fonts.bunny.net/css?family=familjen-grotesk:400" rel="stylesheet" />
<link href="https://fonts.bunny.net/css?family=permanent-marker:400" rel="stylesheet" />
<link href="https://fonts.bunny.net/css?family=allan:400" rel="stylesheet" />
`)
