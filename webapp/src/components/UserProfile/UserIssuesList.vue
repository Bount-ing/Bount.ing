<template>
    <div>
      <h2 class="text-2xl font-semibold mb-4">Issues</h2>
      
      <!-- If there are no issues, display a message -->
      <div v-if="issues.length === 0" class="text-gray-400">
        No issues found.
      </div>
      
      <!-- List of issues -->
      <ul class="space-y-4">
        <li
          v-for="issue in issues"
          :key="issue.id"
          class="p-4 bg-gray-700 rounded-lg hover:bg-gray-600 transition"
        >
          <div class="flex justify-between">
            <h3 class="text-xl font-semibold">{{ issue.title }}</h3>
            <span :class="['text-sm font-semibold', issue.status === 'open' ? 'text-green-400' : 'text-red-400']">
              {{ issue.status }}
            </span>
          </div>
          <p class="text-sm mt-2 text-gray-300">{{ issue.description }}</p>
        </li>
      </ul>
    </div>
  </template>
  
  <script setup>
  // Example issues data (in a real app, this would come from an API or store)
  const issues = [
    {
      id: 1,
      title: 'Issue #1: Bug in payment gateway',
      description: 'There is a bug in the payment gateway where transactions are failing.',
      status: 'open',
    },
    {
      id: 2,
      title: 'Issue #2: UI not responsive on mobile',
      description: 'The app UI is not resizing properly on mobile devices.',
      status: 'closed',
    },
    {
      id: 3,
      title: 'Issue #3: Feature request for new filters',
      description: 'Users requested the ability to filter results by date range.',
      status: 'open',
    },
  ];
  </script>
  