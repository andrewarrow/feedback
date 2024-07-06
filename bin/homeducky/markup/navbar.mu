{{ define "navbar" }}
  div navbar bg-base-200 font-familjen
    div navbar-start 
      div btn btn-ghost text-4xl
        a href=/
          img src=logo.png w-12
    div navbar-center flex hidden md:block
    div navbar-end
      div hidden md:block flex space-x-3
        a href=/frame/routing link link-hover
          Routing
        a href=/frame/sql link link-hover
          SQL
        a href=/frame/migrations link link-hover
          Migrations
        a href=/frame/wasm link link-hover
          WASM
        a href=/frame/faq link link-hover
          FAQ
  {{ end }}
