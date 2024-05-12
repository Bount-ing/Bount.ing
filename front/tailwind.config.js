/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./pages/**/*.vue', './components/**/*.vue', './layouts/**/*.vue', './plugins/**/*.vue', './app.vue'],
  theme: {
    extend: {
      screens: {
        'dark-mode': {'raw': '(prefers-color-scheme: dark)'},
      },
      colors: {
        "primary": "#0D47A1",
        "secondary": "#FBC02D",
        "accent": "#C62828",
        "highlight": "#FBC02D",
        "dark": "#000",
        "soft": "#111",
        "info": "#29B6F6",
        "success": "#2E7D32",
        "warning": "#FFA000",
        "error": "#D32F2F",
      },

      fontFamily: {
        'sans': ['Montserrat', 'sans-serif'],
        'serif': ['Playfair Display', 'serif'],
      },
      spacing: {
        '4': '1rem', '8': '2rem', '12': '3rem', '24': '6rem', '48': '12rem',
        '16': '4rem', '32': '8rem',
      },
      borderRadius: {
        'lg': '0.5rem',
        'xl': '1rem',
      },
      boxShadow: {
        'custom': '0 4px 8px rgba(255, 255, 255, 0.25)',
        'strong': '0 8px 12px rgba(255, 255, 255, 0.35)',
      },
      fontSize: {
        'base': '1rem', 'lg': '1.375rem', 'xl': '1.5rem',
      },
    },
  },
  plugins: [require('@tailwindcss/forms'), require('@tailwindcss/typography')],
}
