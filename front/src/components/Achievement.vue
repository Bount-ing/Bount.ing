<template>
    <div class="achievements grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        <div v-for="achievement in achievements" :key="achievement.id" class="achievement p-4 bg-gray-700 rounded-lg">
            <h3 class="text-xl font-semibold">{{ achievement.title }}</h3>
            <p class="text-gray-300">{{ achievement.description }}</p>
            <p class="text-sm text-gray-400">Date: {{ achievement.date }}</p>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        achievements: {
            type: Array,
            required: true
        }
    }
};
</script>

<style scoped>
/* Add your styles here */
</style>
