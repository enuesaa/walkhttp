import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import generouted from '@generouted/react-router/plugin'

export default defineConfig({
  root: 'web',
  publicDir: "./public",
  plugins: [
    react(),
    generouted({
      source: { routes: './web/src/pages/**/*.tsx', modals: './web/src/pages/**/[+]*.tsx' },
      output: './web/src/router.ts',
    }),
  ],
  resolve: { alias: { '@': '/web/src' } },
})
