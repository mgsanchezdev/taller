/**
 * Enlaza node_modules/next de la raíz en apps/admin-web/node_modules/next
 * para que Next 16 (Turbopack) encuentre el paquete cuando [project] = apps/admin-web.
 * Ejecutar desde la raíz del monorepo (postinstall).
 */
const fs = require('fs');
const path = require('path');

const rootDir = path.resolve(__dirname, '..');
const adminWebDir = path.join(rootDir, 'apps', 'admin-web');
const adminWebNodeModules = path.join(adminWebDir, 'node_modules');
const nextInRoot = path.join(rootDir, 'node_modules', 'next');
const nextLink = path.join(adminWebNodeModules, 'next');

if (!fs.existsSync(nextInRoot)) {
  console.log('link-next-in-admin-web: next no encontrado en raíz, omitiendo.');
  process.exit(0);
}

try {
  if (!fs.existsSync(adminWebNodeModules)) {
    fs.mkdirSync(adminWebNodeModules, { recursive: true });
  }
  if (fs.existsSync(nextLink)) {
    const stat = fs.lstatSync(nextLink);
    if (stat.isSymbolicLink() || stat.isDirectory()) {
      process.exit(0);
    }
    fs.rmSync(nextLink, { recursive: true });
  }
  // Windows: 'junction' no requiere permisos de admin
  const isWin = process.platform === 'win32';
  fs.symlinkSync(nextInRoot, nextLink, isWin ? 'junction' : 'dir');
  console.log('link-next-in-admin-web: next enlazado en apps/admin-web/node_modules/next');
} catch (e) {
  console.warn('link-next-in-admin-web:', e.message);
}
