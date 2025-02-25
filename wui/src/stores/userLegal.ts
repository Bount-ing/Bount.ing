import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from './api'

interface LegalEntity {
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
  confirmedAt: string
}

export const useLegalEntityStore = defineStore('userLegal', () => {
  // State
  const legalData = ref<LegalEntity>({
    isCompany: false,
    documentType: '',
    documentCountry: '',
    documentNumber: '',
    legalName: '',
    legalAddress: '',
    legalCity: '',
    legalZip: '',
    legalState: '',
    legalCountry: '',
    confirmedAt: ''
  })

  // Actions
  function setLegalData(data: LegalEntity) {
    try {
      api.put(`/v1/users/me/legal`, data)
      legalData.value = data
    } catch (error) {
      console.error('Error setting legal data:', error)
    }
  }

  function resetLegalData() {
    legalData.value = {
      isCompany: false,
      documentType: '',
      documentCountry: '',
      documentNumber: '',
      legalName: '',
      legalAddress: '',
      legalCity: '',
      legalZip: '',
      legalState: '',
      legalCountry: '',
      confirmedAt: ''
    }
  }

  async function fetchLegalData() {
    try {
      const response = await api.get(`/v1/users/me/legal`)
      legalData.value.isCompany = response.data.IsCompany
      legalData.value.documentType = response.data.DocumentType
      legalData.value.documentCountry = response.data.DocumentCountry
      legalData.value.documentNumber = response.data.DocumentNumber
      legalData.value.legalName = response.data.LegalName
      legalData.value.legalAddress = response.data.LegalAddress
      legalData.value.legalCity = response.data.LegalCity
      legalData.value.legalZip = response.data.LegalZip
      legalData.value.legalState = response.data.LegalState
      legalData.value.legalCountry = response.data.LegalCountry
    } catch (error) {
      console.error('Error fetching legal data:', error)
    }
  }

  return {
    legalData,
    setLegalData,
    resetLegalData,
    fetchLegalData
  }
})
