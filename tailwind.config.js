/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
  theme: {
    extend: {
      colors: {
        background: "var(--color-background)",
        "card-background": "var(--color-card-background)",
        primary: "var(--color-primary)",
        secondary: "var(--color-secondary)",
        border: "var(--color-border)",
        accent: {
          DEFAULT: "var(--color-accent)",
          hover: "var(--color-accent-hover)",
        },
      },
    },
  },
  plugins: [],
};
