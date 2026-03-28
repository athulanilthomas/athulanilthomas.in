<template>
  <div class="relative flex gap-2 w-full max-w-6xl h-[100vh] sm:h-[85vh] sm:max-h-[900px] sm:bg-glass sm:border border-glass-border rounded-2xl sm:shadow-glass overflow-hidden p-2 sm:p-4">
    <LayoutSidebar />
    <div class="flex-1 flex flex-col min-w-0 gap-2 max-h-full">
      <div
        class="bg-glass-mobile sm:bg-glass border border-glass-border rounded-2xl shadow-glass backdrop-blur-[60px] sm:backdrop-blur-[40px] backdrop-saturate-150 flex-1 flex flex-col min-w-0 overflow-auto"
      >
        <div class="flex-1 p-4 sm:p-6 origin-center overflow-y-auto">
          <slot />
        </div>
      </div>
    </div>
  </div>
</template>

