<template>
  <div class="grid h-screen grid-cols-4 xl:grid-cols-7 2xl:grid-cols-5">
    <Sidebar>
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
      <template #sidebarItems>
        <SidebarItem :itemId="0">
          <Icon name="template" class="w-5 h-5" />
          <p>Feed</p>
        </SidebarItem>
        <SidebarItem :itemId="1">
          <Icon name="search" class="w-5 h-5" />
          <p>Explore</p>
        </SidebarItem>
        <SidebarItem :itemId="2">
          <div class="flex items-center space-x-6">
            <Icon name="bell" class="w-5 h-5" />
            <p>Notifications</p>
          </div>
          <span class="text-sm text-gray-400">2</span>
        </SidebarItem>
        <SidebarItem :itemId="3">
          <Icon name="mail" class="w-5 h-5" />
          <p>Messages</p>
          <span class="text-sm text-gray-400">8</span>
        </SidebarItem>
        <SidebarItem :itemId="4">
          <Icon name="paper-airplane" class="w-5 h-5" />
          <p>Direct</p>
        </SidebarItem>
        <SidebarItem :itemId="5">
          <Icon name="chart-bar" class="w-5 h-5" />
          <p>Stats</p>
        </SidebarItem>
        <SidebarItem :itemId="6">
          <Icon name="cog" class="w-5 h-5" />
          <p>Settings</p>
        </SidebarItem>

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
      </template>
    </Sidebar>

    <div class="container col-span-3 px-16 mx-auto my-12 xl:col-span-5 2xl:col-span-4">
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
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Avatar from '@/components/Avatar.vue';
import Sidebar from '@/components/sidebar/Sidebar.vue';
import SidebarItem from '@/components/sidebar/SidebarItem.vue';
import Icon from '@/components/ui/Icon.vue';
import Button from '@/components/ui/Button.vue';
import Search from '@/components/ui/Search.vue';

export default defineComponent({
  components: {
    Icon,
    Button,
    Avatar,
    Sidebar,
    SidebarItem,
    Search,
  },
});
</script>
