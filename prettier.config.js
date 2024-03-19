/** @type {import('prettier').Config} */
module.exports = {
  trailingComma: 'es5',
  tabWidth: 2,
  semi: false,
  singleQuote: true,

  plugins: [

    require('@tailwindcss/forms'),
  ],

  overrides: [
    {
      files: '.postcssrc',
      options: { parser: 'json' },
    },
  ],
}
