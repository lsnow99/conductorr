import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import viteCompression from "vite-plugin-compression";
import tsconfigPaths from 'vite-tsconfig-paths'
import * as path from "path"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    viteCompression({
      algorithm: "brotliCompress",
      ext: "br",
    }),
    tsconfigPaths({
      extensions: ["ts", "tsx", "js", "jsx", "mjs", "vue"]
    })
  ],
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8282",
        changeOrigin: false,
        secure: false,
        ws: true,
      },
    },
  },
  base: "/",
  build: {
    outDir: "build/dist",
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  }
});
