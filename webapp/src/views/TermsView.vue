<template>
  <section class="bg-secondary-dark container mx-auto px-4 my-6 py-6 rounded-lg text-primary">
    <h1 class="text-4xl font-bold text-center mb-6">Terms of Service</h1>
    <p class="text-lg text-center mb-4">You can view the Terms of Service in the embedded PDF below.</p>
    <div class="pdf-container">
      <iframe
        :src="pdfUrl"
        width="100%"
        height="600px"
        style="border: none;"
      ></iframe>
    </div>
  </section>
</template>

<script>
export default {
  data() {
    return {
      pdfUrl: '/Bount.ing-terms-2025-02-04.pdf'
    };
  }
};
</script>

<style scoped>
.pdf-container {
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
}
</style>
