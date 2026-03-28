<template>
  <PortfolioExperience :data="data" />
</template>

<script setup lang="ts">
import type { ExperienceData } from '~/types/section'

const { data } = await useAsyncData('experience', () =>
  queryCollection('experience').first() as Promise<ExperienceData>
)

const title = 'Experience'
const description = 'Professional experience of Athul Anil Thomas as a Full Stack Developer, including roles at Capgemini and Alokin Software.'

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
