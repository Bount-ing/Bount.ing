<template>
  <footer class="text-primary py-6 min-w-screen w-screen bottom-0 mt-8 pt-8 text-white">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8 text-center">

      <div class="grid grid-cols-1 sm:grid-cols-3 gap-6 text-center">
        <!-- Legal Section -->
        <div>
          <h3 class="text-lg font-semibold mb-2">{{ $t('Legal') }}</h3>
          <router-link to="/terms" class="text-info hover:text-blue-400 block">
            <font-awesome-icon :icon="['far', 'file-lines']" />
            {{ $t('Terms of Service') }}
          </router-link>
          <router-link to="/privacy" class="text-info hover:text-blue-400 block">
            <font-awesome-icon :icon="['fas', 'file-shield']" />
            {{ $t('Privacy Policy') }}
          </router-link>
        </div>

        <!-- Community Section -->
        <div>
          <h3 class="text-lg font-semibold mb-2">{{ $t('Community') }}</h3>
          <a href="https://discord.com/invite/NVV2ww22ek" target="_blank" class="text-info hover:text-blue-400 block">
            <i class="fab fa-discord"></i> {{ $t('Join our Discord') }}
          </a>
          <a href="https://github.com/Bount-ing/Bount.ing" target="_blank" class="text-info hover:text-blue-400 block">
            <i class="fab fa-github"></i> {{ $t('GitHub Repository') }}
          </a>
          <a href="https://github.com/Bount-ing/Bount.ing/issues/new" target="_blank" class="text-info hover:text-blue-400 block">
            <font-awesome-icon :icon="['fas', 'bug']" /> {{ $t('Report an Issue') }}
          </a>
        </div>

        <!-- Contact Section -->
        <div>
          <h3 class="text-lg font-semibold mb-2">{{ $t('Contact') }}</h3>
          <p class="text-info">
            <font-awesome-icon :icon="['fas', 'envelope']" class="mr-2" />
            <a href="mailto:support@bount.ing" class="hover:text-blue-400">support@bount.ing</a>
          </p>
          <p class="text-info">
            <font-awesome-icon :icon="['fas', 'envelope']" class="mr-2" />
            <a href="mailto:contact@bount.ing" class="hover:text-blue-400">contact@bount.ing</a>
          </p>
        </div>
      </div>
      <p class="text-lg font-semibold my-6">© 2025 Bount.ing. {{ $t('All rights reserved.') }}</p>

    </div>
  </footer>
</template>


<script lang="ts">
import { defineComponent } from 'vue'
import { RouterLink } from 'vue-router'

export default defineComponent({
  name: 'Footer',
  components: {
    RouterLink
  }
})
</script>
