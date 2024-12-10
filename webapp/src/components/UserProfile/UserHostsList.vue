<template>
  <div>
    <h2 class="text-2xl font-semibold mb-4">Hosts</h2>
    <div v-if="loading" class="text-gray-400">Loading hosts...</div>
    <div v-else-if="hosts.length === 0" class="text-gray-400">No hosts available.</div>
    <ul v-else>
      <li 
        v-for="host in hosts" 
        :key="host.ID" 
        class="flex items-center justify-between mb-2"
      >
        <HostLine 
          :host="host" 
          class="flex-grow mr-2"
        />
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { api } from '@/stores/api';
import HostLine from './HostLine.vue';


const hosts = ref([]);
const loading = ref(true);

const fetchHosts = async () => {
  try {
    const response = await api.get('/v1/hosts');
    hosts.value = response.data;
  } catch (error) {
    console.error('Error fetching hosts:', error);
  } finally {
    loading.value = false;
  }
};



onMounted(fetchHosts);
</script>