/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      "./index.html",
      "./src/**/*.{vue,js,ts,jsx,tsx}",
    ],
    theme: {
      extend: {
        screens: {
          'dark-mode': {'raw': '(prefers-color-scheme: dark)'},
        },
        colors: {
          // Primary Colors (Turquoise theme)
          'primary': '#008888', // Turquoise
          'primary-dark': '#004d4d', // Darker turquoise
          'primary-light': '#00b3b3', // Lighter turquoise
  
          // Secondary Colors (Carbon and Black theme)
          'secondary': '#191919', // Black
          'secondary-dark': '#000', // Dark gray/Carbon
          'secondary-light': '#666', // Light gray
  
          // Success, Error, Warning, Info
          'success': '#2f855a', // Dark green
          'success-light': '#48bb78', // Light green
          'error': '#9b2c2c', // Dark red
          'error-light': '#f56565', // Light red
          'warning': '#975a16', // Amber
          'warning-light': '#ecc94b', // Bright amber
          'info': '#2c5282', // Deep blue
          'info-light': '#4299e1', // Light blue
  
          // Neutral Colors
          'neutral': '#e1e1e1', // Light gray for neutral backgrounds
          'neutral-dark': '#4a5568', // Dark gray for contrasts
          'neutral-light': '#f7fafc', // Very light gray
  
          // Background and Foreground (Light and Dark)
          'background': '#121212', // Dark background for dark theme
          'background-light': '#ffffff', // Light background for light theme
          'foreground': '#ffffff', // Light text for dark theme
          'foreground-dark': '#121212', // Dark text for light theme
  
          // Specific Dark Mode Background Gradient
          'background-dark': 'linear-gradient(135deg, #1A1A1A 0%, #333333 100%)', // Dark mode gradient
  
          // Specific Light Mode Background Gradient
          'background-light': 'linear-gradient(135deg, #ffffff 0%, #f0f0f0 100%)', // Light mode gradient
        },
        
        fontFamily: {
          'sans': ['Roboto Mono', 'monospace'],
          'serif': ['Philosopher', 'serif'],
          'mono': ['VT323', 'monospace'],
        },
        
        fontSize: {
          'base': '16px', // Standard for readability
          'lg': '18px', // Slightly larger for importance
          'xl': '20px', // Prominent for headings
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
          'custom': '0 2px 4px rgba(0, 0, 0, 0.15)', // Subtle shadow
          'strong': '0 4px 6px rgba(0, 0, 0, 0.20)', // Pronounced shadow
        },
      },
    },
    plugins: [require('@tailwindcss/forms'), require('@tailwindcss/typography')],
    darkMode: 'media', // Automatically switch based on user's system preferences
  }
  