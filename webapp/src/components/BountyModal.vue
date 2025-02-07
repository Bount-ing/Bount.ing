<template>
    <div
      v-if="isOpen"
      class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center mt-16"
    >
      <div class="bg-secondary p-6 rounded-lg w-full max-w-lg max-h-[90vh] overflow-y-auto">
        <h3 class="text-2xl font-semibold mb-6 text-center">Set Bounty for Issue</h3>
  
        <form @submit.prevent="onSubmit">
          <!-- Bounty Amount -->
          <div class="mb-5">
            <label for="bounty" class="block text-sm font-medium mb-2">
              Bounty Amount (in euros)
            </label>
            <input
              v-model="localAmount"
              id="bounty"
              type="number"
              min="10"
              required
              class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
              placeholder="Enter bounty amount"
              aria-label="Bounty Amount"
            />
          </div>
  
          <!-- Start Date -->
          <div class="mb-5">
            <label for="start-date" class="block text-sm font-medium mb-2">Start Date</label>
            <input
              v-model="localStartAt"
              id="start-date"
              type="date"
              required
              class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
              aria-label="Start Date"
            />
          </div>
  
          <!-- End Date -->
          <div class="mb-5">
            <label for="end-date" class="block text-sm font-medium mb-2">End Date</label>
            <input
              v-model="localEndAt"
              id="end-date"
              type="date"
              required
              min="2024-01-01"
              class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
              aria-label="End Date"
            />
          </div>
  
          <!-- Variables for bounty adjustments -->
          <div
            v-if="localVariables.length > 0"
            v-for="(variable, index) in localVariables"
            :key="index"
            class="mb-5"
          >
            <h4 class="text-sm font-semibold mb-2">Variable {{ index + 1 }}</h4>
  
            <div class="mb-3">
              <label class="block text-sm">Direction</label>
              <select
                v-model="variable.direction"
                required
                class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
                aria-label="Variable Direction"
              >
                <option value="increase">Increase</option>
                <option value="decrease">Decrease</option>
              </select>
            </div>
  
            <div class="mb-3">
              <label :for="`variable-amount-${index}`" class="block text-sm">
                Max Amount (in euros)
              </label>
              <input
                v-model="variable.amount"
                :id="`variable-amount-${index}`"
                type="number"
                min="1"
                required
                class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
                placeholder="Amount"
                aria-label="Variable Amount"
              />
            </div>
  
            <div class="mb-3">
              <label :for="`variable-start-date-${index}`" class="block text-sm">
                Start Date
              </label>
              <input
                v-model="variable.startAt"
                :id="`variable-start-date-${index}`"
                type="date"
                required
                class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
                placeholder="Start Date"
                aria-label="Variable Start Date"
              />
            </div>
  
            <div class="mb-3">
              <label :for="`variable-end-date-${index}`" class="block text-sm">
                End Date
              </label>
              <input
                v-model="variable.endAt"
                :id="`variable-end-date-${index}`"
                type="date"
                required
                class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
                placeholder="End Date"
                aria-label="Variable End Date"
              />
            </div>
  
            <!-- Remove Variable Button -->
            <button
              type="button"
              @click="removeVariable(index)"
              class="text-red-500 mt-2 hover:text-red-700 transition-all"
            >
              Remove Variable
            </button>
          </div>
  
          <!-- Add Variable Button -->
          <div class="mb-5">
            <button
              type="button"
              @click="addVariable"
              class="px-6 py-3 bg-secondary-light text-white rounded-lg hover:text-secondary-dark transition-all"
            >
              Add Variable
            </button>
          </div>
  
          <!-- Action Buttons -->
          <div class="flex justify-between mt-6">
            <button
              type="submit"
              class="px-6 py-3 bg-primary text-white rounded-lg hover:bg-primary-light transition-all focus:outline-none focus:ring-2 focus:ring-green-500"
            >
              Set Bounty
            </button>
            <button
              type="button"
              @click="closeModal"
              class="px-6 py-3 bg-secondary-light text-white rounded-lg hover:text-secondary-dark transition-all focus:outline-none focus:ring-2 focus:ring-gray-500"
            >
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, watch } from 'vue'
  
  const props = defineProps({
    isOpen: {
      type: Boolean,
      required: true
    },
    selectedIssue: {
      type: Object,
      default: null
    }
  })
  
  const emit = defineEmits(['close', 'submit'])
  
  const localAmount = ref(0)
  const localStartAt = ref('')
  const localEndAt = ref('')
  const localVariables = ref([])
  
  const addVariable = () => {
    localVariables.value.push({ 
      amount: 0, 
      startAt: '', 
      endAt: '', 
      direction: 'increase' 
    })
  }
  
  const removeVariable = (index) => {
    localVariables.value.splice(index, 1)
  }
  
  const closeModal = () => {
    emit('close')
    resetForm()
  }
  
  const resetForm = () => {
    localAmount.value = 0
    localStartAt.value = ''
    localEndAt.value = ''
    localVariables.value = []
  }
  
  const onSubmit = () => {
    const formattedVariables = localVariables.value.map((variable) => ({
      ...variable,
      startAt: variable.startAt ? new Date(variable.startAt).toISOString() : null,
      endAt: variable.endAt ? new Date(variable.endAt).toISOString() : null
    }))
  
    const bountyData = {
      amount: localAmount.value,
      currency: 'EUR',
      issue_id: props.selectedIssue?.ID,
      issueUrl: props.selectedIssue?.issueUrl,
      issueTitle: props.selectedIssue?.title,
      issueBody: props.selectedIssue?.body,
      startAt: localStartAt.value ? new Date(localStartAt.value).toISOString() : null,
      endAt: localEndAt.value ? new Date(localEndAt.value).toISOString() : null,
      variables: formattedVariables
    }
  
    emit('submit', bountyData)
    closeModal()
  }
  
  // Reset form when modal opens
  watch(() => props.isOpen, (newValue) => {
    if (newValue) {
      resetForm()
    }
  })
  </script>