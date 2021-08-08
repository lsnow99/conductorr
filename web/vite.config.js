import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import viteCompression from 'vite-plugin-compression';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), viteCompression({
    algorithm: 'brotliCompress',
    ext: 'br',
  })],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8282',
        changeOrigin: true,
        secure: false,      
        ws: true,
      }
    }
  },
  base: '/',
  build: {
    outDir: 'build/dist'
  }
})