<template>
  <DefaultLayout>
    <!-- HEADER SECTION -->
    <div class="flex items-center justify-between">
      <!-- SEARCH INPUT FIELD -->
      <Search />

      <!-- RIGHT SECTION -->
      <div class="flex items-center space-x-6">
        <a href="#" class="relative text-gray-600">
          <Icon name="bell" class="w-5 h-5" />
          <div class="absolute top-0 right-0 w-1 h-1 -mt-1 -mr-1 bg-red-600 rounded-full" />
        </a>
        <a href="#" class="relative text-gray-600">
          <Icon name="mail" class="w-5 h-5" />
          <div class="absolute top-0 right-0 w-1 h-1 -mt-1 -mr-1 bg-red-600 rounded-full" />
        </a>
        <Button>
          <Icon name="plus-circle" class="w-6 h-6" />
          <span>Add photo</span>
        </Button>
      </div>
    </div>

    <!-- STORIES SECTION -->
    <div class="mt-10">
      <div class="flex items-center justify-between">
        <h1 class="text-3xl font-bold text-gray-800">Stories</h1>
        <div class="flex items-center space-x-3 cursor-pointer">
          <Icon name="play" class="w-6 h-6 text-gray-300" />
          <p class="text-gray-800">Watch all</p>
        </div>
      </div>
      <div class="flex items-center mt-6 space-x-4">
        <Avatar noActiveStories>
          <div class="absolute inset-0 flex items-center justify-center w-full h-full bg-blue-600 bg-opacity-50 rounded-full">
            <Icon name="plus" class="w-6 h-6 text-white" />
          </div>
        </Avatar>
        <Avatar />
        <Avatar />
        <Avatar noActiveStories />
      </div>
    </div>

    <!-- FEED SECTION -->
    <div class="mt-10">
      <h1 class="text-3xl font-bold text-gray-800">Feed</h1>
      <div class="grid grid-cols-3 gap-6 mt-6">
        <Post />
        <Post />
        <Post />
        <Post />
        <Post />
      </div>
    </div>
  </DefaultLayout>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import DefaultLayout from '@/layout/default.vue';
import Avatar from '@/components/Avatar.vue';
import Post from '@/components/Post.vue';
import { Icon, Button, Search } from '@/components/ui';

export default defineComponent({
  components: {
    DefaultLayout,
    Icon,
    Button,
    Avatar,
    Search,
    Post,
  },
});
</script>
