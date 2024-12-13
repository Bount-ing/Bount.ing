<template>
  <section class="min-h-screen flex flex-col lg:flex-row text-white">
    <!-- Button to toggle Sidebar on mobile -->
    <button
      @click="toggleSidebar"
      class="lg:hidden fixed top-20 left-4 z-50 p-2 bg-primary rounded-md shadow-lg"
      aria-label="Toggle Sidebar"
      :aria-expanded="isSidebarOpen"
    >
      <span>{{ isSidebarOpen ? 'Close' : 'Open' }} Menu</span>
    </button>

    <!-- Left Sidebar Navigation -->
    <div
      :class="[
        'w-64 bg-secondary p-4 pt-32 transition-transform ease-in-out duration-300 transform',
        { 'translate-x-0': isSidebarOpen, '-translate-x-full lg:translate-x-0': !isSidebarOpen }
      ]"
      class="absolute lg:relative top-0 left-0 bottom-0 z-40 lg:z-auto"
      aria-hidden="!isSidebarOpen"
    >
      <div class="flex flex-col space-y-2 pt-2">
        <!-- Sidebar Items (Tabs) -->
        <button
          v-for="tab in tabs"
          :key="tab.name"
          @click="selectTab(tab)"
          :class="[
            'p-4 cursor-pointer text-left rounded-md transition-colors',
            { 'bg-gray-700': currentTab && currentTab.name === tab.name }
          ]"
        >
          {{ tab.name }}
        </button>
      </div>
    </div>

    <!-- Main Content Area -->
    <div class="flex-1 pt-8 pl-4 pr-4 lg:pl-64 transition-all duration-300">
      <div v-if="loading" class="text-center text-2xl font-semibold">Loading...</div>
      <div v-else class="rounded-2xl shadow-xl overflow-hidden p-4 bg-secondary-dark text-white">
        <UserProfile :user="user" />
        
        <!-- Tab Content -->
        <div class="mt-6">
          <keep-alive>
            <component :is="currentTab?.component" :key="currentTab?.name" v-if="currentTab" />
          </keep-alive>
        </div>
      </div>
    </div>
  </section>
</template>


<script setup>
import { ref, watchEffect } from 'vue';
import UserProfile from '../components/UserProfile.vue';
import UserBadgesList from '../components/UserProfile/UserBadgesList.vue';
import UserOrganizationsList from '../components/UserProfile/UserOrganizationsList.vue';
import UserHostsList from '../components/UserProfile/UserHostsList.vue';
import UserRepositoriesList from '../components/UserProfile/UserRepositoriesList.vue';
import UserIssuesList from '../components/UserProfile/UserIssuesList.vue';
import UserBountiesList from '../components/UserProfile/UserBountiesList.vue';
import UserPaymentsList from '../components/UserProfile/UserPaymentsList.vue';
import { useUserStore } from '../stores/user';
import { useI18n } from 'vue-i18n';

const { t, locale } = useI18n();
const userStore = useUserStore();
const user = userStore.user;

// Define tabs
const tabs = ref([]);
const updateTabs = () => {
  tabs.value = [
    { name: t('profile.badges'), component: UserBadgesList },
    { name: t('profile.organizations'), component: UserOrganizationsList },
    { name: t('profile.hosts'), component: UserHostsList },
    { name: t('profile.repositories'), component: UserRepositoriesList },
    { name: t('profile.issues'), component: UserIssuesList },
    { name: t('profile.bounties'), component: UserBountiesList },
    { name: t('profile.payments'), component: UserPaymentsList },
  ];
};

updateTabs();
watchEffect(() => updateTabs());

const currentTab = ref(tabs.value[0] || { name: '', component: null });
const selectTab = (tab) => {
  currentTab.value = tab;
};

const isSidebarOpen = ref(false);
const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value;
};
</script>
