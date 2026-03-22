<template>
  <PortfolioSkills :data="data" />
</template>

<script setup lang="ts">
import type { SkillsData } from '~/types/section'

const { data } = await useAsyncData('skills', () =>
  queryCollection('skills').first() as Promise<SkillsData>
)

useHead({
  title: 'Skills - Athul Anil Thomas',
})
</script>
