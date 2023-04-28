// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    name: 'home',
    path: '/',
    alias: ['/index.html'],
    component: HomeView
  },
  {
    name: 'account',
    path: '/account',
    component: () => import(/* webpackChunkName: "account" */ '../views/AccountView.vue')
  },
  {
    name: 'signup',
    path: '/signup',
    component: () => import(/* webpackChunkName: "signup" */ '../views/SignupView.vue')
  },
  {
    name: 'login',
    path: '/login',
    component: () => import(/* webpackChunkName: "login" */ '../views/LoginView.vue')
  },
  {
    name: 'create',
    path: '/create',
    component: () => import(/* webpackChunkName: "create" */ '../views/CreateView.vue')
  },
  {
    name: 'report',
    path: '/report/:id',
    component: () => import(/* webpackChunkName: "report" */ '../views/ReportView.vue')
  },
  {
    name: 'profile',
    path: '/profile/:id',
    component: () => import(/* webpackChunkName: "profile" */ '../views/ProfileView.vue')
  },
  {
    name: 'NotFound',
    path: '/:pathMatch(.*)*',
    component: () => import('@/views/NotFound.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
