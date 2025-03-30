import { defineConfig } from 'vite'
import solid from 'vite-plugin-solid'
import tailwindcss from '@tailwindcss/vite'


export default defineConfig({
  publicDir: 'src/assets',
  plugins: [solid(), tailwindcss()],
  envDir: '../'
})
