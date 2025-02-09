import axios from 'axios';
import type { 
    AxiosError, 
    AxiosInstance, 
    AxiosRequestConfig, 
    AxiosResponse,
    InternalAxiosRequestConfig
} from 'axios';
import { useUserStore } from '@/stores/user';
import { useNaviStore } from '@/stores/navigation';
import { useNotificationStore } from '@/stores/notification';
import router from '@/router';

interface QueueItem {
    resolve: (value?: unknown) => void;
    reject: (error?: unknown) => void;
    config: InternalAxiosRequestConfig;
}

const ApiBaseURL = import.meta.env.VITE_API_BASE_URL;

let isRefreshing = false;
let failedQueue: QueueItem[] = [];

const processQueue = (error: Error | null, token: string | null = null): void => {
    failedQueue.forEach(({ resolve, reject, config }) => {
        if (error) {
            reject(error);
        } else if (token) {
            config.headers.Authorization = `Bearer ${token}`;
            resolve(axios(config));
        }
    });
    failedQueue = [];
};

export const api: AxiosInstance = axios.create({ 
    baseURL: ApiBaseURL, 
    withCredentials: true 
});

api.interceptors.request.use(
    (config: InternalAxiosRequestConfig): InternalAxiosRequestConfig => {
        const userStore = useUserStore();
        const navi = useNaviStore();

        if (userStore.isLoggedIn && config.headers) {
            const token = localStorage.getItem('token');
            if (token) {
                config.headers['Authorization'] = `Bearer ${token}`;
            }
        }
        navi.SetLoading();
        return config;
    },
    (error: AxiosError): Promise<AxiosError> => {
        const navi = useNaviStore();
        navi.UnsetLoading();
        return Promise.reject(error);
    }
);

api.interceptors.response.use(
    (response: AxiosResponse) => {
        const navi = useNaviStore();
        navi.UnsetLoading();
        return response;
    },
    async (error: AxiosError) => {
        const navi = useNaviStore();
        const userStore = useUserStore();
        const notificationStore = useNotificationStore();
        const originalConfig = error.config as InternalAxiosRequestConfig;

        navi.UnsetLoading();

        // If the error is not 401 or the request was for refresh token
        if (error.response?.status !== 401 || originalConfig.url === '/v1/refresh') {
            return Promise.reject(error);
        }

        if (!isRefreshing) {
            isRefreshing = true;

            try {
                const refreshToken = localStorage.getItem('refreshToken');
                if (!refreshToken) {
                    throw new Error('No refresh token available');
                }

                const response = await axios.post(
                    `${ApiBaseURL}/v1/refresh`,
                    {},
                    {
                        headers: {
                            Cookie: `refreshTkn=${refreshToken}`
                        }
                    }
                );

                const { accessToken, refreshToken: newRefreshToken } = response.data;

                localStorage.setItem('token', accessToken);
                localStorage.setItem('refreshToken', newRefreshToken);

                // Process queue with new token
                processQueue(null, accessToken);
                
                // Retry original request
                originalConfig.headers.Authorization = `Bearer ${accessToken}`;
                return axios(originalConfig);

            } catch (refreshError) {
                processQueue(new Error('Failed to refresh token'));
                userStore.logout();
                router.push('/signin');
                notificationStore.showNotification('Session expired. Please log in again.', 'warning');
                return Promise.reject(refreshError);
            } finally {
                isRefreshing = false;
            }
        }

        // Add failed request to queue
        return new Promise((resolve, reject) => {
            failedQueue.push({
                resolve,
                reject,
                config: originalConfig
            });
        });
    }
);

export default api;