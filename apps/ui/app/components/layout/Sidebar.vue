<template>
  <aside
    :class="[
      'bg-glass border border-glass-border shadow-glass backdrop-blur-[40px] backdrop-saturate-150 flex flex-col shrink-0',
      'fixed inset-y-0 right-0 z-50 w-72 rounded-r-2xl transition-transform duration-300 ease-out',
      'sm:relative sm:inset-auto sm:z-auto sm:rounded-2xl sm:translate-x-0 sm:w-64',
      isOpen ? 'translate-x-0' : 'translate-x-full sm:translate-x-0'
    ]"
  >
    <!-- Header -->
    <div class="p-4 flex items-center gap-3 border-b border-border/30">
      <div class="w-8 h-8 rounded bg-primary/20 flex items-center justify-center text-primary font-bold">
        {{ fullName.charAt(0) }}
      </div>
      <div class="flex-1 min-w-0">
        <h1 class="text-foreground font-semibold text-sm">{{ fullName }}</h1>
        <p class="text-muted-foreground text-xs">{{ headline }}</p>
      </div>
      <button
        type="button"
        @click="isOpen = false"
        class="sm:hidden p-1.5 rounded-md hover:bg-secondary/50 transition-colors"
      >
        <UIcon name="i-heroicons-x-mark" class="w-5 h-5 text-muted-foreground" />
      </button>
    </div>

    <!-- Content -->
    <div class="flex-1 overflow-y-auto p-3">
      <div class="text-xs font-medium text-muted-foreground uppercase tracking-wider mb-3 px-2">
        Explorer
      </div>

      <UNavigationMenu
        orientation="vertical"
        color="neutral"
        variant="link"
        :highlight="false"
        :items="navItems"
        class="w-full"
        :ui="{
          link: 'px-2 py-1.5 text-sm text-muted-foreground hover:text-foreground',
          linkLeadingIcon: 'size-4',
          linkTrailingIcon: 'size-4',
          linkLabel: 'text-sm',
          childLink: 'px-2 py-1.5 text-sm',
          childLinkIcon: 'size-4',
          childLinkLabel: 'text-sm truncate',
        }"
      />
    </div>

    <!-- Overlay for mobile -->
    <div
      v-if="isOpen"
      class="sm:hidden fixed inset-0 z-40 bg-background/80 backdrop-blur-sm"
      @click="isOpen = false"
    />
  </aside>

  <!-- Mobile Toggle Button -->
  <button
    type="button"
    @click="isOpen = true"
    class="sm:hidden absolute top-4 right-4 z-30 p-2.5 bg-glass border border-glass-border rounded-lg shadow-glass-pill backdrop-blur-[40px] backdrop-saturate-150 hover:bg-secondary/50 transition-colors"
  >
    <UIcon name="i-heroicons-chevron-left" class="w-5 h-5 text-foreground" />
  </button>
</template>

<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'

const route = useRoute()

const fullName = 'Athul Anil Thomas'
const headline = 'Full Stack Developer'

const isOpen = ref(false)

const navItems = computed<NavigationMenuItem[]>(() => [
  {
    label: 'about',
    icon: 'i-heroicons-folder',
    defaultOpen: true,
    children: [
      {
        label: 'about.md',
        icon: 'i-heroicons-document-text',
        to: '/about',
        active: route.path === '/about' || route.path === '/',
      },
    ],
  },
  {
    label: 'work',
    icon: 'i-heroicons-folder',
    defaultOpen: true,
    children: [
      {
        label: 'experience.md',
        icon: 'i-heroicons-document-text',
        to: '/experience',
        active: route.path === '/experience',
      },
      {
        label: 'projects.md',
        icon: 'i-heroicons-document-text',
        to: '/projects',
        active: route.path === '/projects',
      },
    ],
  },
  {
    label: 'career',
    icon: 'i-heroicons-folder',
    defaultOpen: true,
    children: [
      {
        label: 'skills.md',
        icon: 'i-heroicons-document-text',
        to: '/skills',
        active: route.path === '/skills',
      },
      {
        label: 'education.md',
        icon: 'i-heroicons-document-text',
        to: '/education',
        active: route.path === '/education',
      },
    ],
  },
])
</script>
