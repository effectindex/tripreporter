import { createApp } from 'vue'
import { defaultConfig, plugin } from '@formkit/vue'
import App from './App.vue'
import router from './router'
import apiClient from '@/api'

const app = createApp(App)
    .use(router)
    .use(plugin, defaultConfig({theme: 'genesis'}))
    .provide('axios', apiClient)

app.mount('#app')
