import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { defaultConfig, plugin } from '@formkit/vue'

createApp(App)
    .use(router)
    .use(plugin, defaultConfig({theme: 'genesis'}))
    .mount('#app')
