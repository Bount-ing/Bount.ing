import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createVuetify } from 'vuetify';
import '@mdi/font/css/materialdesignicons.css';
import 'vuetify/styles';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import  createGtag  from 'vue-gtag-next';
import type GtagPluginOptions from 'vue-gtag-next'


import App from './App.vue'
import router from './router'
import './assets/tailwind.css'
import './assets/datepicker.css'

const app = createApp(App)
const vuetify = createVuetify()

app.use(createPinia())
app.use(vuetify)
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
