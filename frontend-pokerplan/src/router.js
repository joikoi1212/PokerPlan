import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from './dashboard.vue';
import  Main  from './App.vue';

const routes = [
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
  },
  {
    path: '/',
    name: 'Home',
    component: Main,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;