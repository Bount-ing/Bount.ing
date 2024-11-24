<template>
  <div>
    <div v-if="unconfirmedBounties.length">
      <BountyCard
        v-for="bounty in unconfirmedBounties"
        :key="bounty.ID"
        :bounty="bounty"
        @confirm-bounty="finalizeBounty(bounty.ID)"
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
    const unconfirmedBounties = ref([]);
    const baseURL = import.meta.env.VITE_API_BASE_URL;

	  if (!baseURL) {
		  console.error("API base URL is not set.");
		  loading.value = false;
		  return;
	  }


    const fetchBounties = async () => {
      try {
        const response = await axios.get(`${baseURL}/api/v1/admin/bounties/unconfirmed`, { headers: {
          Authorization: userStore.authHeader,
          RefererPolicy: 'origin-when-cross-origin'
        } });
        console.log('Bounties:', response.data);
        console.log('User ID:', userStore.user.userid);
        unconfirmedBounties.value = response.data
      } catch (error) {
        console.error('Error fetching bounties:', error);
      }
    };

    const finalizeBounty = async (bountyId) => {
      if (!bountyId) {
        console.error('Bounty ID is undefined or invalid.');
        return;
      }
      try {
        await axios.put(`${baseURL}/api/v1/admin/bounties/finalize/${bountyId}`, {
          headers: { Authorization: userStore.authHeader },
          RefererPolicy: 'origin-when-cross-origin'
        });
        unconfirmedBounties.value = unconfirmedBounties.value.filter(bounty => bounty.ID !== bountyId);
      } catch (error) {
        console.error('Error confirming bounty:', error);
      }
	}

    const cancelBounty = async (bountyId) => {
      if (!bountyId) {
        console.error('Bounty ID is undefined or invalid.');
        return;
      }
      try {
        await axios.delete(`${baseURL}/api/v1/bounties/${bountyId}`, {
          headers: { Authorization: userStore.authHeader },
          RefererPolicy: 'origin-when-cross-origin'
        });
        unconfirmedBounties.value = unconfirmedBounties.value.filter(bounty => bounty.ID !== bountyId);
      } catch (error) {
        console.error('Error canceling bounty:', error);
      }
    };

    onMounted(fetchBounties);

    return { unconfirmedBounties, cancelBounty, finalizeBounty };
  },
};
</script>
