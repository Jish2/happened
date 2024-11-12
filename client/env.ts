import env from "./.env.json";

export const CLERK_PUBLISHABLE_KEY = env.EXPO_PUBLIC_CLERK_PUBLISHABLE_KEY;

if (!CLERK_PUBLISHABLE_KEY)
  throw new Error("Add CLERK_PUBLISHABLE_KEY to your .env.json file");
