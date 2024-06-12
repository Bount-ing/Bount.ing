import axios from 'axios';
import { ref, computed } from 'vue';
import { defineStore } from "pinia";

interface User {
	userid: string;
	username: string;
	avatar: string;
	userBio?: string;
	fullName?: string;
	email?: string;
	phoneNumber?: string;
	location?: string;
	aboutMe?: string;
	interests?: string[];
	recentPosts?: any[];
	level?: number;
	achievements?: any[];
	activities?: any[];
	bounties?: any[];
	transactions?: any[];
	paymentInfo?: any;
}


export const useUserStore = defineStore('user', () => {
	const u = ref<User | null>(null);
	const loggedIn = ref(false)
	const token = ref<string>('');
	const isLoggedIn = computed<boolean>({
		get() {
			return loggedIn.value || !!localStorage.getItem('token');
		},
		set(v: boolean) {
			loggedIn.value = v
		},
	})
	const user = computed<User | null>({
		get() {
			if (!u.value) {
				const userData = localStorage.getItem('user');
				if (userData) {
					u.value = JSON.parse(userData);
				}
			}
			return u.value;
		},
		set(value: User | null) {
			u.value = value;
		}
	})

	// Getter Header for app API
	const authHeader = computed<string>(() => 'Bearer ' + (localStorage.getItem('token') || ''));

	const authGithubHeader = computed<string>(() => {
	  const storedToken = localStorage.getItem('token');
	  if (storedToken) {
		const u = parseJwt(storedToken);
		token.value = u.access_token;
	  }
	  const t = token.value;
	  return 'Bearer ' + t;
	});

	async function login(jwt: string) {
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

	function parseJwt(token: string) {
		if (!token) throw new Error("Token is required");
		const base64Url = token.split('.')[1];
		const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
		const jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
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
