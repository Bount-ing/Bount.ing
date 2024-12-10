<template>
    <li class="p-4 rounded-md mb-4 bg-secondary">
        <span class="flex flex-row">
            <img :src="host.LogoUrl" alt="Host Image" class="w-16 h-16 rounded-full" />
            <span class="ml-4 flex flex-col">
                <div class="text-lg font-medium">{{ host.Name || "Unnamed Host" }}</div>
                <p class="text-gray-300">{{ host.Address || "No address available" }}</p>
                <div class="text-sm text-gray-500 mt-2">
                    Created at: {{ formattedDate || "Unknown date" }}
                </div>
            </span>
            <button 
                @click="toggleConnection" 
                :class="['ml-auto px-2 py-1 m-0 rounded-lg', isConnected ? 'bg-red-500' : 'bg-primary']">
                {{ isConnected ? 'Disconnect' : 'Connect' }}
            </button>
        </span>
    </li>
</template>

<script setup>
import { ref, computed, defineProps } from 'vue';

// Define props for the component
const props = defineProps({
    host: {
        type: Object,
        required: true,
    },
});

// Reactive state for connection status
const isConnected = ref(false);

// Function to toggle connection status
const toggleConnection = () => {
    isConnected.value = !isConnected.value;
};

// Compute a formatted date with a fallback for missing data
const formattedDate = computed(() => {
    return props.host.CreatedAt ? new Date(props.host.CreatedAt).toLocaleDateString() : null;
});
</script>
