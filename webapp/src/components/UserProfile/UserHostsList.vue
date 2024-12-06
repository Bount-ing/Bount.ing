<template>
    <div>
      <h2 class="text-2xl font-semibold mb-4">Hosts</h2>
      <div v-if="hosts.length === 0" class="text-gray-400">No hosts available.</div>
      <ul v-else>
        <li v-for="host in hosts" :key="host.id" class="bg-gray-700 p-4 rounded-md mb-4">
          <div class="text-lg font-medium">{{ host.name }}</div>
          <p class="text-gray-300">{{ host.description }}</p>
          <div class="text-sm text-gray-500 mt-2">Hosted since: {{ host.date }}</div>
        </li>
      </ul>
    </div>
  </template>
  
  <script setup>
  const hosts = [
    { id: 1, name: 'Host 1', description: 'Description for Host 1', date: '2023-01-15' },
    { id: 2, name: 'Host 2', description: 'Description for Host 2', date: '2022-11-02' },
    { id: 3, name: 'Host 3', description: 'Description for Host 3', date: '2021-08-19' },
  ];
  </script>
  