<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useNotificationStore } from '@/stores/notification.ts'
import { api } from '@/stores/api'
import { useRoute } from 'vue-router'
import { sha256 } from 'js-sha256'

const route = useRoute()
const router = useRouter()
const notificationStore = useNotificationStore()

const email = ref('')
const newPassword = ref('')
const confirmPassword = ref('')
const showPassword = ref(false)
const resetCode = ref('')
const isResetCodeView = ref(false)

const requestPasswordReset = async () => {
  try {
    await api.post('/v1/reset-password-request', { email: email.value })
    notificationStore.showNotification('If the email exists, a reset code will be sent', 'success')
  } catch (error) {
    notificationStore.showNotification('Failed to request password reset', 'error')
  }
}

const resetPassword = async () => {
  console.log('resetPassword')
  if (newPassword.value !== confirmPassword.value) {
    notificationStore.showNotification('Passwords do not match', 'error')
    return
  }

  try {
    await api.post('/v1/reset-password', {
      code: resetCode.value,
      password: sha256(newPassword.value)
    })

    notificationStore.showNotification('Password reset successfully', 'success')
    router.push('/signin')
  } catch (error) {
    console.log(error)
    notificationStore.showNotification('Failed to reset password', 'error')
  }
}

onMounted(() => {
  if (route.params.code) {
    resetCode.value = route.params.code
    isResetCodeView.value = true
  }
})
</script>

<template>
  <div class="max-w-lg max-h-full mx-auto mt-12">
    <div class="shadow-lg rounded-lg p-6 pt-8 mt-8">
      <div class="text-center mb-6">
        <h2 class="text-2xl font-semibold text-primary">
          {{ $t('account.resetPassword.title') }}
        </h2>
        <p class="text-secondary-light">
          {{ $t('account.resetPassword.instructions') }}
        </p>
      </div>

      <form @submit.prevent="isResetCodeView ? resetPassword() : requestPasswordReset()">
        <div class="space-y-4">
          <!-- Email Field (only shown in initial view) -->
          <div v-if="!isResetCodeView" class="relative">
            <input
              type="email"
              v-model="email"
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary bg-secondary"
              :placeholder="$t('account.resetPassword.email')"
              required
            />
          </div>

          <!-- Reset Code and New Password Fields (shown after email submission) -->
          <template v-if="isResetCodeView">
            <div class="relative">
              <input
                v-model="newPassword"
                class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary bg-secondary"
                :type="showPassword ? 'text' : 'password'"
                :placeholder="$t('account.resetPassword.newPassword')"
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

            <div class="relative">
              <input
                v-model="confirmPassword"
                class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-primary bg-secondary"
                :type="showPassword ? 'text' : 'password'"
                :placeholder="$t('account.resetPassword.confirmPassword')"
                required
              />
            </div>
          </template>

          <!-- Submit Button -->
          <div class="mt-6">
            <button
              type="submit"
              class="w-full bg-primary text-white py-2 rounded-lg hover:bg-primary-light focus:outline-none transition-all"
            >
              <i class="fas fa-key mr-2"></i>
              <span>{{
                isResetCodeView
                  ? $t('account.resetPassword.reset')
                  : $t('account.resetPassword.send')
              }}</span>
            </button>
          </div>
        </div>
      </form>

      <div class="mt-6 text-center">
        <router-link to="/signin" class="text-primary hover:text-primary-light">
          {{ $t('account.resetPassword.backToLogin') }}
        </router-link>
      </div>
    </div>
  </div>
</template>
