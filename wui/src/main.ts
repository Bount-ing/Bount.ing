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
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import '@fortawesome/fontawesome-free/css/all.css';




import App from './App.vue'
import router from './router'

const app = createApp(App)

// Pinia Storage
const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);
export default pinia;
app.use(pinia)
app.use(router)
app.use(i18n);

// Google Analytics
const gtagOptions: GtagPluginOptions = {
    property: {
      id: import.meta.env.VITE_GTAG
    },
    // Additional configuration can go here
    appName: 'Bount.ing',
    pageTrackerScreenviewEnabled: true,
  };
app.use(createGtag, gtagOptions);

// Font Awesome Icons
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faHome, faChartBar, faDollarSign, faInfoCircle, faQuestionCircle, faEnvelope, faListCheck } from '@fortawesome/free-solid-svg-icons';

library.add(faHome, faChartBar, faDollarSign, faInfoCircle, faQuestionCircle, faEnvelope, faListCheck);

app.component('font-awesome-icon', FontAwesomeIcon);

// Notifications
import Notifications from '@kyvg/vue3-notification'
app.use(Notifications);

app.mount('#app')
