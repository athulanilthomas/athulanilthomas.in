<template>
  <PortfolioExperience :data="data" />
</template>

<script setup lang="ts">
import type { ExperienceData } from '~/types/section'

const { data } = await useAsyncData('experience', () =>
  queryCollection('experience').first() as Promise<ExperienceData>
)

useHead({
  title: 'Experience - Athul Anil Thomas',
})
</script>
