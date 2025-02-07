<template>
  <div v-if="isOpen" class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center mt-8">
    <div class="bg-secondary p-6 rounded-lg w-full max-w-md text-center">
      <h3 class="text-xl font-semibold mb-4" :class="isError ? 'text-red-500' : 'text-yellow-500'">
        {{ isError ? 'Error' : 'Warning' }}
      </h3>
      <div class="text-sm mb-6">
        <p v-for="(sentence, index) in formattedMessage" :key="index">{{ sentence }}</p>
      </div>
      <button
        @click="closeModal"
        class="px-4 py-2 bg-primary text-white rounded-lg hover:bg-primary-light transition-all focus:outline-none"
      >
        OK
      </button>
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits, computed } from 'vue'

const props = defineProps({
  isOpen: {
    type: Boolean,
    required: true
  },
  message: {
    type: String,
    required: true
  },
  isError: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['close'])  // This defines the emit

// Computed property to split the message into an array of sentences
const formattedMessage = computed(() => {
  return props.message.split('.').filter(Boolean).map(sentence => sentence.trim() + '.')
})

const closeModal = () => {
  console.log("Modal closed")  // Debugging
  emit('close')
}
</script>

<style scoped>
/* z-index: 50 to ensure modal is on top of everything */

.fixed {
  z-index: 9999 !important;  /* Ensure it is on top */
}
</style>
