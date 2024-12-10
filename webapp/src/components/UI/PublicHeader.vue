<template>
  <nav class="border-b border-primary bg-secondary text-primary-light shadow-lg fixed top-0 left-0 w-full z-50">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <!-- Logo and navigation links -->
        <div class="flex items-center">
          <router-link to="/" class="flex items-center">
            <img src="/bount.ing.logo.png" class="h-12 w-12 rounded-3xl" alt="Bount.ing Logo" />
            <span class="ml-3 text-3xl font-bold text-primary tracking-tight">Bount.ing</span>
          </router-link>

        </div>

      <!-- Right Section: Language and Authentication -->
      <div class="flex items-center space-x-2">
        <!-- Language Dropdown Component -->
        <LanguageDropdown :inline="false" />
        <router-link v-if="!isLoggedIn" to="/signin" class="px-2 py-2 rounded-md text-sm font-medium border text-info-light border-info hover:border-info-light">{{ $t('account.signin') }}</router-link>
        <router-link v-if="!isLoggedIn" to="/signup" class="px-2 py-2 rounded-md text-sm font-medium border text-success-light border-success hover:border-success-light">{{ $t('account.signup') }}</router-link>
        <button v-else @click="logout" class="px-2 py-2 rounded-md text-sm font-medium border text-error-light border-error hover:border-error-light">{{ $t('account.logout') }}</button>
      </div>
      
        <div class="hidden md:block" v-if="isLoggedIn">
			<ConnectStripe />
        </div>

        <!-- Mobile menu button -->
        <div class="md:hidden bg-gray-700 p-2 rounded-lg">
          <button @click="isOpen = !isOpen" class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-600">
            <span class="sr-only">{{ $t('Open main menu') }}</span>
            <svg v-if="!isOpen" class="block h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-16 6h16"/>
            </svg>
            <svg v-else class="block h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Mobile Menu -->
    <div v-if="isOpen" class="md:hidden">
      <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
        <router-link v-if="!isLoggedIn" to="/signin" class="block px-3 py-2 rounded-md text-sm font-medium border text-info-light border-info hover:border-info-light">{{ $t('Signin') }}</router-link>
        <router-link v-if="!isLoggedIn" to="/signup" class="block px-3 py-2 rounded-md text-sm font-medium border text-success-light border-success hover:border-success-light">{{ $t('Signup') }}</router-link>
        <button v-else @click="logout" class="block px-3 py-2 rounded-md text-sm font-medium border text-error-light border-error hover:border-error-light">{{ $t('Logout') }}</button>
        <!-- Language Dropdown Component for mobile -->
        <LanguageDropdown :inline="true" />
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useUserStore } from '../../stores/user';
import ConnectStripe from '../StripeConnect.vue'
import LanguageDropdown from '../LanguageDropdown.vue';

const user = useUserStore();
const { isLoggedIn } = storeToRefs(user);
const isOpen = ref(false);

const logout = () => {
  user.logout();
};
</script>

<style scoped>
/* Add your styles here */
</style>
