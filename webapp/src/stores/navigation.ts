import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useNaviStore = defineStore('navi', {
  state: () => ({
    _category: '',
    _lang: '',
    _favorites: [] as string[], // Specify type for _favorites
    loading: ref(false),
  }),
  getters: {
    Favorites(state) {
      return state._favorites
    },
    IsFav(state) {
      return (id: string) => state._favorites.includes(id) // Explicitly type `id`
    },
  },
  actions: {
    _setFavorite(id: string) {
      if (!this._favorites.includes(id)) {
        this._favorites.push(id) // `id` is of type string, matching `_favorites`
      }
    },
    _unsetFavorite(id: string) {
      this._favorites = this._favorites.filter((i) => i !== id) // `id` and `i` are both strings
    },
    ToggleFav(id: string) {
      if (this.IsFav(id)) {
        this._unsetFavorite(id)
      } else {
        this._setFavorite(id)
      }
    },
    SetLoading() {
      this.loading = true
    },
    UnsetLoading() {
      this.loading = false
    },
  },
  persist: {
    storage: localStorage, // Ensure this is configured properly with pinia-plugin-persistedstate
  },
})
