<script setup>
import { ref } from 'vue'
import { api } from '@/stores/api.ts'
import { useNotificationStore } from '@/stores/notification.ts'

const notificationStore = useNotificationStore()

const userMail = ref('')
const userExists = ref(false)
const userCreated = ref(false)
const pdfUrl = '/Bount.ing-terms-2025-02-04.pdf'

const registerUser = async () => {
  try {
    const response = await api.post('/v1/signup', {
      email: userMail.value
    })

    userCreated.value = true
    userExists.value = false
    notificationStore.showNotification("Account created successfully. Please check your email for a verification code.", 'success')
  } catch (error) {
    if (error.response?.status === 409) {
      userExists.value = true
      notificationStore.showNotification("User already exists", 'error')
    } else {
      notificationStore.showNotification("A wild error appeared!. Please try again later.", 'error')
    }
  }
}
</script>

<template>
  <div class="container mx-auto p-6">
    <!-- Registration Form -->
    <div class="mt-8 p-8 rounded-lg shadow-md max-w-2xl mx-auto">
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
          <small class="text-sm text-secondary-light mt-1 block"
            >We will send a verification code to make sure it's correct.</small
          >
          <small v-if="userExists" class="text-sm text-red-500 mt-2">User already exists</small>
        </div>

        <section class="bg-secondary-dark container mx-auto px-4 my-6 py-6 rounded-lg text-primary w-full">
          <h1 class="text-4xl font-bold text-center mb-6">Terms of Service</h1>
          <p class="text-lg text-center mb-4">
            You can view the Terms of Service in the embedded PDF below.
          </p>
          <div class="pdf-container">
            <iframe :src="pdfUrl" width="100%" height="600px" style="border: none"></iframe>
          </div>
        </section>

        <!-- Agreement Checkbox -->
        <div class="mb-4 flex items-center">
          <input
            type="checkbox"
            id="agreement"
            class="mr-2"
            required
          />
          <label for="agreement" class="text-sm text-gray-600">
            I agree to the
            <router-link to="/terms" class="text-primary font-semibold">Terms of Service</router-link>
          </label>

          <small class="text-sm text-red-500 ml-2">* Required</small>
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
          <router-link to="/signin" class="text-primary font-semibold">Sign In</router-link> button
          above.
        </p>
      </div>
    </div>
  </div>
</template>
