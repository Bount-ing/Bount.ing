<template>
  <section class="bg-secondary-dark container mx-auto px-6 py-10 my-8 rounded-lg text-primary">
    <h1 class="text-5xl font-bold text-center mb-10">Bount.ing FAQ</h1>
    
    <p class="text-lg mb-12 text-center">
      Welcome to the Bount.ing FAQ page! Here, you'll find answers to common questions about how our platform works, from verifying contributions to handling disputes. If you have any further questions, feel free to reach out to our support team.
    </p>

    <div class="space-y-8">
      <div class="faq-item">
        <h3 class="text-2xl font-semibold mb-3">General Questions</h3>
        <div class="space-y-4 pl-4">
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">What is Bount.ing?</h4>
            <p class="text-lg mb-4">
              Bount.ing is a Catalan company that started on May 4th, 2024. Bount.ing is a bug bounty platform focusing on GitHub issues, not specifically security bugs. Our goals are to promote open-source development, create earning and learning opportunities, encourage independent and spontaneous initiatives, and establish a compensation system for developer collaboration.
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
              Getting started is simple. Sign up on our platform, connect your GitHub account, and start browsing issues. Whether you want to contribute to projects or sponsor issues, Bount.ing makes it easy to get involved in the open-source community.
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
              Bounties are paid out using Stripe, from the bounty emitter to the bounty claimer. When a bounty is fully resolved, it goes through the stages of being open, claimed, and then marked as resolved after the issue is closed. There is a fee of 4.2% on public repository bounties and 7% on private ones. If the claim is rejected, the transaction won't be processed.
            </p>
          </div>
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">Are there any fees to use Bount.ing?</h4>
            <p class="text-lg mb-4">
              Bount.ing charges a small service fee on bounties to support the platform's maintenance and development. This fee is transparently displayed during the bounty creation process.
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
              We have a robust verification process where we thoroughly review your pull request and the related GitHub issue. Our team ensures that the issue is resolved satisfactorily before crediting your account with the bounty.
            </p>
          </div>
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">What if my pull request is rejected?</h4>
            <p class="text-lg mb-4">
              If your pull request is rejected, you can ask for feedback from the project maintainers to understand what changes are needed. You can then make the necessary updates and resubmit your pull request.
            </p>
          </div>
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">What if there is a dispute over a bounty?</h4>
            <p class="text-lg mb-4">
              In case of a dispute, our moderation team steps in to review the case. We aim to provide a fair resolution based on the evidence and communication from both parties involved.
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
              We prioritize the security of your personal data. Our platform uses industry-standard security measures to protect your information. We do not share your data with third parties without your consent.
            </p>
          </div>
        </div>
      </div>

      <div class="faq-item">
        <h3 class="text-2xl font-semibold mb-3">Support</h3>
        <div class="space-y-4 pl-4">
          <div>
            <h4 class="text-xl font-semibold mb-2 cursor-pointer">What happens if I encounter a technical issue on the platform?</h4>
            <p class="text-lg mb-4">
              If you encounter any technical issues, our support team is here to help. You can reach out to us via our support page or contact us directly through our Discord community for assistance.
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
