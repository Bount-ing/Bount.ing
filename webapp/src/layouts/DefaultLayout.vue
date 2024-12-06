<script setup>
import Header from '@/components/UI/LoggedHeader.vue'
import Footer from '@/components/UI/Footer.vue'
</script>

<template>
	<Header />
	<div class="pt-8">
		<router-view />
	</div>
	<Footer />
</template>
