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
          DEFAULT: "#242424",
        },
        colors: {
          green: colors.emerald,
          yellow: colors.amber,
          purple: colors.violet,
        },
        gray: {
          '600': '#1a1a1a'
        }
      },
    },
  },
  plugins: [],
};
