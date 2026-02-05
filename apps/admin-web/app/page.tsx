'use client';

import Link from 'next/link';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import LoginIcon from '@mui/icons-material/Login';
import DashboardIcon from '@mui/icons-material/Dashboard';

export default function HomePage() {
  return (
    <Box
      sx={{
        minHeight: '100vh',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        p: 3,
        gap: 3,
      }}
    >
      <Typography variant="h4" component="h1" fontWeight="bold" gutterBottom>
        Workshop Platform
      </Typography>
      <Typography color="text.secondary" sx={{ mb: 2 }}>
        Backoffice para gestión de talleres
      </Typography>

      <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 2, justifyContent: 'center', maxWidth: 720 }}>
        <Card sx={{ minWidth: 280, flex: 1 }}>
          <CardContent sx={{ textAlign: 'center', pt: 3, pb: 3 }}>
            <LoginIcon sx={{ fontSize: 48, color: 'primary.main', mb: 1 }} />
            <Typography variant="h6" gutterBottom>
              Login tradicional
            </Typography>
            <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
              Acceso con usuario y contraseña
            </Typography>
            <Button component={Link} href="/login" variant="contained" fullWidth>
              Iniciar sesión
            </Button>
          </CardContent>
        </Card>

        <Card sx={{ minWidth: 280, flex: 1 }}>
          <CardContent sx={{ textAlign: 'center', pt: 3, pb: 3 }}>
            <DashboardIcon sx={{ fontSize: 48, color: 'primary.main', mb: 1 }} />
            <Typography variant="h6" gutterBottom>
              Login moderno
            </Typography>
            <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
              Acceso directo al panel
            </Typography>
            <Button component={Link} href="/login" variant="outlined" fullWidth>
              Acceder con MUI
            </Button>
          </CardContent>
        </Card>
      </Box>

      <Button component={Link} href="/dashboard" color="inherit" size="small">
        Ir al dashboard →
      </Button>
    </Box>
  );
}
