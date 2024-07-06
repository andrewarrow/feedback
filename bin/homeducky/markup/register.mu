div p-0 
  {{ template "navbar" . }}
  div flex justify-center mt-32
    div 
      div text-center
        h1 text-5xl font-bold
          Register
      form mt-6 space-y-3 id=register
        div
          input input input-primary id=email placeholder=email autofocus=true
        div
          input input input-primary id=password placeholder=password
        div flex justify-center
          input type=submit btn btn-primary value=Go
      div mt-3 text-center
        span
          Already have account?
        a href=/core/login underline
          Login
