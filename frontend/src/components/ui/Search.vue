<template>
  <div class="relative w-full max-w-xs 2xl:max-w-lg">
    <Icon name="search" class="absolute inset-y-0 left-0 w-6 h-6 my-auto ml-2 text-gray-300" />
    <input
      placeholder="Search"
      type="text"
      class="w-full py-2 pl-10 pr-4 text-sm text-gray-700 placeholder-gray-300 transition duration-300 ease-in-out bg-gray-100 rounded-lg focus:outline-none focus:shadow-md"
    >
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Icon from './Icon.vue';

export default defineComponent({
  components: {
    Icon,
  },
});
</script>
