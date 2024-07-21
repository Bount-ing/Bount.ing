<template>
  <section class="min-h-screen flex flex-col items-center justify-center p-4 text-white">
    <div v-if="loading" class="text-center text-2xl font-semibold">Loading...</div>
    <div v-else class="w-full max-w-6xl rounded-2xl shadow-xl overflow-hidden bg-gray-800">
      <!-- Tab Content -->
      <div class="p-6">
		  <AdminBountyList />
      </div>
    </div>
  </section>
</template>

<script>
import AdminBountyList from '../components/AdminBountyList.vue';
import { useUserStore } from '../stores/user';

export default {
  components: {
	  AdminBountyList
  },
  setup() {
    const userStore = useUserStore();
    const user = userStore.user;

    return {
      user,
    };
  }
};
</script>
