<template>
  <nav class="border-b border-primary bg-secondary text-primary-light shadow-lg fixed top-0 left-0 w-full z-50">
    <div class="px-4 mx-auto flex items-center justify-between h-16">
      
      <!-- Left Section: Logo and Title -->
      <div class="flex items-center">
        <router-link to="/" class="flex items-center">
          <img src="/bount.ing.logo.png" class="h-16 w-16 rounded-3xl" alt="Bount.ing Logo" />
          <span class="hidden sm:flex  ml-3 text-4xl font-bold text-primary tracking-tight">Bount.ing</span>
        </router-link>
      </div>

      <!-- Center Section: Navigation Icons -->
      <div class="hidden lg:flex justify-center space-x-4 text-sm">
        <router-link  v-if="isLoggedIn" to="/" class="flex flex-col items-center px-3 py-1 rounded-md text-sm font-medium hover:bg-secondary-dark">
          <font-awesome-icon icon="home" class="text-primary-light text-2xl p-1" />
          <span>{{ $t('navigation.home') }}</span>
        </router-link>
        <router-link v-if="isLoggedIn" to="/profile" class="flex flex-col items-center px-3 py-1 rounded-md text-sm font-medium hover:bg-secondary-dark">
          <font-awesome-icon icon="chart-bar" class="text-primary-light text-2xl p-1" />
          <span>{{ $t('navigation.dashboard') }}</span>
        </router-link>
        <router-link v-if="isLoggedIn" to="/bounties" class="flex flex-col items-center px-3 py-1 rounded-md text-sm font-medium hover:bg-secondary-dark">
          <font-awesome-icon icon="list-check" class="text-primary-light text-2xl p-1" />
          <span>{{ $t('navigation.bounties') }}</span>
        </router-link>
        <router-link to="/pricing" class="flex flex-col items-center px-3 py-1 rounded-md text-sm font-medium hover:bg-secondary-dark">
          <font-awesome-icon icon="dollar-sign" class="text-primary-light text-2xl p-1" />
          <span>{{ $t('navigation.pricing') }}</span>
        </router-link>
        <router-link to="/about" class="flex flex-col items-center px-3 py-1 rounded-md text-sm font-medium hover:bg-secondary-dark">
          <font-awesome-icon icon="info-circle" class="text-primary-light text-2xl p-1" />
          <span>{{ $t('navigation.about') }}</span>
        </router-link>
        <router-link to="/help" class="flex flex-col items-center px-3 py-1 rounded-md text-sm font-medium hover:bg-secondary-dark">
          <font-awesome-icon icon="question-circle" class="text-primary-light text-2xl p-1" />
          <span>{{ $t('navigation.faq') }}</span>
        </router-link>
        <router-link to="/contact" class="flex flex-col items-center px-3 py-1 rounded-md text-sm font-medium hover:bg-secondary-dark">
          <font-awesome-icon icon="envelope" class="text-primary-light text-2xl p-1" />
          <span>{{ $t('navigation.contact') }}</span>
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

      <!-- Mobile menu button -->
      <div class="lg:hidden text-primary p-2 rounded-lg">
        <button @click="isOpen = !isOpen" class="inline-flex items-center justify-center p-2 rounded-md hover:text-primary-light hover:bg-secondary-light">
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

    <!-- Mobile Menu -->
    <div v-if="isOpen" class="lg:hidden absolute top-16 left-0 w-full bg-secondary z-40 border-primary-dark border-b-2">
      <div class="flex flex-col space-y-4 px-6 py-4">
  <div class="flex flex-col space-y-4 px-6 py-4">
    <router-link v-if="isLoggedIn" to="/" @click="isOpen = !isOpen" class="flex items-center px-3 py-2 rounded-md text-sm font-medium hover:bg-secondary-dark">
      <font-awesome-icon icon="home" class="text-primary-light text-2xl py-1 px-2" />
      <span class="text-xl">{{ $t('Home') }}</span>
    </router-link>
    <router-link  v-if="isLoggedIn" to="/profile" class="flex items-center px-3 py-2 rounded-md text-sm font-medium hover:bg-secondary-dark">
      <font-awesome-icon icon="chart-bar" class="text-primary-light text-2xl py-1 px-2" />
      <span class="text-xl">{{ $t('Dashboard') }}</span>
    </router-link>
    <router-link to="/pricing" class="flex items-center px-3 py-2 rounded-md text-sm font-medium hover:bg-secondary-dark">
      <font-awesome-icon icon="dollar-sign" class="text-primary-light text-2xl py-1 px-3" />
      <span class="text-xl">{{ $t('Pricing') }}</span>
    </router-link>
    <router-link to="/about" class="flex items-center px-3 py-2 rounded-md text-sm font-medium hover:bg-secondary-dark">
      <font-awesome-icon icon="info-circle" class="text-primary-light text-2xl py-1 px-2" />
      <span class="text-xl">{{ $t('About') }}</span>
    </router-link>
    <router-link to="/help" class="flex items-center px-3 py-2 rounded-md text-sm font-medium hover:bg-secondary-dark">
      <font-awesome-icon icon="question-circle" class="text-primary-light text-2xl py-1 px-2" />
      <span class="text-xl">{{ $t('F.A.Q.') }}</span>
    </router-link>
    <router-link to="/contact" class="flex items-center px-3 py-2 rounded-md text-sm font-medium hover:bg-secondary-dark">
      <font-awesome-icon icon="envelope" class="text-primary-light text-2xl py-1 px-2" />
      <span class="text-xl">{{ $t('Contact') }}</span>
    </router-link>
  </div>
</div>

    </div>
  </nav>
</template>



<script setup lang="ts">
import { ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useAuthStore } from '@/stores/auth';
import LanguageDropdown from '../LanguageDropdown.vue';


const authStore = useAuthStore();
const { isLoggedIn } = storeToRefs(authStore);
const isOpen = ref(false);

const logout = () => {
  authStore.logout();
};
</script>

<style scoped>
font-awesome-icon {
  font-size: 2rem; /* Adjust size as needed */
  width: 2rem;
  height: 2rem;
}
</style>
