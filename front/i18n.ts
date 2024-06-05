// src/i18n.js
import { createI18n } from 'vue-i18n';


// Import language files en, fr, es, eu, ca
import en from './src/locales/en.json';
import fr from './src/locales/fr.json';
import es from './src/locales/es.json';
import eu from './src/locales/eu.json';
import ca from './src/locales/ca.json';
import pt from './src/locales/pt.json';




const messages = {
  en,
  fr,
  es,
  eu,
  ca,
  pt,
};

const i18n = createI18n({
  legacy: false, // Use Composition API mode
  locale: 'en', // set locale
  fallbackLocale: 'en', // set fallback locale
  messages, // set locale messages
});

export default i18n;
