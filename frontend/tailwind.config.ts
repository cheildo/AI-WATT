import type { Config } from 'tailwindcss'

const config: Config = {
  content: ['./index.html', './src/**/*.{ts,tsx}'],
  theme: {
    extend: {
      colors: {
        bg: { DEFAULT: '#F4F6FA', 2: '#EAECF3', 3: '#DDE1EC' },
        white: '#FDFCF9',
        border: { DEFAULT: '#D8D3C6', strong: '#C8C2B0' },
        green: {
          DEFAULT: '#1A3C2A',
          mid: '#2D5C3E',
          light: '#3D7A52',
          bg: '#EBF3EE',
          border: '#A8CDBA',
        },
        gold: { DEFAULT: '#9A6B0A', bg: '#FBF4E4', border: '#E2C46A' },
        teal: { DEFAULT: '#0A7068', bg: '#EAF5F3', border: '#80C9C0' },
        red: { DEFAULT: '#8B2020', bg: '#FBF0F0' },
        text: { 1: '#1C1A14', 2: '#5A5646', 3: '#9A9484' },
      },
      fontFamily: {
        serif: ['Instrument Serif', 'Georgia', 'serif'],
        sans: ['Inter', 'system-ui', 'sans-serif'],
        mono: ['JetBrains Mono', 'monospace'],
      },
    },
  },
  plugins: [],
}

export default config
