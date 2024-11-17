module.exports = {
  root: true,
  ignorePatterns: ["projects/**/*"],
  plugins: ["deprecation"],
  overrides: [
    {
      files: ["*.ts"],
      parserOptions: {
        project: "./tsconfig.json",  // Relative path to the tsconfig.json file
        tsconfigRootDir: __dirname,  // Make sure this points to the directory of your .eslintrc.cjs
        sourceType: "module"
      },
      extends: [
        "eslint:recommended",
        "plugin:@typescript-eslint/recommended",
        "plugin:@angular-eslint/recommended",
        "plugin:@angular-eslint/template/process-inline-templates"
      ],
      rules: {
        "@angular-eslint/directive-selector": [
          "error",
          {
            type: "attribute",
            style: "camelCase"
          }
        ],
        "@angular-eslint/component-selector": [
          "error",
          {
            type: "element",
            prefix: "app",
            style: "kebab-case"
          }
        ],
        "@typescript-eslint/no-explicit-any": "off",
        "@typescript-eslint/no-unused-vars": [
          "error",
          {
            argsIgnorePattern: "^_",       // Ignore unused variables that start with '_'
            varsIgnorePattern: "^_"        // Ignore unused variables in general that start with '_'
          }
        ],
        quotes: ["error", "single", { avoidEscape: true }],
        "deprecation/deprecation": "warn"
      }
    },
    {
      files: ["*.html"],
      extends: [
        "plugin:@angular-eslint/template/recommended",
        "plugin:@angular-eslint/template/accessibility"
      ],
      rules: {}
    }
  ]
};
