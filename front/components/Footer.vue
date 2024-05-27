<template>
  <footer class="bg-secondary-dark text-primary-light py-4  min-w-screen">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-4">
      <p>© 2024 Bount.ing. All rights reserved.</p>
      <p>Follow us on
        <a href="https://twitter.com/yourTwitter" class="underline text-info">Twitter</a>,
        <a href="https://facebook.com/yourFacebook" class="underline text-info">Facebook</a>,
        <a href="https://instagram.com/yourInstagram" class="underline text-info">Instagram</a>
      </p>
      <p>
        <NuxtLink to="/terms" class="underline text-info">Terms of Service</NuxtLink> |
        <NuxtLink to="/privacy" class="underline text-info">Privacy Policy</NuxtLink>
      </p>
      <p>Contact us: <a href="mailto:support@bount.ing" class="underline text-info">support@bount.ing</a></p>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'Footer'
}
</script>
