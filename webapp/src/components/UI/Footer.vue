<template>
  <footer class="text-primary py-4 min-w-screen w-screen bottom-0 ">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-4">
      <p>© 2024 Bount.ing. {{ $t('All rights reserved.') }}</p>
      <p>
        <router-link to="/terms" class="underline text-info">{{ $t('Terms of Service') }}</router-link> |
        <router-link to="/privacy" class="underline text-info">{{ $t('Privacy Policy') }}</router-link>
      </p>
      <p>{{ $t('Contact us') }}: <a href="mailto:support@bount.ing" class="underline text-info">support@bount.ing</a></p>
    </div>
  </footer>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RouterLink } from 'vue-router';

export default defineComponent({
  name: 'Footer',
  components: {
    RouterLink
  }
});
</script>