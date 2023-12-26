/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./**/*.{html,js,php,ts,go}"],
	theme: {
	  extend: {
		backgroundImage: {
		  'swoop': "url('/public/svg/swoop.svg')",
		},
		colors: {
		  "red": "#E51636",
		  "white": "#FFFFFF",
		  "blue": "#0000ff",
		  "gray": "#999999",
		},
		fontFamily: {
		  'playful': ['Chelsea Market', 'cursive'],
		},
	  },
	},
	plugins: [],
  }