<script setup>
import PublicHeader from '@/components/UI/PublicHeader.vue'
import LoggedHeader from '@/components/UI/LoggedHeader.vue'
import Footer from '@/components/UI/Footer.vue'
import NotificationModal from '@/components/NotificationModal.vue'
import { useNotificationStore } from '@/stores/notification.ts'
import { useAuthStore } from '@/stores/auth.ts'
import { storeToRefs } from 'pinia'

const notificationStore = useNotificationStore()

const authStore = useAuthStore()
const { isLoggedIn } = storeToRefs(authStore)

</script>

<template>
  <div>
    <PublicHeader v-if="!isLoggedIn" />
    <LoggedHeader v-else />
    <div class="pt-8">
      <router-view />
    </div>
    <Footer />
    <NotificationModal
      :isOpen="notificationStore.isOpen"
      :message="notificationStore.message"
      :type="notificationStore.type"
      @close="notificationStore.closeNotification"
    />
  </div>
</template>
