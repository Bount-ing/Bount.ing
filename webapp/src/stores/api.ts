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
    (response: AxiosResponse) => {
        const navi = useNaviStore();
        navi.UnsetLoading();
        return response;
    },
    (error: AxiosError) => {
        const navi = useNaviStore();
        const notificationStore = useNotificationStore();
        navi.UnsetLoading();

        if (error.response?.status === 401) {
            notificationStore.showNotification('Your session has expired. Please log in again.', 'warning');
        }
        
        return Promise.reject(error);
    }
);




export default api;