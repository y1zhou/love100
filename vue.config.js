module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:9423',
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  },
  productionSourceMap: false
}
