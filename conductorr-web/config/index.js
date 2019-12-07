module.exports = {
    // ...
    dev: {
      proxyTable: {
        // proxy all requests starting with /api to jsonplaceholder
        '**': {
            target: 'http://127.0.0.1:7080',
            filter: function (pathname, req) {
              return pathname.indexOf('backend') > 0 && req.method === 'GET'
            }
          }
      }
    }
  }