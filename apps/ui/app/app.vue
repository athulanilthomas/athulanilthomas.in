<script setup lang="ts">
useHead({
  meta: [
    { name: 'viewport', content: 'width=device-width, initial-scale=1' },
    { name: 'theme-color', content: '#00dc82' },
    { name: 'author', content: 'Athul Anil Thomas' }
  ],
  link: [
    { rel: 'icon', type: 'image/svg+xml', href: '/favicon.svg' }
  ],
  htmlAttrs: {
    lang: 'en'
  }
})

const title = 'Full Stack Developer'
const description = 'Full Stack Developer specializing in Vue.js, Nuxt, and modern web technologies. Portfolio showcasing projects, experience, and open source contributions.'

useHead({
  titleTemplate: (title) => title ? `${title} | Athul Anil Thomas` : 'Athul Anil Thomas - Full Stack Developer'
})

useSeoMeta({
  title,
  description,
  ogTitle: title,
  ogDescription: description,
  ogImage: '/og-image.png',
  ogType: 'website',
  ogLocale: 'en_US',
  ogSiteName: 'Athul Anil Thomas',
  twitterCreator: '@AthulAnilThoma2',
  twitterSite: '@AthulAnilThoma2',
  twitterCard: 'summary_large_image',
  twitterTitle: title,
  twitterDescription: description,
  twitterImage: '/og-image.png'
})
</script>

<template>
  <UApp>
    <!-- Aurora background — fixed behind everything -->
    <div class="fixed inset-0 w-full h-full overflow-hidden">
      <UiAurora
        :color-stops="['#00dc82', '#66edb8', '#00ad61']"
        :amplitude="1.0"
        :blend="0.5"
        :speed="1.0"
        :intensity="1.0"
      />
    </div>

    <!-- Main content -->
    <div class="relative flex items-center justify-center w-full min-h-screen sm:p-8">
      <NuxtLayout>
        <NuxtPage />
      </NuxtLayout>
    </div>
  </UApp>
</template>

<style scoped>
:global(body) {
  background-color: #f8fafc;
}

:global(.dark body) {
  background-color: #0f172a;
}
</style>
