import type { Config } from 'tailwindcss'

const config: Config = {
  content: ['./index.html', './src/**/*.{ts,tsx}'],
  theme: {
    extend: {
      colors: {
        // AI WATT design tokens
        surface: {
          DEFAULT: '#0D1117',
          card:    '#161B22',
          hover:   '#1C2128',
          border:  '#30363D',
        },
        brand: {
          DEFAULT: '#58A6FF',
          dim:     '#388BFD',
          muted:   '#1F3D5C',
        },
        yield: {
          DEFAULT: '#3FB950',
          muted:   '#1A3A22',
        },
        warn: {
          DEFAULT: '#D29922',
          muted:   '#3A2A0A',
        },
        danger: {
          DEFAULT: '#F85149',
          muted:   '#3D1212',
        },
        text: {
          primary:   '#E6EDF3',
          secondary: '#8B949E',
          muted:     '#484F58',
        },
      },
      fontFamily: {
        sans: ['Inter', 'system-ui', 'sans-serif'],
        mono: ['JetBrains Mono', 'monospace'],
      },
    },
  },
  plugins: [],
}

export default config
