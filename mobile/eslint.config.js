import js from "@eslint/js";
import prettierConfig from "./prettier.config.js";
import prettier from "eslint-plugin-prettier";
import prettierConfigRecommended from "eslint-config-prettier";

export default [
    js.configs.recommended,
    prettierConfigRecommended,
    {
        plugins: {
            prettier,
        },
        rules: {
            "no-unused-vars": "warn",
            "no-undef": "warn",
            semi: ["error", "always"],
            "prettier/prettier": ["error", prettierConfig],
        },
        ignores: ["prettier.config.js", "eslint.config.js"],
    },
];