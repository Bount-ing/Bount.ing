<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user';
import LandingHeader from '@/components/UI/PublicHeader.vue';

//check if user is logged in
const userStore = useUserStore();
const isLoggedIn = ref(userStore.isLoggedIn);
const router = useRouter();

//redirect to dashboard if user is logged in
if (isLoggedIn.value) {
	router.push({ path: '/dashboard' });
}

</script>
<template>
	<div class="fullscreen">
		<section class="user-form-part">
			<div class="user-form-banner">
				<LandingHeader />
				<div class="user-form-content">
					
				</div>
			</div>
			<div class="user-form-category">
				<RouterView />
			</div>
		</section>
	</div>
</template>

<style scoped>

</style>
