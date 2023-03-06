import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { defaultConfig, plugin } from '@formkit/vue'
import App from './App.vue'
import router from './router'
import apiClient from '@/api'

const pinia = createPinia()
const app = createApp(App)
    .use(pinia)
    .use(router)
    .use(plugin, defaultConfig({theme: 'genesis'}))
    .provide('axios', apiClient)

app.mount('#app')
