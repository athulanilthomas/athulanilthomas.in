<template>
  <div class="space-y-6">
    <h1 class="text-2xl font-bold text-foreground">About</h1>

    <div class="relative border border-glass-border rounded-xl h-64 w-full overflow-hidden bg-secondary/10">
      <ClientOnly>
        <LazyUiLiquidEther
          class="absolute inset-0"
          :colors="['#00dc82', '#66edb8', '#00ad61']"
          :auto-demo="true"
          :auto-speed="0.5"
          :auto-intensity="2.2"
          :mouse-force="20"
          :cursor-size="100"
          :resolution="0.5"
        />
      </ClientOnly>
      <NuxtImg
        src="/assets/avatar.png"
        alt="Avatar"
        width="256"
        height="256"
        class="absolute z-10 left-1/2 bottom-0 -translate-x-1/2 w-52 h-52 sm:w-56 sm:h-56 object-contain"
      />
    </div>

    <p v-if="data?.about" class="text-muted-foreground leading-relaxed">{{ data.about }}</p>

    <div class="flex flex-wrap gap-4">
      <UButton icon="i-heroicons-arrow-down-tray" size="sm" color="primary" variant="outline">
        Get Resume
      </UButton>
      <UButton
        icon="i-simple-icons-calendly"
        size="sm"
        color="primary"
        variant="outline"
        to="https://calendly.com/athulanilthomas/30min"
        target="_blank"
      >
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
        :to="`mailto:${data.email}`"
      >
        Email
      </UButton>
      <UButton
        icon="i-simple-icons-github"
        size="sm"
        color="neutral"
        to="https://github.com/athulanilthomas"
        target="_blank"
      >
        GitHub
      </UButton>
      <UButton
        icon="i-simple-icons-linkedin"
        size="sm"
        color="neutral"
        to="https://www.linkedin.com/in/athul-anil-thomas"
        target="_blank"
      >
        LinkedIn
      </UButton>
      <UButton
        icon="i-simple-icons-x"
        size="sm"
        color="neutral"
        to="https://twitter.com/AthulAnilThoma2"
        target="_blank"
      >
        X (Twitter)
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

withDefaults(defineProps<Props>(), {
  data: () => ({})
})
</script>
