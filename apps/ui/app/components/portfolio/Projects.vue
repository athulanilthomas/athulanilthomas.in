<template>
  <div class="space-y-6 min-w-0 overflow-hidden">
    <h1 class="text-2xl font-bold text-foreground">Projects</h1>

    <div v-if="!data?.repositories || data.repositories.length === 0" class="text-muted-foreground">
      No projects yet.
    </div>

    <template v-else>
      <div v-if="ownedRepos.length > 0" class="space-y-3">
        <h2 class="text-sm font-medium text-muted-foreground uppercase tracking-wider">
            Recently Updated
        </h2>
        <div class="grid gap-3">
          <a
            v-for="repo in ownedRepos"
            :key="repo.id"
            :href="repo.url"
            target="_blank"
            rel="noopener noreferrer"
            class="group flex flex-col gap-2 p-3 rounded-lg transition-colors border border-transparent hover:border-primary/30 hover:bg-primary/5 min-w-0 overflow-hidden"
          >
            <div class="flex items-center justify-between gap-4">
              <div class="flex items-center gap-3 min-w-0">
                <NuxtImg
                  :src="repo.ownerAvatarUrl"
                  :alt="repo.nameWithOwner.split('/')[0]"
                  width="24"
                  height="24"
                  class="rounded-full shrink-0"
                />
                <h3 class="font-medium text-foreground group-hover:text-primary transition-colors truncate">
                  {{ repo.nameWithOwner }}
                </h3>
              </div>

              <div class="flex items-center gap-2 text-sm text-muted-foreground shrink-0">
                <UIcon name="i-heroicons-star" class="w-4 h-4" />
                <span>{{ repo.stargazerCount }}</span>
              </div>
            </div>

            <p v-if="repo.description" class="text-sm text-muted-foreground line-clamp-2">
              {{ repo.description }}
            </p>

            <div v-if="repo.primaryLanguage" class="flex items-center gap-2 text-sm text-muted-foreground">
              <span class="w-2 h-2 rounded-full bg-primary" />
              {{ repo.primaryLanguage }}
            </div>
          </a>
        </div>
      </div>

      <div v-if="contributedRepos.length > 0" class="space-y-3">
        <h2 class="text-sm font-medium text-muted-foreground uppercase tracking-wider">
          Contributions
        </h2>
        <div class="grid gap-2">
          <a
            v-for="repo in contributedRepos"
            :key="repo.id"
            :href="repo.url"
            target="_blank"
            rel="noopener noreferrer"
            class="group flex items-center gap-3 p-2 rounded-lg transition-colors hover:bg-secondary/30 min-w-0 overflow-hidden"
          >
            <NuxtImg
              :src="repo.ownerAvatarUrl"
              :alt="repo.nameWithOwner.split('/')[0]"
              width="24"
              height="24"
              class="rounded-full shrink-0"
            />

            <div class="flex items-center justify-between gap-4 flex-1 min-w-0">
              <h3 class="font-medium text-foreground group-hover:text-primary transition-colors truncate">
                {{ repo.nameWithOwner }}
              </h3>

              <div class="flex items-center gap-2 text-sm text-muted-foreground shrink-0 max-sm:hidden">
                <UIcon name="i-heroicons-star" class="w-3 h-3" />
                <span>{{ repo.stargazerCount }}</span>
              </div>
            </div>
          </a>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Repository {
  id: string
  nameWithOwner: string
  description: string | null
  url: string
  stargazerCount: number
  primaryLanguage: string | null
  ownerAvatarUrl: string
  isOwned: boolean
}

interface ProjectsData {
  repositories: Repository[]
}

interface Props {
  data?: ProjectsData
}

const props = withDefaults(defineProps<Props>(), {
  data: () => ({ repositories: [] })
})

const ownedRepos = computed(() => {
  return props.data?.repositories?.filter(r => r.isOwned) || []
})

const contributedRepos = computed(() => {
  return props.data?.repositories?.filter(r => !r.isOwned) || []
})
</script>
