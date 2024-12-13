<script setup>
import { ref, computed } from 'vue'
import { sha256 } from 'js-sha256'
import { useRoute } from 'vue-router'
import { api } from '@/stores/api.ts'
import { useNotificationsStore } from '@/stores/notification.ts'
import { useUserStore } from '@/stores/user.ts'

const notifStore = useNotificationsStore()
const userStore = useUserStore()


const route = useRoute()

const verifCode = ref('')
const validVerifCode = ref(false)

const password = ref('')
const verifyPassword = ref('')

const validPassword = ref(false)
const passwordsMatch = ref(false)

const passwordRegex = /^(?=.*\d)(?=.*[a-zA-Z]).{6,}$/
const passwordMinLen = 6

const passwordRules = () => {
	if (passwordRegex.test(password.value)) {
		validPassword.value = true
	} else {
		validPassword.value = false
	}
}

const passwordRulesHelperMsg = computed(() => {
	return validPassword.value
		? '<span style="color: green">Ok!</span>'
		: '<span style="color:red">Password must be ' +
				passwordMinLen +
				' characters and contain at least 1 number and 1letter</span>'
})

const passwordsMatchHelperMsg = computed(() => {
	return verifyPassword.value == password.value
		? '<span style="color: green">Ok!</span>'
		: '<span style="color:red">Passwords don\'t match'
})

const verifyLink = () => {
	var code = route.params.code

	api.post('/v1/signup/verif/' + code)
		.then((response) => {
			verifCode.value = route.params.verifCode
			validVerifCode.value = true
			notifStore.success('Link verified')
		})
		.catch((error) => {
			console.log(error)
		})
}

const sendPassword = () => {
	let data = {
		password: sha256(password.value),
		code: route.params.code
	}

	api
		.post('/v1/signup/password', data)
		.then((response) => {
			notifStore.success('Password set successfully')
			autoLogin()
		})
		.catch((error) => {
			notifStore.error('Error setting password')
			console.log(error)
		})
}

const autoLogin = async (data) => {
	let mailb64 = route.params.mailb64
	console.log(mailb64)
	let mail = atob(mailb64)
	console.log(mail)
	let loginData = {
		mail: mail,
		password: sha256(password.value)
	}
	await userStore.login(loginData)
	router.push('/profile')
	console.log(loginData)
}

verifyLink()
</script>

<template>
	<div class="user-form-category-btn">
		<ul class="nav nav-tabs">
			<li><router-link to="/signin" class="nav-link">sign in</router-link></li>
			<li><router-link to="/signup" class="nav-link active">sign up</router-link></li>
		</ul>
	</div>
	<div class="tab-pane active" id="login-tab">
		<form v-if="validVerifCode" @submit.prevent="sendPassword">
			<div class="col-12">
				<div class="user-form-title">
					<h2>Define your password</h2>
					<p>Setup your new account in just a few seconds.</p>
				</div>
				<div class="form-group">
					<input
						type="password"
						class="form-control"
						placeholder="Password"
						v-model="password"
						v-on:input="passwordRules()"
					/><button class="form-icon">
						<i class="eye fas fa-eye"></i></button
					><small class="form-alert" v-html="passwordRulesHelperMsg"></small>
				</div>
			</div>
			<div class="col-12">
				<div class="form-group">
					<input
						type="password"
						class="form-control"
						placeholder="Repeat Password"
						v-model="verifyPassword"
					/><button class="form-icon">
						<i class="eye fas fa-eye"></i></button
					><small class="form-alert" v-html="passwordsMatchHelperMsg"></small>
				</div>
			</div>
			<div class="col-12">
				<div class="form-group">
					<button type="submit" class="btn btn-inline">
						<i class="fas fa-user-check"></i><span>Create new account</span>
					</button>
				</div>
			</div>
		</form>
		<div class="col-12" v-else>
			<span style="color: red"> Invalid or expired link </span>
		</div>
	</div>
</template>
