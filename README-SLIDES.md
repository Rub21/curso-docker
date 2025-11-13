# 📊 Guía Rápida de Slides

## ¿Qué son las slides?

Presentaciones en formato **Marp** (Markdown Presentation) para cada módulo del curso.

## 🚀 Inicio Rápido

### Ver slides en VS Code

1. **Instala la extensión Marp:**
   - Abre VS Code
   - Extensiones → Busca "Marp for VS Code"
   - Instala

2. **Abre cualquier slides.md:**
   - Navega a cualquier módulo
   - Abre `slides.md`
   - Presiona `Cmd+Shift+V` (Mac) o `Ctrl+Shift+V` (Windows)

3. **Exporta a PDF/HTML:**
   - `Cmd+Shift+P` → "Marp: Export slide deck"
   - Elige formato (PDF, HTML, PPTX)

### Ver en navegador

```bash
# Instalar Marp CLI
npm install -g @marp-team/marp-cli

# Convertir a HTML
marp 01-primer-contenedor/slides.md -o 01-primer-contenedor/slides.html

# Abrir en navegador
open 01-primer-contenedor/slides.html
```

### Usar el sitio web

```bash
# Abrir index.html directamente
open index.html

# O servir con Python
python3 -m http.server 8000
# Luego abre http://localhost:8000
```

## 📋 Módulos con Slides

- ✅ Módulo 01: Primer Contenedor
- ✅ Módulo 02: Dockerfile Básico
- ✅ Módulo 05: Docker Compose Básico

*Puedes crear slides para los demás módulos siguiendo el mismo formato*

## 💡 Consejos para Presentar

1. **Preparación:**
   - Exporta slides a PDF para respaldo
   - Ten el README abierto para ejercicios
   - Prepara ejemplos en terminal

2. **Durante la clase:**
   - Usa modo presentación (F11 en navegador)
   - Demuestra comandos en vivo
   - Combina slides con práctica

3. **Después:**
   - Comparte slides con estudiantes
   - Comparte código de ejemplo

## 🎨 Personalizar Slides

Edita cualquier `slides.md` y modifica:

```yaml
---
marp: true
theme: default  # Cambia a 'gaia' o 'uncover'
paginate: true
header: 'Tu Título'
---
```

## 📚 Recursos

- [Documentación Marp](https://marp.app/)
- [Temas disponibles](https://github.com/marp-team/marp-themes)
- [Guía completa](./INSTRUCCIONES-SLIDES.md)

---

¡Disfruta presentando! 🎓

