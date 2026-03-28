<template>
  <PortfolioSkills :data="data" />
</template>

<script setup lang="ts">
import type { SkillsData } from '~/types/section'

const { data } = await useAsyncData('skills', () =>
  queryCollection('skills').first() as Promise<SkillsData>
)

const title = 'Skills'
const description = 'Technical skills and expertise of Athul Anil Thomas including Vue.js, Nuxt, TypeScript, Node.js, and more.'

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
