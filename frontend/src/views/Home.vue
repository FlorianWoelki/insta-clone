<template>
  <div class="grid h-screen grid-cols-4">
    <div class="py-8 bg-gray-100">
      <!-- PROFILE HEADER -->
      <div class="px-8 space-y-3">
        <Avatar />
        <div class="space-y-0">
          <h2 class="text-lg font-bold text-center text-gray-800">Kate Lingard</h2>
          <p class="text-sm text-center text-gray-400">@katy69</p>
        </div>
      </div>

      <!-- PROFILE STATS -->
      <div class="grid justify-between grid-cols-3 gap-6 px-8 my-8 text-sm">
        <div class="text-center">
          <p class="font-bold text-gray-800">46</p>
          <p class="text-gray-400">Posts</p>
        </div>
        <div class="text-center">
          <p class="font-bold text-gray-800">2.8k</p>
          <p class="text-gray-400">Followers</p>
        </div>
        <div class="text-center">
          <p class="font-bold text-gray-800">526</p>
          <p class="text-gray-400">Following</p>
        </div>
      </div>

      <!-- SIDEBAR ITEMS -->
      <div class="space-y-5">
        <div class="flex items-center pl-12 space-x-6 text-pink-600 border-r-2 border-pink-600">
          <Icon name="template" class="w-5 h-5" />
          <p class="font-semibold">Feed</p>
        </div>
        <div class="flex items-center pl-12 space-x-6 text-gray-600">
          <Icon name="search" class="w-5 h-5" />
          <p>Explore</p>
        </div>
        <div class="flex items-center pl-12 space-x-6 text-gray-600">
          <div class="flex items-center space-x-6">
            <Icon name="bell" class="w-5 h-5" />
            <p>Notifications</p>
          </div>
          <span class="text-sm text-gray-400">2</span>
        </div>
        <div class="flex items-center pl-12 space-x-6 text-gray-600">
          <Icon name="mail" class="w-5 h-5" />
          <p>Messages</p>
          <span class="text-sm text-gray-400">8</span>
        </div>
        <div class="flex items-center pl-12 space-x-6 text-gray-600">
          <Icon name="paper-airplane" class="w-5 h-5" />
          <p>Direct</p>
        </div>
        <div class="flex items-center pl-12 space-x-6 text-gray-600">
          <Icon name="chart-bar" class="w-5 h-5" />
          <p>Stats</p>
        </div>
        <div class="flex items-center pl-12 space-x-6 text-gray-600">
          <Icon name="cog" class="w-5 h-5" />
          <p>Settings</p>
        </div>

        <div class="flex items-center">
          <div class="flex-grow bg-gray-200 border"></div>
          <div class="p-1 -mr-2 bg-white rounded-full">
            <Icon name="chevron-left" class="w-5 h-5" />
          </div>
        </div>

        <div class="flex items-center pl-12 space-x-6 text-gray-600">
          <Icon name="logout" class="w-5 h-5" />
          <p>Logout</p>
        </div>
      </div>
    </div>

    <div class="container max-w-3xl col-span-3 mx-auto my-12">
      <!-- HEADER SECTION -->
      <div class="flex items-center justify-between">
        <!-- SEARCH INPUT FIELD -->
        <div class="flex items-center w-56 p-2 space-x-2 rounded-lg bg-gray-50">
          <Icon name="search" class="w-6 h-6 my-auto text-gray-300" />
          <input
            placeholder="Search"
            type="text"
            class="w-full pr-2 text-sm text-gray-700 placeholder-gray-300 bg-gray-50 focus:outline-none"
          >
        </div>

        <!-- RIGHT SECTION -->
        <div class="flex items-center space-x-6">
          <div class="relative text-gray-600">
            <Icon name="bell" class="w-5 h-5" />
            <div class="absolute top-0 right-0 w-1 h-1 -mt-1 -mr-1 bg-red-600 rounded-full" />
          </div>
          <div class="relative text-gray-600">
            <Icon name="mail" class="w-5 h-5" />
            <div class="absolute top-0 right-0 w-1 h-1 -mt-1 -mr-1 bg-red-600 rounded-full" />
          </div>
          <button
            type="button"
            class="flex items-center px-4 py-2 space-x-1 font-medium text-white rounded-lg bg-gradient-to-r from-pink-600 to-yellow-500"
          >
            <Icon name="plus-circle" class="w-6 h-6" />
            <span>Add photo</span>
          </button>
        </div>
      </div>

      <!-- STORIES SECTION -->
      <div class="mt-10">
        <h1 class="text-3xl font-bold text-gray-800">Stories</h1>
        <div class="flex items-center mt-6 space-x-4">
          <Avatar noActiveStories>
            <div class="absolute inset-0 flex items-center justify-center w-full h-full bg-blue-600 bg-opacity-50 rounded-full">
              <Icon name="plus" class="w-6 h-6 text-white" />
            </div>
          </Avatar>
          <Avatar />
          <Avatar />
          <Avatar />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Icon from '@/components/ui/Icon.vue';
import Avatar from '@/components/Avatar.vue';

export default defineComponent({
  components: {
    Icon,
    Avatar,
  },
});
</script>
