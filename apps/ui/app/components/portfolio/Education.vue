<template>
  <div class="space-y-6">
    <h1 class="text-2xl font-bold text-foreground">Education</h1>

    <div v-if="data?.educations && data.educations.length === 0" class="text-muted-foreground">
      No education yet.
    </div>

    <div v-else class="space-y-4">
      <div v-for="edu in data?.educations" :key="edu.id" class="flex gap-4">
        <div class="shrink-0">
          <UIcon name="i-heroicons-academic-cap" class="w-12 h-12 text-primary" />
        </div>

        <div class="flex-1">
          <h2 class="text-lg font-semibold text-foreground">{{ edu.school }}</h2>

          <p class="text-sm text-muted-foreground">
            <template v-if="edu.degree">
              {{ edu.degree }} in
            </template>
            {{ edu.areaOfStudy }}
          </p>

          <p class="text-sm text-muted-foreground">
            {{ formatDate(edu.startDate, edu.endDate) }}
          </p>

          <p v-if="edu.description" class="text-sm text-muted-foreground whitespace-pre-line mt-2">
            {{ edu.description }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Education {
  id: string
  school: string
  degree?: string | null
  areaOfStudy: string
  startDate: string
  endDate?: string | null
  description?: string | null
}

interface EducationData {
  educations: Education[]
}

interface Props {
  data?: EducationData
}

withDefaults(defineProps<Props>(), {
  data: () => ({ educations: [] })
})

const formatDate = (start: string, end?: string | null) => {
  if (!end) {
    return `${start} - Present`
  }
  return `${start} - ${end}`
}
</script>
