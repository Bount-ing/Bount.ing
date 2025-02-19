import { defineStore } from "pinia";
import { ref } from "vue";
import { api } from "./api";

interface LegalEntity {

  isCompany: boolean;

  documentType: string;
  documentCountry: string;
  documentNumber: string;

  legalName: string;
  legalAddress: string;
  legalCity: string;
  legalZip: string;
  legalState: string;
  legalCountry: string;
  confirmedAt: string;
}

export const useLegalEntityStore = defineStore("userLegal", () => {
  // State
  const legalData = ref<LegalEntity | null>(null);

  // Actions
  function setLegalData(data: LegalEntity) {
    try {
      api.put(`/v1/users/me/legal`, data);
      legalData.value = data;
    } catch (error) {
      console.error("Error setting legal data:", error);
    }
  }

  function resetLegalData() {
    legalData.value = null;
  }

  async function fetchLegalData() {
    try {
      const response = await fetch(`/api/users/me/legal`);
      if (!response.ok) throw new Error("Failed to fetch legal data");
      legalData.value = await response.json();
    } catch (error) {
      console.error("Error fetching legal data:", error);
    }
  }

  return {
    legalData,
    setLegalData,
    resetLegalData,
    fetchLegalData
  };
});
