# Dockerfile con Multi-Stage (solución)
# Tamaño: ~150MB
# Etapa 1: Build
FROM node:18 AS builder
WORKDIR /app
COPY package.json package-lock.json .
RUN npm ci
COPY . .
RUN npm run build

# Etapa 2: Producción
FROM node:18-alpine
WORKDIR /app
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules
COPY package.json .
CMD ["node", "dist/index.js"]

