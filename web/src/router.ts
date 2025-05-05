import { createRouter, createWebHistory } from 'vue-router'
import HomeView from './views/Home.vue'
import ManageBooks from './views/ManageBooks.vue'
import AddBooks from './views/AddBook.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/manageBooks',
      name: 'manageBooks',
      component: ManageBooks,
    },
    {
      path: '/AddBooks',
      name: 'addBooks',
      component: AddBooks,
    },
  ],
})

export default router
