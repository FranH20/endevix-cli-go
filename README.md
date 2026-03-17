# Service Config CLI

CLI en Go para leer, validar, transformar, filtrar y exportar archivos de configuración en formato JSON o YAML.

## Objetivo

Este proyecto fue creado para practicar Go de manera aplicada, construyendo una herramienta de línea de comandos con casos reales de backend y cloud.

La CLI debe ser capaz de:

- leer archivos JSON y YAML
- validar su estructura
- transformar datos
- filtrar contenido
- exportar resultados

## Qué se practica en este proyecto

Este proyecto está diseñado para reforzar:

- manejo de archivos con `os`, `io`, `bufio`
- modelado con `structs`
- parsing de JSON y YAML
- uso de flags y argumentos
- manejo de errores idiomático en Go
- organización de paquetes
- testing

## Caso de uso

Supongamos que tienes archivos de configuración de servicios o aplicaciones y quieres verificar que cumplan un formato esperado.

Ejemplos de tareas que debería soportar esta herramienta:

- validar que existan campos obligatorios
- detectar configuraciones inválidas
- convertir JSON a YAML o YAML a JSON
- filtrar servicios activos
- exportar el resultado validado o transformado

## Alcance del proyecto

### Versión 1
- leer un archivo JSON
- parsearlo a `structs`
- validar campos obligatorios
- mostrar errores claros en consola

### Versión 2
- soportar YAML
- transformar JSON <-> YAML
- agregar filtros por campos
- exportar resultados a archivo

### Versión 3
- mejorar mensajes de error
- agregar tests
- soportar múltiples archivos
- mejorar experiencia de uso de la CLI

## Posible estructura del proyecto

```bash
service-config-cli/
├── cmd/
│   └── configcli/
│       └── main.go
├── internal/
│   ├── config/
│   │   ├── model.go
│   │   ├── parser.go
│   │   ├── validator.go
│   │   ├── transformer.go
│   │   └── filter.go
│   └── output/
│       └── writer.go
├── testdata/
│   ├── valid.json
│   ├── invalid.json
│   ├── valid.yaml
│   └── invalid.yaml
├── go.mod
├── go.sum
└── README.md