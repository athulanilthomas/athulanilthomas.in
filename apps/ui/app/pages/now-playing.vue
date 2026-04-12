<template>
  <div class="flex flex-col items-center h-full min-h-[400px]">
    <div class="flex-1 flex flex-col items-center justify-center gap-4">
      <Transition name="state" mode="out-in">
        <div v-if="status === 'error'" key="error" class="text-center space-y-2">
          <UIcon name="i-heroicons-exclamation-circle" class="w-10 h-10 text-muted-foreground/50 mx-auto" />
          <p class="text-muted-foreground text-sm">Couldn't reach Spotify</p>
        </div>

        <SpotifyVinyl
          v-else
          key="player"
          :is-playing="status === 'success' && !!nowPlaying"
          :is-loading="status === 'pending' || status === 'idle'"
          :track-name="nowPlaying?.item?.name"
          :artist-name="artists"
          :album-name="nowPlaying?.item?.album?.name"
          :album-image="albumImage"
        />
      </Transition>
    </div>

    <!-- Worker pool stats -->
    <div v-if="showWorkerStats" class="flex items-center gap-3 pb-4 pt-6 opacity-35 hover:opacity-60 transition-opacity duration-500">
      <UTooltip text="Background Go workers powering this live view" :delay-duration="300">
        <div class="flex items-center gap-3 cursor-default">
          <span class="relative flex h-2 w-2 shrink-0">
            <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-500 opacity-75" />
            <span class="relative inline-flex h-2 w-2 rounded-full bg-emerald-500" />
          </span>
          <UBadge color="neutral" variant="soft" size="sm" leading-icon="i-simple-icons-go">
            {{ goroutineCount }} goroutines
          </UBadge>
        </div>
      </UTooltip>
    </div>
  </div>
</template>

<style scoped>
.state-enter-active {
  transition: opacity 0.3s ease, transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}
.state-leave-active {
  transition: opacity 0.2s ease, transform 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}
.state-enter-from {
  opacity: 0;
  transform: translateY(8px);
}
.state-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}
</style>

<script setup lang="ts">
interface SpotifyImage {
  url: string
  height: number
  width: number
}

interface SpotifyArtist {
  name: string
}

interface SpotifyAlbum {
  name: string
  images: SpotifyImage[]
}

interface SpotifyTrack {
  name: string
  album: SpotifyAlbum
  artists: SpotifyArtist[]
}

interface NowPlaying {
  playing: boolean
  item: SpotifyTrack
  progress_ms: number
}

interface WorkerStats {
  goroutines: number
}

const { data: nowPlaying, status } = useLazyFetch<NowPlaying>('/api/now-playing')

const { data: workerStats } = useLazyFetch<WorkerStats | null>('/api/worker-stats')

const showWorkerStats = computed(() => !!workerStats.value && !!nowPlaying.value)
const goroutineCount = computed(() => workerStats.value?.goroutines ?? 0)

const artists = computed(() =>
  nowPlaying.value?.item?.artists?.map(a => a.name).join(', ')
)

const albumImage = computed(() => {
  const images = nowPlaying.value?.item?.album?.images
  if (!images?.length) return undefined
  // Prefer a ~300px image, fallback to first
  return (images.find(i => i.width >= 200 && i.width <= 400) ?? images[0])?.url
})

const title = 'Now Playing'
const description = 'See what Athul is currently listening to on Spotify.'

useHead({ title, meta: [{ name: 'robots', content: 'noindex' }] })

useSeoMeta({
  title,
  description,
  ogTitle: title,
  ogDescription: description,
  robots: 'noindex, nofollow, noimageindex'
})
</script>
