<template>
  <div v-if="isOpen" class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center mt-8">
    <div class="bg-secondary p-6 rounded-lg w-full max-w-md text-center">
      <h3 class="text-xl font-semibold mb-4" :class="typeClasses">
        {{ typeLabels[type] }}
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
  isOpen: { type: Boolean, required: true },
  message: { type: String, required: true },
  type: { type: String, default: 'info' }
})

const emit = defineEmits(['close'])

const typeLabels = {
  info: '💡 ~ Notice ~ 💡',
  success: '✅ ~ Success ~ ✅',
  warning: '⚠️ ~ Warning ~ ⚠️',
  error: '❌ ~ Error ~ ❌'
};

const typeClasses = computed(() => {
  return {
    info: 'text-blue-500',
    success: 'text-green-500',
    warning: 'text-yellow-500',
    error: 'text-red-500'
  }[props.type]
})

const formattedMessage = computed(() => {
  if (typeof props.message !== 'string') {
    console.error("Invalid message format:", props.message);
    return [String(props.message)]; // Convert to string to avoid errors
  }
  return props.message.split('.').filter(Boolean).map(sentence => sentence.trim() + '.');
});
const closeModal = () => {
  console.log("Modal closed")
  emit('close')
}
</script>

<style scoped>
.fixed {
  z-index: 9999 !important;
}
</style>
