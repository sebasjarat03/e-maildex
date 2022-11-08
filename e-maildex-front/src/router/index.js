import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import IndexerView from '../views/IndexerView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    redirect : "/index",
    component: HomeView
  },
  {
    path : '/index',
    name : 'Indexer',
    component : IndexerView

  }
  
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
