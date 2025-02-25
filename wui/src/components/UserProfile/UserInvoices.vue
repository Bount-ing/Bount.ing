<template>
  <div class="p-6 bg-secondary-dark rounded-xl shadow-lg text-white">
    <h2 class="text-2xl font-semibold mb-4">{{ t('profile.legal_data') }}</h2>

    <form @submit.prevent="handleSubmit">
      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="isCompany">
          {{ t('profile.legal_data_business_type') }}
        </label>
        <select
          v-model="formData.isCompany"
          id="isCompany"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          @change="updateCompanyFields"
        >
          <option :value="true">{{ t('profile.legal_data_company') }}</option>
          <option :value="false">{{ t('profile.legal_data_individual') }}</option>
        </select>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="documentType">
          {{ t('profile.legal_data_tax_id_type') }}
        </label>
        <select
          v-model="formData.documentType"
          id="documentType"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          @change="handleDocumentTypeChange"
        >
          <option v-for="type in availableDocumentTypes" :key="type.value" :value="type.value">
            {{ type.label }}
          </option>
        </select>
        <small class="text-gray-400">{{ currentDocumentTypeDescription }}</small>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="documentNumber">
          {{ t('profile.legal_data_tax_id') }}
        </label>
        <input
          v-model="formData.documentNumber"
          id="documentNumber"
          type="text"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          required
          @input="onDocumentNumberInput"
        />
        <div v-if="documentError" class="text-red-500 text-sm mt-1">{{ documentError }}</div>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="legalName">
          {{ t('profile.legal_data_business_name') }}
        </label>
        <input
          v-model="formData.legalName"
          id="legalName"
          type="text"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          required
        />
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="legalAddress">
          {{ t('profile.legal_data_address') }}
        </label>
        <input
          v-model="formData.legalAddress"
          id="legalAddress"
          type="text"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          required
        />
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div class="mb-4">
          <label class="block text-sm font-medium mb-1" for="legalCity">
            {{ t('profile.legal_data_city') }}
          </label>
          <input
            v-model="formData.legalCity"
            id="legalCity"
            type="text"
            class="w-full p-2 rounded-md bg-gray-800 text-white"
            required
          />
        </div>
        <div class="mb-4">
          <label class="block text-sm font-medium mb-1" for="legalZip">
            {{ t('profile.legal_data_postal_code') }}
          </label>
          <input
            v-model="formData.legalZip"
            id="legalZip"
            type="text"
            class="w-full p-2 rounded-md bg-gray-800 text-white"
            required
          />
        </div>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="legalState">
          {{ t('profile.legal_data_state') }}
        </label>
        <input
          v-model="formData.legalState"
          id="legalState"
          type="text"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          required
        />        
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="legalCountry">
          {{ t('profile.legal_data_country') }}
        </label>
        <select
          v-model="formData.legalCountry"
          id="legalCountry"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          required
        >
          <option v-for="country in countries" :key="country.value" :value="country.value">
            {{ country.label }}
          </option>
        </select>
      </div>

      <div class="mb-4">
        <input v-model="formData.confirmed" id="confirmed" type="checkbox" class="mr-2" required />
        <label class="text-sm font-medium" for="confirmed">
          {{ t('profile.legal_data_confirm') }}
        </label>
      </div>

      <button
        type="submit"
        class="bg-primary px-4 py-2 rounded-md shadow-md text-white font-semibold"
        :disabled="isSubmitting"
      >
        {{ t('profile.save') }}
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useVuelidate } from '@vuelidate/core'
import { required, helpers } from '@vuelidate/validators'
import { useLegalEntityStore } from '@/stores/userLegal'
import { useI18n } from 'vue-i18n'
import { useNotificationStore } from '@/stores/notification'

const notificationStore = useNotificationStore()

// Types
interface FormData {
  isCompany: boolean
  documentType: string
  documentCountry: string
  documentNumber: string
  legalName: string
  legalAddress: string
  legalCity: string
  legalZip: string
  legalState: string
  legalCountry: string
  confirmed: boolean
}

// Composables
const { t } = useI18n()
const userLegalStore = useLegalEntityStore()

// State
const isSubmitting = ref(false)
const documentError = ref('')
const confirmed = ref(false)

const formData = ref<FormData>({
  isCompany: false,
  documentType: '',
  documentCountry: 'ES',
  documentNumber: '',
  legalName: '',
  legalAddress: '',
  legalCity: '',
  legalZip: '',
  legalState: '',
  legalCountry: 'ES',
  confirmed: false
})

// Validation rules
const rules = computed(() => ({
  documentType: { required: helpers.withMessage(t('validation.required'), required) },
  documentNumber: { required: helpers.withMessage(t('validation.required'), required) },
  legalName: { required: helpers.withMessage(t('validation.required'), required) },
  legalAddress: { required: helpers.withMessage(t('validation.required'), required) },
  legalCity: { required: helpers.withMessage(t('validation.required'), required) },
  legalZip: { required: helpers.withMessage(t('validation.required'), required) },
  legalState: { required: helpers.withMessage(t('validation.required'), required) },
  legalCountry: { required: helpers.withMessage(t('validation.required'), required) },
  confirmed: { required: helpers.withMessage(t('validation.must_confirm'), required) }
}))

const v$ = useVuelidate(rules, formData)

// Computed
const countries = computed(() => [
  { value: 'ES', label: 'Spain' },
  { value: 'FR', label: 'France' },
  { value: 'DE', label: 'Germany' }
])

const availableDocumentTypes = computed(() => {
  const types = formData.value.isCompany
    ? [
        { value: 'NIF', label: t('profile.document_type.nif') },
        { value: 'CIF', label: t('profile.document_type.cif') },
        { value: 'OTHER', label: t('profile.document_type.other') }
      ]
    : [
        { value: 'DNI', label: t('profile.document_type.dni') },
        { value: 'NIE', label: t('profile.document_type.nie') },
        { value: 'PASSPORT', label: t('profile.document_type.passport') },
        { value: 'OTHER', label: t('profile.document_type.other') }
      ]
  return types
})

const currentDocumentTypeDescription = computed(() =>
  t(`profile.document_type_description.${formData.value.documentType.toLowerCase()}`)
)

const documentNumberPlaceholder = computed(() =>
  t(`profile.document_number_placeholder.${formData.value.documentType.toLowerCase()}`)
)

const isValidDocument = computed(() => {
  //Check for empty fields
  if (!formData.value.documentType) {
    return false
  }
  if (!formData.value.documentNumber || formData.value.documentNumber === '') {
    return false
  }
  if (!formData.value.legalName || formData.value.legalName === '') {
    return false
  }
  if (!formData.value.legalAddress || formData.value.legalAddress === '') {
    return false
  }
  if (!formData.value.legalCity || formData.value.legalCity === '') {
    return false
  }
  if (!formData.value.legalZip || formData.value.legalZip === '') {
    return false
  }
  if (!formData.value.legalState || formData.value.legalState === '') {
    return false
  }
  if (!formData.value.legalCountry || formData.value.legalCountry === '') {
    return false
  }
  if (!formData.value.confirmed) {
    return false
  }
  return true
})

const isFormValid = computed(() => isValidDocument.value)

// Methods
const updateCompanyFields = () => {
  formData.value.documentType = formData.value.isCompany
    ? formData.value.legalCountry === 'ES'
      ? 'NIF'
      : 'OTHER'
    : formData.value.legalCountry === 'ES'
      ? 'DNI'
      : 'OTHER'
}

const validateDocument = () => {
  return isValidDocument.value
}

const onCountryChange = (country: string) => {
  formData.value.documentCountry = country
  updateCompanyFields()
}

const handleDocumentTypeChange = () => {
  formData.value.documentNumber = ''
  documentError.value = ''
}

const onDocumentNumberInput = () => {
  validateDocument()
}

const handleSubmit = async () => {
  isSubmitting.value = true
  try {
    await userLegalStore.setLegalData({
      ...formData.value,
      confirmedAt: new Date().toISOString()
    })
    notificationStore.showNotification('profile.data_saved', 'success')
  } catch (error) {
    notificationStore.showNotification('profile.error_saving_data', 'error')
    console.error(error)
  } finally {
    isSubmitting.value = false
    formData.value.confirmed = false
  }
}

// Lifecycle
onMounted(async () => {
  await userLegalStore.fetchLegalData()
  if (userLegalStore.legalData) {
    formData.value = { ...formData.value, ...userLegalStore.legalData }
  }
})

watch(() => formData.value.legalCountry, onCountryChange)
</script>
