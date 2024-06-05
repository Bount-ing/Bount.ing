// src/i18n.js
import { createI18n } from 'vue-i18n';


// Import language files
import en from './src/locales/en.json';
import fr from './src/locales/fr.json';

const messages = {
  en,
  fr
};

const i18n = createI18n({
  legacy: false, // Use Composition API mode
  locale: 'en', // set locale
  fallbackLocale: 'en', // set fallback locale
  messages, // set locale messages
});

export default i18n;
