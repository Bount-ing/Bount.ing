<script setup>
import LandingHeader from '@/components/UI/PublicHeader.vue';
import { useErrorStore } from '@/stores/errors';
import ErrorModal from '@/components/ErrorModal.vue';

const errorStore = useErrorStore();
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
		<ErrorModal
			:isOpen="errorStore.isOpen"
			:message="errorStore.message"
			:isError="errorStore.isError"
			@close="errorStore.closeError"
		/>
	</div>
</template>

<style scoped>

</style>
