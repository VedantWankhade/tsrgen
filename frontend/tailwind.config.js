/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.vue'],
  theme: {
    extend: {},
    fontFamily: {
      maple: ['Maple Mono'],
      inter: ['Inter']
    }
  },
  //   safelist: [
  //     {
  //       pattern: /bg-ctp-*|text-ctp-*|border-ctp-*|from-ctp-*|to-ctp-*/
  //     }
  //   ],
  plugins: [
    // eslint-disable-next-line no-undef
    require('@catppuccin/tailwindcss')({
      prefix: 'ctp',
      defaultFlavour: 'frappe'
    })
  ]
}
