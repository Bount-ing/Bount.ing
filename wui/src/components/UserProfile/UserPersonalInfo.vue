<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue';
import { storeToRefs } from 'pinia';
import type { User, EditableUserFields } from '@/types/user';
import { DEFAULT_USER_VALUES } from '@/types/user';
import { useUserStore } from '@/stores/user';
import router from '@/router';

const userStore = useUserStore();
const { user } = storeToRefs(userStore);

const isLoading = ref(true);
const error = ref<string | null>(null);
const isEditing = ref(false);
const editedUser = ref<EditableUserFields>(DEFAULT_USER_VALUES);

function startEditing() {
  editedUser.value = {
    username: user.value?.username || '',
    fullName: user.value?.fullName || '',
    phoneNumber: user.value?.phoneNumber || '',
    location: user.value?.location || '',
    bio: user.value?.bio || '',
    avatar: user.value?.avatar || '',
  };
  isEditing.value = true;
}

async function saveChanges() {
  try {
    if (!user.value) return;
    
    // You'll need to implement this method in your user store
    await userStore.updateProfileInfo(editedUser.value);
    isEditing.value = false;
  } catch (err) {
    error.value = 'Failed to update profile. Please try again.';
  }
}

function cancelEditing() {
  isEditing.value = false;
}

onMounted(async () => {
  try {
    await userStore.getUserInfo();
  } catch (err) {
    error.value = 'Unable to load user data. Please refresh the page.';
  } finally {
    isLoading.value = false;
  }
  if (!user.value) router.push('/signin');
});
</script>

<template>
  <div class="max-w-4xl mx-auto p-6">
    <div class="mb-8">
      <h1 class="text-2xl font-bold">Account Settings</h1>
      <p class="text-gray-400">Manage your account settings and preferences</p>
    </div>

    <div v-if="isLoading" class="text-center py-8">
      <p>Loading your profile...</p>
    </div>

    <div v-else-if="error" class="bg-red-900/50 text-red-200 p-4 rounded-lg">
      {{ error }}
    </div>

    <div v-else-if="user?.id" class="space-y-6">
      <section class="bg-gray-800 rounded-lg shadow-md overflow-hidden">
        <div class="p-6 border-b border-gray-700">
          <h2 class="text-xl font-semibold">Profile Information</h2>
          <p class="text-sm text-gray-400">Your public profile information</p>
        </div>

        <div class="p-6">
          <div class="flex items-start space-x-6 mb-8">
            <img :src="user?.avatar || '/default-avatar.png'" class="w-24 h-24 rounded-full object-cover" alt="Profile avatar" />
            <div class="flex-1">
              <div v-if="!isEditing" class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <p class="text-sm text-gray-400">Username</p>
                  <p class="font-medium">{{ user.username || 'Not set' }}</p>
                </div>
                <div>
                  <p class="text-sm text-gray-400">Email</p>
                  <p class="font-medium">{{ user.email || 'Not set' }}</p>
                </div>
                <div>
                  <p class="text-sm text-gray-400">Full Name</p>
                  <p class="font-medium">{{ user.fullName || 'Not set' }}</p>
                </div>
              </div>
              
              <form v-else @submit.prevent="saveChanges" class="space-y-4">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div>
                    <label for="username" class="block text-sm text-gray-400">Username</label>
                    <input
                      id="username"
                      v-model="editedUser.username"
                      type="text"
                      class="mt-1 block w-full px-3 py-2 bg-gray-700 border border-gray-600 rounded-md text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                  </div>
                  
                  <div>
                    <label for="full_name" class="block text-sm text-gray-400">Full Name</label>
                    <input
                      id="full_name"
                      v-model="editedUser.fullName"
                      type="text"
                      class="mt-1 block w-full px-3 py-2 bg-gray-700 border border-gray-600 rounded-md text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                  </div>
                </div>
                
                <div class="flex space-x-4">
                  <button
                    type="submit"
                    class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                  >
                    Save Changes
                  </button>
                  
                  <button
                    type="button"
                    @click="cancelEditing"
                    class="px-4 py-2 bg-gray-700 text-white rounded-md hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
                  >
                    Cancel
                  </button>
                </div>
              </form>
            </div>
          </div>

          <div v-if="!isEditing" class="flex justify-end">
            <button
              @click="startEditing"
              class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            >
              Edit Profile
            </button>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>