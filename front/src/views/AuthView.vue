<template>
	<div class="container mx-auto px-4 py-12">
	  <h1 class="text-4xl font-bold text-center mb-6 text-gray-800">Login with GitHub</h1>
	  <div class="flex justify-center">
		  Redirecting...
	  </div>
	</div>
</template>

<script setup>
	import { useRoute, useRouter } from 'vue-router';
	import { useUserStore } from '../stores/user';

	const route = useRoute();
	const router = useRouter();

	useUserStore().login(route.query.token).then(() => {
		router.push('/')
	});
</script>

<style scoped>
	.cursor-pointer {
		cursor: pointer;
	}
</style>

