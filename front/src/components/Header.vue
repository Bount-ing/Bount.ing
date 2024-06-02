<template>
  <nav class="border-b border-primary bg-secondary-dark text-primary-light shadow-lg fixed top-0 left-0 w-full z-50">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <!-- Logo and navigation links -->
        <div class="flex items-center">
          <router-link to="/" class="flex items-center">
            <img src="/bount.ing.logo.png" class="h-12 w-12 bg-black rounded-3xl" alt="Bount.ing Logo" />
            <span class="ml-3 text-3xl font-bold text-primary tracking-tight">Bount.ing</span>
          </router-link>
          <div class="hidden md:flex space-x-4 ml-10">
            <router-link to="/" class="px-3 py-2 rounded-md text-sm font-medium bg-gray-900 hover:bg-gray-700">Home</router-link>
            <router-link to="/dashboard" class="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-700">Dashboard</router-link>
            <router-link to="/profile" class="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-700">Profile</router-link>
            <router-link to="/about" class="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-700">About</router-link>
            <router-link to="/help" class="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-700">F.A.Q.</router-link>
            <router-link to="/contact" class="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-700">Contact</router-link>
          </div>
        </div>

        <!-- Authentication buttons -->
        <div class="hidden md:block">
          <router-link v-if="!isLoggedIn" to="/login" class="px-3 py-2 rounded-md text-sm font-medium border text-success-light border-success hover:border-success-light">Login</router-link>
          <button v-else @click="logout" class="px-3 py-2 rounded-md text-sm font-medium border text-error-light border-error hover:border-error-light">Logout</button>
        </div>

        <!-- Mobile menu button -->
        <div class="md:hidden bg-gray-700 p-2 rounded-lg">
          <button @click="isOpen = !isOpen" class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-600">
            <span class="sr-only">Open main menu</span>
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
        <router-link to="/" class="block px-3 py-2 rounded-md text-base font-medium bg-gray-900 hover:bg-gray-700">Home</router-link>
        <router-link to="/dashboard" class="block px-3 py-2 rounded-md text-base font-medium hover:bg-gray-700">Dashboard</router-link>
        <router-link to="/profile" class="block px-3 py-2 rounded-md text-base font-medium hover:bg-gray-700">Profile</router-link>
        <router-link to="/about" class="block px-3 py-2 rounded-md text-base font-medium hover:bg-gray-700">About</router-link>
        <router-link to="/help" class="block px-3 py-2 rounded-md text-base font-medium hover:bg-gray-700">F.A.Q.</router-link>
        <router-link to="/contact" class="block px-3 py-2 rounded-md text-base font-medium hover:bg-gray-700">Contact</router-link>
        <router-link v-if="!isLoggedIn" to="/login" class="block px-3 py-2 rounded-md text-sm font-medium border text-success-light border-success hover:border-success-light">Login</router-link>
        <button v-else @click="logout" class="block px-3 py-2 rounded-md text-sm font-medium border text-error-light border-error hover:border-error-light">Logout</button>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useUserStore } from '../stores/user';

const user = useUserStore()
const { isLoggedIn } = storeToRefs(user)
const isOpen = ref(false);

const logout = () => {
  user.logout();
  isLoggedIn.value = false;
};

</script>

<style scoped>
/* Add your styles here */
</style>
