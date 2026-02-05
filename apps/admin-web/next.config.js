/** @type {import('next').NextConfig} */
const path = require('path');

const nextConfig = {
  transpilePackages: ['@workshop-platform/api-client', '@workshop-platform/types'],
  output: 'standalone',
  turbopack: {
    root: path.resolve(__dirname, '..', '..'),
  },
};

module.exports = nextConfig;
