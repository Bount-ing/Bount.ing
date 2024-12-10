<template>
  <div>
    <h2 class="text-2xl font-semibold mb-4">Hosts</h2>
    <div v-if="loading" class="text-gray-400">Loading hosts...</div>
    <div v-else-if="hosts.length === 0" class="text-gray-400">No hosts available.</div>
    <ul v-else>
      <HostLine v-for="host in hosts" :key="host.ID" :host="host" />
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { api } from '@/stores/api'; // Import the API client
import HostLine from './HostLine.vue'; // Import the sub-component

// Reactive state variables
const hosts = ref([]);
const loading = ref(true);

// Fetch hosts data on component mount
const fetchHosts = async () => {
  try {
    const response = await api.get('/v1/hosts'); // Replace with your API endpoint
    hosts.value = response.data; // Assuming the API returns an array of hosts
  } catch (error) {
    console.error('Error fetching hosts:', error);
  } finally {
    loading.value = false;
  }
};

// Use the lifecycle hook to fetch data
onMounted(fetchHosts);
</script>
