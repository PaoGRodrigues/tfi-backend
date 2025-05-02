# Trabajo Final Integrador - Backend

## Introducción
Este repositorio contiene el backend del trabajo final **Análisis, visibilidad de tráfico y seguridad para usuarios finales en redes hogareñas**

## Documentación
Dirigirse a **este link** para obtener información sobre los endpoints disponibles y cómo se utiliza cada uno.

## Instalación
### Prerequisitos
* Configuración de la Raspberry Pi (ver XXXX).
* Instalación y configuración de la herramienta de monitoreo (ver XXXX).
* Instalación de git en la Raspberry Pi (ver XXXX).

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
* Copiar los comandos que están en el archivo **scripts/db.sql** de este repositorio.
* Ejecutar SQlite3 vía terminal. Pegar los comandos y apretar Enter.

### Ejecución de la aplicación
`Nota: Se debe ejecutar como root debido a que se modifican IPTables.`
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

## Ejecución de pruebas