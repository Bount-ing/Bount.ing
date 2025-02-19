<template>
  <div class="p-6 bg-secondary-dark rounded-xl shadow-lg text-white">
    <h2 class="text-2xl font-semibold mb-4">{{ t('profile.legal_data') }}</h2>

    <form @submit.prevent="saveLegalData">
      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="businessType">{{
          t('profile.legal_data_business_type')
        }}</label>
        <select
          v-model="legalData.businessType"
          id="businessType"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          @change="updateBusinessTypeFields"
        >
          <option value="company">{{ t('profile.legal_data_company') }}</option>
          <option value="individual">{{ t('profile.legal_data_individual') }}</option>
        </select>
      </div>

      <div class="mb-4" v-if="legalData.businessType === 'company'">
        <label class="block text-sm font-medium mb-1" for="businessName">{{
          t('profile.legal_data_business_name')
        }}</label>
        <input
          v-model="legalData.businessName"
          id="businessName"
          type="text"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          required
        />
      </div>

      <div class="mb-4" v-if="legalData.businessType === 'company'">
        <label class="block text-sm font-medium mb-1" for="vatNumber">{{
          t('profile.legal_data_vat_number')
        }}</label>
        <input
          v-model="legalData.vatNumber"
          id="vatNumber"
          type="text"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          required
        />
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="taxIdType">{{
          t('profile.legal_data_tax_id_type')
        }}</label>
        <select
          v-model="legalData.taxIdType"
          id="taxIdType"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
        >
          <option value="NIF">NIF</option>
          <option value="NIE">NIE</option>
          <option value="DNI">DNI</option>
          <option value="Passport">{{ t('profile.legal_data_passport') }}</option>
          <option value="Other">{{ t('profile.legal_data_other') }}</option>
        </select>
        <small class="text-gray-400">{{ getTaxIdTypeDescription() }}</small>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="taxId">{{
          t('profile.legal_data_tax_id')
        }}</label>
        <input
          v-model="legalData.taxId"
          id="taxId"
          type="text"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          required
        />
        <div v-if="taxIdError" class="text-red-500 text-sm mt-1">{{ taxIdError }}</div>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="email">{{
          t('profile.legal_data_email')
        }}</label>
        <input
          v-model="legalData.email"
          id="email"
          type="email"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          required
        />
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="phone">{{
          t('profile.legal_data_phone')
        }}</label>
        <input
          v-model="phoneInput"
          @input="formatPhoneNumber"
          id="phone"
          type="tel"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          required
        />
        <div v-if="phoneError" class="text-red-500 text-sm mt-1">{{ phoneError }}</div>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="address">{{
          t('profile.legal_data_address')
        }}</label>
        <input
          v-model="legalData.address"
          id="address"
          type="text"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          required
        />
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div class="mb-4">
          <label class="block text-sm font-medium mb-1" for="city">{{
            t('profile.legal_data_city')
          }}</label>
          <input
            v-model="legalData.city"
            id="city"
            type="text"
            class="w-full p-2 rounded-md bg-gray-800 text-white"
            required
          />
        </div>

        <div class="mb-4">
          <label class="block text-sm font-medium mb-1" for="postalCode">{{
            t('profile.legal_data_postal_code')
          }}</label>
          <input
            v-model="legalData.postalCode"
            id="postalCode"
            type="text"
            class="w-full p-2 rounded-md bg-gray-800 text-white"
            required
          />
        </div>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium mb-1" for="country">{{
          t('profile.legal_data_country')
        }}</label>
        <select
          v-model="legalData.country"
          id="country"
          class="w-full p-2 rounded-md bg-gray-800 text-white"
          required
          @change="onCountryChange"
        >
          <option v-for="country in countries" :key="country.code" :value="country.code">
            {{ country.name }}
          </option>
        </select>
      </div>

      <div class="mb-4">
        <input v-model="legalData.confirmed" id="confirmed" type="checkbox" class="mr-2" required />
        <label class="text-sm font-medium" for="confirmed">{{
          t('profile.legal_data_confirm')
        }}</label>
      </div>

      <button
        type="submit"
        class="bg-primary px-4 py-2 rounded-md shadow-md text-white font-semibold"
      >
        {{ t('profile.save') }}
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useUserStore } from '@/stores/user'
import { useI18n } from 'vue-i18n'
import { parsePhoneNumberFromString, isValidPhoneNumber } from 'libphonenumber-js'

const phoneInput = ref('');
const phoneError = ref('');

const { t } = useI18n()
const userStore = useUserStore()

const countries = ref([
  { code: 'ES', name: 'Spain' },
  { code: 'FR', name: 'France' },
  { code: 'DE', name: 'Germany' }
  // Add more countries as needed
])

const legalData = ref({
  businessType: userStore.user?.legalData?.businessType || 'individual',
  businessName: userStore.user?.legalData?.businessName || '',
  vatNumber: userStore.user?.legalData?.vatNumber || '',
  taxIdType: userStore.user?.legalData?.taxIdType || 'DNI',
  taxId: userStore.user?.legalData?.taxId || '',
  email: userStore.user?.legalData?.email || '',
  phone: userStore.user?.legalData?.phone || '',
  address: userStore.user?.legalData?.address || '',
  city: userStore.user?.legalData?.city || '',
  postalCode: userStore.user?.legalData?.postalCode || '',
  country: userStore.user?.legalData?.country || 'ES',
  confirmed: false
})

const taxIdError = ref('')

const updateBusinessTypeFields = () => {
  if (legalData.value.businessType === 'individual') {
    legalData.value.businessName = ''
    legalData.value.vatNumber = ''
  }
}

const getTaxIdTypeDescription = () => {
  switch (legalData.value.taxIdType) {
    case 'NIF':
      return t('profile.tax_id_description.nif')
    case 'NIE':
      return t('profile.tax_id_description.nie')
    case 'DNI':
      return t('profile.tax_id_description.dni')
    case 'Passport':
      return t('profile.tax_id_description.passport')
    default:
      return ''
  }
}

const onCountryChange = () => {
  // Reset tax ID type based on country if needed
  if (legalData.value.country === 'ES') {
    legalData.value.taxIdType = legalData.value.businessType === 'company' ? 'NIF' : 'DNI'
  } else {
    legalData.value.taxIdType = 'Other'
  }
}

const validateTaxId = () => {
  taxIdError.value = ''

  if (legalData.value.country === 'ES') {
    if (legalData.value.taxIdType === 'NIF' || legalData.value.taxIdType === 'DNI') {
      const dniRegex = /^[0-9]{8}[A-Z]$/
      if (!dniRegex.test(legalData.value.taxId)) {
        taxIdError.value = t('profile.error_invalid_dni_format')
        return false
      }
    } else if (legalData.value.taxIdType === 'NIE') {
      const nieRegex = /^[XYZ][0-9]{7}[A-Z]$/
      if (!nieRegex.test(legalData.value.taxId)) {
        taxIdError.value = t('profile.error_invalid_nie_format')
        return false
      }
    }
  }
  return true
}

const validateEmail = () => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(legalData.value.email)
}

const saveLegalData = async () => {
  if (!legalData.value.confirmed) {
    alert(t('profile.legal_data_must_confirm'))
    return
  }

  if (!validateTaxId()) {
    return
  }

  if (!validateEmail()) {
    alert(t('profile.error_invalid_email'))
    return
  }

  try {
    // Add timestamp for when data was confirmed
    const dataToSave = {
      ...legalData.value,
      confirmedAt: new Date().toISOString()
    }

    await userStore.updateUserLegalData(dataToSave)
    alert(t('profile.legal_data_saved'))
  } catch (error) {
    alert(t('profile.error_saving_data'))
    console.error(error)
  }
}

const formatPhoneNumber = () => {
  phoneError.value = '';
  
  if (!phoneInput.value) {
    legalData.value.phone = '';
    return;
  }
  
  try {
    const phoneNumber = parsePhoneNumberFromString(phoneInput.value, legalData.value.country);
    if (phoneNumber) {
      if (isValidPhoneNumber(phoneInput.value, legalData.value.country)) {
        legalData.value.phone = phoneNumber.format('E.164');
      } else {
        phoneError.value = t('profile.error_invalid_phone');
      }
    }
  } catch (error) {
    phoneError.value = t('profile.error_invalid_phone');
  }
};

onMounted(() => {
  // Pre-fill email if available from user store
  if (userStore.user?.email && !legalData.value.email) {
    legalData.value.email = userStore.user.email
  }
})

watch(() => legalData.value.country, () => {
  formatPhoneNumber();
});
</script>
