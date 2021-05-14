import { createRouter, createWebHistory } from 'vue-router'
import Main from '../views/Main.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Main,
    children: [
      {
        path: '',
        name: 'NewsFeed',
        component: () => import('../views/NewsFeed.vue'),
      },
      {
        path: 'notif',
        name: 'Notification',
        component: () => import('../views/Notification.vue'),
      },
      {
        path: 'profile',
        name: 'Profile',
        component: () => import('../views/Profile.vue'),
      },
      {
        path: 'photo',
        name: 'Photo',
        component: () => import('../views/Photos.vue'),
      },
    ],
    meta: {
      needsAuth: true,
    },
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
  },
  {
    path: '/signup',
    name: 'SignUp',
    component: () => import('../views/SignUp.vue'),
  },
  { path: '/:pathMatch(.*)*', name: 'not-found', redirect: '/' },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
