<template>
  <PortfolioProjects :data="merged" />
</template>

<script setup lang="ts">
import type { ProjectsData, Repository } from '~/types/section'

const { data: content } = await useAsyncData('projects', () =>
  queryCollection('projects').first() as Promise<ProjectsData>
)

const { data: repos } = await useAsyncData('repos', async () => {
  return await $fetch<ProjectsData>('/api/repos')
})

const merged = computed<ProjectsData>(() => {
  const contentRepos = content.value?.repositories ?? []
  const apiRepos = repos.value?.repositories ?? []

  const repoMap = new Map<string, Repository>()

  for (const repo of contentRepos) {
    repoMap.set(repo.nameWithOwner, repo)
  }

  for (const repo of apiRepos) {
    repoMap.set(repo.nameWithOwner, { ...repoMap.get(repo.nameWithOwner), ...repo })
  }

  return { repositories: [...repoMap.values()] }
})

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
