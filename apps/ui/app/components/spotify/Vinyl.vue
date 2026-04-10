<template>
  <div class="player-container">
    <!-- Record player -->
    <div class="player">
      <!-- Album sleeve -->
      <div :class="['sleeve', { 'sleeve-shifted': isPlaying, 'sleeve-idle': !isPlaying }]">
        <div class="sleeve-front">
          <NuxtImg v-if="albumImage" :src="albumImage" :alt="albumName" loading="lazy" width="300" height="300" format="webp" class="sleeve-art" />
          <div v-else class="sleeve-placeholder">
            <UIcon name="i-simple-icons-spotify" class="sleeve-placeholder-icon" />
          </div>
        </div>
      </div>

      <!-- Vinyl record -->
      <div :class="['vinyl-area', { 'vinyl-out': isPlaying }]">
        <div :class="['vinyl', { spinning: isPlaying }]">
          <!-- Grooves -->
          <div class="vinyl-grooves">
            <div v-for="i in 50" :key="i" class="groove" :style="grooveStyle(i)" />
          </div>
          <!-- Sheen -->
          <div class="vinyl-sheen" />
          <!-- Center label -->
          <div class="vinyl-label">
            <NuxtImg v-if="albumImage" :src="albumImage" :alt="albumName" loading="lazy" width="100" height="100" format="webp" class="label-art" />
            <div v-else class="label-fallback" />
            <div class="label-hole" />
          </div>
        </div>
      </div>

      <!-- Tonearm -->
      <div :class="['tonearm', { playing: isPlaying }]">
        <div class="tonearm-base" />
        <div class="tonearm-arm">
          <div class="tonearm-head" />
        </div>
      </div>
    </div>

    <!-- Song info -->
    <div v-if="isPlaying" class="mt-6 text-center space-y-1.5">
      <p class="text-base font-semibold text-foreground truncate max-w-[260px] mx-auto">{{ trackName }}</p>
      <p class="text-sm text-muted-foreground truncate max-w-[240px] mx-auto">{{ artistName }}</p>
      <p class="text-xs text-muted-foreground/60 truncate max-w-[200px] mx-auto">{{ albumName }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  isPlaying: boolean
  trackName?: string
  artistName?: string
  albumName?: string
  albumImage?: string
}>()

function grooveStyle(i: number) {
  const size = 96 - i * 1.6
  return { width: `${size}%`, height: `${size}%` }
}
</script>

<style scoped>
.player-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.player {
  position: relative;
  width: 280px;
  height: 220px;
}

@media (min-width: 640px) {
  .player {
    width: 340px;
    height: 260px;
  }
}

/* ── Album sleeve ──────────────────────────────── */
.sleeve {
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 200px;
  height: 200px;
  z-index: 3;
  transition: transform 0.8s cubic-bezier(0.4, 0, 0.2, 1);
}

.sleeve-shifted {
  transform: translateY(-50%) translateX(-40px);
}

@media (min-width: 640px) {
  .sleeve {
    width: 240px;
    height: 240px;
  }
  .sleeve-shifted {
    transform: translateY(-50%) translateX(-50px);
  }
}

.sleeve-front {
  width: 100%;
  height: 100%;
  border-radius: 6px;
  overflow: hidden;
  box-shadow:
    0 4px 24px rgba(0, 0, 0, 0.5),
    0 0 0 1px rgba(255, 255, 255, 0.06);
}

.sleeve-art {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.sleeve-placeholder {
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #0a0a0a 0%, #0d1f15 50%, #0a1a10 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}

.sleeve-placeholder-icon {
  width: 40%;
  height: 40%;
  color: rgba(0, 220, 130, 0.5);
}

/* ── Vinyl record ──────────────────────────────── */
.vinyl-area {
  position: absolute;
  left: 50px;
  top: 50%;
  transform: translateY(-50%);
  width: 200px;
  height: 200px;
  z-index: 2;
  transition: transform 0.8s cubic-bezier(0.4, 0, 0.2, 1);
}

.vinyl-out {
  transform: translateY(-50%) translateX(60px);
}

@media (min-width: 640px) {
  .vinyl-area {
    left: 60px;
    width: 240px;
    height: 240px;
  }
  .vinyl-out {
    transform: translateY(-50%) translateX(70px);
  }
}

.vinyl {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  background: radial-gradient(
    ellipse at center,
    #1a1a1a 0%,
    #0d0d0d 25%,
    #141414 35%,
    #080808 50%,
    #111 60%,
    #0a0a0a 75%,
    #050505 100%
  );
  box-shadow:
    0 0 0 2px rgba(40, 40, 40, 0.7),
    0 8px 30px rgba(0, 0, 0, 0.6);
}

.vinyl.spinning {
  animation: spin 1.8s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

/* Grooves */
.vinyl-grooves {
  position: absolute;
  inset: 0;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.groove {
  position: absolute;
  border-radius: 50%;
  border: 1px solid rgba(255, 255, 255, 0.025);
}

/* Sheen */
.vinyl-sheen {
  position: absolute;
  inset: 0;
  border-radius: 50%;
  background: linear-gradient(
    135deg,
    transparent 25%,
    rgba(255, 255, 255, 0.03) 40%,
    rgba(255, 255, 255, 0.06) 50%,
    rgba(255, 255, 255, 0.03) 60%,
    transparent 75%
  );
  pointer-events: none;
}

/* Center label */
.vinyl-label {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 36%;
  height: 36%;
  border-radius: 50%;
  overflow: hidden;
  box-shadow:
    0 0 0 2px rgba(50, 50, 50, 0.5),
    inset 0 0 8px rgba(0, 0, 0, 0.3);
}

.label-art {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 50%;
}

.label-fallback {
  width: 100%;
  height: 100%;
  background: radial-gradient(
    circle,
    var(--primary) 0%,
    oklch(0.45 0.15 160) 100%
  );
}

.label-hole {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 10%;
  height: 10%;
  border-radius: 50%;
  background: #0a0a0a;
  box-shadow: inset 0 0 3px rgba(0, 0, 0, 0.8);
}

/* ── Tonearm ───────────────────────────────────── */
.tonearm {
  position: absolute;
  top: 8px;
  right: 20px;
  z-index: 4;
  width: 16px;
  height: 16px;
}

@media (min-width: 640px) {
  .tonearm {
    right: 10px;
  }
}

.tonearm-base {
  position: absolute;
  width: 16px;
  height: 16px;
  border-radius: 50%;
  background: radial-gradient(circle, #555 0%, #333 60%, #222 100%);
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.5);
  z-index: 2;
}

.tonearm-arm {
  position: absolute;
  top: 6px;
  left: 6px;
  width: 3px;
  height: 120px;
  background: linear-gradient(180deg, #666 0%, #444 50%, #555 100%);
  border-radius: 2px;
  transform-origin: top center;
  transform: rotate(-30deg);
  transition: transform 0.8s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 1px 2px 4px rgba(0, 0, 0, 0.4);
}

.tonearm.playing .tonearm-arm {
  transform: rotate(8deg);
}

@media (min-width: 640px) {
  .tonearm-arm {
    height: 145px;
  }
}

.tonearm-head {
  position: absolute;
  bottom: -2px;
  left: 50%;
  transform: translateX(-50%);
  width: 8px;
  height: 14px;
  background: linear-gradient(180deg, #555, #333);
  border-radius: 1px 1px 2px 2px;
}

/* ── Idle state ────────────────────────────────── */
.sleeve-idle .sleeve-front {
  animation: idle-glow 4s ease-in-out infinite;
}

.sleeve-idle .sleeve-placeholder-icon {
  animation: icon-breathe 4s ease-in-out infinite;
}

@keyframes idle-glow {
  0%, 100% {
    box-shadow:
      0 4px 20px rgba(0, 0, 0, 0.5),
      0 0 0 1px rgba(255, 255, 255, 0.04);
  }
  50% {
    box-shadow:
      0 4px 20px rgba(0, 0, 0, 0.4),
      0 0 15px rgba(0, 220, 130, 0.08),
      0 0 0 1px rgba(0, 220, 130, 0.06);
  }
}

@keyframes icon-breathe {
  0%, 100% {
    opacity: 0.5;
  }
  50% {
    opacity: 0.7;
  }
}
</style>
