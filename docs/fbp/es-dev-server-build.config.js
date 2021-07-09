// eslint-disable-next-line import/no-extraneous-dependencies
const proxy = require('koa-proxies');

module.exports = {
  port: 8480,
  hostname: 'localhost',
  http2: false,
  watch: true,
  nodeResolve: true,
  appIndex: 'index.html',
  preserveSymlinks: true,
  moduleDirs: ['node_modules', 'custom-modules'],
  middlewares: [
    proxy('/api', {
      target: 'http://localhost:8481',
      changeOrigin: true,
      rewrite: path => path.replace(/^\/api\//, '/'),
    }),
  ],
};
