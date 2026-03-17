### 1. Modelo de datos

Primero define la estructura que representa la configuración.

Ejemplo de cosas que podría tener una config:

* nombre del servicio
* puerto
* entorno
* si está activo o no
* lista de dependencias
* variables de entorno

No necesitas empezar con algo complejo. Lo importante es que la estructura te permita practicar validación y transformación.

## Ejemplo de archivo JSON

```json
{
  "service_name": "payments-api",
  "port": 8080,
  "environment": "dev",
  "enabled": true,
  "dependencies": ["postgres", "redis"]
}
```

## Ejemplo de archivo YAML

```yaml
service_name: payments-api
port: 8080
environment: dev
enabled: true
dependencies:
  - postgres
  - redis
```

## Funcionalidades esperadas

### Leer archivo

La CLI debe recibir la ruta de un archivo y detectar o asumir el formato.

Ejemplo:

```bash
go run ./cmd/configcli --file ./testdata/valid.json
```

### Validar estructura

Debe revisar reglas como:

* `service_name` no puede estar vacío
* `port` debe ser mayor que 0
* `environment` debe tener un valor permitido como `dev`, `qa`, `prod`
* `dependencies` puede ser opcional, según tu diseño

### Transformar datos

Ejemplos de transformación:

* convertir JSON a YAML
* convertir YAML a JSON
* normalizar nombres
* ordenar dependencias

### Filtrar contenido

Ejemplos:

* mostrar solo servicios activos
* mostrar solo configs de `prod`
* mostrar configs que dependan de `redis`

### Exportar resultados

Ejemplos:

* guardar el archivo transformado
* guardar solo las configuraciones filtradas
* exportar validaciones a otro archivo

## Posibles flags

Puedes empezar con flags como estas:

* `--file`: ruta del archivo de entrada
* `--format`: formato de salida (`json` o `yaml`)
* `--validate`: valida la estructura
* `--filter-env`: filtra por entorno
* `--only-enabled`: muestra solo servicios habilitados
* `--out`: ruta del archivo de salida

Ejemplos:

```bash
go run ./cmd/configcli --file ./testdata/valid.json --validate
```

```bash
go run ./cmd/configcli --file ./testdata/valid.yaml --format json
```

```bash
go run ./cmd/configcli --file ./testdata/services.yaml --filter-env prod --only-enabled
```

## Flujo recomendado de implementación

### Paso 1

Crear el proyecto y definir el `struct` principal.

Preguntas guía:

* ¿Qué campos mínimos tendrá una config?
* ¿Qué campos serán obligatorios?
* ¿Qué tipos usaré?

### Paso 2

Implementar la lectura del archivo.

Preguntas guía:

* ¿Voy a leer todo el archivo de una vez o línea por línea?
* ¿Cómo manejaré errores si el archivo no existe?
* ¿Cómo validaré que la ruta fue enviada correctamente?

### Paso 3

Parsear JSON a `structs`.

Preguntas guía:

* ¿Cómo mapeo los nombres del JSON al struct?
* ¿Qué errores pueden aparecer si el archivo está mal formado?

### Paso 4

Agregar validación.

Preguntas guía:

* ¿Qué reglas son de formato?
* ¿Qué reglas son de negocio?
* ¿Devuelvo el primer error o una lista de errores?

### Paso 5

Agregar soporte YAML.

Preguntas guía:

* ¿Qué cambia entre parsear JSON y YAML?
* ¿Puedo reutilizar parte de la lógica?

### Paso 6

Agregar transformación y exportación.

Preguntas guía:

* ¿La salida se imprime en consola o se guarda en archivo?
* ¿Cómo decidir el formato de salida?
* ¿Qué parte del código se encarga solo de escribir resultados?

### Paso 7

Agregar filtros.

Preguntas guía:

* ¿Los filtros aplican sobre una sola config o una lista?
* ¿Qué pasa si no encuentra resultados?

### Paso 8

Escribir tests.

Preguntas guía:

* ¿Qué funciones son fáciles de testear?
* ¿Qué casos válidos e inválidos debo cubrir?
* ¿Puedo usar archivos de ejemplo en `testdata/`?

## Reglas de validación sugeridas

Puedes comenzar con estas:

* `service_name` es obligatorio
* `port` debe ser mayor que 0 y menor que 65536
* `environment` debe ser uno de: `dev`, `qa`, `prod`
* si `enabled` es `true`, entonces el `service_name` debe existir
* no debe haber dependencias duplicadas

## Errores que deberías manejar bien

* archivo no encontrado
* formato inválido
* JSON o YAML mal formado
* campos obligatorios faltantes
* valor de puerto inválido
* formato de salida no soportado
* error al escribir archivo de salida

## Ejemplo de salida esperada

### Validación exitosa

```text
Config válida
service_name: payments-api
environment: dev
port: 8080
```

### Validación con errores

```text
Config inválida:
- service_name es obligatorio
- port debe ser mayor que 0
- environment debe ser uno de: dev, qa, prod
```

## Testing

Debes agregar tests al menos para:

* lectura de archivos
* parsing
* validación
* filtros
* transformación

Casos mínimos recomendados:

* archivo válido JSON
* archivo inválido JSON
* archivo válido YAML
* archivo con campos faltantes
* archivo con valores inválidos
* exportación correcta
* filtro sin resultados

## Librerías que podrías investigar

Para JSON puedes usar la librería estándar.

Para YAML, puedes investigar una librería popular de YAML para Go.

La idea no es copiar todo, sino buscar solo lo necesario para:

* parsear YAML
* serializar YAML

## Cómo trabajar este proyecto para aprender de verdad

No intentes hacerlo completo en un solo día.

Trabájalo en partes pequeñas:

### Día 1

* crear proyecto
* definir `struct`
* leer archivo JSON

### Día 2

* parsear JSON
* imprimir contenido
* manejar errores básicos

### Día 3

* agregar validación
* mostrar errores amigables

### Día 4

* agregar soporte YAML

### Día 5

* agregar transformación JSON <-> YAML

### Día 6

* agregar filtros

### Día 7

* exportar resultados
* escribir tests

## Qué debes evitar

* meter demasiadas capas desde el inicio
* crear interfaces sin necesitarlas
* mezclar parsing, validación y output en un solo archivo
* ignorar errores
* hacer una arquitectura demasiado grande para un proyecto inicial

## Qué deberías aprender al terminar

Al finalizar este proyecto deberías poder:

* leer y procesar archivos en Go
* modelar datos con `structs`
* trabajar con JSON y YAML
* usar flags en una CLI
* manejar errores de forma idiomática
* escribir tests básicos
* organizar un proyecto Go de forma clara

## Posibles mejoras futuras

Cuando termines la primera versión, puedes agregar:

* soporte para múltiples archivos
* validación con reglas más complejas
* salida coloreada en consola
* resumen de errores por archivo
* integración con Docker
* pipeline de CI
* benchmark simple
* documentación de uso más completa

## Estado actual

* [ ] lectura de archivos
* [ ] parsing JSON
* [ ] parsing YAML
* [ ] validación
* [ ] filtros
* [ ] transformación
* [ ] exportación
* [ ] tests

## Notas personales

Usa esta sección para anotar cosas que vas aprendiendo:

* errores comunes
* decisiones de diseño
* feedback de PRs
* cosas que harías distinto en la siguiente versión

## Comandos útiles

Inicializar módulo:

```bash
go mod init service-config-cli
```

Ejecutar proyecto:

```bash
go run ./cmd/configcli --file ./testdata/valid.json --validate
```

Ejecutar tests:

```bash
go test ./...
```

## Enfoque de aprendizaje

La meta de este proyecto no es solo “hacer que funcione”, sino practicar cómo pensar soluciones en Go.

Regla recomendada:

1. primero pensar el problema
2. dividirlo en partes
3. implementar una sola parte a la vez
4. pedir ayuda solo para desbloquearte
5. refactorizar después

## Autor

Proyecto de práctica para reforzar Go con enfoque backend/cloud.

````

## Mi recomendación
Para que ese README te sirva de verdad, te conviene agregarle 3 cosas tuyas:

1. **MVP exacto de la semana**
2. **reglas de validación que tú vas a implementar**
3. **ejemplos reales de input/output**

También te dejo una versión más corta y profesional para la parte inicial del README, por si quieres que se vea más “GitHub-ready”:

```md
# Service Config CLI

Herramienta de línea de comandos desarrollada en Go para leer, validar, transformar, filtrar y exportar archivos de configuración en formato JSON y YAML.

## Objetivo

Practicar fundamentos de Go aplicados a un caso real de backend/cloud:

- lectura de archivos
- modelado con `structs`
- parsing JSON/YAML
- validación de datos
- uso de flags
- manejo de errores
- testing

## Funcionalidades
- lectura de archivos JSON/YAML
- validación de estructura
- transformación de formato
- filtrado de contenido
- exportación de resultados
````

Puedo también dejarte el README **más profesional y limpio para subirlo directo a GitHub**, o uno **más didáctico tipo guía de estudio**, según prefieras.
