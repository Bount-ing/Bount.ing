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

api.interceptors.response.use(
    (response: AxiosResponse): AxiosResponse => {
        const navi = useNaviStore();
        navi.UnsetLoading();
        return response;
    },
    async (error: AxiosError): Promise<AxiosError | AxiosResponse> => {
        const navi = useNaviStore();
        const userStore = useUserStore();
        const originalRequest = error.config as RetryConfig;

        // Handle 401 errors
        if (error.response?.status === 401) {
            // If refresh token fails or user is logged in but token is invalid
            if (originalRequest.url?.includes('/refresh') || originalRequest._retry) {
                userStore.logout();
                navi.UnsetLoading();
                router.push('/signin');
                return Promise.reject(error);
            }

            if (!isRefreshing) {
                isRefreshing = true;
                originalRequest._retry = true;

                try {
                    await userStore.refreshJwt();
                    isRefreshing = false;
                    
                    // Process queued requests
                    processQueue(null, userStore.authHeader);
                    
                    // Retry original request with new token
                    if (originalRequest.headers) {
                        originalRequest.headers['Authorization'] = userStore.authHeader;
                    }
                    return api(originalRequest);
                } catch (refreshError) {
                    isRefreshing = false;
                    processQueue(refreshError as Error, null);
                    userStore.logout();
                    navi.UnsetLoading();
                    router.push('/signin');
                    return Promise.reject(refreshError);
                }
            } else {
                // Queue failed requests while refresh is in progress
                return new Promise((resolve, reject) => {
                    failedQueue.push({ resolve, reject });
                }).then(() => {
                    if (originalRequest.headers) {
                        originalRequest.headers['Authorization'] = userStore.authHeader;
                    }
                    return api(originalRequest);
                }).catch(err => {
                    return Promise.reject(err);
                });
            }
        }

        navi.UnsetLoading();
        return Promise.reject(error);
    }
);

export default api;