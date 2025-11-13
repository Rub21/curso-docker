# 📝 Instrucciones: Notas de Orador

## ¿Qué son las Notas de Orador?

Las notas de orador son guías que te ayudan a presentar cada slide. **No se muestran en la proyección**, solo las ves tú en una pantalla separada.

## 🖥️ Configuración de 3 Pantallas

### Pantalla 1: Proyección (Para la audiencia)
- Abre `viewer.html?module=XX-modulo` en modo pantalla completa
- Esta es la que proyectas

### Pantalla 2: Notas de Orador (Solo para ti)
- Haz clic en el botón **"📝 Notas"** en el visor de slides
- Se abrirá en una nueva ventana/pestaña
- Colócala en tu segunda pantalla
- Se sincroniza automáticamente con las slides

### Pantalla 3: Terminal/Código (Opcional)
- Terminal para demostraciones
- Editor de código si necesitas
- O cualquier otra herramienta

## 🚀 Cómo Usar

### Paso 1: Abrir las Slides
```bash
npm start
# Navega a cualquier módulo y haz clic en "Ver Slides"
```

### Paso 2: Abrir las Notas
1. En el visor de slides, haz clic en el botón **"📝 Notas"** (arriba a la izquierda)
2. Se abrirá `speaker-notes.html` en una nueva ventana
3. Coloca esta ventana en tu segunda pantalla

### Paso 3: Sincronización Automática
- Cuando avances una slide (→ o Espacio), las notas avanzan automáticamente
- Las notas se sincronizan usando `localStorage`
- Funciona incluso si las ventanas están en diferentes pantallas

## 📋 Formato de las Notas

Cada módulo tiene un archivo `speaker-notes.md` con el formato:

```markdown
## Slide 1: Portada
- Punto 1 a mencionar
- Punto 2 a mencionar
- **Importante**: Algo destacado

## Slide 2: Título
- Explicación breve
- **Demostrar**: Algo que mostrar
```

## ✏️ Personalizar las Notas

Puedes editar cualquier `speaker-notes.md` y agregar:
- Puntos clave a mencionar
- Ejemplos adicionales
- Recordatorios
- Tiempos estimados
- Preguntas para hacer

## 🎯 Características

✅ **Sincronización automática**: Avanzan con las slides  
✅ **Navegación independiente**: Puedes navegar manualmente si quieres  
✅ **Formato Markdown**: Fácil de editar  
✅ **Diseño oscuro**: No distrae en presentación  
✅ **Contador de slides**: Sabes en qué slide estás  

## ⌨️ Controles

### En las Slides:
- `→` o `Espacio`: Siguiente slide
- `←`: Anterior slide
- `F`: Pantalla completa

### En las Notas:
- `→` o `Espacio`: Siguiente nota
- `←`: Anterior nota
- Botones: También puedes usar los botones

## 💡 Consejos

1. **Preparación**: Revisa las notas antes de presentar
2. **Personaliza**: Agrega tus propios puntos
3. **Práctica**: Prueba la sincronización antes
4. **Backup**: Ten las notas impresas por si acaso
5. **Flexibilidad**: No tienes que seguir las notas al pie de la letra

## 🔧 Solución de Problemas

### Las notas no se sincronizan
- Asegúrate de que ambas ventanas están en el mismo navegador
- Verifica que `localStorage` no esté bloqueado
- Refresca ambas ventanas

### No encuentro el botón de Notas
- Está arriba a la izquierda, al lado del botón "Volver"
- Si no aparece, verifica que el módulo tenga `speaker-notes.md`

### Las notas están desfasadas
- Haz clic en "Anterior" y luego "Siguiente" para resincronizar
- O refresca la ventana de notas

## 📝 Crear Notas para un Nuevo Módulo

1. Crea `XX-modulo/speaker-notes.md`
2. Usa el formato:
```markdown
# Notas de Orador - Módulo XX: Nombre

## Slide 1: Título
- Punto 1
- Punto 2

## Slide 2: Título
- Punto 1
- Punto 2
```

3. Una sección `## Slide N:` por cada slide
4. Las notas se cargarán automáticamente

---

¡Disfruta presentando con tus notas! 🎤

