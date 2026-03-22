<template>
  <div class="space-y-6">
    <h1 class="text-2xl font-bold text-foreground">Skills</h1>

    <div v-if="data?.skills && data.skills.length === 0" class="text-muted-foreground">
      No skills yet.
    </div>

    <div v-else class="grid grid-cols-1 sm:grid-cols-3 gap-y-3">
      <div v-for="skill in data?.skills" :key="skill.id" class="flex items-center gap-3">
        <img
          :src="skill.logoUrl"
          :alt="`${skill.title} logo`"
          width="24"
          height="24"
          class="shrink-0 rounded"
        />
        <span class="font-medium text-foreground truncate">{{ skill.title }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Skill {
  id: string
  title: string
  yearsOfExperience: number
  logoUrl: string
}

interface SkillsData {
  skills: Skill[]
  maxYears: number
}

interface Props {
  data?: SkillsData
}

withDefaults(defineProps<Props>(), {
  data: () => ({ skills: [], maxYears: 0 })
})
</script>
