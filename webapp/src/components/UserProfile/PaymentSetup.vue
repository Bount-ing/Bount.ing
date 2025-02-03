<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useUserStore } from '../../stores/user'
import { api } from '@/stores/api'

const userStore = useUserStore()
const stripe = ref(null)
const elements = ref(null)
const cardElement = ref(null)
const error = ref('')
const isProcessing = ref(false)
const isAdding = ref(false)
const cardholderName = ref('')
const savedCards = ref([])
const paymentMethods = ref([])

onMounted(async () => {
  try {
    // Load Stripe.js if not already loaded
    if (!window.Stripe) {
      const script = document.createElement('script')
      script.src = 'https://js.stripe.com/v3/'
      script.async = true
      document.head.appendChild(script)
      await new Promise((resolve, reject) => {
        script.onload = resolve
        script.onerror = reject
      })
    }

    // Initialize Stripe with error handling
    const publicKey = import.meta.env.VITE_STRIPE_PUBLIC_KEY
    if (!publicKey) {
      throw new Error('Stripe public key is not configured')
    }
    stripe.value = Stripe(publicKey)
    await loadSavedCards()
  } catch (err) {
    error.value = 'Failed to initialize payment system: ' + err.message
    console.error('Stripe initialization error:', err)
  }
})

onUnmounted(() => {
  if (cardElement.value) {
    cardElement.value.destroy()
  }
})

async function loadSavedCards() {
  try {
    const response = await api.get('/v1/payment-methods')
    console.log('Saved cards:', response.data.paymentMethods)
    savedCards.value = response.data.paymentMethods
    paymentMethods.value = savedCards.value.filter((card) => card.Type === 'card')
  } catch (err) {
    console.error('Load saved cards error:', err)
    error.value = 'Failed to load saved cards: ' + err.message
  }
}

function startAddingCard() {
  isAdding.value = true
  error.value = ''

  setTimeout(() => {
    try {
      if (!stripe.value) {
        throw new Error('Payment form is not ready.')
      }

      elements.value = stripe.value.elements()

      cardElement.value = elements.value.create('card', {
        style: {
          base: {
            color: '#fff',
            '::placeholder': { color: '#aab7c4' },
            fontSize: '16px'
          }
        }
      })

      cardElement.value.mount('#card-element')

      cardElement.value.on('change', (event) => {
        error.value = event.error ? event.error.message : ''
      })
    } catch (err) {
      error.value = err.message
      console.error('Start adding card error:', err)
    }
  }, 100)
}

function cancelAddingCard() {
  isAdding.value = false
  cardholderName.value = ''
  error.value = ''
  if (cardElement.value) {
    cardElement.value.unmount()
    cardElement.value = null
  }
}

async function handleSubmit() {
  if (!stripe.value || !cardElement.value) {
    error.value = 'Payment form is not ready yet.'
    return
  }

  if (!cardholderName.value.trim()) {
    error.value = 'Cardholder name is required.'
    return
  }

  isProcessing.value = true
  error.value = ''

  try {
    // Create PaymentMethod
    const { paymentMethod, error: paymentMethodError } = await stripe.value.createPaymentMethod({
      type: 'card',
      card: cardElement.value,
      billing_details: {
        name: cardholderName.value.trim()
      }
    })

    if (paymentMethodError) {
      throw new Error(paymentMethodError.message)
    }

    // Get SetupIntent
    const setupResponse = await api.post('/v1/create-setup-intent')
    const { clientSecret } = setupResponse.data

    if (!clientSecret) {
      throw new Error('Failed to get setup intent')
    }

    // Confirm SetupIntent
    const { error: setupError, setupIntent } = await stripe.value.confirmCardSetup(clientSecret, {
      payment_method: paymentMethod.id
    })

    if (setupError) {
      throw new Error(setupError.message)
    }

    // Handle success
    if (setupIntent.status === 'succeeded') {
      await api.post('/v1/confirm-setup', {
        setupIntentId: setupIntent.id,
        paymentMethodId: setupIntent.payment_method
      })

      await loadSavedCards()
      cancelAddingCard()
    } else {
      throw new Error('Setup intent did not succeed')
    }
  } catch (err) {
    console.error('Submit payment error:', err)
    error.value = err.message || 'Failed to set up payment method'
  } finally {
    isProcessing.value = false
  }
}

async function removeCard(paymentMethodId) {
  if (!paymentMethodId) {
    error.value = 'Invalid payment method'
    return
  }

  try {
    await api.delete(`/v1/payment-methods/${paymentMethodId}`)
    await loadSavedCards()
  } catch (err) {
    console.error('Remove card error:', err)
    error.value = 'Failed to remove card: ' + err.message
  }
}
</script>

<template>
  <div class="max-w-2xl mx-auto p-6 bg-secondary-dark rounded-lg">
    <div v-if="error" class="text-red-400 text-sm mb-4">
      {{ error }}
    </div>

    <!-- Saved Cards Section -->
    <div v-if="savedCards && savedCards.length > 0">
      <h3 class="text-xl font-semibold mb-4">Saved Cards</h3>
      <ul class="space-y-4">
        <li
          v-for="card in savedCards"
          :key="card.ID"
          class="flex items-center justify-between p-4 bg-secondary rounded-lg space-y-2"
        >
          <div class="flex items-center space-x-4">
            <div class="p-2 bg-primary rounded">
              <span v-if="card.Brand === 'visa'" class="text-lg">💳</span>
              <span v-else-if="card.Brand === 'mastercard'" class="text-lg">💳</span>
              <span v-else class="text-lg">💳</span>
            </div>
            <div>
              <p class="font-medium">•••• •••• •••• {{ card.Last4 }}</p>
              <p class="text-sm text-gray-400">Expires {{ card.ExpMonth }}/{{ card.ExpYear }}</p>
            </div>
          </div>
          <button
            @click="removeCard(card.ID)"
            :disabled="isProcessing"
            class="px-3 py-1 text-sm text-red-400 hover:text-red-300 transition-colors"
          >
            Remove
          </button>
        </li>
      </ul>
    </div>

    <!-- Add New Card Section -->
    <div v-if="!isAdding">
      <button
        @click="startAddingCard"
        :disabled="isProcessing"
        class="w-full py-3 px-4 bg-primary text-white rounded-lg hover:bg-primary-dark transition-colors my-4"
      >
        + Add New Card
      </button>
    </div>

    <div v-else class="mt-6 p-6 bg-secondary rounded-lg">
      <div class="mb-4">
        <label for="cardholder-name" class="block text-sm font-medium mb-2">Cardholder Name</label>
        <input
          id="cardholder-name"
          v-model="cardholderName"
          type="text"
          required
          :disabled="isProcessing"
          class="w-full px-4 py-2 bg-secondary border border-gray-600 rounded-lg focus:outline-none focus:border-primary"
        />
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-2">Card Details</label>
        <div id="card-element" class="p-4 bg-secondary border border-gray-600 rounded-lg"></div>
      </div>

      <div class="flex space-x-4">
        <button
          @click="handleSubmit"
          :disabled="isProcessing"
          class="flex-1 py-3 px-4 bg-primary text-white rounded-lg hover:bg-primary-dark transition-colors disabled:opacity-50"
        >
          {{ isProcessing ? 'Processing...' : 'Save Card' }}
        </button>
        <button
          @click="cancelAddingCard"
          :disabled="isProcessing"
          class="flex-1 py-3 px-4 bg-gray-700 text-white rounded-lg hover:bg-gray-600 transition-colors disabled:opacity-50"
        >
          Cancel
        </button>
      </div>
    </div>
  </div>
</template>
