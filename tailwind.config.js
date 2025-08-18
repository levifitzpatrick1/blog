/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/templates/**/*.templ"],
  safelist: ["prose", "lg:prose-xl", "prose-invert"],
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
      typography: ({ theme }) => ({
        DEFAULT: {
          css: {
            "--tw-prose-body": theme("colors.secondary"),
            "--tw-prose-headings": theme("colors.primary"),
            "--tw-prose-lead": theme("colors.secondary"),
            "--tw-prose-links": theme("colors.accent.DEFAULT"),
            "--tw-prose-bold": theme("colors.primary"),
            "--tw-prose-counters": theme("colors.secondary"),
            "--tw-prose-bullets": theme("colors.border"),
            "--tw-prose-hr": theme("colors.border"),
            "--tw-prose-quotes": theme("colors.primary"),
            "--tw-prose-quote-borders": theme("colors.border"),
            "--tw-prose-captions": theme("colors.secondary"),
            "--tw-prose-code": theme("colors.primary"),
            "--tw-prose-pre-code": theme("colors.primary"),
            "--tw-prose-pre-bg": theme("colors.card-background"),
            "--tw-prose-th-borders": theme("colors.border"),
            "--tw-prose-td-borders": theme("colors.border"),
          },
        },
      }),
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
