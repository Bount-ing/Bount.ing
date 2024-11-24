import './assets/main.css'
import './assets/tailwind.css'

import '@fontsource/roboto-mono'; 
import '@fontsource/philosopher'; 
import '@fontsource/vt323'; 

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import i18n from '../i18n';
import  createGtag  from 'vue-gtag-next';
import type GtagPluginOptions from 'vue-gtag-next'


import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(i18n);

const gtagOptions: GtagPluginOptions = {
    property: {
      id: import.meta.env.VITE_GTAG
    },
    // Additional configuration can go here
    appName: 'Bount.ing',
    pageTrackerScreenviewEnabled: true,
  };
app.use(createGtag, gtagOptions);

app.mount('#app')
