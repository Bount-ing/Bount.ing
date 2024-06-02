<template>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <img
        v-for="src in imageSources"
        :key="src"
        :src="src"
        alt="GitHub Stats"
      />
    </div>
  </template>
  
  <script>
  export default {
    props: {
      username: {
        type: String,
        required: true,
      },
    },
    computed: {
      imageSources() {
        return [
          `https://github-readme-stats.vercel.app/api?username=${this.username}&show_icons=true&theme=radical`,
          `https://github-readme-stats.vercel.app/api/top-langs/?username=${this.username}&layout=compact&theme=radical`,
          `https://github-readme-streak-stats.herokuapp.com/?user=${this.username}&theme=radical`,
          `https://komarev.com/ghpvc/?username=${this.username}&color=blueviolet&style=flat-square`,
          `https://github-profile-trophy.vercel.app/?username=${this.username}&theme=radical&margin-w=15`,
          `https://metrics.lecoq.io/${this.username}?template=classic&base.header=0&base.activity=0&base.community=0&base.repositories=0&base.metadata=0&languages=1&isocalendar=1&languages.ignored=html%2Ccss%2Cscss&languages.limit=8&languages.colors=github&isocalendar.duration=half-year&config.timezone=America%2FNew_York`,
          `https://github-readme-stats.vercel.app/api/wakatime?username=${this.username}&theme=radical`,
          `https://github-contributor-stats.vercel.app/api?username=${this.username}&theme=radical`,
          `https://github-profile-summary-cards.vercel.app/api/cards/repos-per-language?username=${this.username}&theme=radical`,
          `https://github-profile-summary-cards.vercel.app/api/cards/most-commit-language?username=${this.username}&theme=radical`,
          `https://github-readme-stats.vercel.app/api/gist?id=yourgistid&theme=radical`,
          `https://quotes-github-readme.vercel.app/api?type=horizontal&theme=radical`,
          `https://github-readme-stats.vercel.app/api?username=${this.username}&show_icons=true&hide=contribs,prs&cache_seconds=86400&theme=radical`,
          `https://img.shields.io/badge/Product%20Hunt-Follow%20Me-red`,
          `https://img.shields.io/badge/Sponsor-${this.username}-ff69b4?style=flat-square&logo=github-sponsors`,
          `https://img.shields.io/badge/Fun%20Fact-${this.username}-blueviolet?style=flat-square&logo=github`,
          `https://badgen.net/badge/stats/awesome/purple?icon=github`,
          `https://img.shields.io/badge/commits-${this.username}-orange?style=flat-square&logo=github`,
          `https://img.shields.io/badge/Discussions-${this.username}-yellow?style=flat-square&logo=github`,
          `https://img.shields.io/badge/Actions-${this.username}-blue?style=flat-square&logo=github`,
          `https://img.shields.io/github/followers/${this.username}?style=social`,
          `https://img.shields.io/github/stars/${this.username}?style=social`,
        ];
      },
    },
  };
  </script>
  
  <style scoped>
  img {
    max-width: 100%;
    height: auto;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  }
  </style>
  