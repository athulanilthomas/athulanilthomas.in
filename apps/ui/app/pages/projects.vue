<template>
  <PortfolioProjects :data="data" />
</template>

<script setup lang="ts">
import type { ProjectsData } from '~/types/section'

const { data } = await useAsyncData('projects', () =>
  queryCollection('projects').first() as Promise<ProjectsData>
)

const title = 'Projects'
const description = 'Open source projects and contributions by Athul Anil Thomas, including TypeScript SDKs, portfolio websites, and community contributions.'

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
