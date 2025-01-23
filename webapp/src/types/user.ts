// src/types/user.ts

// User-related interfaces
export interface EditableUserFields {
    fullName: string;
    Email: string;
    phoneNumber: string;
    location: string;
    userBio: string;
    avatar: string;
  }
  
  // Default values and constants
  export const DEFAULT_USER_VALUES: EditableUserFields = {
    fullName: '',
    Email: '',
    phoneNumber: '',
    location: '',
    userBio: '',
    avatar: '',
  };
  
  // You can also add other user-related types and constants here
  export interface UserResponse {
    // API response type
    id: number;
    username: string;
    // ... other fields
  }
  
  export interface UserUpdatePayload {
    // API request type
    id: number;
    // ... fields that can be updated
  }