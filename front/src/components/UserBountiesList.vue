<template>
  <div>
    <div v-if="filteredBounties.length">
      <BountyCard
        v-for="bounty in filteredBounties"
        :key="bounty.ID"
        :bounty="bounty"
        @cancel-bounty="cancelBounty(bounty.ID)"
      />
    </div>
    <p v-else class="text-gray-400">{{ $t('No bounties found.') }}</p>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useUserStore } from '../stores/user';
import BountyListItem from '../components/BountyListItem.vue'; // Assuming the card component is extracted

export default {
  components: {
    BountyCard: BountyListItem,
  },
  setup() {
    const userStore = useUserStore();
    const filteredBounties = ref([]);
    const baseURL = import.meta.env.VITE_API_BASE_URL;

    const fetchBounties = async () => {
      try {
        const response = await axios.get(`${baseURL}/api/v1/bounties/`, {
          headers: { Authorization: userStore.authHeader },
        });
        console.log('Bounties:', response.data);
        console.log('User ID:', userStore.user.userid);
        filteredBounties.value = response.data.filter(bounty => bounty.owner_id == userStore.user.userid);
      } catch (error) {
        console.error('Error fetching bounties:', error);
      }
    };

    const cancelBounty = async (bountyId) => {
      if (!bountyId) {
        console.error('Bounty ID is undefined or invalid.');
        return;
      }
      try {
        await axios.delete(`${baseURL}/api/v1/bounties/${bountyId}`, {
          headers: { Authorization: userStore.authHeader },
        });
        filteredBounties.value = filteredBounties.value.filter(bounty => bounty.ID !== bountyId);
      } catch (error) {
        console.error('Error canceling bounty:', error);
      }
    };

    onMounted(fetchBounties);

    return { filteredBounties, cancelBounty };
  },
};
</script>
