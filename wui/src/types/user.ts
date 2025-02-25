import type { GithubUser } from './github';

export interface User {
  id: number;
  username?: string | null | undefined;
  avatar: string;
  bio?: string;
  fullName?: string;
  email?: string;
  phoneNumber?: string;
  location?: string;
  bearerToken?: string;
  refreshJwt?: string;
  aboutMe?: string;
  interests?: string[];
  recentPosts?: any[];
  level?: number;
  achievements?: any[];
  activities?: any[];
  bounties?: any[];
  transactions?: any[];
  paymentInfo?: any;
  isLoggedIn?: boolean;
  authGithubHeader?: string;
  githubUser?: GithubUser;
}

export interface EditableUserFields {
  username: string | null;
  fullName: string | null;
  phoneNumber: string | null;
  location: string | null;
  bio: string | null;
  avatar: string | null;
}

export const DEFAULT_USER_VALUES: EditableUserFields = {
  fullName: '',
  username: '',
  phoneNumber: '',
  location: '',
  bio: '',
  avatar: 'http://localhost:3000/default-avatar.png',
};