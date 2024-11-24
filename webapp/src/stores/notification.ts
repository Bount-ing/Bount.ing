import { defineStore } from 'pinia';

type Notification = {
    message: string;
    style: 'error' | 'warning' | 'success' | 'default';
    type?: string; // Added because `success` and `error` methods use `type`
    timeOut?: number;
    index?: number;
};

export const useNotificationsStore = defineStore('notifications', {
    state: () => ({
        _notifs: [] as Notification[], // Explicitly typed
        _styles: ['error', 'warning', 'success', 'default'] as const, // Read-only array of styles
    }),

    getters: {
        notifs: (state) => {
            return state._notifs; // Use `state` instead of `this`
        },
    },

    actions: {
        add(notif: Partial<Notification>) { // Allow partial notifications (not all fields required)
            if (!notif.message) {
                notif.message = 'Error'; // Default message
            }
            if (!notif.style) {
                notif.style = 'success'; // Default style
            }

            const notification: Notification = {
                ...notif,
                message: notif.message,
                style: notif.style,
            };

            this._notifs.push(notification);

            setTimeout(() => {
                this.clear(notification);
            }, notif.timeOut || 10000); // Default timeout
        },

        remove(notif: Notification) {
            this._notifs = this._notifs.filter((n) => n !== notif); // Remove specific notification
        },

        clear(notif: Notification | number) {
            if (typeof notif === 'object') {
                this.remove(notif);
            } else {
                this._notifs.splice(notif, 1); // Remove by index
            }
        },

        clearAll() {
            this.$reset(); // Clear all notifications
        },

        success(msg: string) {
            this.add({ type: 'success', message: msg, style: 'success' });
        },

        error(msg: string) {
            this.add({ type: 'error', message: msg, style: 'error' });
        },
    },
});
