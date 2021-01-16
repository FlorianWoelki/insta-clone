<template>
  <div
    class="flex items-center pl-12 space-x-6"
    :class="{
      'text-pink-600 border-pink-600 border-r-2': isActive,
      'text-gray-600': !isActive,
    }"
  >
    <slot />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, inject } from 'vue';

export default defineComponent({
  props: {
    itemId: {
      type: Number,
      required: true,
    },
  },
  setup(props) {
    const isActive = computed(() => inject('sharedState').value.activeItem === props.itemId);

    return {
      isActive,
    };
  },
});
</script>
