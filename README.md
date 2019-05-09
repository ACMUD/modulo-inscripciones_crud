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

el ejecutarse el api esta se podra visializar en el localhost y en el puerto seteado en el caso por defecto `localhost:8080`

para poder hacer pruebas de querys ingresar a `localhost:8080/swagger` , alli se encontrara por cada tabla 5 metodos:

<details>
<summary>Get</summary>
sin usar filtro se obtendran todos los datos de la tabla:

- para hacer un filtro o un ordenamiento o limitacion alli apareceran las casillas.
- luego se le da al boton de jecutar y muestra los resultados en formato JSON
- adicionalmente mostrara la direccion a usar cuando se quiera hacer la consulta desde alguna aplicacion externa.
- ejemplo : query : `http://localhost:8080/v1/dia/?query=Activo:true` . para este caso en endponit es `dia` es decir el nombre de la tabla a la cual se hace la consulta y la segunda parte de este endpoint es cuando se hace un filtro en ese caso el endpoint completo es : `dia/?query=Activo:true`
- la respuesta dada por el Api es :
```javascript
[
    {
        "Id": 1,
        "Nombre": "Lunes",
        "Activo": true
    },
    {
        "Id": 2,
        "Nombre": "Martes",
        "Activo": true
    },
    {
        "Id": 3,
        "Nombre": "Miercoles",
        "Activo": true
    },
    {
        "Id": 4,
        "Nombre": "Jueves",
        "Activo": true
    },
    {
        "Id": 5,
        "Nombre": "Viernes",
        "Activo": true
    }
]
```
</details>

<details>
<summary>Post</summary>
para este se debe de mandar los datos en formato JSON y si ocurre algun error de datos el Api lo informara , si todo sucede normal mandara un `200` y la respuesta de la query mandara o un `201` o los datos ingresados de regreso, de estamanera se sabra que el post fue exitoso

- exiten casos en los cuales una clase tiene una llave foranea, para estos casos de tiene que agregar el id de este como un sub elemento del JSON 

- Ejemplo de como mandar los datos (tablas sin FK) :
```javascript
DatosGrupoInteres = {
    Abreviacion: 'GISAC'
    Activo:	true
    // se puede omitir el dato Descripcion ya que no es requerido, no dejarlo vacio, simplemente no agregarlo al JSON
    Descripcion: 'en este grupo se da introduccion a la segurudad informatica' 
    Nombre:	'Grupo de interes en seguridad informatica'
}
```

- Ejemplo de como mandar los datos (tablas CON FK) :
```javascript
DatosPersona = {
    Apellidos:	perez
    Codigo:	20101020000
    EstadoUsuario: {
        Id:	1
    }
    Nombres:	pepito
    Telefono:	1111111
    }
}
```
</details>

<details>
<summary>Get por ID</summary>
este por lo general no se usa ya que esto mismo se puede colocar en el otro metodo get,el id seria un elemto del filtro, adicionalmente este no trae la escalabilidad entre tablas.
</details>

<details>
<summary>Put</summary>
este es muy similar al post la unica diferencia es que este necesita que en el JSON o en la peticion se mande el Id del elemnto a actualizar.

Adicionalmente si solo se quieren actualizar pocos datos y no mandar todo en la peticion se puede solo mandar en el JSON el Id del Objeto a actualizar y los datos de este que se desean actualizar.

- Ejemplo de como mandar los datos (tablas CON FK) :
```javascript
DatosPersona = {
    Apellidos:	perez
    Codigo:	20101020000
    EstadoUsuario: {
        Id:	1
    }
    Nombres:	pepito
    Telefono:	1111111
    Id: 3
    }
}
```
</details>

<details>
<summary>Delete</summary>
para el Delete solo se necesita mandar el Id, pero este metodo no se tiene recomendado usar ya que por ello se crearon las tablas de `estados` o los valores de `activo` en la base de datos
</details>


documentacion realizada por [BOTOOM](https://github.com/BOTOOM)