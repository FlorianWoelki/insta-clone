<template>
  <component :is="iconData" />
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue';

export default defineComponent({
  props: {
    name: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    /* eslint-disable-next-line */
    const iconData = computed(() => require(`@/assets/icons/${props.name}.svg?raw`).default);
    return { iconData };
  },
});
</script>
