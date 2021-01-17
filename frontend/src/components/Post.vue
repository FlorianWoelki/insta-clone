<template>
  <div class="space-y-2">
    <img class="w-full cursor-pointer rounded-2xl" src="https://placekitten.com/300/300" alt="feed-image">
    <div class="flex items-center justify-between">
      <div class="flex items-center space-x-2 cursor-pointer">
        <Avatar :imageSize="6" />
        <p class="text-sm font-bold text-gray-800">Dom_Hill</p>
      </div>
      <div class="flex items-center space-x-3 text-gray-700">
        <div class="flex items-center space-x-1 cursor-pointer">
          <Icon name="heart" class="w-5 h-5" />
          <span class="text-sm">5.2k</span>
        </div>
        <div class="flex items-center space-x-1 cursor-pointer">
          <Icon name="chat" class="w-5 h-5" />
          <span class="text-sm">38</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Icon from './ui/Icon.vue';
import Avatar from './Avatar.vue';

export default defineComponent({
  components: {
    Icon,
    Avatar,
  },
});
</script>
