/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.{html,js,templ}",
    "./notes/**/*.{html,js,templ}",
    "./js-playground/**/*.{html,js,templ}",
  ],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["night"]
  }
}