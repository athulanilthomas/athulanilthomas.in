<script setup lang="ts">
useHead({
  meta: [
    { name: 'viewport', content: 'width=device-width, initial-scale=1' },
    { name: 'theme-color', content: '#00DC82' },
    { name: 'description', content: 'Athul Anil Thomas - Full Stack Developer Portfolio' }
  ],
  link: [
    { rel: 'icon', href: '/favicon.ico' }
  ],
  htmlAttrs: {
    lang: 'en'
  }
})

const title = 'Athul Anil Thomas - Full Stack Developer'
const description = 'Exploring modern web development, open source contributions, and technical writing.'

useSeoMeta({
  title,
  description,
  ogTitle: title,
  ogDescription: description,
  ogImage: '/og-image.png',
  twitterCreator: '@yourtwitterhandle',
  twitterCard: 'summary_large_image'
})
</script>

<template>
  <UApp>
    <!-- Aurora background — fixed behind everything -->
    <div class="fixed inset-0 w-full h-full overflow-hidden">
      <UiAurora
        :color-stops="['#00DC82', '#052E16', '#00DC82']"
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
