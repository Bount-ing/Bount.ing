<template>
  <section class="bg-secondary-dark container mx-auto px-6 py-10 my-8 rounded-lg text-primary">
    <h1 class="text-5xl font-bold text-center mb-10">Bount.ing FAQ</h1>

    <p class="text-lg mb-12 text-center">
      Welcome to the Bount.ing FAQ page! Here, you'll find answers to common questions about how our platform works,
      from verifying contributions to handling disputes. If you have any further questions, feel free to reach out to
      our support team.
    </p>

    <div class="space-y-8">
      <div class="faq-item">
        <h3 class="text-2xl font-semibold mb-3">General Questions</h3>
        <div class="space-y-4 pl-4">
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">What is Bount.ing?</h4>
            <p class="text-lg mb-4">
              Bount.ing is a Catalan company that started on May 4th, 2024. Bount.ing is a bug bounty platform focusing
              on GitHub issues, not specifically security bugs. Our goals are to promote open-source development, create
              earning and learning opportunities, encourage independent and spontaneous initiatives, and establish a
              compensation system for developer collaboration.
            </p>
          </div>
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">Who can participate on Bount.ing?</h4>
            <p class="text-lg mb-4">
              Any organization, animal, and autognomes are welcomed to participate on the platform!
            </p>
          </div>
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">How do I get started with Bount.ing?</h4>
            <p class="text-lg mb-4">
              Getting started is simple. Sign up on our platform, connect your GitHub account, and start browsing
              issues. Whether you want to contribute to projects or sponsor issues, Bount.ing makes it easy to get
              involved in the open-source community.
            </p>
          </div>
        </div>
      </div>

      <div class="faq-item">
        <h3 class="text-2xl font-semibold mb-3">Payments and Fees</h3>
        <div class="space-y-4 pl-4">
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">How are bounties paid out?</h4>
            <p class="text-lg mb-4">
              Bounties are paid out using Stripe, from the bounty emitter to the bounty claimer. When a bounty is fully
              resolved, it goes through the stages of being open, claimed, and then marked as resolved after the issue
              is closed. There is a fee of 4.2% on public repository bounties and 10% on private ones. If the claim is
              rejected, the transaction won't be processed.
            </p>
          </div>
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">Are there any fees to use Bount.ing?</h4>
            <p class="text-lg mb-4">
              Bount.ing charges a small service fee on bounties to support the platform's maintenance and development.
              This fee is transparently displayed during the bounty creation process.
            </p>
          </div>
        </div>
      </div>

      <div class="faq-item">
        <h3 class="text-2xl font-semibold mb-3">Verification and Disputes</h3>
        <div class="space-y-4 pl-4">
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">How is my contribution verified?</h4>
            <p class="text-lg mb-4">
              We have a robust verification process where we thoroughly review your pull request and the related GitHub
              issue. Our team ensures that the issue is resolved satisfactorily before crediting your account with the
              bounty.
            </p>
          </div>
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">What if my pull request is rejected?</h4>
            <p class="text-lg mb-4">
              If your pull request is rejected, you can ask for feedback from the project maintainers to understand what
              changes are needed. You can then make the necessary updates and resubmit your pull request.
            </p>
          </div>
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">What if there is a dispute over a bounty?</h4>
            <p class="text-lg mb-4">
              In case of a dispute, our moderation team steps in to review the case. We aim to provide a fair resolution
              based on the evidence and communication from both parties involved.
            </p>
          </div>
        </div>
      </div>

      <div class="faq-item">
        <h3 class="text-2xl font-semibold mb-3">Security and Privacy</h3>
        <div class="space-y-4 pl-4">
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">Is my personal data safe?</h4>
            <p class="text-lg mb-4">
              We prioritize the security of your personal data. Our platform uses industry-standard security measures to
              protect your information. We do not share your data with third parties without your consent.
            </p>
          </div>
        </div>
      </div>

      <div class="faq-item">
        <h3 class="text-2xl font-semibold mb-3">Support</h3>
        <div class="space-y-4 pl-4">
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">What happens if I encounter a technical issue on the
              platform?</h4>
            <p class="text-lg mb-4">
              If you encounter any technical issues, our support team is here to help. You can reach out to us via our
              support page or contact us directly through our Discord community for assistance.
            </p>
            <div class="flex justify-start mx-4">
              <a href="https://discord.gg/NVV2ww22ek" target="_blank"
                class="text-primary hover:text-primary-light transition border rounded-2xl px-2 border-primary hover:border-primary-light mx-2">
                <span
                  class="inline-flex items-center space-x-2 text-xl text-primary hover:text-primary-light transition  rounded-2xl py-2 px-2 font-bold">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
                    <path
                      d="M20.317 4.369a19.791 19.791 0 0 0-4.885-1.621.074.074 0 0 0-.078.037 13.875 13.875 0 0 0-.601 1.233 19.918 19.918 0 0 0-5.506 0 13.875 13.875 0 0 0-.601-1.233.077.077 0 0 0-.078-.037c-1.666.37-3.344.919-4.885 1.621a.069.069 0 0 0-.032.027C.533 9.053.093 13.594.385 18.074a.074.074 0 0 0 .031.051 19.875 19.875 0 0 0 5.989 3.05.073.073 0 0 0 .083-.026 13.564 13.564 0 0 0 1.107-1.789.074.074 0 0 0-.041-.107 12.35 12.35 0 0 1-1.746-.841.073.073 0 0 1-.007-.12c.116-.088.232-.18.344-.271a.075.075 0 0 1 .079-.012c3.684 1.672 7.658 1.672 11.317 0a.075.075 0 0 1 .081.01c.112.092.229.184.346.272a.073.073 0 0 1-.007.12 12.351 12.351 0 0 1-1.747.841.073.073 0 0 0-.04.107c.361.637.753 1.239 1.107 1.789a.073.073 0 0 0 .083.026 19.86 19.86 0 0 0 5.988-3.05.072.072 0 0 0 .031-.051c.307-4.503-.092-9.014-3.884-13.678a.07.07 0 0 0-.033-.027zm-12.33 9.139c-1.052 0-1.912-.967-1.912-2.156 0-1.19.847-2.157 1.912-2.157 1.061 0 1.918.967 1.912 2.157 0 1.19-.847 2.156-1.912 2.156zm7.027 0c-1.052 0-1.912-.967-1.912-2.156 0-1.19.847-2.157 1.912-2.157 1.061 0 1.918.967 1.912 2.157 0 1.19-.851 2.156-1.912 2.156z" />
                  </svg>
                  <span class="font-bold ">Get support on Discord</span>
                </span>
              </a>
              <a href="https://github.com/Bount-ing/Bount.ing/issues/new" target="_blank"
                class="transition border rounded-2xl px-2 border-warning hover:border-warning-light mx-2">
                <span
                  class="inline-flex items-center space-x-2 text-xl text-warning hover:text-warning-light transition  rounded-2xl py-2 px-2 font-bold">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="currentColor" viewBox="0 0 496 512">
                    <path
                      d="M165.9 397.4c0 2-2.3 3.6-5.2 3.6-3.3 .3-5.6-1.3-5.6-3.6 0-2 2.3-3.6 5.2-3.6 3-.3 5.6 1.3 5.6 3.6zm-31.1-4.5c-.7 2 1.3 4.3 4.3 4.9 2.6 1 5.6 0 6.2-2s-1.3-4.3-4.3-5.2c-2.6-.7-5.5 .3-6.2 2.3zm44.2-1.7c-2.9 .7-4.9 2.6-4.6 4.9 .3 2 2.9 3.3 5.9 2.6 2.9-.7 4.9-2.6 4.6-4.6-.3-1.9-3-3.2-5.9-2.9zM244.8 8C106.1 8 0 113.3 0 252c0 110.9 69.8 205.8 169.5 239.2 12.8 2.3 17.3-5.6 17.3-12.1 0-6.2-.3-40.4-.3-61.4 0 0-70 15-84.7-29.8 0 0-11.4-29.1-27.8-36.6 0 0-22.9-15.7 1.6-15.4 0 0 24.9 2 38.6 25.8 21.9 38.6 58.6 27.5 72.9 20.9 2.3-16 8.8-27.1 16-33.7-55.9-6.2-112.3-14.3-112.3-110.5 0-27.5 7.6-41.3 23.6-58.9-2.6-6.5-11.1-33.3 2.6-67.9 20.9-6.5 69 27 69 27 20-5.6 41.5-8.5 62.8-8.5s42.8 2.9 62.8 8.5c0 0 48.1-33.6 69-27 13.7 34.7 5.2 61.4 2.6 67.9 16 17.7 25.8 31.5 25.8 58.9 0 96.5-58.9 104.2-114.8 110.5 9.2 7.9 17 22.9 17 46.4 0 33.7-.3 75.4-.3 83.6 0 6.5 4.6 14.4 17.3 12.1C428.2 457.8 496 362.9 496 252 496 113.3 383.5 8 244.8 8zM97.2 352.9c-1.3 1-1 3.3 .7 5.2 1.6 1.6 3.9 2.3 5.2 1 1.3-1 1-3.3-.7-5.2-1.6-1.6-3.9-2.3-5.2-1zm-10.8-8.1c-.7 1.3 .3 2.9 2.3 3.9 1.6 1 3.6 .7 4.3-.7 .7-1.3-.3-2.9-2.3-3.9-2-.6-3.6-.3-4.3 .7zm32.4 35.6c-1.6 1.3-1 4.3 1.3 6.2 2.3 2.3 5.2 2.6 6.5 1 1.3-1.3 .7-4.3-1.3-6.2-2.2-2.3-5.2-2.6-6.5-1zm-11.4-14.7c-1.6 1-1.6 3.6 0 5.9 1.6 2.3 4.3 3.3 5.6 2.3 1.6-1.3 1.6-3.9 0-6.2-1.4-2.3-4-3.3-5.6-2z" />
                  </svg>
                  <span class="font-bold ">Open a public issue on github</span>
                </span>
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
