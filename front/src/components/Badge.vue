<template>
    <div class="badge flex items-center space-x-2">
        <span class="text-lg font-bold">Level {{ level }}</span>
        <img :src="badgeIcon" alt="Level Badge" class="w-8 h-8" />
    </div>
</template>

<script>
export default {
    props: {
        level: {
            type: Number,
            required: true
        }
    },
    computed: {
        badgeIcon() {
            if (this.level < 5) return 'path/to/bronze-badge.png';
            if (this.level < 10) return 'path/to/silver-badge.png';
            return 'path/to/gold-badge.png';
        }
    }
};
</script>

<style scoped>
/* Add your styles here */
</style>
