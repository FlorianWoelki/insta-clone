<template>
  <div
    class="flex items-center pl-12 space-x-6 cursor-pointer"
    :class="{
      'text-pink-600 border-pink-600 border-r-2 font-semibold': isActive,
      'text-gray-600': !isActive,
    }"
    @click="toggle"
  >
    <slot />
  </div>
</template>

<script lang="ts">
import {
  defineComponent, computed, inject, ref,
} from 'vue';
import { SidebarSharedState } from '../../models/SidebarSharedState';

export default defineComponent({
  props: {
    itemId: {
      type: Number,
      required: true,
    },
  },
  setup(props) {
    const sharedState = ref(inject('sharedState') as SidebarSharedState);
    const isActive = computed(() => sharedState.value.activeItem === props.itemId);

    const toggle = () => {
      sharedState.value.activeItem = props.itemId;
    };

    return {
      isActive,
      toggle,
    };
  },
});
</script>
