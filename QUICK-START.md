# 🚀 Inicio Rápido

## Ver el Curso en el Navegador

### Opción 1: Con npm (Recomendado)

```bash
# Instalar dependencias (solo la primera vez)
npm install

# Iniciar servidor y abrir navegador
npm start

# O sin abrir automáticamente
npm run serve
```

El sitio se abrirá automáticamente en `http://localhost:8000`

### Opción 2: Con Python

```bash
python3 -m http.server 8000
# Luego abre http://localhost:8000 en tu navegador
```

### Opción 3: Abrir directamente

```bash
# macOS
open index.html

# Linux
xdg-open index.html

# Windows
start index.html
```

## Ver Slides

### En VS Code

1. Instala la extensión "Marp for VS Code"
2. Abre cualquier `slides.md`
3. Presiona `Cmd+Shift+V` (Mac) o `Ctrl+Shift+V` (Windows)

### Exportar Slides

```bash
# Convertir todas las slides a HTML
npm run slides:html

# Convertir todas a PDF
npm run slides:pdf
```

## Estructura

```
curso-docker/
├── index.html          # 🏠 Página principal (abre esto)
├── 01-primer-contenedor/
│   ├── README.md       # 📖 Documentación completa
│   └── slides.md       # 📊 Slides para presentar
├── 02-dockerfile-basico/
│   └── ...
└── ...
```

## Comandos Disponibles

| Comando | Descripción |
|---------|-------------|
| `npm start` | Inicia servidor y abre navegador |
| `npm run serve` | Inicia servidor sin abrir navegador |
| `npm run slides:html` | Convierte slides a HTML |
| `npm run slides:pdf` | Convierte slides a PDF |

---

¡Disfruta del curso! 🐳

