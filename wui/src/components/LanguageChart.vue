<template>
    <div>
        <canvas id="languageChart"></canvas>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { useUserStore } from '../stores/user';
import { Chart, registerables } from 'chart.js';

Chart.register(...registerables);

export default {
    setup() {
        const userStore = useUserStore();
        const languageData = ref(null);

        const fetchLanguageData = async () => {
            if (userStore.user) {
                const repos = await fetch(`https://api.github.com/users/${userStore.user.username}/repos`, {
                    headers: {
                        Authorization: userStore.authGithubHeader
                    }
                }).then(res => res.json());

                const languagePromises = repos.map(repo =>
                    fetch(repo.languages_url, {
                        headers: {
                            Authorization: userStore.authGithubHeader
                        }
                    }).then(res => res.json())
                );

                const languages = await Promise.all(languagePromises);
                const languageCounts = languages.reduce((acc, langObj) => {
                    for (const [lang, count] of Object.entries(langObj)) {
                        if (acc[lang]) {
                            acc[lang] += count;
                        } else {
                            acc[lang] = count;
                        }
                    }
                    return acc;
                }, {});

                languageData.value = languageCounts;
                createChart(languageCounts);
            }
        };

        const createChart = (data) => {
            const ctx = document.getElementById('languageChart').getContext('2d');
            new Chart(ctx, {
                type: 'pie',
                data: {
                    labels: Object.keys(data),
                    datasets: [{
                        data: Object.values(data),
                        backgroundColor: ['#FF6384', '#36A2EB', '#FFCE56', '#4BC0C0', '#9966FF', '#FF9F40']
                    }]
                }
            });
        };

        onMounted(() => {
            fetchLanguageData();
        });

        return {
            languageData
        };
    }
};
</script>

<style scoped>
/* Add your styles here */
</style>
