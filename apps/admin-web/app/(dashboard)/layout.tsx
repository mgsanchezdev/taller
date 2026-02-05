import Link from 'next/link';

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div style={{ display: 'flex', minHeight: '100vh' }}>
      <aside
        style={{
          width: '220px',
          background: 'var(--surface)',
          borderRight: '1px solid var(--border)',
          padding: '1.5rem',
        }}
      >
        <nav style={{ display: 'flex', flexDirection: 'column', gap: '0.5rem' }}>
          <Link href="/dashboard" style={{ padding: '0.5rem', borderRadius: '6px' }}>
            Dashboard
          </Link>
          <Link href="/vehicles" style={{ padding: '0.5rem', borderRadius: '6px' }}>
            Vehículos
          </Link>
          <Link href="/entries" style={{ padding: '0.5rem', borderRadius: '6px' }}>
            Entradas
          </Link>
          <Link href="/inspections" style={{ padding: '0.5rem', borderRadius: '6px' }}>
            Inspecciones
          </Link>
          <Link href="/repairs" style={{ padding: '0.5rem', borderRadius: '6px' }}>
            Reparaciones
          </Link>
          <Link href="/reports" style={{ padding: '0.5rem', borderRadius: '6px' }}>
            Reportes
          </Link>
          <Link href="/settings" style={{ padding: '0.5rem', borderRadius: '6px' }}>
            Configuración
          </Link>
        </nav>
      </aside>
      <main style={{ flex: 1, padding: '2rem' }}>{children}</main>
    </div>
  );
}
