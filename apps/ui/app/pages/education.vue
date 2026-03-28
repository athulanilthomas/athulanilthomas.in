<template>
  <PortfolioEducation :data="data" />
</template>

<script setup lang="ts">
import type { EducationData } from '~/types/section'

const { data } = await useAsyncData('education', () =>
  queryCollection('education').first() as Promise<EducationData>
)

const title = 'Education'
const description = 'Educational background of Athul Anil Thomas, including degrees and qualifications in computer science and engineering.'

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
