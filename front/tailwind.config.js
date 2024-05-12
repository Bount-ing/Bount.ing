/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./pages/**/*.vue', './components/**/*.vue', './layouts/**/*.vue', './plugins/**/*.vue'],
  content: ['./pages/**/*.vue', './components/**/*.vue', './layouts/**/*.vue', './plugins/**/*.vue'],
  theme: {
    extend: {},
    extend: {
      colors: {  // Define the color scheme for the application
        'primary': '#007aff',  // Apple blue
        'secondary': '#5856d6',
        'danger': '#ff3b30',
        'background': '#f9fafb',
        'text': '#333333',
      },
      fontFamily: {
        'sans': ['Helvetica Neue', 'Arial', 'sans-serif'],
        'serif': ['Times New Roman', 'serif'],
      },
      spacing: {
        '4': '1rem', '8': '2rem', '16': '4rem', '32': '8rem', '64': '16rem', // Extended spacings for responsive design
      },
    },
  },
  plugins: [],
  plugins: [require('@tailwindcss/forms'), require('@tailwindcss/typography')], // Include necessary plugins for forms and typography
}

