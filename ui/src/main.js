// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import '@formkit/themes/genesis'
import '@formkit/pro/genesis'
import '@formkit/addons/css/multistep'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { defaultConfig, plugin } from '@formkit/vue'
import { createProPlugin, inputs } from '@formkit/pro'
import { createMultiStepPlugin } from '@formkit/addons'
import App from './App.vue'
import router from './router'
import apiClient from '@/api'

const formkitPro = createProPlugin(process.env.VUE_APP_FORMKIT_API_KEY, inputs)
const pinia = createPinia()
const app = createApp(App)
    .use(pinia)
    .use(router)
    .use(plugin, defaultConfig({ theme: 'genesis', plugins: [formkitPro, createMultiStepPlugin()] }))
    .provide('axios', apiClient)
    .provide('router', router)

app.mount('#app')
