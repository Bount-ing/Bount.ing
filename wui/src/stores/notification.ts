import { defineStore } from 'pinia'
import { ref, nextTick } from 'vue'

export const useNotificationStore = defineStore('notification', () => {
  const isOpen = ref(false)
  const message = ref('')
  const type = ref('info') // "info", "success", "warning", "error"

  const showNotification = async (msg: string, msgType = 'info') => {
    message.value = msg
    type.value = msgType
    await nextTick()
    isOpen.value = true
  }

  const closeNotification = () => {
    console.log('🔔 Closing notification modal')
    isOpen.value = false
  }

  return { isOpen, message, type, showNotification, closeNotification }
})
