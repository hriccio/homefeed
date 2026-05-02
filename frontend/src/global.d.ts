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
          CreateNotePost: (
            feedSlug: string,
            title: string,
            body: string,
          ) => Promise<unknown>;
        };
      };
    };
  }
}

export {};
