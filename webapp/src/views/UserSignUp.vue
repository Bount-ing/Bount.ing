<script setup>
import { ref } from 'vue'
import { notify } from "@kyvg/vue3-notification" // Changed this line
import { api } from '@/stores/api.ts'

const userMail = ref('')
const userExists = ref(false)
const userCreated = ref(false)

const registerUser = async () => {
  try {
    const response = await api.post('/v1/signup', {
      email: userMail.value
    })
    
    userCreated.value = true
    userExists.value = false
    notify({
      title: "Success!",
      text: "Your account has been created successfully",
      type: "success",
      duration: 5000
    })
  } catch (error) {
    if (error.response?.status === 409) {
      userExists.value = true
      notify({
        title: "Account Exists",
        text: "An account with this email already exists",
        type: "error",
        duration: 5000
      })
    } else {
      notify({
        title: "Error",
        text: "Failed to create account. Please try again",
        type: "error",
        duration: 5000
      })
    }
  }
}
</script>

<template>
  <div class="container mx-auto p-6">
    <!-- Tab Links -->
    <notifications position="top right" />  <!-- Added this line -->

    <!-- Registration Form -->
    <div class="mt-8 p-8 rounded-lg shadow-md max-w-md mx-auto">
      <div class="mb-6">
        <h2 class="text-2xl font-bold text-primary">Register</h2>
        <p class="text-secondary-light">Setup a new account in a minute.</p>
      </div>

      <form @submit.prevent="registerUser">
        <!-- Email Input -->
        <div class="mb-4">
          <label for="email" class="block text-gray-700 font-medium mb-2">Email</label>
          <input
            type="email"
            id="email"
            v-model="userMail"
            class="w-full p-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-primary bg-secondary"
            placeholder="Enter your email"
            required
          />
          <small class="text-sm text-secondary-light mt-1 block">We will send a verification code to make sure it's correct.</small>
          <small v-if="userExists" class="text-sm text-red-500 mt-2">User already exists</small>
        </div>

        <!-- Submit Button -->
        <div class="mb-4">
          <button
            type="submit"
            class="w-full bg-primary text-white font-semibold py-3 rounded-md hover:bg-primary-light focus:outline-none focus:ring-2 focus:ring-primary"
          >
            <i class="fas fa-user-check mr-2"></i><span>Create new account</span>
          </button>
        </div>
      </form>

      <!-- Direction to Sign In -->
      <div class="mt-4 text-center">
        <p class="text-gray-600">
          Already have an account? Click on the 
          <router-link to="/signin" class="text-primary font-semibold">Sign In</router-link> button above.
        </p>
      </div>
    </div>
  </div>
</template>