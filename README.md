# 🛡️ Go-Network & Financial Bot

Sistemas de auditoría de red local y monitoreo de indicadores financieros desarrollado en **Go**. El proyecto implementa concurrencia de alto nivel para tareas de red e integración con APIs externas para la obtención de datos en tiempo real.

## ⚙️ Arquitectura del Sistema

El proyecto está diseñado bajo un modelo de **Arquitectura Multicapa**, separando la gestión de eventos de la lógica de negocio para asegurar la escalabilidad.



* **`main.go`**: Punto de entrada del sistema. Inicializa el bot y gestiona las señales de interrupción.
* **`manager.go`**: Orquestador central (Command Router). Procesa los `updates` de la API de Telegram y delega las peticiones a los servicios correspondientes.
* **`/handlers`**: Controladores de interfaz. Gestionan la lógica de respuesta y el formateo de mensajes (Markdown/HTML).
* **`/services`**: Capa de lógica de negocio. Contiene los clientes de red (ARP/Nmap) y los consumidores de APIs financieras.

---

## 🛠️ Especificaciones Técnicas

### 1. Módulo de Networking y Ciberseguridad
* **Escaneo Concurrente**: Uso de `sync.WaitGroup` y `Goroutines` para realizar barridos de red (Ping Sweeps) no bloqueantes.
* **Análisis de Capa 2 (ARP)**: Extracción y limpieza de la tabla de direcciones físicas del host mediante expresiones regulares (`regexp`).
* **Fingerprinting de Dispositivos**: Lógica de identificación basada en el escaneo de puertos críticos (TCP Scan) para diferenciar activos (PCs vs Móviles).
* **Integración Nmap**: Capacidad de orquestar escaneos profundos mediante la ejecución de binarios externos en entornos Linux.



### 2. Módulo de Inteligencia Financiera
* **Binance P2P Integration**: Consumo de API REST para la monitorización de tipos de cambio USDT/VES en el mercado paralelo.
* **BCV API.dolarvzla**: Sistema de consulta de indicadores oficiales del Banco Central de Venezuela.
* **Sanitización de Datos**: Implementación de filtros de codificación para asegurar compatibilidad UTF-8 en reportes financieros.

---

## 🚀 Instalación y Despliegue

### Requisitos
* Go 1.20+
* Token de Telegram Bot API

### Configuración
1. Clonar el repositorio.
2. Configurar las variables de entorno para el token del bot.
3. Compilar y ejecutar