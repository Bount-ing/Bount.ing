import 'vue-i18n';

declare module 'vue-i18n' {
  export interface DefineLocaleMessage {
    hello: string;
    welcome: string;
  }

  export interface DefineDateTimeFormat {}

  export interface DefineNumberFormat {}
}
