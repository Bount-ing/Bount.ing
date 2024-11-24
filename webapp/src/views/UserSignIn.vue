<script setup>
import { ref } from 'vue'
import { useUserStore } from '@/stores/user.ts'
import { api } from '@/stores/api.ts'
import { sha256 } from 'js-sha256'
import { useRouter } from 'vue-router'

	const router = useRouter()
const showPassword = ref(false)
const mail = ref('')
const password = ref('')
const userStore = useUserStore()

const userLogin = async () => {
	let data = {
		mail: mail.value,
		password: sha256(password.value)
	}

	await userStore.login(data)
	router.push('/user')
	console.log(data)
}
</script>
<template>
	<div class="user-form-category-btn">
		<ul class="nav nav-tabs">
			<li><router-link to="/signin" class="nav-link active">sign in</router-link></li>
			<li><router-link to="/signup" class="nav-link">sign up</router-link></li>
		</ul>
	</div>
	<div class="tab-pane active" id="login-tab">
		<div class="user-form-title">
			<h2>{{ $t('content.userSignin.welcome') }}</h2>
			<p>{{ $t('content.userSignin.userCreds') }}</p>
		</div>
		<form data-bitwarden-watching="1">
			<div class="row">
				<div class="col-12">
					<div class="form-group">
						<input
							type="text"
							v-model="mail"
							class="form-control"
							:placeholder="$t('content.userSignin.mailField')"
						/>
					</div>
				</div>
				<div class="col-12">
					<div class="form-group">
						<input
							v-model="password"
							class="form-control"
							id="pass"
							:placeholder="$t('content.userSignin.pwdField')"
							:type="showPassword ? 'text' : 'password'"
						/><button
							type="button"
							class="form-icon"
							style="color: rgba(0, 0, 0, 0.5)"
							@click="() => (showPassword = !showPassword)"
						>
							<font-awesome-icon icon="eye" />
						</button>
					</div>
				</div>
				<div class="col-6">
					<div class="form-group">
						<div class="custom-control custom-checkbox">
							<input type="checkbox" class="custom-control-input" id="signin-check" /><label
								class="custom-control-label"
								for="signin-check"
								>Remember me</label
							>
						</div>
					</div>
				</div>
				<div class="col-6">
					<div class="form-group text-right">
						<a href="#" class="form-forgot">{{ $t('content.userSignin.forgotPwd') }} </a>
					</div>
				</div>
				<div class="col-12">
					<div class="form-group">
						<button type="button" class="btn btn-inline" @click="userLogin()">
							<i class="fas fa-unlock"></i><span>{{ $t('content.userSignin.enter') }}</span>
						</button>
					</div>
				</div>
			</div>
		</form>
		<div class="user-form-direction">
			<p>
				<router-link to="/signup" v-html="$t('content.userSignin.noAccountMsg')"> </router-link>
			</p>
		</div>
	</div>
</template>
