/// <reference types="vite/client" />

declare global {
  interface Window {
    go?: {
      main?: {
        App?: {
          InitializeWorkspace: () => Promise<unknown>;
          ImportFolder: (
            sourcePath: string,
            feedSlug: string,
          ) => Promise<unknown>;
        };
      };
    };
  }
}

export {};
