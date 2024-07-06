{{ define "navbar" }}
  div navbar bg-base-200 font-familjen
    div navbar-start 
      div btn btn-ghost text-4xl
        a href=/
          img src=logo.png w-12
    div navbar-center flex hidden md:block
    div navbar-end
      div hidden md:block flex space-x-3
        {{ if .user }}
          a href=/ id=logout
            Logout
        {{ else }}
          a href=/core/login link link-hover
            Login
        {{ end }}
  {{ end }}
