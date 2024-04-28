import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import BountyDetail from '../views/BountyDetail.vue'
import CreateBounty from '../views/CreateBounty.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/bounty/:id',
    name: 'BountyDetail',
    component: BountyDetail,
    props: true
  },
  {
    path: '/create',
    name: 'CreateBounty',
    component: CreateBounty
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router

