# Workshop Platform (Taller)

Monorepo para la plataforma de gestión de talleres: recepción, inspección, reparaciones y reportes.

## Estructura

- **apps/admin-web** – Backoffice Next.js (App Router)
- **apps/desktop** – Cliente Tauri (reutiliza admin-web)
- **apps/mobile** – App React Native (Expo)
- **backend** – API en Go (Clean Architecture)
- **packages/types** – Tipos compartidos TS
- **packages/api-client** – SDK TypeScript para la API
- **db/migrations** – Migraciones PostgreSQL
- **infra** – Docker y Nginx para desarrollo/producción

## Requisitos

- Node.js 18+
- Go 1.21+
- PostgreSQL 15+
- Docker (opcional, para entorno local)

## Desarrollo local

### 1. Base de datos y backend

```bash
# Con Docker
cd infra/docker && docker-compose up -d

# Backend Go
cd backend && go mod download && go run ./cmd/api
```

### 2. Admin Web (Next.js)

```bash
npm install
npm run dev:admin
```

Admin: http://localhost:3000  
API: http://localhost:8080

### 3. Variables de entorno

Copiar `.env.example` a `.env` en la raíz y en `apps/admin-web` y ajustar valores.

## Scripts

| Script         | Descripción              |
|----------------|--------------------------|
| `npm run dev:admin`   | Levanta admin-web en dev |
| `npm run build:admin` | Build de admin-web       |
| `npm run dev:desktop` | Tauri desktop            |
| `npm run dev:mobile`  | Expo mobile              |

## Licencia

Privado – Fulltronik Adrian
