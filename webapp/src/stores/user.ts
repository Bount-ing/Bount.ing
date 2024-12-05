import axios from 'axios';
import { ref, computed } from 'vue';
import { defineStore } from "pinia";
import { api } from './api';

interface User {
    userid: string;
    username: string;
    avatar: string;
    userBio?: string;
    fullName?: string;
    email?: string;
    phoneNumber?: string;
    location?: string;
    bearerToken?: string;
    refreshJwt?: string;
    aboutMe?: string;
    interests?: string[];
    recentPosts?: any[];
    level?: number;
    achievements?: any[];
    activities?: any[];
    bounties?: any[];
    transactions?: any[];
    paymentInfo?: any;
    isLoggedIn?: boolean;
    authGithubHeader?: string;
}

interface LoginCredentials {
    username: string;
    password: string;
}

export const useUserStore = defineStore('user', () => {
    const user = ref<User | null>(null);
    const loggedIn = ref<boolean>(false);
    const token = ref<string>('');

    const isLoggedIn = computed<boolean>(() => {
        return loggedIn.value || !!localStorage.getItem('token');
    });

    const authHeader = computed<string>(() => {
        const storedToken = localStorage.getItem('token');
        return `Bearer ${storedToken || ''}`;
    });

    const authGithubHeader = computed<string>(() => {
        const storedGithubToken = localStorage.getItem('githubToken'); // Assuming you store the GitHub token separately
        return `Bearer ${storedGithubToken || ''}`;
    });
    

    async function githubLogin(jwt: string): Promise<void> {
        try {
            localStorage.setItem('token', jwt);
            const parsedJwt = parseJwt(jwt);
            if (!parsedJwt) throw new Error("Invalid JWT format");
    
            token.value = parsedJwt.access_token;
    
            const response = await axios.get('https://api.github.com/user', {
                headers: { Authorization: `Bearer ${jwt}` },
            });
    
            const userData: User = {
                userid: parsedJwt.user_id,
                username: response.data.login,
                avatar: response.data.avatar_url || 'default-image.jpg',
                authGithubHeader: `Bearer ${jwt}`, // Save the GitHub-specific token here
            };
    
            localStorage.setItem('githubToken', jwt); // Save GitHub token separately
            localStorage.setItem('user', JSON.stringify(userData));
            user.value = userData;
            loggedIn.value = true;
        } catch (error) {
            console.error('GitHub login failed:', error);
        }
    }
    

    async function refreshJwt(): Promise<boolean> {
        const storedToken = localStorage.getItem('token');
        if (!storedToken) {
            logout();
            return false;
        }

        try {
            const response = await axios.post('https://api.example.com/auth/refresh', {}, {
                headers: { Authorization: `Bearer ${storedToken}` },
            });

            const newToken = response.data.token;
            if (!newToken) {
                logout();
                return false;
            }

            localStorage.setItem('token', newToken);
            token.value = newToken;
            loggedIn.value = true;
            return true;
        } catch (error) {
            console.error('Token refresh failed:', error);
            logout();
            return false;
        }
    }

    async function login(creds: LoginCredentials): Promise<void> {
        try {
            const response = await api.post('/v1/signin', creds);
            const { accessToken, refreshToken } = response.data;
            console.log(response.data);

            if (accessToken && refreshToken) {
                localStorage.setItem('token', accessToken);
                localStorage.setItem('refreshToken', refreshToken);
                token.value = accessToken;
                loggedIn.value = true;

                await getUserInfo(); // Replace with an actual function fetching user data
            } else {
                throw new Error('Missing tokens in response');
            }
        } catch (error) {
            console.error('Login failed:', error);
            throw new Error('Invalid credentials or login error');
        }
    }

    function parseJwt(token: string): Record<string, any> | null {
        try {
            const base64Url = token.split('.')[1];
            const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
            const jsonPayload = atob(base64)
                .split('')
                .map((c) => `%${('00' + c.charCodeAt(0).toString(16)).slice(-2)}`)
                .join('');
            return JSON.parse(decodeURIComponent(jsonPayload));
        } catch (error) {
            console.error('Failed to parse JWT:', error);
            return null;
        }
    }

    function logout(): void {
        if (!loggedIn.value) return;

        user.value = null;
        loggedIn.value = false;
        token.value = '';

        localStorage.removeItem('token');
        localStorage.removeItem('refreshToken');
        localStorage.removeItem('user');
    }

    async function getUserInfo(): Promise<void> {
        // Mock implementation; replace with actual API call.
        try {
            const response = await axios.get('/auth/user', {
                headers: { Authorization: authHeader.value },
            });
            user.value = response.data;
            localStorage.setItem('user', JSON.stringify(response.data));
        } catch (error) {
            console.error('Failed to fetch user info:', error);
        }
    }

    return {
        user,
        loggedIn,
        isLoggedIn,
        authHeader,
        token,
        authGithubHeader,
        githubLogin,
        refreshJwt,
        login,
        logout,
        getUserInfo,
    };
});
