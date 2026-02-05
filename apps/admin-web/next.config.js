/** @type {import('next').NextConfig} */
const nextConfig = {
  transpilePackages: ['@workshop-platform/api-client', '@workshop-platform/types'],
  output: 'standalone',
};

module.exports = nextConfig;
