div p-0 
  {{ template "navbar" . }}
  div flex justify-center mt-32
    div 
      div text-center
        h1 text-5xl font-bold
          Login
      form mt-6 space-y-3 id=login
        div
          input input input-primary id=email placeholder=email autofocus=true
        div
          input input input-primary id=password placeholder=password
        div flex space-x-6 justify-center
          div
            input type=submit btn btn-primary value=Go
      div mt-3 data-theme=light p-2 rounded-lg text-center
        span
          Need an account?
        a href=/core/register underline
          Register
      div hidden mt-6 text-sm flex justify-center
        a href=/core/forgot underline
          Password Help
