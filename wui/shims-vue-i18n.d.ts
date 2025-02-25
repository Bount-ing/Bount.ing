// shims-vue-i18n.d.ts
import 'vue-i18n';

declare module 'vue-i18n' {
  // Optional: Add typing for the $t function globally
  interface VueI18n {
    $t: (key: string, ...args: any[]) => string;
  }
}

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $t: (key: string, ...args: any[]) => string;
  }
}

