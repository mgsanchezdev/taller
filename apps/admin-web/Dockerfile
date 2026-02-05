# Build from monorepo root: docker build -f apps/admin-web/Dockerfile .
FROM node:20-alpine AS deps
WORKDIR /workspace
COPY package.json package-lock.json* ./
COPY apps/admin-web/package.json ./apps/admin-web/
COPY packages/types/package.json ./packages/types/
COPY packages/api-client/package.json ./packages/api-client/
RUN npm ci

FROM node:20-alpine AS builder
WORKDIR /workspace
COPY --from=deps /workspace/node_modules ./node_modules
COPY . .
RUN npm run build --workspace=admin-web

FROM node:20-alpine AS runner
WORKDIR /app
ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
COPY --from=builder /workspace/apps/admin-web/public ./public
COPY --from=builder /workspace/apps/admin-web/.next/standalone ./
COPY --from=builder /workspace/apps/admin-web/.next/static ./.next/static
EXPOSE 3000
CMD ["node", "server.js"]
