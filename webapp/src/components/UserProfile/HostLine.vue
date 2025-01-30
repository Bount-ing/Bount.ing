<template>
  <li class="p-4 rounded-md mb-4 bg-secondary">
    <span class="flex flex-row items-center">
      <img :src="host.LogoUrl" alt="Host Image" class="w-16 h-16 rounded-full" />
      <span class="ml-4 flex flex-col flex-grow">
        <div class="text-lg font-medium">{{ hostName }}</div>
        <p class="text-gray-300">{{ hostAddress }}</p>
        <div class="text-sm text-gray-500 mt-2">
          Created at: {{ formattedDate || 'Unknown date' }}
        </div>
      </span>
      <button @click="syncGithubData" class="mr-2 px-2 py-1 m-0 rounded-lg bg-blue-500 text-white">
        Sync
      </button>
      <button @click="toggleConnection" :class="buttonClass">
        {{ buttonLabel }}
      </button>
    </span>
  </li>
</template>

<script setup>
import { computed, defineProps, defineEmits } from 'vue'
import { useUserStore } from '@/stores/user'
import { api } from '@/stores/api'

// Define props and emits for the component
const props = defineProps({
  host: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['host-action'])
const userStore = useUserStore()

// Compute host name with fallback
const hostName = computed(() => props.host.Name || 'Unnamed Host')

// Compute host address with fallback
const hostAddress = computed(() => props.host.Address || 'No address available')

// Compute formatted date
const formattedDate = computed(() => {
  return props.host.CreatedAt ? formatDate(props.host.CreatedAt) : null
})

// Compute button class based on connection status
const buttonClass = computed(() => {
  console.log('Hosts:', userStore.hosts)
  const host = userStore.hosts[props.host.Address]
  return ['ml-auto px-2 py-1 m-0 rounded-lg', host?.connected ? 'bg-red-500' : 'bg-primary']
})

// Compute button label based on connection status
const buttonLabel = computed(() =>
  userStore.hosts[props.host.Address]?.connected ? 'Disconnect' : 'Connect'
)

// Helper function to format date
const formatDate = (date) => {
  return new Date(date).toLocaleDateString()
}

// Function to toggle connection status
const toggleConnection = () => {
  if (props.host.Address.includes('github.com')) {
    if (!props.host.connected) {
      initiateGitHubOAuth() // Trigger OAuth flow if not already connected
    } else {
      const actionType = 'disconnect'
      userStore[actionType + 'Github'](props.host)
      emit('host-action', { type: actionType, host: props.host })
    }
  } else if (props.host.Address.includes('stripe.com')) {
    if (!props.host.connected) {
      initiateStripeOAuth() // Trigger OAuth flow if not already connected
    } else {
      const actionType = 'disconnect'
      userStore[actionType + 'Stripe'](props.host)
      emit('host-action', { type: actionType, host: props.host })
    }
  }
}

const initiateGitHubOAuth = () => {
  const clientId = import.meta.env.VITE_GITHUB_CLIENT_ID
  const redirectUri = import.meta.env.VITE_GITHUB_REDIRECT_URI

  const oauthProvider = 'https://github.com'
  //encode base 64
  const encodedOauthProvider = btoa(oauthProvider)
  //get the state from the server
  api.get(`/v1/oauth/${encodedOauthProvider}`).then((response) => {
    console.log('Response:', response)
    const state = response.data.state
    const scope = 'read:user user:email read:org'
    const authUrl = `https://github.com/login/oauth/authorize?client_id=${clientId}&redirect_uri=${redirectUri}&scope=${scope}&state=${state}`

    // Store state in session for verification
    // document.cookie = `oauth_state=${state}; path=/; Secure; SameSite=None; HttpOnly`;
    document.cookie = `oauth_state=${state}; path=/; SameSite=None; HttpOnly`

    //send a random cookie to the server
    // Redirect to GitHub for OAuth
    window.location.href = authUrl
  })
}

const initiateStripeOAuth = () => {
  const clientId = import.meta.env.VITE_STRIPE_CLIENT_ID
  const redirectUri = import.meta.env.VITE_STRIPE_REDIRECT_URI

  const oauthProvider = 'https://stripe.com'
  //encode base 64
  const encodedOauthProvider = btoa(oauthProvider)
  //get the state from the server
  api.get(`/v1/oauth/${encodedOauthProvider}`).then((response) => {
    console.log('Response:', response)
    const state = response.data.state
    const scope = 'read_write';
    const authUrl = `https://connect.stripe.com/oauth/authorize?response_type=code&client_id=${clientId}&scope=${scope}&state=${state}`

    // Store state in session for verification
    // document.cookie = `oauth_state=${state}; path=/; Secure; SameSite=None; HttpOnly`;
    document.cookie = `oauth_state=${state}; path=/; SameSite=None; HttpOnly`

    //send a random cookie to the server
    // Redirect to Stripe for OAuth
    window.location.href = authUrl
  })
}

// Function to sync GitHub data
const syncGithubData = () => {
  const token = userStore.hosts['https://github.com'].token
  userStore.syncGithubData(token)
}
</script>
