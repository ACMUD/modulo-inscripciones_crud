# Api modulo de inscripciones ACM-UD

Api para la gestion de inscripcion de estudiantes o personas externas a la universidad en la participacion de grupos de insteres y eventos desarrollados por el grupo ACM - UD

## Requerimientos

<details>
<summary>Postgres</summary>
Ejemplo de instalacion de postgres para manjaro , si usted tiene otra distribucion o SO diferente por favor buscar el como instalarlo.

```javascript
sudo pacman -S postgresql
```

```javascript
sudo su postgres -l # or sudo -u postgres -i
initdb --locale $LANG -E UTF8 -D '/var/lib/postgres/data/'
exit
```
</details>

<details>
<summary>Go version 1.12</summary>
Ejemplo de instalacion de Go para manjaro , si usted tiene otra distribucion o SO diferente por favor buscar el como instalarlo.

```javascript
sudo pacman -S go
```

en otros otras distribuciones sera usar `apt-get` o `yum install` , pero el paquete se encuentra como ``go`` o ``golang``

Luego se configuraran las variables de entorno.`

Luego se configuraran las variables de entorno.

```javascript
sudo nano /etc/profile.d/goenv.sh
```

Dentro del archivo colocar lo siguiente

```javascript
export GOROOT=/usr/lib/go    
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

en la carpeta `home` se debe crear la carpeta `go`

Luego para recargar las varialbes de entorno se ejecuta:

```javascript
source /etc/profile.d/goenv.sh
```

</details>

<details>
<summary>Beego y Bee</summary>
Se necesita instalar beego y bee para su corrcta ejecucion.

```javascript
go get -u github.io/astaxie/beego
```

```javascript
go get -u github.io/beego/bee
```
</details>

---
### Descargar el proyecto

no sejecutar `git clone` ya que puede generar problemas si se clona en cualquier directorio, para evitar ello se usara `go get` con ello automaticamente se creara el siguiente directorio `go/src/github.com/ACMUD/modulo-inscripciones_crud` en esta carpeta estara el proyecto listo para su ejecucion.

```javascript
go get github.com/ACMUD/modulo-inscripciones_crud
```
---
### Ejecucucion

Para la ejecucion del Api se creo un archivo llamado `ejec.sh` en este estan las sentencias para la correcta ejecucion del api.

El archivo debe tener permisos de ejecucion, si no los tiene, ejecutar dentro de la carpeta del proyecto:

```javascript
chmod 777 ./eject.sh
```
<details>
<summary>variables del bash</summary>
En el bash se encuentran la siguientes variables:

- INSCRIPCION_DB
    - nombre de la base de datos, este se cambia en caso de que el script sea ejecutado en una base de datos diferente al valor ya seteado en el bash.
- INSCRIPCION_PASSDB
    - password de la base de datos o conexion.
- INSCRIPCION_URLS
    - url de conexion a la base de datos, en este caso se tiene por defecto localhost, generalmente se puede dejar tal como esta.
- INSCRIPCION_USERDB
    - usuario con permisos para la base de datos.
- INSCRIPCION_HTTP_PORT
    - puerto de ejecucion el api, por defecto esta el 8080, pero este se puede cambiar a preferencia.

</details>