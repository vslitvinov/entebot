module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
       colors: {
        'eblack': '#35404e',
        'herobg': '#35444e',
      },
      backgroundImage: {
        'hero': "url('../public/images/bg.jpg')",
      }
    },
  },
  plugins: [],
}