import { rule } from "postcss/lib/postcss";

export default {
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,
  server: {
    host: 'localhost', // Default: localhost
    port: 3000 // You can specify the port here as well if needed
  },

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'Bount.ing',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
  '~/assets/css/main.css'
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/tailwindcss
    '@nuxtjs/tailwindcss',
    'nuxt-gtag',
    '@nuxtjs/auth-next',
    '@nuxtjs/dotenv',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    '@nuxtjs/axios',
    'nuxt-gtag',
    '@nuxtjs/dotenv',
    '@nuxtjs/auth-next',
    '@nuxtjs/tailwindcss',
    '@tailwindcss/forms',
    '@tailwindcss/typography'
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // Workaround to avoid enforcing hard-coded localhost:3000: https://github.com/nuxt-community/axios-module/issues/308
    baseURL: '/',
  },

  gtag: {
    id: 'G-JBW6WFECS1', // Replace G-XXXXXXX with your Measurement ID
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  },
  auth: {
    strategies: {
      github: {
        clientId: process.env.GITHUB_CLIENT_ID,
        clientSecret: process.env.GITHUB_CLIENT_SECRET,
        redirectUri: process.env.GITHUB_REDIRECT_URI,

        codeChallengeMethod: 'S256',
        responseType: 'token',
        scope: ['read:user', 'user:email', 'repo', 'read:org'],
        tokenType: 'Bearer',
      },
    },
    redirect: {
      login: '/login',
      logout: '/',
      callback: '/auth/github/callback',
      home: '/'
    },
    user: {
      property: 'user',
      // autoFetch: true
    },
    endpoints: {
      login: { url: '/api/auth/login', method: 'post' },
      logout: { url: '/api/auth/logout', method: 'post' },
      user: { url: '/api/auth/user', method: 'get' }
    }
  },
  eslint: {
    rules: {
      'no-console': 'off',
      'no-unused-vars': 'off'
    }
  }
}
