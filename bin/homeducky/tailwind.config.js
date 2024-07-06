/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['views/*.html',],
  theme: {
    extend: {
      colors: {
        'cream': '#EFDECD',
        'lime': '#8FBC8F',
        'a-blue': '#4A88EE',
        'a-dark': '#00364d',
        'a-good': '#00364d'
      },
      fontFamily: {
        pragmatica: ['Pragmatica'],
        familjen: ['Familjen Grotesk'],
      },
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["light", "dark", "luxury", "sunset"],
  },
}
    