<template>
  <div class="space-y-6">
    <h1 class="text-2xl font-bold text-foreground">Experience</h1>

    <div v-if="data?.experiences && data.experiences.length === 0" class="text-muted-foreground">
      No experience yet.
    </div>

    <div v-else class="space-y-4">
      <div v-for="(exp, index) in data?.experiences" :key="exp.id">
        <hr v-if="index > 0" class="border-border/50 mb-8" />

        <div class="flex gap-4">
          <div class="shrink-0">
            <img
              :src="exp.logoUrl"
              :alt="`${exp.company} logo`"
              width="48"
              height="48"
              class="rounded-lg bg-secondary"
            />
          </div>

          <div class="space-y-2 min-w-0">
            <h2 class="text-lg font-semibold text-foreground">{{ exp.position }}</h2>

            <div class="flex flex-wrap items-center gap-x-2 text-sm text-muted-foreground">
              <span class="font-medium text-foreground">{{ exp.company }}</span>
              <template v-if="exp.location">
                <span>·</span>
                <span>{{ exp.location }}</span>
              </template>
              <span>|</span>
              <span>{{ formatDate(exp.startDate, exp.endDate) }}</span>
            </div>

            <p class="text-sm text-muted-foreground whitespace-pre-line">{{ exp.description }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Experience {
  id: string
  company: string
  position: string
  location?: string | null
  startDate: string
  endDate?: string | null
  description?: string | null
  logoUrl: string
}

interface ExperienceData {
  experiences: Experience[]
}

interface Props {
  data?: ExperienceData
}

withDefaults(defineProps<Props>(), {
  data: () => ({ experiences: [] })
})

const formatDate = (start: string, end?: string | null) => {
  if (!end) {
    return `${start} - Present`
  }
  return `${start} - ${end}`
}
</script>
