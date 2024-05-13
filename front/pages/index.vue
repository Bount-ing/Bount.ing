<template>
  <section class="min-h-screen flex items-center justify-center p-6">
    <div class="w-full max-w-5xl bg-white rounded-2xl shadow-xl overflow-hidden">
      <div class="px-6 md:px-10 lg:px-16 py-8 space-y-6">
        <h1 class="text-3xl md:text-4xl font-bold text-gray-900 text-center">Empower Open Source Innovation</h1>
        <p class="text-md md:text-lg text-gray-700 text-center">
          Enhance your coding skills, foster innovation, and contribute to the growth of open source projects while earning rewards.
        </p>

        <!-- Tabs for different sections -->
        <div class="bg-gray-50 p-4 rounded-lg">
          <nav class="flex justify-center space-x-2">
            <button
              class="flex-grow py-3 text-sm md:text-base font-semibold text-center text-blue-600 bg-white rounded-md shadow-sm"
              :class="{'bg-blue-100': currentTab !== 'whyInvest'}"
              @click="currentTab = 'whyInvest'">
              Why Invest
            </button>
            <button
              class="flex-grow py-3 text-sm md:text-base font-semibold text-center text-blue-600 bg-white rounded-md shadow-sm"
              :class="{'bg-blue-100': currentTab !== 'careerGrowth'}"
              @click="currentTab = 'careerGrowth'">
              Career Growth
            </button>
          </nav>

          <div class="mt-6 p-6 bg-white rounded-md shadow">
            <ul v-if="currentTab === 'whyInvest'" class="space-y-3 text-gray-800">
              <li v-for="item in reasonsContent" :key="item" class="flex items-start">
                <span class="flex-shrink-0 mt-1">
                  <svg class="h-6 w-6 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                </span>
                <p class="ml-3 text-sm md:text-base">{{ item }}</p>
              </li>
            </ul>
            <ul v-else-if="currentTab === 'careerGrowth'" class="space-y-3 text-gray-800">
              <li v-for="item in careerContent" :key="item" class="flex items-start">
                <span class="flex-shrink-0 mt-1">
                  <svg class="h-6 w-6 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                </span>
                <p class="ml-3 text-sm md:text-base">{{ item }}</p>
              </li>
            </ul>
          </div>
        </div>

        <!-- Call to action -->
        <div class="text-center">
          <button class="bg-blue-600 hover:bg-blue-700 text-white font-semibold py-3 px-10 rounded-lg transition duration-300">
            Start Funding Today
          </button>
        </div>
      </div>
    </div>
  </section>
</template>


<script>
export default {
  data() {
    return {
      currentTab: 'whyInvest',
      "reasonsContent": [
    "Accelerate progress: Fund specific issues to speed up development of features and bug fixes.",
    "Attract top talent: Support the project financially to attract skilled developers who deliver robust, high-quality contributions.",
    "Empower the community: Help sustain a vibrant community of developers by financially supporting regular contributors.",
    "Spark innovation: Encourage experimental solutions and new features by funding challenging projects that push boundaries.",
    "Prioritize development: Ensure critical features are developed first by directing funds to issues deemed most essential by the community.",
    "Be a leader in open source: Enhance your company's reputation as a champion of open innovation and collaboration.",
    "Discover hidden gems: Identify and potentially recruit top developers who have a proven track record of solving problems in your technology area.",
    "Shape the future: Gain influence over the development direction of projects you care about by actively participating in the community.",
    "Mitigate security risks: By contributing to open source projects, you can help improve the security posture of software you (and potentially millions of others) rely on. Open source allows for transparent review and collaboration, leading to faster identification and patching of vulnerabilities."
  ],
  "careerContent": [
    "Level up your skills: Gain valuable real-world software development experience through active project contributions.",
    "Get rewarded: Earn recognition and potentially financial rewards while honing your coding abilities by working on real projects.",
    "Learn from the best: Collaborate with and learn from experienced developers, and contribute to the mentoring of newcomers in the community.",
    "Make a lasting impact:  Ensure the long-term health and longevity of vital open-source projects that power the software we all use.",
    "Shape the software you use: Directly influence the development direction of projects that matter to you and impact their functionality.",
    "Build your portfolio: Showcase your contributions to a large user base, potentially boosting your resume and career prospects.",
    "Expand your network: Connect with talented developers from around the world and build a network of professionals in your field."
  ],
      loading: false
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
