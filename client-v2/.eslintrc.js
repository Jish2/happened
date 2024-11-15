// https://docs.expo.dev/guides/using-eslint/
module.exports = {
  extends: [
    "expo",
    "prettier",
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
  ],
  parser: "@typescript-eslint/parser",
  ignorePatterns: ["/dist/*", "/.expo/*", "/node_modules/*"],
  plugins: ["prettier", "@typescript-eslint"],
  rules: {
    "prettier/prettier": "off",
    "@typescript-eslint/no-require-imports": "off",
    "@typescript-eslint/no-unused-vars": "warn",
    "import/no-unresolved": "off",
  },
};
