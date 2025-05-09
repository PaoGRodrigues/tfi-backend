# Trabajo Final Integrador - Backend

![Go](https://img.shields.io/badge/code-Golang-blue?logo=go)

## Introducción
Este repositorio contiene el backend del trabajo final **Análisis, visibilidad de tráfico y seguridad para usuarios finales en redes hogareñas**

Esta aplicación es la intermediaria entre la herramienta de monitoreo de tráfico y el frontend. Su función principal es obtener información y ofrecer las funcionalidades que la herramienta no posee en su versión _Community_. 

Entre las funcionalidades se encuentran
* Obtención y guardado de flujos de tráfico
* Obtención y guardado de _hosts_
* Cantidad de tráfico por país y por destino
* Bloqueo de tráfico
* Envío de notificaciones

## Documentación
Dirigirse a **este link** para obtener información sobre los endpoints disponibles y cómo se utiliza cada uno.

## Instalación
### Prerrequisitos
Esta solución está pensada para ser desplegada en una Raspberry Pi.
* Configuración del entorno (ver [TFI](https://github.com/PaoGRodrigues/tfi)).
* Instalación y configuración de la herramienta de monitoreo (ver [TFI](https://github.com/PaoGRodrigues/tfi)).
* Instalar Git
* Instalar Golang

### Descarga y build
* Hacer un clone del repositorio en la Raspberry Pi.
``` 
$ git clone https://github.com/PaoGRodrigues/tfi-backend
```
* Moverse dentro del directorio y ejecutar el build
``` 
$ cd ~/tfi-backend/app
$ go build .
```

### Configuración de la base de datos
* Instalar SQlite3
``` 
$ sudo apt install sqlite3
```
* Crear el archivo donde se guardarán los datos.
``` 
$ sqlite3 file.sqlite
```
* Copiar los comandos que están en el archivo [db.sql](/scripts/db.sql) de este repositorio.
* Ejecutar SQlite3 vía terminal. Pegar los comandos y apretar Enter.

Con estos pasos ya está creada la Base de datos.

### Ejecución de la aplicación
`Nota: El backend debe ejecutarse como root debido a que se modifican las IPTables.`
```
$ sudo su
$ cd /home/pi/tfi-backend
$ ./app/app -s=prod -ip=192.168.0.13 -pr=3000 -u=Admin -p=Admin -db=file.sqlite
```

#### Parámetros
| Parámetro | Descripción |
|----------|-------------|
| `-s`     | Entorno de ejecución. Si se omite, se ejecuta en un entorno de prueba, lo que genera datos de prueba. Para ejecutar en un entorno que consuma la herramienta de monitoreo indicar `prod`. |
| `-ip`    | Dirección IP de la herramienta. En este caso, se debe indicar IP de la interfaz Ethernet de la Raspberry Pi. |
| `-pr`    | Puerto donde la herrammienta escucha los requests. |
| `-u`     | Usuario de la herramienta. |
| `-p`     | Contraseña del usuario de la herramienta. |
| `-db`    | Ruta al archivo de base de datos SQLite a utilizar. |

### Configuración del bot de Telegram
Para configurar el bot de Telegram, ver [Configuración del bot de Telegram](docs/TelegramConfig.md)

## Documentación de la API
Para consultar la documentación de la api:
1. Sin el backend en ejecución
    * Clonar el repositorio.
    * En el root del repositorio, ejecutar ``` $ python3 -m http.server 8080```
    * Para acceder a la documentación desde el navegador: http://localhost:8080/docs/swagger-ui/
2. Con el backend en ejecución en la Raspberry Pi
    * Seguir los pasos anteriores indicados en este Readme para la configuración de la Raspberry Pi, la herramienta y el backend. 
    * Antes de ejecutar la aplicación, en el atributo _server_, reemplazar _url-base_ por la IP actual de la Raspberry Pi en el archivo [.yml](docs/openapi_trabajo_final.yaml).
    ```
    servers:
        - url: "http://url_base:8080"
        - description: _IP_ de la Raspberry Pi. 
    ```
    * Una vez ejecutado el backend, se puede encontrar la documentación en el endpoint http://**IP Raspberry**:8080/swagger desde el navegador. También se pueden consultar el resto de los endpoints desde el Swagger.

## Ejecución de pruebas
Esta colección de pruebas está pensada para ejecutarse desde una computadora conectada a la misma red de la Raspberry donde está corriendo el backend.

* Ejecutar el backend en la Raspberry Pi como _root_. 
* Importar la colección a Postman o Postman Web, como se indica en [este enlace](https://learning.postman.com/docs/getting-started/importing-and-exporting/importing-data/).
* Reemplazar la variable base_url en las variables de la colección por la IP de la Raspberry Pi.
* Se pueden correr las pruebas individuales en cada endpoint o bien ejecutar todas las pruebas a la vez desde el Runner de Postman. 
* Por defecto, se ejecutarán las pruebas que den Status code: 200. Algunas pruebas tienen la opción de ejecutarse enviando bodies incorrectos. Para eso, modificar la el valor de la variable de la collección test_env ponerla como error.
