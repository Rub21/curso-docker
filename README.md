# Curso de Docker - De Básico a Avanzado

Bienvenido al curso completo de Docker. Este curso está diseñado para llevarte desde los conceptos básicos hasta técnicas avanzadas de producción.

## 🎯 Formas de Usar Este Curso

### 📖 Opción 1: Navegación Web
Abre `index.html` en tu navegador para una interfaz visual con todos los módulos.

**Opción rápida con npm:**
```bash
npm start
# O simplemente
npm run dev
```

Esto abrirá automáticamente el navegador en `http://localhost:8000`

**Opción manual:**
```bash
open index.html  # macOS
xdg-open index.html  # Linux
start index.html  # Windows
```

### 📊 Opción 2: Slides para Presentaciones
Cada módulo incluye slides en formato **Marp** (Markdown) que puedes:
- Ver en VS Code con la extensión [Marp](https://marketplace.visualstudio.com/items?itemName=marp-team.marp-vscode)
- Exportar a PDF, HTML o PowerPoint
- Usar para dictar el curso

**Ver instrucciones detalladas:** [INSTRUCCIONES-SLIDES.md](./INSTRUCCIONES-SLIDES.md)

### 📝 Opción 3: Notas de Orador
Cada módulo incluye notas de orador que se sincronizan automáticamente con las slides.

**Características:**
- Se muestran en una ventana separada (no en la proyección)
- Se sincronizan automáticamente cuando avanzas las slides
- Perfecto para presentar con múltiples pantallas

**Ver instrucciones:** [INSTRUCCIONES-NOTAS-ORADOR.md](./INSTRUCCIONES-NOTAS-ORADOR.md)

### 📖 Opción 4: READMEs Tradicionales
Cada módulo tiene un README.md completo con toda la información y ejercicios.

## 📚 Estructura del Curso

### [00 - Conceptos Básicos](./00-conceptos-basicos/)
Comprende los conceptos fundamentales de Docker antes de empezar.

**Temas:**
- ¿Qué es Docker y para qué sirve?
- Conceptos: Imagen, Contenedor, Registry
- Ventajas y casos de uso
- Arquitectura básica

### [01 - Primer Contenedor Sencillo](./01-primer-contenedor/)
Aprende a ejecutar tu primer contenedor Docker y los comandos fundamentales.

**Temas:**
- Comandos básicos: `docker run`, `docker ps`, `docker stop`
- Diferencia entre imagen y contenedor
- Ejecutar contenedores interactivos y en segundo plano

### [02 - Dockerfile Básico](./02-dockerfile-basico/)
Crea tus propias imágenes Docker usando Dockerfile.

**Temas:**
- Sintaxis básica de Dockerfile
- Comandos: FROM, RUN, COPY, CMD, EXPOSE
- Construir y ejecutar imágenes personalizadas

### [03 - Volúmenes](./03-volumenes/)
Aprende a persistir datos y compartir archivos entre host y contenedor.

**Temas:**
- Bind mounts vs Named volumes
- Persistir datos de bases de datos
- Desarrollo con hot-reload

### [04 - Redes](./04-redes/)
Gestiona redes Docker para conectar contenedores entre sí.

**Temas:**
- Tipos de redes: bridge, host, none
- Redes personalizadas
- Comunicación entre contenedores

### [05 - Docker Compose Básico](./05-docker-compose-basico/)
Orquesta múltiples contenedores con un solo archivo.

**Temas:**
- Sintaxis de docker-compose.yml
- Servicios, volúmenes y redes
- Comandos básicos de Docker Compose

### [06 - Multi-Stage Builds](./06-multi-stage-builds/)
Crea imágenes Docker más pequeñas y optimizadas.

**Temas:**
- Múltiples etapas en Dockerfile
- Reducir tamaño de imágenes
- Separar dependencias de build y runtime

### [07 - Docker Compose Avanzado](./07-docker-compose-avanzado/)
Configuraciones avanzadas para aplicaciones complejas.

**Temas:**
- Override files y entornos múltiples
- Healthchecks y dependencias
- Escalado de servicios
- Profiles

### [08 - Optimización de Imágenes](./08-optimizacion-imagenes/)
Técnicas avanzadas para optimizar tamaño y rendimiento.

**Temas:**
- Reducir tamaño de imágenes
- Mejorar tiempos de build
- Optimizar capas y caché
- .dockerignore

### [09 - Healthchecks y Logging](./09-healthchecks-logging/)
Implementa healthchecks y gestiona logs en contenedores.

**Temas:**
- Healthchecks en Dockerfiles y Docker Compose
- Gestionar logs de contenedores
- Drivers de logging y rotación

### [10 - Producción y Mejores Prácticas](./10-produccion-mejores-practicas/)
Mejores prácticas para ejecutar Docker en producción.

**Temas:**
- Seguridad en contenedores
- Gestión de secrets
- Estrategias de despliegue
- Monitoreo y observabilidad
- Backup y recuperación

## 🚀 Cómo usar este curso

1. **Sigue el orden**: Los módulos están diseñados para ser completados secuencialmente.
2. **Practica**: Ejecuta todos los comandos y ejemplos.
3. **Experimenta**: Modifica los ejemplos y prueba variaciones.
4. **Lee los READMEs**: Cada módulo tiene documentación detallada.

## 📋 Requisitos Previos

- Conocimientos básicos de línea de comandos (bash/zsh)
- Docker instalado en tu sistema
- Editor de texto o IDE

## 🔧 Instalación de Docker

### macOS
```bash
# Usando Homebrew
brew install --cask docker

# O descarga desde: https://www.docker.com/products/docker-desktop
```

### Linux (Ubuntu/Debian)
```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER
```

### Windows
Descarga Docker Desktop desde: https://www.docker.com/products/docker-desktop

## ✅ Verificar Instalación

```bash
docker --version
docker-compose --version
docker run hello-world
```

## 📖 Recursos Adicionales

- [Documentación oficial de Docker](https://docs.docker.com/)
- [Docker Hub](https://hub.docker.com/)
- [Best Practices](https://docs.docker.com/develop/dev-best-practices/)

## 🎯 Objetivos del Curso

Al finalizar este curso serás capaz de:

- ✅ Crear y gestionar contenedores Docker
- ✅ Construir imágenes personalizadas con Dockerfile
- ✅ Gestionar volúmenes y redes
- ✅ Orquestar aplicaciones con Docker Compose
- ✅ Optimizar imágenes para producción
- ✅ Implementar healthchecks y logging
- ✅ Aplicar mejores prácticas de seguridad
- ✅ Desplegar aplicaciones en producción

## 🤝 Contribuciones

Si encuentras errores o tienes sugerencias, ¡son bienvenidas!

## 📝 Licencia

Este curso es de uso educativo. Siéntete libre de usarlo y modificarlo según tus necesidades.

---

**¡Comienza con el [Módulo 01](./01-primer-contenedor/) y disfruta aprendiendo Docker!** 🐳

