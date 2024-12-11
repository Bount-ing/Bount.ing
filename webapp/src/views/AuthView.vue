<template>
	<div class="container mx-auto px-4 py-12">
	  <h1 class="text-4xl font-bold text-center mb-6 text-gray-800">Login with GitHub</h1>
	  <div class="flex justify-center">
		  Redirecting...
	  </div>
	</div>
</template>

<script setup>
	import { onMounted } from 'vue';
	import { useRoute, useRouter } from 'vue-router';
	import { useUserStore } from '../stores/user';

	const route = useRoute();
	const router = useRouter();
	const userStore = useUserStore();

	onMounted(async () => {
		const github_token = route.query.github_token;
		console.log('Query:', route.query);
		if (github_token) {
			try {
				userStore.hosts["https://github.com"] = {
					connected: true,
					token: github_token,
				};
				userStore.syncGithubData(github_token);
				router.push('/');
			} catch (error) {
				console.error('Login failed:', error);
				// Handle the error (e.g., show a notification to the user)
			}
		} else {
			console.error('No token found in the query parameters');
			// Handle the case where the token is missing
		}
	});
</script>

<style scoped>
	.cursor-pointer {
		cursor: pointer;
	}
</style>
