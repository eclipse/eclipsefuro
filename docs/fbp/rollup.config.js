/* eslint-disable */
import merge from 'deepmerge';
// use createSpaConfig for bundling a Single Page App
import { createSpaConfig } from '@open-wc/building-rollup';

const workboxConfig = require('./workbox-config.js');
const { generateSW } = require('rollup-plugin-workbox');

import copy from 'rollup-plugin-copy';

// use createBasicConfig to do regular JS to JS bundling
// import { createBasicConfig } from '@open-wc/building-rollup';

const baseConfig = createSpaConfig({
  // use the outputdir option to modify where files are output
  // outputDir: 'dist',

  // if you need to support older browsers, such as IE11, set the legacyBuild
  // option to generate an additional build just for this browser
  // legacyBuild: true,

  // development mode creates a non-minified build for debugging or development
  developmentMode: process.env.ROLLUP_WATCH === 'true',

  // set to true to inject the service worker registration into your index.html
  injectServiceWorker: false,
  // we use a separate workbox config (workbox-config.js)
  workbox: false,
});

const copyConf = merge(baseConfig, {
  plugins: [
    copy({
      targets: [
        { src: 'assets/**/*', dest: 'dist/assets' },
        { src: 'manifest.webmanifest', dest: 'dist' },
        { src: 'robots.txt', dest: 'dist' },
      ],
      // set flatten to false to preserve folder structure
      flatten: false,
    }),
  ],
  // alternatively, you can use your JS as entrypoint for rollup and
  // optionally set a HTML template manually
  // input: './app.js',
});

export default merge(copyConf, {
  // if you use createSpaConfig, you can use your index.html as entrypoint,
  // any <script type="module"> inside will be bundled by rollup
  input: './index.html',

  plugins: [generateSW(workboxConfig)],
  // alternatively, you can use your JS as entrypoint for rollup and
  // optionally set a HTML template manually
  // input: './app.js',
});
