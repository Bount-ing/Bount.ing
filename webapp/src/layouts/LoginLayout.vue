<script setup>
import LandingHeader from '@/components/UI/PublicHeader.vue';


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
