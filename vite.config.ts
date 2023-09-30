import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'node:path'

export default defineConfig({
  plugins: [react()],
  root: './web',
  resolve: {
    alias: {
      '@/': path.join(__dirname, './web/'),
    },
  },
  build: {
    outDir: '../build',
    rollupOptions: {
      output: {
        entryFileNames: '[name].js',
        assetFileNames: 'assets/[name][extname]',
      },
    },
  },
})
