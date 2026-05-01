/// <reference types="vite/client" />

declare global {
  interface Window {
    go?: {
      main?: {
        App?: {
          InitializeWorkspace: () => Promise<unknown>;
        };
      };
    };
  }
}

export {};

