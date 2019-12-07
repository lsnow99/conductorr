module.exports = {
    devServer: {
      proxy: {
        '/backend': {
          target: 'http://localhost:7080',
          ws: true,
          changeOrigin: true
        }
      }
    }
  }