const fs = require('fs');
const path = require('path');

const content = '[]';
const locations = [
  path.join(__dirname, '..', 'node_modules', 'next', 'dist', 'lib', 'default-transpiled-packages.json'),
  path.join(__dirname, '..', 'apps', 'admin-web', 'node_modules', 'next', 'dist', 'lib', 'default-transpiled-packages.json'),
];

for (const filePath of locations) {
  const dir = path.dirname(filePath);
  if (!fs.existsSync(path.join(dir, '..', '..', 'package.json'))) continue; // next not installed here
  try {
    fs.mkdirSync(dir, { recursive: true });
    fs.writeFileSync(filePath, content);
    console.log('Created:', filePath);
  } catch (e) {
    // ignore
  }
}
