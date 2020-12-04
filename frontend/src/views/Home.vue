<template>
  <div class="grid h-screen grid-cols-4">
    <div class="py-8 bg-gray-100">
      <!-- PROFILE HEADER -->
      <div class="px-8 space-y-3">
        <div class="flex items-center justify-center">
          <div class="inline-block p-1 rounded-full bg-gradient-to-tr from-yellow-400 to-pink-600">
            <a class="block p-1 bg-white rounded-full" href="#">
              <img class="w-20 h-20 rounded-full" src="https://placekitten.com/200/200" alt="pb-avatar">
            </a>
          </div>
        </div>
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
    <div class="col-span-3">
      <h1>test</h1>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Icon from '@/components/ui/Icon.vue';

export default defineComponent({
  components: {
    Icon,
  },
});
</script>
