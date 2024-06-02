<template>
    <section class="min-h-screen flex flex-col items-center justify-center p-4 text-white">
        <div v-if="loading" class="text-center text-2xl font-semibold">Loading...</div>
        <div v-else class="w-full max-w-6xl rounded-2xl shadow-xl overflow-hidden bg-gray-800">
            <!-- Profile Section -->
            <div class="flex flex-col items-center p-8 bg-black-900">
                <img :src="user.avatar" alt="Profile Picture" class="w-24 h-24 rounded-full border-2 border-gray-500"/>
                <h1 class="text-2xl font-semibold mt-2">{{ user.username }}</h1>
                <p class="text-gray-400 mt-1">{{ user.userBio }}</p>
                <button class="mt-3 px-3 py-1 bg-blue-600 text-white rounded-lg focus:outline-none">Edit Profile</button>
                <div class="mt-3 flex space-x-3">
                    <Badge :level="user.level" />
                    <Achievement :achievements="user.achievements" />
                </div>
            </div>

            <!-- Tab Navigation -->
            <div class="flex justify-around bg-gray-700 text-gray-300">
                <button @click="currentTab = 'personal'" :class="tabClass('personal')">Personal Info</button>
                <button @click="currentTab = 'activity'" :class="tabClass('activity')">Activity</button>
                <button @click="currentTab = 'github'" :class="tabClass('github')">GitHub Stats</button>
                <button @click="currentTab = 'bounties'" :class="tabClass('bounties')">Bounties</button>
                <button @click="currentTab = 'transactions'" :class="tabClass('transactions')">Transactions</button>
                <button @click="currentTab = 'payment'" :class="tabClass('payment')">Payment Info</button>
            </div>

            <!-- Tab Content -->
            <div class="p-6">
                <div v-if="currentTab === 'personal'">
                    <PersonalInfo :user="user" />
                    <AboutMe :aboutMe="user.aboutMe" />
                    <Interests :interests="user.interests" />
                </div>

                <div v-if="currentTab === 'activity'">
                    <RecentPosts :recentPosts="user.recentPosts" />
                    <ActivityFeed :activities="user.activities" />
                </div>

                <div v-if="currentTab === 'github'" class="flex flex-col items-center">
                    <h2 class="text-xl font-semibold mb-3">GitHub Stats</h2>
                    <img :src="`https://github-readme-stats.vercel.app/api?username=${user.username}&show_icons=true&theme=radical`" alt="GitHub Stats" class="mb-3" />
                    <img :src="`https://github-readme-stats.vercel.app/api/top-langs/?username=${user.username}&layout=compact&theme=radical`" alt="Top Languages" class="mb-3" />
                    <img :src="`https://github-profile-trophy.vercel.app/?username=${user.username}`" alt="Profile Trophy" class="mb-3" />
                    <img :src="`https://activity-graph.herokuapp.com/graph?username=${user.username}&theme=rogue`" alt="Contribution Graph" class="mb-3" />
                </div>

                <div v-if="currentTab === 'bounties'">
                    <h2 class="text-xl font-semibold mb-3">Bounties</h2>
                    <UserBountiesList :bounties="user.bounties" />
                </div>

                <div v-if="currentTab === 'transactions'">
                    <h2 class="text-xl font-semibold mb-3">Transaction History</h2>
                    <UserTransactionsHistory :transactions="user.transactions" />
                </div>

                <div v-if="currentTab === 'payment'">
                    <h2 class="text-xl font-semibold mb-3">Payment Information</h2>
                    <UserPaymentInformation :paymentInfo="user.paymentInfo" />
                </div>
            </div>
        </div>

        <!-- Additional Sections -->
        <IssuesList v-if="issues.length > 0" :issues="issues" />
    </section>
</template>

<script>
import PersonalInfo from '../components/PersonalInfo.vue';
import AboutMe from '../components/AboutMe.vue';
import Interests from '../components/Interests.vue';
import RecentPosts from '../components/RecentPosts.vue';
import IssuesList from '../components/IssuesList.vue';
import Badge from '../components/Badge.vue';
import Achievement from '../components/Achievement.vue';
import ActivityFeed from '../components/ActivityFeed.vue';
import UserBountiesList from '../components/UserBountiesList.vue';
import UserTransactionsHistory from '../components/UserTransactionsHistory.vue';
import UserPaymentInformation from '../components/UserPaymentInformation.vue';
import { ref, computed } from 'vue';
import { useUserStore } from '../stores/user';

export default {
    components: {
        PersonalInfo,
        AboutMe,
        Interests,
        RecentPosts,
        IssuesList,
        Badge,
        Achievement,
        ActivityFeed,
        UserBountiesList,
        UserTransactionsHistory,
        UserPaymentInformation
    },
    setup() {
        const userStore = useUserStore();
        const user = computed(() => userStore.user);
        const loading = ref(true);
        const issues = ref([]);
        const repos = ref([]);
        const currentTab = ref('personal');

        const fetchUserData = async () => {
            if (user.value) {
                try {
                    const response = await fetch(`https://api.github.com/users/${user.value.username}`, {
                        headers: {
                            Authorization: userStore.authGithubHeader
                        }
                    });
                    const userData = await response.json();

                    userStore.setUser({
                        username: userData.login,
                        avatar: userData.avatar_url,
                        userBio: userData.bio,
                        fullName: userData.name,
                        email: userData.email || 'Not Available',
                        phoneNumber: 'Not Available',
                        location: userData.location,
                        aboutMe: userData.bio,
                        interests: ['Coding', 'Music', 'Travel', 'Photography'], // Example static data
                        recentPosts: [
                            { id: 1, content: 'Just finished a new project!', timestamp: '2 hours ago' },
                            { id: 2, content: 'Learning Vue.js is fun!', timestamp: '1 day ago' }
                        ],
                        level: calculateLevel(userData.public_repos, userData.followers),
                        achievements: fetchAchievements(userData),
                        activities: fetchRecentActivities(userData),
                        bounties: fetchBounties(userData),
                        transactions: fetchTransactions(userData),
                        paymentInfo: fetchPaymentInfo(userData)
                    });

                    const reposResponse = await fetch(`https://api.github.com/users/${user.value.username}/repos`, {
                        headers: {
                            Authorization: userStore.authGithubHeader
                        }
                    });
                    repos.value = await reposResponse.json();
                } catch (error) {
                    console.error("Error fetching user data:", error);
                } finally {
                    loading.value = false;
                }
            }
        };

        const calculateLevel = (repos, followers) => {
            return Math.floor(repos / 10) + Math.floor(followers / 50);
        };

        const fetchAchievements = (userData) => {
            // Logic to fetch or generate achievements based on user data
            return [
                { id: 1, title: 'First Repository', description: 'Created your first repository', date: '2021-01-01' },
                { id: 2, title: '100 Commits', description: 'Made 100 commits', date: '2022-01-01' }
            ];
        };

        const fetchRecentActivities = (userData) => {
            // Logic to fetch or generate recent activities
            return [
                { id: 1, action: 'Pushed to repository', repo: 'sample-repo', date: '2 hours ago' },
                { id: 2, action: 'Opened an issue', repo: 'another-repo', date: '1 day ago' }
            ];
        };

        const fetchBounties = (userData) => {
            // Logic to fetch user bounties
            return [
                { id: 1, title: 'Bug Fix', description: 'Fix the bug in project X', amount: '$100', status: 'Completed' },
                { id: 2, title: 'Feature Implementation', description: 'Implement feature Y in project Z', amount: '$200', status: 'In Progress' }
            ];
        };

        const fetchTransactions = (userData) => {
            // Logic to fetch user transactions
            return [
                { id: 1, type: 'Deposit', amount: '$500', date: '2023-01-01' },
                { id: 2, type: 'Withdrawal', amount: '$200', date: '2023-01-15' }
            ];
        };

        const fetchPaymentInfo = (userData) => {
            // Logic to fetch payment information
            return {
                cardNumber: '**** **** **** 1234',
                expiryDate: '12/25',
                cardHolderName: userData.name
            };
        };

        const tabClass = (tab) => {
            return {
                'px-4 py-2 border-b-2': true,
                'border-blue-500 text-blue-500': currentTab.value === tab,
                'border-transparent': currentTab.value !== tab
            };
        };

        fetchUserData();

        return {
            user,
            loading,
            issues,
            repos,
            currentTab,
            tabClass
        };
    }
};
</script>

<style scoped>
/* Add your styles here */
.logo {
    font-size: 1.5rem;
    font-weight: bold;
}
</style>
