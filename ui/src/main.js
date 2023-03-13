import '@formkit/themes/genesis'
import '@formkit/pro/genesis'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { defaultConfig, plugin } from '@formkit/vue'
import { createProPlugin, inputs } from '@formkit/pro'
import App from './App.vue'
import router from './router'
import apiClient from '@/api'

const formkitPro = createProPlugin(process.env.FORMKIT_API_KEY, inputs)
const pinia = createPinia()
const app = createApp(App)
    .use(pinia)
    .use(router)
    .use(plugin, defaultConfig({ plugins: [formkitPro()] }))
    .provide('axios', apiClient)

app.mount('#app')
