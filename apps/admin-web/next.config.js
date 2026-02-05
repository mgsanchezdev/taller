/** @type {import('next').NextConfig} */
const path = require('path');

const rootDir = path.resolve(__dirname, '..', '..');

const nextConfig = {
  transpilePackages: [
    '@workshop-platform/api-client',
    '@workshop-platform/types',
    '@mui/material',
    '@mui/icons-material',
    '@emotion/react',
    '@emotion/styled',
  ],
  output: 'standalone',
  turbopack: {
    root: rootDir,
  },
  webpack: (config) => {
    // Resolver MUI y otras dependencias desde la raíz del monorepo
    config.resolve.modules = [
      path.resolve(__dirname, 'node_modules'),
      path.resolve(rootDir, 'node_modules'),
      'node_modules',
    ];
    return config;
  },
};

module.exports = nextConfig;
