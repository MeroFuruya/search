// @ts-check
import { defineConfig } from 'astro/config';
import tailwindcss from "@tailwindcss/vite";
import { loadEnv } from "vite";
import { fileURLToPath } from 'url';
import { dirname, join } from 'path';


const { PUBLIC_API_URL } = loadEnv(import.meta.env.NODE_ENV, join(dirname(fileURLToPath(import.meta.url)), '..'), ['PUBLIC']);

// https://astro.build/config
export default defineConfig({
  vite: {
    plugins: [tailwindcss()]
  }
});
