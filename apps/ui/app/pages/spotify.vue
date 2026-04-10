<template>
  <div class="flex flex-col items-center justify-center h-full min-h-[400px] gap-4">
    <div v-if="status === 'pending' || status === 'idle'" class="text-muted-foreground text-sm">
      Loading...
    </div>

    <div v-else-if="status === 'error'" class="text-center space-y-2">
      <UIcon name="i-heroicons-exclamation-circle" class="w-10 h-10 text-muted-foreground/50 mx-auto" />
      <p class="text-muted-foreground text-sm">Couldn't reach Spotify</p>
    </div>

    <template v-else>
      <SpotifyVinyl
        v-if="nowPlaying"
        :is-playing="true"
        :track-name="nowPlaying.item?.name"
        :artist-name="artists"
        :album-name="nowPlaying.item?.album?.name"
        :album-image="albumImage"
      />

      <div v-else class="text-center space-y-4">
        <SpotifyVinyl :is-playing="false" />
        <div class="space-y-2">
          <p class="text-muted-foreground/80 text-sm font-medium">It's quiet right now</p>
          <p class="text-muted-foreground/40 text-xs">Check back later to see what's spinning</p>
        </div>
      </div>
    </template>
  </div>
</template>

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

const { data: nowPlaying, status } = useLazyFetch<NowPlaying>('/api/now-playing')

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

useHead({ title })

useSeoMeta({
  title,
  description,
  ogTitle: title,
  ogDescription: description,
  robots: 'noindex, nofollow, noimageindex'
})
</script>
