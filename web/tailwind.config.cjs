/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{svelte,html,js}"],
  theme: {
    extend: {},
  },
  plugins: [
    require("daisyui"),
    require('@tailwindcss/typography')
  ],

  daisyui: {
    themes: [
      {
        light: {
          ...require("daisyui/src/colors/themes")["[data-theme=business]"],
        },
      },
    ]
  }
}
