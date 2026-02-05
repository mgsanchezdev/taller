import Link from 'next/link';

export default function HomePage() {
  return (
    <main style={{ padding: '2rem', textAlign: 'center' }}>
      <h1 style={{ marginBottom: '1rem' }}>Workshop Platform</h1>
      <p style={{ color: 'var(--muted)', marginBottom: '2rem' }}>
        Backoffice para gestión de talleres
      </p>
      <Link
        href="/login"
        style={{
          display: 'inline-block',
          padding: '0.75rem 1.5rem',
          background: 'var(--accent)',
          color: 'var(--bg)',
          borderRadius: '8px',
          fontWeight: 600,
        }}
      >
        Iniciar sesión
      </Link>
      {' · '}
      <Link href="/dashboard">Dashboard</Link>
    </main>
  );
}
