<template>
    <div :class="{'relative': !inline}">
      <button @click="toggleDropdown" class="px-3 py-2 rounded-md text-sm font-medium ">
        <span :class="'flag fi fi-squared fi-' + currentFlag"></span>
      </button>
      <div v-if="isDropdownOpen" :class="dropdownClass">
        <a @click="setLanguage('en')" class="block px-4 text-lg py-2 text-sm text-primary hover:bg-secondary-dark">
          <span class="flag fi fi-squared fi-gb"></span> {{ $t('language.english') }}
        </a>
        <a @click="setLanguage('es')" class="block px-4 text-lg py-2 text-sm text-primary hover:bg-secondary-dark">
          <span class="flag fi fi-squared fi-es"></span> {{ $t('language.spanish') }}
        </a>
        <a @click="setLanguage('fr')" class="block px-4 text-lg py-2 text-sm text-primary hover:bg-secondary-dark">
          <span class="flag fi fi-squared fi-fr"></span> {{ $t('language.french') }}
        </a>
        <a @click="setLanguage('ca')" class="block px-4 text-lg py-2 text-sm text-primary hover:bg-secondary-dark">
          <span class="flag fi fi-squared fi-es-ct"></span> {{ $t('language.catalan') }}
        </a>
        <a @click="setLanguage('pt')" class="block px-4 text-lg py-2 text-sm text-primary hover:bg-secondary-dark">
          <span class="flag fi fi-squared fi-pt"></span> {{ $t('language.portuguese') }}
        </a>
        <a @click="setLanguage('eu')" class="block px-4 text-lg py-2 text-sm text-primary hover:bg-secondary-dark">
          <span class="flag fi fi-squared fi-es-pv"></span> {{ $t('language.euskera') }}
        </a>
        <!-- Add more languages as needed -->
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, computed, watch } from 'vue';
  import { useI18n } from 'vue-i18n';
  
  const props = defineProps({
    inline: {
      type: Boolean,
      default: false
    }
  });
  
  const isDropdownOpen = ref(false);
  
  const toggleDropdown = () => {
    isDropdownOpen.value = !isDropdownOpen.value;
  };
  
  const { locale } = useI18n();
  
  const setLanguage = (language: string) => {
    locale.value = language; // Change the language
    isDropdownOpen.value = false; // Close the dropdown after selection
  };
  
  const dropdownClass = computed(() => {
    return props.inline ? 'text-primary mt-2 w-full bg-secondary border border-primary-dark rounded-md shadow-lg' : 'absolute right-0 mt-2 w-48 border bg-secondary rounded-md shadow-lg text-primary border-primary-dark';
  });
  
  const currentFlag = ref('eu'); // Default flag
  
  // Update the current flag based on the selected language
  watch(locale, (newLocale) => {
    switch (newLocale) {
      case 'en':
        currentFlag.value = 'gb';
        break;
      case 'es':
        currentFlag.value = 'es';
        break;
      case 'fr':
        currentFlag.value = 'fr';
        break;
      case 'ca':
        currentFlag.value = 'es-ct';
        break;
      case 'pt':
        currentFlag.value = 'pt';
        break;
      case 'eu':
        currentFlag.value = 'es-pv';
        break;
      // Add more cases for other languages
      default:
        currentFlag.value = 'eu'; // Default to English flag
        break;
    }
  }, {immediate: true });
  </script>
  
  <style scoped>
  @import url('https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/7.2.1/css/flag-icons.min.css');
  
  </style>
  