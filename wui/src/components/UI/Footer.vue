<template>
<!-- Footer -->
<footer class="bg-secondary text-white py-8">
  <div class="container mx-auto px-4">
    <!-- Top section with logo and social links -->
    <div class="flex flex-col md:flex-row justify-between items-center">
      <div class="mb-4 md:mb-0">
        <h3 class="text-xl font-bold">Bount.ing</h3>
        <p class="text-gray-400">Sustainable open source for everyone</p>
      </div>
      <div class="flex space-x-4">
        <!-- GitHub -->
        <a href="https://github.com/Bount-ing/Bount.ing" target="_blank" class="text-gray-400 hover:text-white transition">
          <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
            <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
          </svg>
        </a>
        <!-- Discord -->
        <a href="https://discord.gg/NVV2ww22ek" target="_blank" class="text-gray-400 hover:text-white transition">
          <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
            <path d="M20.317 4.369a19.791 19.791 0 0 0-4.885-1.621.074.074 0 0 0-.078.037 13.875 13.875 0 0 0-.601 1.233 19.918 19.918 0 0 0-5.506 0 13.875 13.875 0 0 0-.601-1.233.077.077 0 0 0-.078-.037c-1.666.37-3.344.919-4.885 1.621a.069.069 0 0 0-.032.027C.533 9.053.093 13.594.385 18.074a.074.074 0 0 0 .031.051 19.875 19.875 0 0 0 5.989 3.05.073.073 0 0 0 .083-.026 13.564 13.564 0 0 0 1.107-1.789.074.074 0 0 0-.041-.107 12.35 12.35 0 0 1-1.746-.841.073.073 0 0 1-.007-.12c.116-.088.232-.18.344-.271a.075.075 0 0 1 .079-.012c3.684 1.672 7.658 1.672 11.317 0a.075.075 0 0 1 .081.01c.112.092.229.184.346.272a.073.073 0 0 1-.007.12 12.351 12.351 0 0 1-1.747.841.073.073 0 0 0-.04.107c.361.637.753 1.239 1.107 1.789a.073.073 0 0 0 .083.026 19.86 19.86 0 0 0 5.988-3.05.072.072 0 0 0 .031-.051c.307-4.503-.092-9.014-3.884-13.678a.07.07 0 0 0-.033-.027zm-12.33 9.139c-1.052 0-1.912-.967-1.912-2.156 0-1.19.847-2.157 1.912-2.157 1.061 0 1.918.967 1.912 2.157 0 1.19-.847 2.156-1.912 2.156zm7.027 0c-1.052 0-1.912-.967-1.912-2.156 0-1.19.847-2.157 1.912-2.157 1.061 0 1.918.967 1.912 2.157 0 1.19-.851 2.156-1.912 2.156z" />
          </svg>
        </a>
        <!-- Twitter/X -->
        <a href="#" class="text-gray-400 hover:text-white transition">
          <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
            <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723 10.016 10.016 0 01-3.127 1.195A4.92 4.92 0 0016.953 2a4.927 4.927 0 00-4.926 4.926c0 .386.044.762.126 1.124a13.995 13.995 0 01-10.155-5.147 4.9 4.9 0 00-.667 2.475c0 1.71.87 3.214 2.19 4.097a4.886 4.886 0 01-2.23-.616v.061a4.925 4.925 0 003.95 4.828 4.847 4.847 0 01-1.297.174 4.92 4.92 0 01-.926-.086 4.931 4.931 0 004.596 3.42 9.885 9.885 0 01-6.116 2.105c-.398 0-.79-.023-1.175-.069a13.948 13.948 0 007.55 2.212c9.058 0 14.01-7.502 14.01-14.01 0-.214-.005-.426-.015-.639A9.93 9.93 0 0024 4.59z"/>
          </svg>
        </a>
        <!-- LinkedIn -->
        <a href="https://www.linkedin.com/company/bount-ing/" target="_blank" class="text-gray-400 hover:text-white transition">
          <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
            <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
          </svg>
        </a>
      </div>
    </div>
    
    <!-- Middle section with Legal and Contact only -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-8">
      <!-- Legal Section -->
      <div class="text-center md:text-left">
        <h3 class="text-lg font-semibold mb-3">Legal</h3>
        <a href="/terms" class="text-gray-400 hover:text-white transition block mb-2">
          <svg class="h-4 w-4 inline mr-2" fill="currentColor" viewBox="0 0 384 512">
            <path d="M320 464c8.8 0 16-7.2 16-16V160H256c-17.7 0-32-14.3-32-32V48H64c-8.8 0-16 7.2-16 16V448c0 8.8 7.2 16 16 16H320zM0 64C0 28.7 28.7 0 64 0H229.5c17 0 33.3 6.7 45.3 18.7l90.5 90.5c12 12 18.7 28.3 18.7 45.3V448c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V64z"/>
          </svg>
          Terms of Service
        </a>
        <a href="/privacy" class="text-gray-400 hover:text-white transition block mb-2">
          <svg class="h-4 w-4 inline mr-2" fill="currentColor" viewBox="0 0 512 512">
            <path d="M184 48H328c4.4 0 8 3.6 8 8V96H176V56c0-4.4 3.6-8 8-8zm-56 8V96H64C28.7 96 0 124.7 0 160v256c0 35.3 28.7 64 64 64H448c35.3 0 64-28.7 64-64V160c0-35.3-28.7-64-64-64H384V56c0-30.9-25.1-56-56-56H184c-30.9 0-56 25.1-56 56zM192 128H320c35.3 0 64 28.7 64 64v64c0 35.3-28.7 64-64 64H192c-35.3 0-64-28.7-64-64V192c0-35.3 28.7-64 64-64z"/>
          </svg>
          Privacy Policy
        </a>
      </div>
      
      <!-- Contact Section (now includes issue reporting) -->
      <div class="text-center md:text-right">
        <h3 class="text-lg font-semibold mb-3">Contact</h3>
        <p class="text-gray-400 mb-2">
          <svg class="h-4 w-4 inline mr-2" fill="currentColor" viewBox="0 0 512 512">
            <path d="M48 64C21.5 64 0 85.5 0 112c0 15.1 7.1 29.3 19.2 38.4L236.8 313.6c11.4 8.5 27 8.5 38.4 0L492.8 150.4c12.1-9.1 19.2-23.3 19.2-38.4c0-26.5-21.5-48-48-48H48zM0 176V384c0 35.3 28.7 64 64 64H448c35.3 0 64-28.7 64-64V176L294.4 339.2c-22.8 17.1-54 17.1-76.8 0L0 176z"/>
          </svg>
          <a href="mailto:support@bount.ing" class="hover:text-white transition">support@bount.ing</a>
        </p>
        <p class="text-gray-400 mb-2">
          <svg class="h-4 w-4 inline mr-2" fill="currentColor" viewBox="0 0 512 512">
            <path d="M48 64C21.5 64 0 85.5 0 112c0 15.1 7.1 29.3 19.2 38.4L236.8 313.6c11.4 8.5 27 8.5 38.4 0L492.8 150.4c12.1-9.1 19.2-23.3 19.2-38.4c0-26.5-21.5-48-48-48H48zM0 176V384c0 35.3 28.7 64 64 64H448c35.3 0 64-28.7 64-64V176L294.4 339.2c-22.8 17.1-54 17.1-76.8 0L0 176z"/>
          </svg>
          <a href="mailto:contact@bount.ing" class="hover:text-white transition">contact@bount.ing</a>
        </p>
        <p class="text-gray-400">
          <svg class="h-4 w-4 inline mr-2" fill="currentColor" viewBox="0 0 512 512">
            <path d="M256 0c53 0 96 43 96 96v3.6c0 15.7-12.7 28.4-28.4 28.4H188.4c-15.7 0-28.4-12.7-28.4-28.4V96c0-53 43-96 96-96zM41.4 105.4c12.5-12.5 32.8-12.5 45.3 0l64 64c.7 .7 1.3 1.4 1.9 2.1c14.2-7.3 30.4-11.5 47.8-11.5H312c17.3 0 33.5 4.1 47.8 11.5c.6-.7 1.2-1.4 1.9-2.1l64-64c12.5-12.5 32.8-12.5 45.3 0s12.5 32.8 0 45.3l-64 64c-.7 .7-1.4 1.3-2.1 1.9c6.2 12 10.1 25.3 11.5 39.6V480c0 17.7-14.3 32-32 32H128c-17.7 0-32-14.3-32-32V256.7c1.4-14.3 5.3-27.6 11.5-39.6c-.7-.6-1.4-1.2-2.1-1.9l-64-64c-12.5-12.5-12.5-32.8 0-45.3zM160 240c-8.8 0-16 7.2-16 16v32c0 8.8 7.2 16 16 16h32c8.8 0 16-7.2 16-16V256c0-8.8-7.2-16-16-16H160zm112 16c0-8.8 7.2-16 16-16h32c8.8 0 16 7.2 16 16v32c0 8.8-7.2 16-16 16H288c-8.8 0-16-7.2-16-16V256zM256 368v32c0 8.8-7.2 16-16 16H208c-8.8 0-16-7.2-16-16V368c0-8.8 7.2-16 16-16h32c8.8 0 16 7.2 16 16zm48-16c8.8 0 16 7.2 16 16v32c0 8.8-7.2 16-16 16H272c-8.8 0-16-7.2-16-16V368c0-8.8 7.2-16 16-16h32z"/>
          </svg>
          <a href="https://github.com/Bount-ing/Bount.ing/issues/new" target="_blank" class="hover:text-white transition">Report an Issue</a>
        </p>
      </div>
    </div>
    
    <!-- Bottom copyright section -->
    <div class="mt-8 text-center md:text-left text-gray-400">
      <p>&copy; 2025 Bount.ing - All rights reserved.</p>
    </div>
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
