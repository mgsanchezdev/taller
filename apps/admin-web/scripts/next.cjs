#!/usr/bin/env node
/* eslint-disable no-console */
const path = require('path');
const { spawnSync } = require('child_process');

const args = process.argv.slice(2);
const cwd = process.cwd();

function resolveNextBin() {
  // Prefer resolución desde el directorio actual (apps/admin-web)
  const candidates = [
    cwd,
    // Fallback: raíz del monorepo (../../)
    path.resolve(cwd, '..', '..'),
  ];

  for (const base of candidates) {
    try {
      return require.resolve('next/dist/bin/next', { paths: [base] });
    } catch {
      // continue
    }
  }
  return null;
}

const nextBin = resolveNextBin();
if (!nextBin) {
  console.error(
    'No se encontró Next.js. Ejecuta `npm install` en la raíz del monorepo.'
  );
  process.exit(1);
}

const res = spawnSync(process.execPath, [nextBin, ...args], {
  stdio: 'inherit',
  cwd,
});

process.exit(typeof res.status === 'number' ? res.status : 1);

