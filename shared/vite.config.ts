import * as path from "path";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {

  const LIB_MODE = {
    lib: {
      entry: path.resolve(__dirname, "src/lib.ts"),
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
  }

  const WEB_MODE = {
    outDir: "build/dist",
  }

  const buildMode = mode == 'web' ? WEB_MODE : LIB_MODE;

  return {
    plugins: [vue()],
    base: "/",
    build: buildMode,
    resolve: {
      alias: {
        '@': path.resolve(__dirname, './src'),
      },
    }
  };
});
