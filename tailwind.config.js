/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		"./**/*.{templ,html}",
		"./**/*.templ"  // Catch all Templ files
	],
	theme: {
		extend: {
			fontFamily: {
				"atma": ["Atma", "serif"]
			}
		},
	},
	plugins: [],
};
