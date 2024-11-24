import axios from 'axios';
import { useUserStore } from '@/stores/user';
import { useNaviStore } from '@/stores/navigation';

const ApiBaseURL = import.meta.env.VITE_API_BASE_URL;

const axiosMethods = {
    get: axios.get,
    post: axios.post,
    delete: axios.delete,
    patch: axios.patch,
};

export const api = axios.create({ baseURL: ApiBaseURL, withCredentials: true });

api.interceptors.request.use(
    (config) => {
        const userStore = useUserStore();
        const navi = useNaviStore();

        if (userStore.isLoggedIn && config.headers) {
            config.headers['Authorization'] = `Bearer ${userStore.bearerToken}`;
        }
        navi.SetLoading();
        return config;
    },
    (err) => {
        return Promise.reject(err);
    }
);

api.interceptors.response.use(
    (response) => {
        const navi = useNaviStore();
        navi.UnsetLoading();
        return response;
    },
    async (err) => {
        const navi = useNaviStore();
        const userStore = useUserStore();

        const { config, response } = err;

        if (response?.status === 401 && !config?._retry) {
            config._retry = true;
            const refreshError = await userStore.refreshJwt();
            if (refreshError) {
                return Promise.reject(refreshError);
            }
            return await api(config);
        }

        navi.UnsetLoading();
        return Promise.reject(err);
    }
);
