# 📊 Instrucciones para Usar las Slides

## ¿Qué son las slides?

He creado presentaciones en formato **Marp** (Markdown Presentation Ecosystem) que te permiten:

- ✅ Ver las slides directamente en VS Code
- ✅ Exportar a PDF, HTML o PowerPoint
- ✅ Editar fácilmente (es Markdown)
- ✅ Presentar de forma profesional

## Opción 1: Ver en VS Code (Recomendado)

### Instalar la extensión Marp

1. Abre VS Code
2. Ve a Extensiones (Cmd+Shift+X / Ctrl+Shift+X)
3. Busca "Marp for VS Code"
4. Instala la extensión de `marp-team`

### Usar las slides

1. Abre cualquier archivo `slides.md` en VS Code
2. Presiona `Cmd+Shift+P` (Mac) o `Ctrl+Shift+P` (Windows/Linux)
3. Escribe "Marp: Export slide deck"
4. Elige el formato:
   - **HTML**: Para ver en navegador
   - **PDF**: Para imprimir o compartir
   - **PPTX**: Para PowerPoint

### Ver preview en VS Code

1. Abre `slides.md`
2. Presiona `Cmd+Shift+V` (Mac) o `Ctrl+Shift+V` (Windows/Linux)
3. O haz clic en el ícono de preview en la esquina superior derecha

## Opción 2: Ver en Navegador (HTML)

### Usando Marp CLI

```bash
# Instalar Marp CLI globalmente
npm install -g @marp-team/marp-cli

# Convertir slides a HTML
marp 01-primer-contenedor/slides.md -o 01-primer-contenedor/slides.html

# Ver en navegador
open 01-primer-contenedor/slides.html
```

### Convertir todas las slides

```bash
# Desde el directorio raíz del curso
for dir in */; do
  if [ -f "${dir}slides.md" ]; then
    marp "${dir}slides.md" -o "${dir}slides.html"
  fi
done
```

## Opción 3: Presentar con Reveal.js

Si prefieres una presentación más interactiva, puedes usar reveal.js:

```bash
# Instalar reveal-md
npm install -g reveal-md

# Presentar slides
reveal-md 01-primer-contenedor/slides.md
```

## Opción 4: Sitio Web de Navegación

He creado un `index.html` que puedes abrir directamente en tu navegador:

```bash
# Abrir en navegador
open index.html  # macOS
xdg-open index.html  # Linux
start index.html  # Windows
```

O simplemente haz doble clic en `index.html`

## Estructura de las Slides

Cada módulo tiene un archivo `slides.md` con:

- Portada del módulo
- Objetivos y temas
- Contenido estructurado
- Ejemplos de código
- Práctica
- Preguntas

## Personalizar las Slides

Puedes editar cualquier `slides.md` y cambiar:

- **Tema**: Cambia `theme: default` por otros temas disponibles
- **Estilo**: Modifica la sección `style:` en el frontmatter
- **Contenido**: Agrega o quita slides según necesites

## Ejemplo de Uso en Clase

1. **Preparación**:
   - Abre las slides en VS Code
   - Exporta a PDF o HTML
   - Ten el README abierto para referencia

2. **Durante la clase**:
   - Presenta las slides (PDF o HTML)
   - Demuestra los comandos en terminal
   - Usa el README para ejercicios prácticos

3. **Después de la clase**:
   - Comparte las slides con los estudiantes
   - Comparte el código de ejemplo

## Consejos

- ✅ Usa modo presentación en VS Code (F11)
- ✅ Exporta a PDF para imprimir o compartir
- ✅ Combina slides con demostraciones en vivo
- ✅ Usa el README para ejercicios prácticos
- ✅ Personaliza las slides según tu audiencia

## Recursos Adicionales

- [Documentación de Marp](https://marp.app/)
- [Temas de Marp](https://github.com/marp-team/marp-themes)
- [Marp CLI](https://github.com/marp-team/marp-cli)

---

¡Disfruta presentando el curso! 🎓

