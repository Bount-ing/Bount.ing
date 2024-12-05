<script setup>
import { ref } from 'vue'
import { api } from '@/stores/api.ts'
import { useNotificationsStore } from '@/stores/notification.ts'
const notifStore = useNotificationsStore()

const userMail = ref('')

const userExists = ref(false)
const userCreated = ref(false)

const registerUser = () => {
	api.post('/v1/signup', {
		email: userMail.value
	})
		.then((response) => {
			userCreated.value = true
			notifStore.success('User created, please check your mail for the verification code')
		})
		.catch((error) => {
			if (error.response) {
				if (error.response.status == 409) {
					userExists.value = true
				}
			}
			notifStore.error('Error creating user')

		})
}
</script>
<template>
	<div class="user-form-category-btn">
		<ul class="nav nav-tabs">
			<li><router-link to="/signin" class="nav-link">sign in</router-link></li>
			<li><router-link to="/signup" class="nav-link active">sign up</router-link></li>
		</ul>
	</div>
	<div class="tab-pane active" id="register-tab">
		<div class="user-form-title">
			<h2>Register</h2>
			<p>Setup a new account in a minute.</p>
		</div>
		<form data-bitwarden-watching="1" @submit.prevent="registerUser" action="#">
			<div class="row">
				<div class="col-12">
					<div class="form-group">
						<input type="email" class="form-control" placeholder="Email" v-model="userMail" /><small
							class="form-alert"
							>Please provide a valid mail adress, we will send a verification code to make sure
							it's correct</small
						>
						<small v-if="userExists" style="color: var(--red)">User already exists </small>
					</div>
				</div>
				<div class="col-12">
					<div class="form-group">
						<button type="button" @click="registerUser()" class="btn btn-inline">
							<i class="fas fa-user-check"></i><span>Create new account</span>
						</button>
					</div>
				</div>
			</div>
		</form>
		<div class="tab-pane active">
			<div class="user-form-direction">
				<p>
					Already have an account? click on the
					<router-link to="/userSignin/signin">(sign in)</router-link> button above.
				</p>
			</div>
		</div>
	</div>
</template>
