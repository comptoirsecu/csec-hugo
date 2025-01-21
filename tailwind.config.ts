import typography from '@tailwindcss/typography';
import type { Config } from 'tailwindcss';

export default {
  content: ['./src/**/*.{html,js,svelte,ts}', './node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'],

  theme: {
    extend: {},
    fontFamily: {
      'titres': ['roboto slab', '"Helvetica Neue"', 'Helvetica', 'Roboto', 'Arial', 'sans-serif'],
    }
  },

  plugins: [typography]
} satisfies Config;
