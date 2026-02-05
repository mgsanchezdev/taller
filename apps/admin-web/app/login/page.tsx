'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { login, configureApi } from '@workshop-platform/api-client';

export default function LoginPage() {
  const router = useRouter();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [workshopId, setWorkshopId] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError('');
    setLoading(true);
    try {
      const res = await login({ email, password, workshop_id: workshopId });
      if (typeof window !== 'undefined') {
        localStorage.setItem('token', res.token);
        localStorage.setItem('user', JSON.stringify(res.user));
      }
      configureApi({ getToken: () => res.token });
      router.push('/dashboard');
      router.refresh();
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'Error al iniciar sesión');
    } finally {
      setLoading(false);
    }
  }

  return (
    <main
      style={{
        minHeight: '100vh',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        padding: '1rem',
      }}
    >
      <div
        style={{
          width: '100%',
          maxWidth: '400px',
          background: 'var(--surface)',
          border: '1px solid var(--border)',
          borderRadius: '12px',
          padding: '2rem',
        }}
      >
        <h1 style={{ marginBottom: '1.5rem', fontSize: '1.5rem' }}>
          Iniciar sesión
        </h1>
        <form onSubmit={handleSubmit}>
          <label style={{ display: 'block', marginBottom: '0.5rem', color: 'var(--muted)', fontSize: '0.875rem' }}>
            Workshop ID
          </label>
          <input
            type="text"
            value={workshopId}
            onChange={(e) => setWorkshopId(e.target.value)}
            required
            style={{
              width: '100%',
              padding: '0.75rem',
              marginBottom: '1rem',
              background: 'var(--bg)',
              border: '1px solid var(--border)',
              borderRadius: '8px',
              color: 'var(--text)',
            }}
          />
          <label style={{ display: 'block', marginBottom: '0.5rem', color: 'var(--muted)', fontSize: '0.875rem' }}>
            Email
          </label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            style={{
              width: '100%',
              padding: '0.75rem',
              marginBottom: '1rem',
              background: 'var(--bg)',
              border: '1px solid var(--border)',
              borderRadius: '8px',
              color: 'var(--text)',
            }}
          />
          <label style={{ display: 'block', marginBottom: '0.5rem', color: 'var(--muted)', fontSize: '0.875rem' }}>
            Contraseña
          </label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            style={{
              width: '100%',
              padding: '0.75rem',
              marginBottom: '1rem',
              background: 'var(--bg)',
              border: '1px solid var(--border)',
              borderRadius: '8px',
              color: 'var(--text)',
            }}
          />
          {error && (
            <p style={{ color: '#ef4444', fontSize: '0.875rem', marginBottom: '1rem' }}>
              {error}
            </p>
          )}
          <button
            type="submit"
            disabled={loading}
            style={{
              width: '100%',
              padding: '0.75rem',
              background: 'var(--accent)',
              color: 'var(--bg)',
              border: 'none',
              borderRadius: '8px',
              fontWeight: 600,
              cursor: loading ? 'not-allowed' : 'pointer',
            }}
          >
            {loading ? 'Entrando...' : 'Entrar'}
          </button>
        </form>
      </div>
    </main>
  );
}
