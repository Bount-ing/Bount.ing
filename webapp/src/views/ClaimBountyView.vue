<template>
  <div class="container mx-auto p-6">
    <div class="rounded-lg shadow-lg p-6">
      <!-- Page Header -->
      <h1 class="text-2xl font-bold text-primary mb-4">Claim Issue's Resolution</h1>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-4">
        Loading issue details...
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">
        {{ error }}
      </div>

      <!-- Content when data is loaded -->
      <template v-else>
        <!-- Issue Details -->
        <div class="bg-secondary p-4 rounded-md shadow-md mb-6">
          <div class="flex items-start gap-4">
            <img 
              v-if="issue?.avatarUrl" 
              :src="issue.avatarUrl" 
              :alt="issue.Title"
              class="w-12 h-12 rounded-full"
            />
            <div>
              <h2 class="text-lg font-semibold text-primary-light">{{ issue?.Title }}</h2>
              <p class="text-gray-300 mt-2" v-if="issue?.Description">
                {{ issue.Description || "No description provided." }}
              </p>
              <a 
                v-if="issue?.issueUrl" 
                :href="issue.issueUrl" 
                target="_blank" 
                rel="noopener noreferrer"
                class="text-primary hover:text-primary-dark mt-2 inline-block"
              >
                View Original Issue
              </a>
            </div>
          </div>
        </div>

        <!-- Displaying Bounties -->
        <div v-if="bounties?.length > 0" class="mt-6">
          <h2 class="text-lg font-semibold text-primary-light mb-4">Available Bounties</h2>
          <div class="space-y-4">
            <div 
              v-for="bounty in bounties" 
              :key="bounty.ID" 
              class="bg-secondary rounded-lg p-4 shadow-md"
            >
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <div class="flex items-baseline gap-2 mb-2">
                    <span class="text-2xl font-bold text-primary">
                      {{ bounty.amount }} {{ bounty.currency }}
                    </span>
                  </div>
                  <div class="text-gray-300 space-y-1">
                    <div class="flex items-center gap-2">
                      <span class="font-medium">Start Date:</span>
                      {{ formatDate(bounty.startAt) }}
                    </div>
                    <div class="flex items-center gap-2">
                      <span class="font-medium">End Date:</span>
                      {{ formatDate(bounty.endAt) }}
                    </div>
                    <div class="flex items-center gap-2">
                      <span class="font-medium">Status:</span>
                      <span class="capitalize">{{ bounty.status || 'Open' }}</span>
                    </div>
                  </div>
                </div>
                <div class="text-gray-300">
                  <div class="mb-2">
                    <span class="font-medium">Bounty ID:</span> #{{ bounty.ID }}
                  </div>
                  <div class="mb-2">
                    <span class="font-medium">Created:</span> {{ formatDate(bounty.CreatedAt) }}
                  </div>
                  <div v-if="bounty.finalized_at !== '0001-01-01T00:00:00Z'">
                    <span class="font-medium">Finalized:</span> {{ formatDate(bounty.finalized_at) }}
                  </div>
                </div>
              </div>

              <!-- Bounty Claimer and Owner Status -->
              <div class="mt-4">
                <h3 class="text-lg font-semibold text-primary-light">Bounty Claimer and Owner Status</h3>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-4">
                  <!-- Bounty Claimer Status -->
                  <div>
                    <h4 class="text-md font-semibold text-primary">Bounty Claimer</h4>
                    <div class="space-y-2">
                      <div class="flex items-center gap-2">
                        <span class="font-medium">PR has been found:</span>
                        <span v-if="bounty.claimer?.prFound" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                      <div class="flex items-center gap-2">
                        <span class="font-medium">PR links to the issue:</span>
                        <span v-if="bounty.claimer?.prLinked" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                      <div class="flex items-center gap-2">
                        <span class="font-medium">PR has been accepted:</span>
                        <span v-if="bounty.claimer?.prAccepted" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                      <div class="flex items-center gap-2">
                        <span class="font-medium">Issue has been closed:</span>
                        <span v-if="bounty.claimer?.issueClosed" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                    </div>
                  </div>

                  <!-- Bounty Owner Status -->
                  <div>
                    <h4 class="text-md font-semibold text-primary">Bounty Owner</h4>
                    <div class="space-y-2">
                      <div class="flex items-center gap-2">
                        <span class="font-medium">PR has been found:</span>
                        <span v-if="bounty.owner?.prFound" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                      <div class="flex items-center gap-2">
                        <span class="font-medium">PR links to the issue:</span>
                        <span v-if="bounty.owner?.prLinked" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                      <div class="flex items-center gap-2">
                        <span class="font-medium">PR has been accepted:</span>
                        <span v-if="bounty.owner?.prAccepted" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                      <div class="flex items-center gap-2">
                        <span class="font-medium">Issue has been closed:</span>
                        <span v-if="bounty.owner?.issueClosed" class="text-green-500">✔</span>
                        <span v-else class="text-red-500">❌</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Claim Form -->
        <form @submit.prevent="submitClaim" class="mt-6">
          <div class="mb-4">
            <label for="claimDetails" class="block text-gray-300 font-semibold mb-2">
              Claim Details
            </label>
            <textarea
              id="claimDetails"
              v-model="claimDetails"
              class="w-full bg-secondary-dark border border-gray-600 rounded-lg p-3 focus:ring-primary focus:border-primary text-gray-200"
              placeholder="Provide details about your claim..."
              rows="5"
            ></textarea>
          </div>

          <div class="flex justify-end">
            <button
              type="submit"
              :disabled="submitting"
              class="bg-success text-white py-2 px-4 rounded-lg shadow-md hover:bg-success-dark transition duration-200 disabled:opacity-50"
            >
              {{ submitting ? 'Submitting...' : 'Submit Claim' }}
            </button>
          </div>
        </form>
      </template>
    </div>
  </div>
</template>



<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { api } from '@/stores/api';

const route = useRoute();

// Add debug data ref
const debugData = ref({});

// Rest of the refs...
const issue = ref(null);
const bounties = ref([]);
const claimDetails = ref('');
const loading = ref(true);
const error = ref(null);
const submitting = ref(false);

// Format date helper remains the same...

// Updated fetch function with detailed logging
const fetchIssueDataAndBounties = async () => {
  const id = route.params.id;
  
  if (!id) {
    error.value = 'Missing issue ID';
    loading.value = false;
    return;
  }

  try {
    loading.value = true;
    error.value = null;

    const response = await api.get(`/v1/issues/${id}/bounties`);
    
    console.log('Raw API Response:', response);
    console.log('Response Data:', response.data);
    
    // Store debug data
    debugData.value = {
      responseData: response.data,
      responseStatus: response.status,
      endpoint: `/v1/issues/${id}/bounties`
    };

    if (!response?.data) {
      throw new Error('Invalid response from server');
    }

    // Log the structure of the response
    console.log('Response structure:', {
      hasIssue: !!response.data.issue,
      hasBounties: !!response.data.bounties,
      dataKeys: Object.keys(response.data)
    });

    issue.value = response.data;
    if (response.data.Bounties) {
      bounties.value = response.data.Bounties;
    } else if (Array.isArray(response.data)) {
      bounties.value = response.data;
    } else {
      console.warn('Unexpected bounties data structure:', response.data);
    }

    console.log('Processed Data:', {
      issue: issue.value,
      bounties: bounties.value
    });

  } catch (err) {
    console.error('Error details:', {
      message: err.message,
      response: err.response,
      stack: err.stack
    });
    error.value = 'Failed to load issue details. Please try again later.';
  } finally {
    loading.value = false;
  }
};

const formatDate = (date) => {
  if (!date) return 'N/A';
  return new Date(date).toLocaleString();
};


onMounted(() => {
  fetchIssueDataAndBounties();
});
</script>