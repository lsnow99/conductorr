module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: 'class', // or 'media' or 'class'
  theme: {
    extend: {
      width: {
        '36r': '36rem'
      }
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
