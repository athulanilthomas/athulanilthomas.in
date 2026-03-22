<template>
  <PortfolioProjects :data="data" />
</template>

<script setup lang="ts">
import type { ProjectsData } from '~/types/section'

const { data } = await useAsyncData('projects', () =>
  queryCollection('projects').first() as Promise<ProjectsData>
)

useHead({
  title: 'Projects - Athul Anil Thomas',
})
</script>
