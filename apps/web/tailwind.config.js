/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{ts,tsx}"],
  theme: {
    extend: {},
    fontFamily: {
      sans: ["SourceSans3", "sans-serif"],
    },
  },
  corePlugins: {
    // We provide our own css resets
    preflight: false,
  },
  plugins: [],
}
