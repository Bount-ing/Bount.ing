<template>
  <section class="min-h-screen flex flex-col items-center justify-center p-4 text-white">
    <div v-if="loading" class="text-center text-2xl font-semibold">Loading...</div>
    <div v-else class="w-full max-w-6xl rounded-2xl shadow-xl overflow-hidden bg-gray-800">
      <UserProfile :user="user" />
      
      <!-- Tab Navigation -->
      <div class="flex justify-around bg-gray-700 text-gray-300">
        <button
          v-for="tab in tabs"
          :key="tab.name"
          @click="selectTab(tab)"
          :class="{ 'bg-gray-800': currentTab && currentTab.name === tab.name }"
          class="p-4 cursor-pointer"
        >
          {{ tab.name }}
        </button>
      </div>

      <!-- Tab Content -->
      <div class="p-6">
        <keep-alive>
          <component :is="currentTab?.component" :key="currentTab?.name" v-if="currentTab" />
        </keep-alive>
      </div>
    </div>
  </section>
</template>

<script>
import UserProfile from '../components/UserProfile.vue';
import UserBountiesList from '../components/UserBountiesList.vue';
import UserTransactionsHistory from '../components/UserTransactionsHistory.vue';
import UserPaymentInformation from '../components/UserPaymentInformation.vue';
import GitHubImages from '../components/GitHubImages.vue';
import { ref } from 'vue';
import { useUserStore } from '../stores/user';

export default {
  components: {
    UserProfile,
    UserBountiesList,
    UserTransactionsHistory,
    UserPaymentInformation,
    GitHubImages,
  },
  setup() {
    const userStore = useUserStore();
    const user = userStore.user;
    const tabs = ref([
      { name: 'Bounties', component: 'UserBountiesList' },
      { name: 'Transactions', component: 'UserTransactionsHistory' },
      { name: 'Payment Info', component: 'UserPaymentInformation' },
      { name: 'GitHub Images', component: 'GitHubImages' },
    ]);

    const currentTab = ref(tabs.value[0] || { name: '', component: 'FallbackComponent' });

    const selectTab = (tab) => {
      currentTab.value = tab;
    };

    return {
      user,
      tabs,
      currentTab,
      selectTab,
    };
  }
};
</script>
