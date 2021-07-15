import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Notary from '../views/notary/notary.vue'
import NotaryIndex from '../views/notary/index.vue'
import Allocation from '../views/allocation'
import Client from '../views/client'
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/notary',
    component: Notary,
    children: [
      {
        path: '',
        name: 'NotaryIndex',
        component: NotaryIndex,
      },
      {
        path: 'allocation',
        name: 'Allocation',
        component: Allocation,
      }
    ]
  },
  {
    path: '/client',
    name: 'Client',
    component: Client,
  },
]

const router = new VueRouter({
  routes
})

export default router
