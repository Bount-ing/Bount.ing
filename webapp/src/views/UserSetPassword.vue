<script setup>
import { ref, computed } from 'vue'
import { sha256 } from 'js-sha256'
import { useRoute } from 'vue-router'
import { api } from '@/stores/api.ts'
import { notify } from "@kyvg/vue3-notification" 
import { useUserStore } from '@/stores/user.ts'
import  router  from '@/router'

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
		? '<span style="color: green">✔ Ok!</span>'
		: `<ul style="color: red; padding-left: 20px; list-style-type: disc;">
				<li>At least ${passwordMinLen} characters long</li>
				<li>Contains at least 1 number</li>
				<li>Contains at least 1 letter</li>
			</ul>`;
});

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
			notify({
				title: "Success!",
				text: "Link verified successfully",
				type: "success",
				duration: 5000
			})
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
			notify(
				{
					title: "Success!",
					text: "Password set successfully",
					type: "success",
					duration: 5000
				}
			)
			autoLogin()
		})
		.catch((error) => {
			notify(
				{
					title: "Error",
					text: "Failed to set password",
					type: "error",
					duration: 5000
				}
			)
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
	<div class="flex flex-col items-center justify-center min-h-screen p-4">
	  <div class="w-full max-w-md">
		<div class="tab-pane active" id="login-tab">
		  <form v-if="validVerifCode" @submit.prevent="sendPassword" class="bg-secondary p-6 rounded-lg shadow-md">
			<div class="text-center mb-4">
			  <h2 class="text-xl font-semibold">Define your password</h2>
			  <p class="text-gray-500">Setup your new account in just a few seconds.</p>
			</div>
			<div class="mb-4">
			  <input
				type="password"
				class="w-full p-3 border rounded-lg text-primary bg-black"
				placeholder="Password"
				v-model="password"
				v-on:input="passwordRules()"
			  />
			  <small class="text-red-500" v-html="passwordRulesHelperMsg"></small>
			</div>
			<div class="mb-4">
			  <input
				type="password"
				class="w-full p-3 border rounded-lg text-primary bg-black"
				placeholder="Repeat Password"
				v-model="verifyPassword"
			  />
			  <small class="text-red-500" v-html="passwordsMatchHelperMsg"></small>
			</div>
			<div>
			  <button type="submit" class="w-full bg-primary text-white p-3 rounded-lg hover:bg-primary-light">
				<i class="fas fa-user-check"></i>
				<span class="ml-2">Create new account</span>
			  </button>
			</div>
		  </form>
		  <div class="text-center text-red-500 mt-4" v-else>
			<span>Invalid or expired link</span>
		  </div>
		</div>
	  </div>
	</div>
  </template>
  