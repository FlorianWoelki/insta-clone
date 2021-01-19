import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import Home from '../views/Home.vue';
import Direct from '../views/Direct.vue';
import Explore from '../views/Explore.vue';
import Messages from '../views/Messages.vue';
import Notifications from '../views/Notifications.vue';
import Settings from '../views/Settings.vue';
import Stats from '../views/Stats.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/direct',
    name: 'Direct',
    component: Direct,
  },
  {
    path: '/explore',
    name: 'Explore',
    component: Explore,
  },
  {
    path: '/messages',
    name: 'Message',
    component: Messages,
  },
  {
    path: '/notifications',
    name: 'Notifications',
    component: Notifications,
  },
  {
    path: '/settings',
    name: 'Settings',
    component: Settings,
  },
  {
    path: '/stats',
    name: 'Stats',
    component: Stats,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
