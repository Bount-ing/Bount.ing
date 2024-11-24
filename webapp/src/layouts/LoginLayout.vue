<script setup>
</script>
<template>
	<div class="fullscreen">
		<section class="user-form-part">
			<div class="user-form-banner">
				<div class="user-form-content">
				</div>
			</div>
			<div class="user-form-category">
				<router-view />
			</div>
		</section>
	</div>
</template>

<style scoped>
.previous-btn {
	background-color: var(--white);
	width: 35px;
	height: 35px;
	border-radius: 100%;
	display: block;
}

.fullscreen {
	z-index: 99;
	background: var(--chalk);
	position: absolute;
	top: 0;
	left: 0;
	bottom: 0;
	right: 0;
}

.user-form-banner::before {
	background-color: hsla(343, 57%, 30%, 1);
	background-image: radial-gradient(at 40% 20%, hsla(343, 70%, 51%, 1) 0px, transparent 50%),
		radial-gradient(at 80% 100%, hsla(359, 32%, 72%, 1) 0px, transparent 50%),
		radial-gradient(at 80% 50%, hsla(340, 69%, 36%, 1) 0px, transparent 50%),
		radial-gradient(at 0% 50%, hsla(355, 29%, 54%, 1) 0px, transparent 50%),
		radial-gradient(at 0% 100%, hsla(22, 100%, 77%, 1) 0px, transparent 50%);
}
</style>
