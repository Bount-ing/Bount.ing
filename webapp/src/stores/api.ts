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
import {useErrorStore} from '@/stores/errors';
import { nextTick } from 'vue';
import router from '@/router';  // Changed from named to default import


interface RetryConfig extends InternalAxiosRequestConfig {
    _retry?: boolean;
}

interface QueueItem {
    resolve: (value?: unknown) => void;
    reject: (error?: unknown) => void;
}

const ApiBaseURL = import.meta.env.VITE_API_BASE_URL;

// Type-safe axios methods
export const axiosMethods = {
    get: <T = any>(url: string, config?: AxiosRequestConfig) => 
        axios.get<T>(url, config),
    post: <T = any>(url: string, data?: any, config?: AxiosRequestConfig) => 
        axios.post<T>(url, data, config),
    delete: <T = any>(url: string, config?: AxiosRequestConfig) => 
        axios.delete<T>(url, config),
    patch: <T = any>(url: string, data?: any, config?: AxiosRequestConfig) => 
        axios.patch<T>(url, data, config),
};

export const api: AxiosInstance = axios.create({ 
    baseURL: ApiBaseURL, 
    withCredentials: true 
});

// Keep track of refresh token request to prevent multiple simultaneous refreshes
let isRefreshing = false;
let failedQueue: QueueItem[] = [];

const processQueue = (error: Error | null, token: string | null = null): void => {
    failedQueue.forEach(prom => {
        if (error) {
            prom.reject(error);
        } else {
            prom.resolve(token);
        }
    });
    failedQueue = [];
};

api.interceptors.request.use(
    (config: InternalAxiosRequestConfig): InternalAxiosRequestConfig => {
        const userStore = useUserStore();
        const navi = useNaviStore();

        if (userStore.isLoggedIn && config.headers) {
            config.headers['Authorization'] = `${userStore.authHeader}`;
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



export default api;