<template>
  <div class="mb-8">
    <h2 class="text-2xl font-semibold mb-4">Import GitHub Issue</h2>
    <div class="bg-secondary p-6 rounded-lg shadow-lg">
      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <label for="issueUrl" class="block text-sm font-medium mb-2">GitHub Issue URL</label>
          <input
            v-model="issueUrl"
            id="issueUrl"
            type="url"
            required
            placeholder="https://github.com/owner/repo/issues/123"
            class="border px-4 py-2 w-full rounded-lg bg-secondary focus:ring-2 focus:ring-green-500"
          />
        </div>

        <!-- Error Alert -->
        <div
          v-if="error"
          class="bg-red-500 bg-opacity-10 border border-red-500 text-red-500 px-4 py-2 rounded-lg"
        >
          {{ error }}
        </div>

        <button
          type="submit"
          :disabled="loading"
          class="px-6 py-3 bg-primary text-white rounded-lg hover:bg-primary-light transition-all focus:outline-none focus:ring-2 focus:ring-green-500 disabled:opacity-50"
        >
          {{ loading ? 'Loading...' : 'Import Issue' }}
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { api } from '@/stores/api'
import { useUserStore } from '@/stores/user'
import { useErrorStore } from '@/stores/errors'

const userStore = useUserStore()
const errorStore = useErrorStore()

const emit = defineEmits(['issueImported'])

const issueUrl = ref('')
const loading = ref(false)
const error = ref('')

const validateGitHubUrl = (url) => {
  const githubIssueRegex = /^https:\/\/github\.com\/[\w-]+\/[\w.-]+\/issues\/\d+$/;
  return githubIssueRegex.test(url)
}

const handleSubmit = async () => {
  error.value = ''
  loading.value = true

  try {
    // Validate GitHub URL format
    if (!validateGitHubUrl(issueUrl.value)) {
      throw new Error('Please enter a valid GitHub issue URL')
    }

    // Encode the URL in base64
    const encodedUrl = btoa(issueUrl.value)
    
    // Send request to backend with encoded URL in path
    const { data: importedIssue } = await api.post(`/v1/import-issue-by-url/${encodedUrl}`)

    // Update the store with the new issue if needed
    if (userStore.importedIssues) {
      const existingIssueIndex = userStore.importedIssues.findIndex(
        issue => issue.URL === issueUrl.value
      )

      if (existingIssueIndex >= 0) {
        // Update existing issue
        userStore.importedIssues[existingIssueIndex] = importedIssue
      } else {
        // Add new issue
        userStore.importedIssues = [...userStore.importedIssues, importedIssue];
      }
    }

    // Emit event to parent component
    emit('issueImported', importedIssue)

    // Clear form
    issueUrl.value = ''
    
  } catch (err) {
    // Handle error from the API response
    if (err.response?.data?.error) {
      error.value = err.response.data.error
    } else {
      error.value = err.message || 'Failed to import GitHub issue'
    }
  } finally {
    loading.value = false
  }
}
</script>