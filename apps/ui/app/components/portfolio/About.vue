<template>
  <div class="space-y-6">
    <h1 class="text-2xl font-bold text-foreground">About</h1>

    <div class="relative border border-glass-border rounded-xl h-64 w-full overflow-hidden bg-secondary/10">
      <ClientOnly>
        <UiLiquidEther
          class="absolute inset-0"
          :colors="['#00DC82', '#75EDAE', '#00C16A']"
          :auto-demo="true"
          :auto-speed="0.5"
          :auto-intensity="2.2"
          :mouse-force="20"
          :cursor-size="100"
          :resolution="0.5"
        />
      </ClientOnly>
      <div class="absolute z-10 left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 text-white text-center">
        <div class="w-32 h-32 rounded-full bg-primary/20 backdrop-blur-sm flex items-center justify-center">
          <span class="text-4xl font-bold text-primary">Your</span>
        </div>
      </div>
    </div>

    <p v-if="data?.about" class="text-muted-foreground leading-relaxed">{{ data.about }}</p>

    <div class="flex flex-wrap gap-4">
      <UButton icon="i-heroicons-arrow-down-tray" size="sm" color="primary" variant="outline">
        Get Resume
      </UButton>
      <UButton icon="i-heroicons-calendar" size="sm" color="primary" variant="outline">
        Book Session
      </UButton>
    </div>

    <p class="text-muted-foreground">
      I'm always interested in hearing about new projects and opportunities.
    </p>

    <p class="text-muted-foreground">Feel free to reach out to me via:</p>

    <div class="flex flex-wrap gap-3">
      <UButton 
        v-if="data?.email"
        icon="i-heroicons-envelope"
        size="sm" 
        color="primary"
        @click="openEmail"
      >
        Email
      </UButton>
      <UButton icon="i-heroicons-globe-alt" size="sm" color="gray">
        GitHub
      </UButton>
      <UButton icon="i-heroicons-window" size="sm" color="gray">
        LinkedIn
      </UButton>
      <UButton icon="i-heroicons-command-line" size="sm" color="gray">
        Toptal
      </UButton>
    </div>
  </div>
</template>

<script setup lang="ts">
interface AboutData {
  about?: string
  email?: string
  location?: string
}

interface Props {
  data?: AboutData
}

const props = withDefaults(defineProps<Props>(), {
  data: () => ({})
})

const openEmail = () => {
  if (props.data?.email) {
    window.location.href = `mailto:${props.data.email}`
  }
}
</script>
