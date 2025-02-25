<template>
	<button @click="ConnectStripe" class="px-3 py-2 rounded-md text-sm font-medium border text-error-light border-error hover:border-error-light">Connect Stripe</button>
</template>

<script lang="ts">
	import { defineComponent } from 'vue';

  export default defineComponent({
	setup() {
	  const ConnectStripe = () => {
		  const clientId = import.meta.env.VITE_STRIPE_CLIENT_ID;
		  const redirectUri = import.meta.env.VITE_STRIPE_REDIRECT_URI;
		  const authUrl = `https://connect.stripe.com/oauth/authorize?response_type=code&client_id=${clientId}&scope=read_write&redirect_uri=${redirectUri}`;

		  window.location.href = authUrl;
	  };

	  return {
		ConnectStripe
	  };
	},
  });
</script>

