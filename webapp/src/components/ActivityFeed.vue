<template>
    <div class="activity-feed space-y-4">
        <div v-for="activity in activities" :key="activity.id" class="activity p-4 bg-gray-700 rounded-lg">
            <p class="text-gray-300"><span class="font-semibold">{{ activity.action }}</span> in <a :href="activity.repoUrl" class="text-blue-400 hover:underline">{{ activity.repo }}</a></p>
            <p class="text-sm text-gray-400">{{ activity.date }}</p>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        activities: {
            type: Array,
            required: true
        }
    }
};
</script>

<style scoped>
/* Add your styles here */
</style>
