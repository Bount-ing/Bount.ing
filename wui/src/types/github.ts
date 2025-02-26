export interface GithubUser {
    id: number;
    login: string;
    avatar_url: string;
  }
  
  export interface Issue {
    id: number;
    number: number;
    title: string;
    body: string;
    state: string;
    created_at: string;
    updated_at: string;
    closed_at: string;
    labels: {
      id: number;
      name: string;
      color: string;
    }[];
    assignees: {
      login: string;
    }[];
    repo_avatar?: string;
  }
  
  export interface Repo {
    id: number;
    name: string;
    full_name: string;
    owner: {
      login: string;
      avatar_url: string;
    };
    html_url: string;
    description: string;
    stargazers_count: number;
    watchers_count: number;
    forks_count: number;
    open_issues_count: number;
    license: {
      name: string;
    };
    created_at: string;
    updated_at: string;
    pushed_at: string;
    issues?: Issue[];
  }
  
  export interface HostData {
    connected: boolean;
    token: string | null;
    userInfo: any | null;
    repos: Repo[];
  }