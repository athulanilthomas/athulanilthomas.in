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

useHead({
  title: 'About - Athul Anil Thomas',
})
</script>
