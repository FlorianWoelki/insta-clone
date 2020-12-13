<template>
  <div class="flex items-center justify-center">
    <div
      class="inline-block p-0.5 rounded-full bg-gradient-to-tr from-yellow-400 to-pink-600"
      :class="
        noActiveStories ? '' : ''
      "
    >
      <a class="relative block p-1 bg-white rounded-full" href="#">
        <img
          :class="`w-${imageSize} h-${imageSize} rounded-full`"
          src="https://placekitten.com/200/200"
          alt="pb-avatar"
        >
        <slot></slot>
      </a>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    noActiveStories: Boolean,
    imageSize: {
      type: Number,
      default: 16,
    },
  },
});
</script>
