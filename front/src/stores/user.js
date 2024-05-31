import axios from 'axios';
import { ref, computed } from 'vue';
import { defineStore } from "pinia";

export const useUserStore = defineStore('user', () => {
	const u = ref(null)
	const loggedIn = ref(false)
	const token = ref('')
	const isLoggedIn = computed({
		get() {
			return loggedIn.value || !!localStorage.getItem('token');
		},
		set(v) {
			loggedIn.value = v
		},
	})
	const user = computed({
		get() {
			if (!u.value) {
				u.value = JSON.parse(localStorage.getItem('user'));
			}
			return u.value
		},
	})

	// Getter Header for app API
	const authHeader = computed(() => 'Bearer ' + localStorage.getItem('token'))

	// Header for Github API
	const authGithubHeader = computed(() => {
		if (!token.value) {
			const u = parseJwt(localStorage.getItem('token'));
			token.value = u.access_token
		}
		const t = token.value
		return 'Bearer ' + t
	})

	async function login(jwt) {
		localStorage.setItem('token', jwt)
		const u = parseJwt(jwt);
		token.value = u.access_token;
		// Fetch User Data
		const response = await axios.get('https://api.github.com/user', {
			headers: { Authorization: authGithubHeader.value },
		});
		localStorage.setItem('user', JSON.stringify({
			userid: u.user_id,
			username: response.data.login,
			avatar: response.data.avatar_url || 'default-image.jpg'
		}))
		loggedIn.value = true
	}

	function parseJwt(token) {
		var base64Url = token.split('.')[1];
		var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
		var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
			return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
		}).join(''));

		return JSON.parse(jsonPayload);
}

	function logout() {
		if (!isLoggedIn.value) return;
		user.value = null
		loggedIn.value = false
		localStorage.removeItem('token')
	}

	return { user, loggedIn, isLoggedIn, token, login, logout, authHeader, authGithubHeader }
});
