<template>
  <PortfolioEducation :data="data" />
</template>

<script setup lang="ts">
import type { EducationData } from '~/types/section'

const { data } = await useAsyncData('education', () =>
  queryCollection('education').first() as Promise<EducationData>
)

useHead({
  title: 'Education - Athul Anil Thomas',
})
</script>
