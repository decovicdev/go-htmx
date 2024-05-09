/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html"],
  theme: {
    extend: {
      colors: {
        "accent-primary": "var(--color-accent-primary)",
        "accent-secondary": "var(--color-accent-secondary)",
        "shade": "var(--color-shade)",
        "shade-2": "var(--color-shade-2)",
        "green": "var(--color-green)",
      }
    },
  },
  plugins: [],
}

