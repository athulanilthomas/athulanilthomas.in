<template>
  <div class="ui-aurora-container">
    <BackgroundAurora
      :color-stops="['#00DC82', '#00DC82', '#020420', '#00DC82']"
      :amplitude="1"
      :blend="0.5"
      :speed="1.0"
      :intensity="1.0"
      class="w-full h-full"
    />
    <slot />
  </div>
</template>

<style scoped>
  .ui-aurora-container {
    width: 100%;
    height: 100vh;
    position: relative;
    overflow: hidden;
  }
</style>