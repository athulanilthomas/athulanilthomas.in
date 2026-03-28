// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    pageTransition: { name: 'page', mode: 'out-in' }
  },

  modules: [
    '@nuxt/eslint',
    '@nuxt/ui',
    '@nuxt/content',
    '@nuxt/image'
  ],

  image: {
    domains: ['avatars.githubusercontent.com']
  },

  runtimeConfig: {
    apiBase: '',
  },

  devtools: {
    enabled: true
  },

  vite: {
    optimizeDeps: {
      include: ['ogl', 'three']
    }
  },

  css: ['~/assets/css/main.css'],

  routeRules: {
    '/': { prerender: true },
    '/about': { prerender: true },
    '/experience': { prerender: true },
    '/skills': { prerender: true },
    '/education': { prerender: true },
    '/projects': { prerender: true }
  },

  compatibilityDate: '2025-01-15',

  eslint: {
    config: {
      stylistic: {
        commaDangle: 'never',
        braceStyle: '1tbs'
      }
    }
  }
})
