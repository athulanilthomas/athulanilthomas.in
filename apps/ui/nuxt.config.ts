// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    pageTransition: { name: 'page', mode: 'out-in' }
  },

  modules: [
    '@nuxt/eslint',
    '@nuxt/ui',
    '@nuxt/content',
    '@nuxt/image',
    '@nuxtjs/seo'
  ],

  site: {
    url: 'https://athulanilthomas.in',
    name: 'Athul Anil Thomas',
    description: 'Full Stack Developer specializing in Vue.js, Nuxt, and modern web technologies. Portfolio showcasing projects, experience, and open source contributions.',
    defaultLocale: 'en'
  },

  ogImage: {
    enabled: false
  },

  sitemap: {
    enabled: true
  },

  robots: {
    allow: '/',
    sitemap: '/sitemap.xml'
  },

  schemaOrg: {
    identity: {
      type: 'Person',
      name: 'Athul Anil Thomas',
      url: 'https://athulanilthomas.in',
      image: 'https://avatars.githubusercontent.com/u/30122216',
      jobTitle: 'Full Stack Developer',
      sameAs: [
        'https://github.com/athulanilthomas',
        'https://www.linkedin.com/in/athul-anil-thomas',
        'https://twitter.com/AthulAnilThoma2'
      ]
    }
  },

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
