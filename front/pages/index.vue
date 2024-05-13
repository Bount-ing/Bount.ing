<template>
  <section class="min-h-screen bg-gradient-to-br from-blue-50 to-blue-200 flex items-center justify-center p-6">
    <div class="max-w-4xl bg-white rounded-lg shadow-xl overflow-hidden">
      <div class="p-6 sm:p-10 space-y-6">
        <h1 class="text-3xl font-bold text-gray-800 text-center">Invest in Open Source Issues</h1>

        <p class="text-gray-600 text-lg text-center">
          Discover the power of community-driven development by funding issues. Enhance your skills, support innovation, and earn while contributing.
        </p>

        <!-- Reasons accordion -->
        <div class="space-y-4">
          <reason-accordion title="Why Put Money into Open Source" :content="reasonsContent"></reason-accordion>
          <reason-accordion title="Earning While Learning" :content="educationContent"></reason-accordion>
        </div>

        <!-- Call to action -->
        <div class="text-center">
          <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg transition duration-300">
            Get Started
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'FundingPage',
  components: {
    ReasonAccordion: () => import('../components/ReasonAccordion') // Lazy-loaded component
  },
  data() {
    return {
      reasonsContent: [
        "Speed up enhancements and innovations by funding specific issues.",
        "Attract skilled developers who deliver robust, high-quality contributions.",
        "Help sustain the community and empower regular contributors financially.",
        "Encourage experimental solutions and new features by funding challenging projects.",
        "Ensure important features are developed by directing funds to critical issues.",
        "Enhance your company's reputation as a supporter of open innovation.",
        "Identify and recruit top developers who are proven problem solvers in your technology area."
      ],
      educationContent: [
        "Gain real-world software development skills through active project contributions.",
        "Earn rewards while improving your coding skills and working on real projects.",
        "Learn from experienced developers and contribute to the mentoring of newcomers.",
        "Ensure the longevity of vital open-source projects through your contributions.",
        "Directly impact the development direction of projects that matter to you."
      ]
    };
  },
  name: 'Home',
  methods: {
    loginWithGithub() {
      console.log("Attempting to log in with GitHub");
      this.$auth.loginWith('github').then(() => {
        console.log("Logged in with GitHub");
      });
    }
  }
}
</script>
