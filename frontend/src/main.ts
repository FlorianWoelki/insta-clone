import { createApp } from 'vue';
import autoImport from '@flowy/vue-auto-import';
import App from './App.vue';
import router from './router';
import store from './store';
import '@/assets/styles/tailwind.css';

createApp(App)
  .use(autoImport)
  .use(store)
  .use(router)
  .mount('#app');
