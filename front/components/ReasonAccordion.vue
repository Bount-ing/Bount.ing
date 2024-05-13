<template>
  <div>
    <button @click="toggle" class="w-full py-4 px-6 bg-blue-100 rounded-md text-left text-lg font-semibold text-blue-700 flex justify-between items-center hover:bg-blue-200 transition duration-300">
      <span>{{ title }}</span>
      <span :class="spanClass">{{ isOpen ? '-' : '+' }}</span>
    </button>
    <div v-if="isOpen" class="mt-2 px-6 py-4 bg-blue-50 border border-blue-200 rounded-md">
      <ul class="list-disc space-y-2 pl-5 text-gray-600">
        <li v-for="(item, index) in content" :key="index">{{ item }}</li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ReasonAccordion',
  props: {
    title: String,
    content: Array
  },
  data() {
    return {
      isOpen: false
    };
  },
  methods: {
    toggle() {
      this.isOpen = !this.isOpen;
    }
  },
  computed: {
    spanClass() {
      // Properly compute class with condition to control the rotation of the icon
      return {
        'text-xl font-bold transition-transform duration-300 transform': true,
        'rotate-180': this.isOpen
      };
    }
  }
};
</script>

<style scoped>
/* Additional styling if needed */
</style>
