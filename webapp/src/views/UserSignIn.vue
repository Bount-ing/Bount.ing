<script setup>
import { ref, onMounted } from 'vue'
import { useUserStore } from '@/stores/user.ts'
import { api } from '@/stores/api.ts'
import { sha256 } from 'js-sha256'
import { useRouter } from 'vue-router'
import { useNotificationStore } from '@/stores/notification.ts'

const router = useRouter()
const userStore = useUserStore()
const notificationStore = useNotificationStore()


const showPassword = ref(false)
const mail = ref('')
const password = ref('')

const userLogin = async () => {
  let data = {
    mail: mail.value,
    password: sha256(password.value)
  }

  try {
    await userStore.login(data)
    router.push('/profile')
  } catch (error) {
    notificationStore.showNotification("Invalid credentials", "error")
  }
}
</script>
<template>
  <div class="max-w-lg max-h-full mx-auto mt-12">
    <div class="shadow-lg rounded-lg p-6 pt-8 mt-8">
      <div class="text-center mb-6">
        <h2 class="text-2xl font-semibold text-primary">
          {{ $t('account.userCredentials.welcome') }}
        </h2>
        <p class="text-secondary-light">{{ $t('account.userCredentials.enterCredentials') }}</p>
      </div>

      <form>
        <div class="space-y-4">
          <!-- Email Field -->
          <div class="relative">
            <input
              type="email"
              v-model="mail"
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary bg-secondary"
              :placeholder="$t('account.userCredentials.email')"
              required
            />
          </div>

          <!-- Password Field -->
          <div class="relative">
            <input
              v-model="password"
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary bg-secondary"
              :type="showPassword ? 'text' : 'password'"
              :placeholder="$t('account.userCredentials.password')"
              required
            />
            <button
              type="button"
              class="absolute right-3 top-2 text-gray-500"
              @click="showPassword = !showPassword"
            >
              <font-awesome-icon icon="eye" />
            </button>
          </div>

          <!-- Remember Me & Forgot Password -->
          <div class="flex justify-between items-center">
            <div class="flex items-center space-x-2">
              <input
                type="checkbox"
                class="h-4 w-4 bg-secondary text-primary border-primary"
                id="signin-check"
              />
              <label for="signin-check" class="text-gray-600">{{
                $t('account.userCredentials.rememberMe')
              }}</label>
            </div>
            <div>
              <a href="#" class="text-sm text-primary hover:text-primary-light">{{
                $t('account.userCredentials.forgotPassword')
              }}</a>
            </div>
          </div>

          <!-- Submit Button -->
          <div class="mt-6">
            <button
              type="button"
              @click="userLogin()"
              class="w-full bg-primary text-white py-2 rounded-lg hover:bg-primary-light focus:outline-none transition-all"
            >
              <i class="fas fa-unlock mr-2"></i>
              <span>{{ $t('account.userCredentials.enter') }}</span>
            </button>
          </div>
        </div>
      </form>

      <div class="mt-6 text-center">
        <p class="text-gray-600">
          <router-link to="/signup" class="text-primary hover:text-primary-light">
            {{ $t('account.userCredentials.noAccount') }}
          </router-link>
        </p>
      </div>
    </div>
  </div>
</template>
