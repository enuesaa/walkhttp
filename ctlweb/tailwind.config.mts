import { type Config } from 'tailwindcss'

export default {
  content: [
    './src/**/*.tsx',
  ],
  theme: {
    extend: {
      colors: {
        gray: '#374151'
      },
    },
  },
  plugins: [],
} satisfies Config
