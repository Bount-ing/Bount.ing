/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./pages/**/*.vue', './components/**/*.vue', './layouts/**/*.vue', './plugins/**/*.vue', './app.vue'],
  theme: {
    extend: {
      screens: {
        'dark-mode': {'raw': '(prefers-color-scheme: dark)'},
      },
      colors: {
        'primary': '#006666',
        'primary-dark': '#004d4d',
        'primary-light': '#FFF',

        'secondary': '#002D0B', // Neutral, sophisticated gray
        'secondary-dark': '#001800', // Darker gray for contrast
        'secondary-light': '#014421', // Light gray for soft accents

        'success': '#2f855a', // Dark green, for success and confirmation
        'success-light': '#48bb78', // Light green, for positive feedback
        'error': '#9b2c2c', // Dark red, elegant and alerting
        'error-light': '#f56565', // Light red, for critical alerts
        'warning': '#975a16', // Refined amber, precise and clear
        'warning-light': '#ecc94b', // Bright amber, for warnings and notices
        'info': '#2c5282', // Deep blue for clarity and trust
        'info-light': '#4299e1', // Light blue for information and guidance

        'neutral': '#e1e1e1', // Soft, clean gray for neutral backgrounds
        'neutral-dark': '#4a5568', // Solid, dark gray for strong contrasts
        'neutral-light': '#f7fafc', // Very light gray, pure and clean

        'background': '#ffffff', // Pure white background, clean and bright
        'foreground': '#1a202c', // Almost black, for sharp and precise text
      },

      fontFamily: {
        'sans': ['Roboto Condensed', 'sans-serif'], // More modern, clean lines, good for dynamic content
        'serif': ['Merriweather', 'serif'], // Elegant but readable for longer texts
      },
      fontSize: {
        'base': '16px', // Standard for readability
        'lg': '18px', // Slightly larger for importance
        'xl': '20px', // Prominent for headings and important calls to action
      },
      lineHeight: {
        'normal': '1.6', // Optimal reading ease
        'heading': '1.3', // Tight for impactful headings
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
        'custom': '0 2px 4px rgba(0, 0, 0, 0.15)', // Subtle shadow for textural depth
        'strong': '0 4px 6px rgba(0, 0, 0, 0.20)', // More pronounced for critical interface elements
      },
      gradients: {
        'background': 'linear-gradient(135deg, #242424 0%, #000000 100%)', // Stealthy gradient for dramatic sections
        'background-dark': 'linear-gradient(135deg, #1A1A1A 0%, #333333 100%)', // Dark mode gradient that enhances the theme
      }
    },
  },
  plugins: [require('@tailwindcss/forms'), require('@tailwindcss/typography')],
  darkMode: 'media', // Automatically switch based on user's system preferences
}
