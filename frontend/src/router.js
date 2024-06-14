import { createRouter, createWebHistory } from 'vue-router'
import HomeView from './views/HomeView.vue'
import GenerateView from './views/GenerateView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/generate',
      name: 'generate',
      component: GenerateView
    }
  ]
})

export default router
