const path = require("path");
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
  return {
    plugins: [vue()],
    base: "/",
    resolve: command === 'build'?{mainFields: ["mainLib"]}:undefined,
    build: {
      lib: {
        entry: path.resolve(__dirname, "src/lib.js"),
        name: "conductorr-lib",
        fileName: (format) => `conductorr-lib.${format}.js`,
      },
      rollupOptions: {
        external: ["vue"],
        output: {
          globals: {
            vue: "Vue",
          },
        },
      },
      outDir: "dist",
    },
  };
});
