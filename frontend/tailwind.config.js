const colors = require('tailwindcss/colors')

module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  darkMode: "class", // or 'media' or 'class'
  theme: {
    extend: {
      width: {
        "36r": "36rem",
        "4-5r": "4.5rem",
      },
      colors: {
        darkgray: {
          DEFAULT: "#22272e",
        },
        colors: {
          green: colors.emerald,
          yellow: colors.amber,
          purple: colors.violet,
        },
      },
    },
  },
  plugins: [],
};
