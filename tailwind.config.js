/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./notes/**/*.{html,js,templ}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["night"]
  }
}