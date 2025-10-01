/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'bookstore-primary': '#2d5a4d',
        'bookstore-secondary': '#5fe9bc',
        viridian: {
          100: '#d4f2eb',
          200: '#a8e5d7',
          300: '#7bd8c3',
          400: '#4fcbaf',
          500: '#23be9b',
          600: '#1c987c',
          700: '#15725d',
          800: '#0e4c3e',
          900: '#07261f',
        },
      },
      fontFamily: {
        'sans': ['Prompt', 'sans-serif'],
      }
    },
  },
  plugins: [],
}
