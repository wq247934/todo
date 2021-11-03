import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Todos from '../views/Todos.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    children:[
        {
            path:'/:tag',
            name:'Todos',
            component:Todos
        }
    ]
  },
]

const router = new VueRouter({
//   mode: 'hash',
//   base: process.env.BASE_URL,
  routes
})

export default router
