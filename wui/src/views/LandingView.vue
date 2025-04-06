<template>
  <section class="min-h-screen">
    <!-- Hero Section -->
    <div class="container mx-auto px-4 py-12">
      <div class="flex flex-col items-center text-center mb-12">
        <h1 class="text-4xl md:text-5xl font-bold mb-6">
          Fix, fund & sustain open source software
        </h1>
        <p class="text-xl text-primary max-w-2xl mb-8">
          Connect with projects that matter, contribute to the ecosystem, and build a more
          sustainable future for open source.
        </p>

        <!-- CTA Buttons -->
        <div class="flex flex-col sm:flex-row gap-4 mb-8">
          <router-link
          class="bg-primary text-white py-2 px-6 rounded-lg text-lg hover:bg-primary-light transition"
            to="/signin"
          >
            Sign In
          </router-link>
          <router-link
          class="bg-success text-white py-2 px-6 rounded-lg text-lg hover:bg-success-light transition"
            to="/signup"
          >
            Sign Up
          </router-link>
        </div>

        <!-- Secondary Actions -->
        <div class="flex flex-col sm:flex-row gap-4">
          <a
            href="https://github.com/Bount-ing/Bount.ing"
            class="flex items-center justify-center space-x-2 border border-gray-300 rounded-lg py-2 px-4 hover:text-primary hover:border-primary transition"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"
              />
            </svg>
            <span>Check Repo</span>
          </a>

          <a
            href="https://discord.gg/NVV2ww22ek"
            target="_blank"
            class="flex items-center justify-center space-x-2 border border-primary text-primary rounded-lg py-2 px-4 hover:bg-primary hover:text-white transition"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                d="M20.317 4.369a19.791 19.791 0 0 0-4.885-1.621.074.074 0 0 0-.078.037 13.875 13.875 0 0 0-.601 1.233 19.918 19.918 0 0 0-5.506 0 13.875 13.875 0 0 0-.601-1.233.077.077 0 0 0-.078-.037c-1.666.37-3.344.919-4.885 1.621a.069.069 0 0 0-.032.027C.533 9.053.093 13.594.385 18.074a.074.074 0 0 0 .031.051 19.875 19.875 0 0 0 5.989 3.05.073.073 0 0 0 .083-.026 13.564 13.564 0 0 0 1.107-1.789.074.074 0 0 0-.041-.107 12.35 12.35 0 0 1-1.746-.841.073.073 0 0 1-.007-.12c.116-.088.232-.18.344-.271a.075.075 0 0 1 .079-.012c3.684 1.672 7.658 1.672 11.317 0a.075.075 0 0 1 .081.01c.112.092.229.184.346.272a.073.073 0 0 1-.007.12 12.351 12.351 0 0 1-1.747.841.073.073 0 0 0-.04.107c.361.637.753 1.239 1.107 1.789a.073.073 0 0 0 .083.026 19.86 19.86 0 0 0 5.988-3.05.072.072 0 0 0 .031-.051c.307-4.503-.092-9.014-3.884-13.678a.07.07 0 0 0-.033-.027zm-12.33 9.139c-1.052 0-1.912-.967-1.912-2.156 0-1.19.847-2.157 1.912-2.157 1.061 0 1.918.967 1.912 2.157 0 1.19-.847 2.156-1.912 2.156zm7.027 0c-1.052 0-1.912-.967-1.912-2.156 0-1.19.847-2.157 1.912-2.157 1.061 0 1.918.967 1.912 2.157 0 1.19-.851 2.156-1.912 2.156z"
              />
            </svg>
            <span>Join Discord</span>
          </a>
        </div>
      </div>
    </div>

    <!-- Tabs Section -->
    <div class="bg-secondary shadow-md">
      <div class="container mx-auto px-4">
        <div class="flex border-b border-secondary-light">
          <button
            @click="activeTab = 'sponsor'"
            :class="[
              'py-4 px-6 text-lg font-medium transition-colors duration-200',
              activeTab === 'sponsor'
                ? 'text-primary border-b-2 border-primary'
                : 'text-secondary-light hover:text-primary-light'
            ]"
          >
            Sponsor
          </button>
          <button
            @click="activeTab = 'developers'"
            :class="[
              'py-4 px-6 text-lg font-medium transition-colors duration-200',
              activeTab === 'developers'
                ? 'text-primary border-b-2 border-primary'
                : 'text-secondary-light hover:text-primary-light'
            ]"
          >
            Developers
          </button>
        </div>
      </div>
    </div>

    <!-- Tab Content -->
    <div class="container mx-auto px-4 py-12">
      <div v-if="activeTab === 'sponsor'" class="grid md:grid-cols-2 gap-8">
        <div>
          <h2 class="text-3xl font-bold mb-4">Fund the projects you rely on</h2>
          <p class="text-lg text-gray-600 mb-6">
            Support maintainers building the tools that power your business. Create sustainable
            models for open source development.
          </p>
          <ul class="space-y-3">
            <li class="flex items-start">
              <svg
                class="h-6 w-6 text-success mr-2"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span>Support projects with monthly sponsorships</span>
            </li>
            <li class="flex items-start">
              <svg
                class="h-6 w-6 text-success mr-2"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span>Fund specific features or bug fixes</span>
            </li>
            <li class="flex items-start">
              <svg
                class="h-6 w-6 text-success mr-2"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span>Ensure the security and maintenance of critical dependencies</span>
            </li>
          </ul>
        </div>
        <div class="flex items-center justify-center">
          <div class="text-center">
            <svg
              class="h-32 w-32 text-success mx-auto mb-4"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="1"
                d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
            <h3 class="text-xl font-semibold">Sustainable Funding</h3>
            <p class="text-gray-600">Direct your resources where they matter most</p>
          </div>
        </div>
      </div>

      <div v-if="activeTab === 'developers'" class="grid md:grid-cols-2 gap-8">
        <div>
          <h2 class="text-3xl font-bold mb-4">Get rewarded for your contributions</h2>
          <p class="text-lg text-gray-600 mb-6">
            Work on projects you're passionate about while getting compensated for your time and
            expertise.
          </p>
          <ul class="space-y-3">
            <li class="flex items-start">
              <svg
                class="h-6 w-6 text-success mr-2"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span>Find bounties that match your skills</span>
            </li>
            <li class="flex items-start">
              <svg
                class="h-6 w-6 text-success mr-2"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span>Build your reputation in the open source community</span>
            </li>
            <li class="flex items-start">
              <svg
                class="h-6 w-6 text-success mr-2"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span>Collaborate with maintainers to improve critical software</span>
            </li>
          </ul>
        </div>
        <div class="flex items-center justify-center">
          <div class="text-center">
            <svg
              class="h-32 w-32 text-info mx-auto mb-4"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="1"
                d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"
              />
            </svg>
            <h3 class="text-xl font-semibold">Developer Contributions</h3>
            <p class="text-gray-600">Code, fix, and improve open source projects</p>
          </div>
        </div>
      </div>
    </div>

    <!-- FAQ Section -->
    <div class="py-12">
      <div class="container mx-auto px-4">
        <h2 class="text-3xl font-bold text-center mb-8 text-primary">Frequently Asked Questions</h2>

        <div class="max-w-3xl mx-auto space-y-4">
          <div
            v-for="(faq, index) in faqs"
            :key="index"
            class="border border-primary-dark rounded-lg overflow-hidden"
          >
            <button
              @click="toggleFaq(index)"
              class="flex justify-between items-center w-full p-4 text-left bg-secondary-dark hover:bg-secondary transition"
            >
              <span class="font-medium text-lg text-primary">{{ faq.question }}</span>
              <svg
                class="h-5 w-5 transform transition-transform duration-200"
                :class="openFaq === index ? 'rotate-180' : ''"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M19 9l-7 7-7-7"
                />
              </svg>
            </button>
            <div v-show="openFaq === index" class="p-4 border-t border-primary-dark">
              <p class="text-secondary-light">{{ faq.answer }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Discord CTA -->
    <div class="py-12">
      <div class="container mx-auto px-4 text-center">
        <h2 class="text-3xl font-semibold text-secondary-light mb-4">
          Any doubts? Ask on Discord!
        </h2>
        <div class="flex justify-center mx-4">
          <a
            href="https://discord.gg/NVV2ww22ek"
            target="_blank"
            class="text-primary hover:text-primary-light transition border rounded-2xl px-2 border-primary hover:border-primary-light"
          >
            <span
              class="inline-flex items-center space-x-2 text-xl text-primary hover:text-primary-light transition rounded-2xl py-2 px-2 font-bold"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  d="M20.317 4.369a19.791 19.791 0 0 0-4.885-1.621.074.074 0 0 0-.078.037 13.875 13.875 0 0 0-.601 1.233 19.918 19.918 0 0 0-5.506 0 13.875 13.875 0 0 0-.601-1.233.077.077 0 0 0-.078-.037c-1.666.37-3.344.919-4.885 1.621a.069.069 0 0 0-.032.027C.533 9.053.093 13.594.385 18.074a.074.074 0 0 0 .031.051 19.875 19.875 0 0 0 5.989 3.05.073.073 0 0 0 .083-.026 13.564 13.564 0 0 0 1.107-1.789.074.074 0 0 0-.041-.107 12.35 12.35 0 0 1-1.746-.841.073.073 0 0 1-.007-.12c.116-.088.232-.18.344-.271a.075.075 0 0 1 .079-.012c3.684 1.672 7.658 1.672 11.317 0a.075.075 0 0 1 .081.01c.112.092.229.184.346.272a.073.073 0 0 1-.007.12 12.351 12.351 0 0 1-1.747.841.073.073 0 0 0-.04.107c.361.637.753 1.239 1.107 1.789a.073.073 0 0 0 .083.026 19.86 19.86 0 0 0 5.988-3.05.072.072 0 0 0 .031-.051c.307-4.503-.092-9.014-3.884-13.678a.07.07 0 0 0-.033-.027zm-12.33 9.139c-1.052 0-1.912-.967-1.912-2.156 0-1.19.847-2.157 1.912-2.157 1.061 0 1.918.967 1.912 2.157 0 1.19-.847 2.156-1.912 2.156zm7.027 0c-1.052 0-1.912-.967-1.912-2.156 0-1.19.847-2.157 1.912-2.157 1.061 0 1.918.967 1.912 2.157 0 1.19-.851 2.156-1.912 2.156z"
                />
              </svg>
              <span class="font-bold">Join us on Discord!</span>
            </span>
          </a>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'LandingView',
  setup() {
    const activeTab = ref('sponsor')
    const openFaq = ref(null)

    const faqs = ref([
      // General
      {
        question: 'What is Bount.ing?',
        answer:
          'Bount.ing is an open-source bounty platform where users can sponsor and/or resolve GitHub issues. Sponsors offer financial rewards for issue resolution, while developers work to solve them.'
      },
      {
        question: 'How does Bount.ing work?',
        answer:
          'Sponsors post bounties on GitHub issues by setting a reward. Developers solve the issue, submit a pull request, and receive payment upon approval by the sponsor.'
      },
      {
        question: 'Who can use Bount.ing?',
        answer:
          'Anyone with a GitHub account can use Bount.ing — whether to sponsor issues or to work on them as a developer.'
      },

      // Bounties
      {
        question: 'How do I create a bounty?',
        answer:
          'Find a GitHub issue, define a bounty amount, and publish it. The reward is charged at resolution, including our platform fee, VAT and Stripe processing fees.'
      },
      {
        question: 'Can I create a bounty for any GitHub issue?',
        answer:
          'Yes, as long as the issue is public and not restricted by the repository maintainers.'
      },
      {
        question: 'How are bounties priced?',
        answer:
          'Sponsors set the bounty amount, and Bount.ing charges a commission at the time of bounty payout.'
      },
      {
        question: 'What happens if multiple developers work on the same bounty?',
        answer:
          'The sponsor chooses which pull request to accept. Only the selected developer gets paid.'
      },
      {
        question: 'What if my pull request is rejected?',
        answer:
          'Sponsors are responsible for choosing a winner. If your PR is rejected, you will not receive the bounty, but you may still negotiate with the sponsor.'
      },

      // Payments
      {
        question: 'How are payments processed on Bount.ing?',
        answer:
          'Payments are processed directly through Stripe once a sponsor approves the solution. Bount.ing does not hold any funds.'
      },
      {
        question: 'What payment methods are supported?',
        answer:
          'Payments are processed through Stripe, supporting credit cards, bank transfers, and other methods enabled by Stripe.'
      },
      {
        question: 'When do I receive my payment?',
        answer:
          'Once a bounty is approved, Stripe processes the payment to the solver immediately. Payout time depends on Stripe’s processing times.'
      },
      {
        question: 'Are there any fees associated with payments?',
        answer:
          'Bount.ing charges a commission to the developer at bounty resolution. The comission fees includes VAT and Stripe processing fees.'
      },

      // Tips
      {
        question: 'Can I tip a developer for an issue without a bounty?',
        answer:
          'No, Bount.ing only supports payments for bounties. However, you can create a bounty for any issue you want to support.'
      },

      // Pricing
      {
        question: 'Is there a fee to use Bount.ing?',
        answer:
          'Creating an account is free. A commission is charged to developers when solving a bounty.'
      },
      {
        question: 'Are there subscription plans available?',
        answer:
          'Currently, Bount.ing does not offer subscription plans. It operates on a per-bounty commission model.'
      }
    ])

    const toggleFaq = (index) => {
      openFaq.value = openFaq.value === index ? null : index
    }

    return {
      activeTab,
      faqs,
      openFaq,
      toggleFaq
    }
  }
})
</script>

<style scoped>
/* Custom transitions for FAQ */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
