/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
	"./**/*.{templ,html,js,jsx,ts,tsx}",
    "**/*.templ"  // Catch all Templ files
  ],
  theme: {
    extend: {
    },
  },
  plugins: [],
};
