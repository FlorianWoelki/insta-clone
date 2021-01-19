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
        <SidebarItem :itemId="0" @click="$router.push('/')">
          <Icon name="template" class="w-5 h-5" />
          <p>Feed</p>
        </SidebarItem>
        <SidebarItem :itemId="1" @click="$router.push('explore')">
          <Icon name="search" class="w-5 h-5" />
          <p>Explore</p>
        </SidebarItem>
        <SidebarItem :itemId="2" @click="$router.push('notifications')">
          <div class="flex items-center space-x-6">
            <Icon name="bell" class="w-5 h-5" />
            <p>Notifications</p>
          </div>
          <span class="text-sm text-gray-400">2</span>
        </SidebarItem>
        <SidebarItem :itemId="3" @click="$router.push('messages')">
          <Icon name="mail" class="w-5 h-5" />
          <p>Messages</p>
          <span class="text-sm text-gray-400">8</span>
        </SidebarItem>
        <SidebarItem :itemId="4" @click="$router.push('direct')">
          <Icon name="paper-airplane" class="w-5 h-5" />
          <p>Direct</p>
        </SidebarItem>
        <SidebarItem :itemId="5" @click="$router.push('stats')">
          <Icon name="chart-bar" class="w-5 h-5" />
          <p>Stats</p>
        </SidebarItem>
        <SidebarItem :itemId="6" @click="$router.push('settings')">
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
      <slot></slot>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { Sidebar, SidebarItem } from '@/components/sidebar';
import Avatar from '@/components/Avatar.vue';
import { Icon } from '@/components/ui';

export default defineComponent({
  components: {
    Sidebar,
    SidebarItem,
    Icon,
    Avatar,
  },
});
</script>
