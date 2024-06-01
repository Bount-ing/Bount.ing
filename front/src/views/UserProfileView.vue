<template>
    <section class="min-h-screen flex flex-col items-center justify-center p-4 text-white">
        <div v-if="loading" class="text-center text-2xl font-semibold">Loading...</div>
        <div v-else class="w-full max-w-6xl rounded-2xl shadow-xl overflow-hidden ">
            <!-- Profile Section -->
            <div class="flex flex-col items-center p-8 ">
                <img :src="user.avatar" alt="Profile Picture" class="w-32 h-32 rounded-full border-4 border-gray-500"/>
                <h1 class="text-3xl font-semibold mt-4">{{ user.username }}</h1>
                <p class="text-gray-300 mt-2">{{ user.userBio }}</p>
                <button class="mt-4 px-4 py-2  text-white rounded-lg focus:outline-none">Edit Profile</button>
                <div class="mt-4 flex space-x-4">
                    <Badge :level="user.level" />
                    <Achievement :achievements="user.achievements" />
                </div>
            </div>

            <!-- Details Section -->
            <div class="p-8">
                <PersonalInfo :user="user" />
                <AboutMe :aboutMe="user.aboutMe" />
                <Interests :interests="user.interests" />
            </div>

            <!-- Activity Section -->
            <div class="p-8">
                <RecentPosts :recentPosts="user.recentPosts" />
                <div class="mt-6 flex flex-col items-center">
                    <h2 class="text-2xl font-semibold mb-4">Recent Activities</h2>
                    <ActivityFeed :activities="user.activities" />
                </div>
            </div>

            <!-- GitHub Stats Section -->
            <div class="p-8 flex flex-col items-center">
                <h2 class="text-2xl font-semibold mb-4">GitHub Stats</h2>
                <img :src="`https://github-readme-stats.vercel.app/api?username=${user.username}&show_icons=true&theme=radical`" alt="GitHub Stats" class="mb-4" />
                <img :src="`https://github-readme-stats.vercel.app/api/top-langs/?username=${user.username}&layout=compact&theme=radical`" alt="Top Languages" class="mb-4" />
                <img :src="`https://github-profile-trophy.vercel.app/?username=${user.username}`" alt="Profile Trophy" class="mb-4" />
                <img :src="`https://activity-graph.herokuapp.com/graph?username=${user.username}&theme=rogue`" alt="Contribution Graph" class="mb-4" />
            </div>

            <!-- Repositories Section -->
            <div class="p-8">
                <h2 class="text-2xl font-semibold mb-4">Repositories</h2>
                <div v-if="repos.length">
                    <div v-for="repo in repos" :key="repo.id" class="mb-4 p-4 rounded-lg">
                        <h3 class="text-xl font-semibold">{{ repo.name }}</h3>
                        <p class="text-gray-300">{{ repo.description }}</p>
                        <a :href="repo.html_url" class="text-blue-400 hover:underline">View Repository</a>
                    </div>
                </div>
                <div v-else>
                    <p class="text-gray-400">No repositories found.</p>
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
        ActivityFeed
    },
    setup() {
        const userStore = useUserStore();
        const user = computed(() => userStore.user);
        const loading = ref(true);
        const issues = ref([]);
        const repos = ref([]);

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
                        activities: fetchRecentActivities(userData)
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

        fetchUserData();

        return {
            user,
            loading,
            issues,
            repos
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
