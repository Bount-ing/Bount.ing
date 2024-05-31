import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import  createGtag  from 'vue-gtag-next';
import type GtagPluginOptions from 'vue-gtag-next'


import App from './App.vue'
import router from './router'
import './assets/tailwind.css'


const app = createApp(App)

app.use(createPinia())
app.use(router)


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
