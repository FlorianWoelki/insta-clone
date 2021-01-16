<template>
  <div class="py-8 bg-gray-100 xl:col-span-2 2xl:col-span-1">
    <slot />

    <div class="space-y-5">
      <slot name="sidebarItems" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, provide } from 'vue';

export default defineComponent({
  setup() {
    const sharedState = ref({
      activeItem: 0,
    });
    provide('sharedState', sharedState);

    return {
      sharedState,
    };
  },
});
</script>
