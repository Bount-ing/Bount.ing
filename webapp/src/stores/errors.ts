import { defineStore } from 'pinia'
import { ref, nextTick } from 'vue'

export const useErrorStore = defineStore('error', () => {
  const isOpen = ref(false)
  const message = ref('')
  const isError = ref(true)

  const showError = async (msg: string, errorType = true) => {
    message.value = msg
    isError.value = errorType
    await nextTick() // Ensures reactivity kicks in
    isOpen.value = true
  }

  const closeError = () => {
    console.log('❌ Closing error modal') // Debugging
    isOpen.value = false
  }

  return { isOpen, message, isError, showError, closeError }
})
