import type { Metadata } from 'next';
import './globals.css';

export const metadata: Metadata = {
  title: 'Workshop Platform - Admin',
  description: 'Backoffice para gestión de talleres',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="es">
      <body>{children}</body>
    </html>
  );
}
