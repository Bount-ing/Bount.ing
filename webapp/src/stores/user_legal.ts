import { defineStore } from "pinia";
import { ref } from "vue";

interface UserLegalData {
  taxId: string;
  documentType: string;
  documentCountry: string;

  isCompany: boolean;
  companyName: string;
  companyAddress: string;
  companyCity: string;
  companyState: string;
  companyZip: string;
  companyCountry: string;
}

export const useUserLegalStore = defineStore("userLegal", () => {
  // State
  const legalData = ref<UserLegalData | null>(null);

  // Actions
  function setLegalData(data: UserLegalData) {
    legalData.value = data;
  }

  function resetLegalData() {
    legalData.value = null;
  }

  async function fetchLegalData(userId: string) {
    try {
      const response = await fetch(`/api/users/${userId}/legal`);
      if (!response.ok) throw new Error("Failed to fetch legal data");
      legalData.value = await response.json();
    } catch (error) {
      console.error("Error fetching legal data:", error);
    }
  }

  return { legalData, setLegalData, resetLegalData, fetchLegalData };
});
