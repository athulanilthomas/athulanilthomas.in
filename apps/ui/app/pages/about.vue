<template>
  <PortfolioAbout :data="data" />
</template>

<script setup lang="ts">
import type { AboutData } from '~/types/section'

definePageMeta({
  alias: ['/'],
})

const { data } = await useAsyncData('about', () =>
  queryCollection('about').first() as Promise<AboutData>
)

const title = 'About'
const description = 'Learn about Athul Anil Thomas, a Full Stack Developer specializing in Vue.js, Nuxt, and modern web technologies based in India.'

useHead({ title })

useSeoMeta({
  title,
  description,
  ogTitle: title,
  ogDescription: description
})

useSchemaOrg([
  defineWebPage({
    name: title,
    description
  })
])
</script>
