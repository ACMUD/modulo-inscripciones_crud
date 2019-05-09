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

en otros otras distribuciones sera usar `apt-get` o `yum install` , pero el paquete se encuentra como ``go``
</details>


<!-- ## algo -->

<!-- ```javascript
go get github.com/ACMUD/modulo-inscripciones_crud
``` -->
