<script setup lang="ts">
import { ref, computed, watch } from 'vue';
import { storeToRefs } from 'pinia';
import { useUserStore } from '@/stores/user';
import type { EditableUserFields } from '@/types/user';
import { DEFAULT_USER_VALUES } from '@/types/user';

const userStore = useUserStore();
const { user } = storeToRefs(userStore);

// Create a computed property to safely access user data with defaults
const safeUser = computed(() => {
  if (!user.value) {
    return DEFAULT_USER_VALUES;
  }

  return {
    fullName: user.value.fullName ?? DEFAULT_USER_VALUES.fullName,
    email: user.value.email ?? DEFAULT_USER_VALUES.email,
    phoneNumber: user.value.phoneNumber ?? DEFAULT_USER_VALUES.phoneNumber,
    location: user.value.location ?? DEFAULT_USER_VALUES.location,
    userBio: user.value.userBio ?? DEFAULT_USER_VALUES.userBio,
    avatar: user.value.avatar ?? DEFAULT_USER_VALUES.avatar,
  };
});

// Initialize editable user with safe values
const editableUser = ref<EditableUserFields>({ ...safeUser.value });

// Safe reset function
const resetForm = () => {
  editableUser.value = { ...safeUser.value };
};

const isUpdating = ref(false);

// Safe update function
const updateUserInfo = async () => {
  if (!user.value) {
    console.error('No user data available');
    return;
  }

  try {
    isUpdating.value = true;
    
    const updatedData = {
      ...user.value,
      ...editableUser.value
    };

    await userStore.updateUser(updatedData);
    // Show success message
  } catch (error) {
    console.error('Failed to update user:', error);
    // Handle error
  } finally {
    isUpdating.value = false;
  }
};

// Watch for user changes safely
watch(() => user.value, (newUser) => {
  if (newUser) {
    resetForm();
  }
}, { deep: true });
</script>

<template>
  <div class="max-w-4xl mx-auto p-6">
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-2xl font-bold">Account Settings</h1>
      <p class="text-gray-400">Manage your account settings and preferences</p>
    </div>

    <!-- Show settings only if we have a valid user ID -->
    <div v-if="user?.id" class="space-y-6">
      <!-- Rest of your settings form -->
      <section class="bg-gray-800 rounded-lg shadow-md overflow-hidden">
        <div class="p-6 border-b border-gray-700">
          <h2 class="text-xl font-semibold">Profile Information</h2>
          <p class="text-sm text-gray-400">Your public profile information</p>
        </div>

        <div class="p-6">
          <!-- Avatar and Basic Info -->
          <div class="flex items-start space-x-6 mb-8">
            <div class="flex-shrink-0">
              <img 
                :src="safeUser.avatar || '/default-avatar.png'" 
                class="w-24 h-24 rounded-full object-cover"
                alt="Profile avatar"
              />
              <button class="mt-2 text-sm text-blue-400 hover:text-blue-300">
                Change Avatar
              </button>
            </div>
            
            <!-- Non-editable Information -->
            <div class="flex-grow">
              <div class="grid grid-cols-2 gap-4">
                <div class="space-y-1">
                  <p class="text-sm text-gray-400">Username</p>
                  <p class="font-medium">{{ user.username || 'Not set' }}</p>
                </div>
                <div class="space-y-1">
                  <p class="text-sm text-gray-400">User ID</p>
                  <p class="font-medium">{{ user.id || 'N/A' }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Editable Information Form -->
          <form @submit.prevent="updateUserInfo" class="space-y-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <!-- Personal Information -->
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium mb-1">Full Name</label>
                  <input 
                    v-model="editableUser.fullName" 
                    type="text" 
                    class="w-full p-2.5 rounded-md bg-gray-700 border border-gray-600 focus:border-blue-500 focus:ring-1 focus:ring-blue-500"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium mb-1">Email</label>
                  <input 
                    v-model="editableUser.email" 
                    type="email" 
                    class="w-full p-2.5 rounded-md bg-gray-700 border border-gray-600 focus:border-blue-500 focus:ring-1 focus:ring-blue-500"
                  />
                </div>
                <!-- ... other fields ... -->
              </div>
            </div>

            <!-- Save Button -->
            <div class="flex justify-end space-x-3">
              <button 
                type="button"
                @click="resetForm"
                class="px-4 py-2 text-sm font-medium text-gray-300 hover:text-white bg-gray-700 rounded-md"
              >
                Cancel
              </button>
              <button 
                type="submit"
                class="px-4 py-2 text-sm font-medium text-white bg-blue-600 hover:bg-blue-500 rounded-md disabled:opacity-50"
                :disabled="isUpdating"
              >
                {{ isUpdating ? 'Saving...' : 'Save Changes' }}
              </button>
            </div>
          </form>
        </div>
      </section>
    </div>

    <!-- Loading State -->
    <div v-else-if="isUpdating" class="text-center py-8">
      <p>Loading user data...</p>
    </div>

    <!-- Error State -->
    <div v-else class="text-center py-8 text-red-400">
      <p>Unable to load user data. Please try refreshing the page.</p>
    </div>
  </div>
</template>