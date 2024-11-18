import { type Config } from 'tailwindcss'

export default {
  content: [
    './src/**/*.tsx',
  ],
  theme: {
    extend: {
      colors: {
        white: '#fafafa',
        black: '#1a1a1a',
        grayer: '#dddddd',
        gray: '#374151',
        green: '#008000',
      },
    },
  },
  plugins: [],
} satisfies Config
