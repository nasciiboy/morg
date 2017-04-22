la aleatoriadad llevo a programar morg en golang (aunque creo que ya se como
portarlo a c) de momento a instalar go

## instalar go (version GNU)

primero ve a la direccion https://golang.org/dl/ y clica el enlace de descarga
donde diga algo como (al dia de escribir esto)
<em>go1.8.linux-amd64.tar.gz</em> (Cambia el <em>amd64</em> por la arquitectura
de tu equipo)

Si todo va bien tendras el comprimido en `$HOME/Downloads`, abrimos un terminal
y vamos hay

### segun la guia oficial

pasa a modo root, y ejecuta

``` sh
tar -C /usr/local -xzf go*.tar.gz
rm go*.tar.gz
```

esto crea la carpeta `/usr/local/go` que en su interior contiene todo lo
necesario para ejecutar go y sus cosillas. Si la curiosidad te puede ve a dicha
carpeta y observa que regalos contiene (me parecio interesante la estructura de
los articulos en la carpeta`blog`, al parecer los de google tambien estan
ideando su lenguaje de marcas ligeras).

Ahora bien, la carpeta `/usr/local/go` tiene un directorio `bin` donde se
encuetra el compilador/interprete/formateador y otras cosas del lenguaje.

Da la casualidad, que nuestro shell no sabe eso, asi que hay que informarle
donde buscar estos binarios, para que luego, de forma comoda se pueda escribir
`go algo` y se ejecute sin mas.

pero antes, de configurar `$PATH`... go necesita una ruta donde colocara el
codigo de los proyectos, dependencias y demas cosas que descarguemos, esto se
establece en la variable de entorno `GOPATH`, en mi caso

``` sh
export GOPATH=/home/nasciiboy/go
```

ahora cuando le pida a `go`, que me instale algo, lo colacara dentro de la
carpeta <span class="file" >go</span> en mi
home... peero para ejecutar los binarios que se creen en dicha direccion tambien
debe agregarse esta direccion a `$PATH`. Vamos con ello

``` sh
export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin:
```

es decir, al final tenemos

``` sh
export GOPATH=/home/nasciiboy/go
export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin:
```

obviamente, queremos que los cambios sean permanentes, asi que agregamos la
configuracion en `~/.bashrc` o similar

### opcion personalizada

no me fio mucho de google, para reducir los daños, prefiero colocar los binarios
y todo lo demas en mi `$HOME`

sin pasar a root, ni naaah

``` sh
# se asume estar en ~/Downloads
tar -C . -xzf go*.tar.gz
rm go*.tar.gz
mv go ../.go
```

con esto dejamos los binarios y cosas de go en la carpeta `.go` en nuestro
`$HOME` configuramos variables de entorno

``` sh
export GOROOT=/home/nasciiboy/.go
export GOPATH=/home/nasciiboy/go
export PATH=$PATH:$GOPATH/bin:$GOROOT/bin/
```

(`GOROOT` segun la pagina de instalacion de golang, se necesita para
instalaciones a medida como en este caso)

Si utilizas `fish` en lugar de `bash`, como en mi caso, la configuracion se
agrega de la siguiente manera

``` sh
set --export GOROOT /home/nasciiboy/.go
set --export GOPATH /home/nasciiboy/go
set --export PATH   $PATH $GOPATH/bin $GOROOT/bin/
```

en  `~/.config/fish/config.fish`

## morg

vamos al meollo del asunto, ya con `go` en el sistema, instalamos morg:

``` sh
go get -v github.com/nasciiboy/morg
```

si ya tenemos morg y queremos actualizar

``` sh
go get -u -v github.com/nasciiboy/morg
```

(nota: no se para que sea la `v`, asi lo vi en un proyecto... y asi lo dejo)

para quien no lo sepa, morg es un lenguaje de marcas ligeras que pretende
dominar el mundo, terminar con el empleo, y otras cosas (probablemente demasiado
ambisiosas)

el sistema de documentacion aun esta en fase inicial y no es mi proposito ni
intencion definir o programar todos los aspectos, pero vamos a lo basico

### primer paso

creamos en nuestro editor de confiansa un fichero con (o sin) terminacion
`.morg`, para el caso  *ejemplo.morg*

en su interior ponemos algo de texto, como `hola morg`

o con un `echo`

```
echo "hola morg" > ejemplo.morg
```

y exportamos a html

``` sh
morg export ejemplo.morg
```

listo, se genera un fichero de nombre `ejemplo.html` con nuestro documento.

Agregemos mas cosas. morg utiliza **comandos** para configurar el documento:

```
..title    >  Morg, ejemplo y otras cosas
..author   >  nasciiboy
..lang     >  es
..date     >  2016
..tags     >  regex regexp c-prog algorithm
..licence  >  GNU FDL v1.3
..style    >  worg-data/worg.css
..options  >  toc
```

los comandos de configuracion tienen la estructura `..` (dos puntos) `comando`
(el comando) y `>` (mayor que) y luego sigue `argumento`, que es todo lo que se
encuentre en la misma linea de la declaracion, aunque puede extenderse por
varias lineas, siempre y cuando no se dejen lineas en blanco y se cumpla una
indentacion de dos espacios, como en

```
..title > este es un tutulo de documento muy muy
  extenso, como para ocupar multiples lineas
  y mas
```

un comando de configuracion finaliza cuando se encuentra la primer linea en
blanco o se incumple con la indentacion (en caso de que el argumento abarque
multiples lineas)

Ademas los comandos de configuracion deben colocarse al inicio del documento,
sin dejar espacios en blanco entre el inicio de la linea y la declaracion del
comando. Este <q>bloque de configuracion</q> terminara en cuanto aparesca el
primer parrafo o encabezado

el comando `style` sirve para establecer la ruta a una hoja de estilos, para el
caso un css. De momento [este](../worg-data.zip) es el que utilizo, puedes
encontrar la version *de desarrollo* dentro del codigo fuente, si vas a
`~/go/src/github.com/nasciiboy/morg/biskana/testdata/worg-data/worg.css`

si no se incluye un titulo (`..title > titulo`) en el documento, se asigna el
nombre del fichero como el titulo.

de momento el comando `options` permite tres cosas:

- Si colocamos `toc` se creara una tabla de contenido (indice)

- Con `highlight` se incluira el enlace a un script para resaltar el codigo
  fuente, este apunta a una carpeta de nombre `highlight` que debe encontrarse
  en la misma ruta que el fichero resultado

  Para optener tu carpeta `highlight` ve
  a [esta](https://highlightjs.org/download/) pagina y especifica que lenguajes
  deseas, descarga, descomprime y coloca en la misma carpeta del resultado, o en
  `~/go/src/github.com/nasciiboy/morg/biskana/testdata/` encortraras la carpeta
  `highlight` con todos los lenguajes disponibles, copia y mueve a tu directorio
  de trabajo

  Coloque un estilo por defecto en el codigo fuente, la forma mas sencilla de
  modificarlo es hacelo el fichero resultado (a mano), los estilos de resaltado
  disponibles estan dentro de `highlight/styles`

- Si por el contrario si (tienes instalado y) quieres que el resaltado se
  incluya con etiquetas utilizando `pygments`, debes utilazar esta opcion en
  options `..options > pygments`

  encontraras tambien dentro del [css](../worg-data.zip), etiquetas para
  establecer el estilo del resaltado.

  Puedes opteren mas estilos css para
  pygments [aqui](https://github.com/jwarby/jekyll-pygments-themes/), descarga
  el estilo, y agrega un comando `..style > direccion/a/mi/pygments.css`

### estructura

la estructura es fundamental en un lenguaje de marcas que se precie. En el caso
de morg (basado en org, md reestructure text, texinfo y otras cosas) se opta
por dividir el documento en secciones o encabezados

La sintaxis de un encabezado es muy simple:

```
* Encabezado principal
** Segundo nivel
*** 3er. nivel

    primer parrafo del 3er. nivel

    segundo parrafo del 3er. nivel

*** otro 3er. nivel

    mas texto

* Otro encabezado principal
```

(como vez, he indentado el texto que pertenece a cada encabezado (no es
necesario, pero a mi me gusta que se vea asi).

Toda linea que inicia con uno o varios `*` seguido por al menos un espacio en
blanco, es un encabezado.

El texto (en la misma linea) donde esta el o los `*` sera el titulo del encabezado.

El nivel del encabezado depende del numero de `*` al inicio de la linea, es
decir

- `*` == `h1`
- `**` == `h2`
- `***` == `h3`
- `****` == `h4`
- `*****` == `h5`

(en realidad, al exportar a html, los encabezados se desplazan en uno, el titulo
del documento es el que se corresponde con `h1`, pero lo dejo asi para evitar
esfuerzo mental)

A que encabezado pertenece el texto? al primero que aparezca por encima de
el. Si no tiene un encabezado encima pertenece al <q>titulo</q> del documento

por cierto, el titulo de los encabezados (al igual que el argumento de los
comandos de configuracion) puede extenderse en varias lineas siempre y cuando se
respete la indentacion

```
** encabezada dividido en multiples lineas
   por ser demasiado extenso y a modo
   de ejemplo
```


por cierto, todos los encabezados generan una referencia dentro del documento,
donde se substituyen los espacios en blanco por `-` (menos), por ejemplo

```
** El Raptor
```

generara

```
<h2 id="El-Raptor" >El Raptor</h2>
```

Si deseamos especificar el contenido de la referencia, se utiliza la sintaxis

```
** referencia personalizada <> texto del encabezado
```

que genera

```
<h2 id="referencia-personalizada" >texto del encabezado</h2>
```

### enfasis, bold, links, el comando `@`

otros lenguajes de marcas ligeras utilizan marcas varias, el problema con este
enfoque radica en naturaleza exotica de las marcas, que los limitan a dos o tres
formas de resaltado, antes de recurrir a etiquetas html (¡blasfemia!).

luego de meditar por 15 dias, 10 noches, ver texinfo y realizar una
investigacion con fondos provenientes del gobierno (nada extraordinario, solo un
par de millones) se descubrio *la sintaxis* de marcas perfecta

```
@x()
@x[]
@x{}
@x<>
```

donde `@` indica <em>a continuacion contenido especial</em>, `x` es un caracter
ascii imprimible, que describe el comando o accion a aplicar al contenido
delimitado por `{…}`, `(…)`, `<…>` o `[…]`

<em>por que una `@`?</em> Fuera de algun lenguaje exotico o el correo, podria
ser el signo menos utilizado y mas aun con la estructura `@x{}`

<em>y la `x`?</em> Un caracter ascii imprimible. Si hemos de necesitar mas
marcas que los caracteres ascii, mal vamos.

de momento, puedes utilizar las marcas de resaltado


```
- A :: acronym
- a :: abbr
- b :: bold
- c :: code
- e :: emph
- f :: file
- ' :: ‘samp’
- " :: cite
- i :: italic
- k :: kbd
- ^ :: sup
- _ :: sub
- m :: math
- q :: quote
- - :: —exp—
- s :: strike
- u :: underline
- v :: verbatim
- $ :: command
```


`^` sirve para generar superindices y `_` para los subindices

<em>y los `({[<>]})`?</em> Mas opciones, mas diversion.

Segun sea el contexto `{}` o `()` podrian requerir el *escape* de algun
caracter. Para minimizar la inclucion de signos extraños, los delimitadores se
aplican deacuerdo a la necesidad y gusto del <q>creador</q>.

importante resaltar que el signo `@` sirve tambien para escapar caracteres,
cuando no se cumple la secuencia `@x()`, por ejemplo, `@(` se substituye por
`(`, `@}` por `}`. (nota: dentro del texto o un comado `@` para que aparesca `@`
hay que colocar `@@`)

los comandos `@` pueden utilizarse en cualquier sitio, inclusive dentro del
titulo del documento, encabezados, tablas y demas, (solo no colocar en otros
comandos de configuracion que no sean `title` no contemplo esta sitiacion y no
tiene sentido...

los <q>comandos</q> `@` soportan anidamiento de otros comandos como en `@b(bold
@e(emph @i(italic)))`. En caso de no encontrar el operador de cierre de un
comando, la accion de este tendra efecto hasta el final del parrafo o encabezado
en cuestion

por si esto no fuera suficiente tenemos un comando especifico para el manejo de
enlaces

```
- l :: link
```

donde `@l(https://direccion)`, convierte en un enlace a su contenido, asi:

```
<a href="https://direccion" >https://direccion</a>
```

para que aparesca un texto distinto al de la direccion seria
`@l(https://direccion<>texto distinto)`


```
<a href="https://direccion" >texto distinto</a>
```

he? (movimiento de manos y aparece un arcoiris) la magia del comando `@`. En
realidad todos los comandos `@` tienen la estructura

```
@x(custom<>contenido)
```

donde `custom` es un parametro personalizado y opcional, y el `contenido` es el
contenido del comando, el cual siempre debe proporcionarse

Cuando un comando, por ejemplo, un enlace enlace, requiere un `custom` y este no se ha
especificado, se genera apartir del `contenido`, extrayendo las marcas expeciales
dejando unicamente el texto.

Cuando el comando no requiere de un `custom` y este es proporcionado, el comando
o lo ignora o se utiliza como identificador o etiqueta segun sea el caso (nota:
esto esta a consideracion y podria requerir sintaxis adicional)

Cuando el comando esta dentro de otro comando, el comando interno *pasa* su
contenido al comando externo, por ejemplo:

```
@l(objetivo<>lo que sea con @e(enfasis con @b(algo<>bold)))
```

genera

```
<a href="objetivo" >lo que sea con <em>enfasis con <b>bold</b></em></a>
```
y

```
@l(lo que sea con @e(enfasis con @b(algo<>bold)))
```

genera

```
<a href="lo-que-sea-con-enfasis-con-bold" >lo que sea con <em>enfasis con <b>bold</b></em></a>
```

y finalmente

```
@l(@b(bold)<>lo que sea)
```

genera

```
<a href="bold" >lo que sea</a>
```

el comando

```
- t :: target
```

sirve para crear una referencia dentro del documento. Por ejemplo

```
@t(objetivo interno)

@l(#objetivo interno)
```

donde `@t` declara el objetivo interno en el documento, y `@l` enlaza a dicho
objetivo. Recordar que  para los encabezados, el exportador genera  el <q>objetivo</q>
automaticamente

Cuando se exporta la referencia `@l(#referencia)` se elimina el signo `#` del
texto, es decir, genera

```
<a href="#referencia" >referencia</a>
```

### listados

```
- lista desordenada

  contenido de elemento

+ lista desordenada

1. lista ordenada numericamente

1) lista ordenada numericamente

a. lista ordenada alfabeticamente

a) lista ordenada alfabeticamente

   contenido de elemento a
```

El contenido de una lista debe indentarse segun la seccion de la que forme
parte. Se permite anidar listas dentro de otras listas, asi como otro tipo de
elementos del formato

### dialogos

```
> "Dialogo, Lorem ipsum dolor sit amet, consectetur adipiscing
  elit, sed eiusmod tempor incidunt ut labore et dolore magna
  aliqua. Ut enim ad minim veniam."
```

los dialogos tienen la mismas normas que una lista.

#### definiciones

```
- elemento :: definicion Lorem ipsum dolor sit amet, consectetur
  adipiscing elit, sed eiusmod tempor incidunt ut labore et
  dolore magna aliqua. Ut enim ad minim veniam.

+ elemento :: definicion Lorem ipsum dolor sit amet, consectetur
  adipiscing elit, sed eiusmod tempor incidunt ut labore et
  dolore magna aliqua. Ut enim ad minim veniam.
```

#### about's

en realidad no se como nombrar estos elementos, asi que por ahora se llaman
<q>acerca de</q> o about's. Son comunes en muchos libros, por lo tienen sintaxis
propia

```
:: NOTA ::  Lorem ipsum dolor sit amet, consectetur
   adipiscing elit, sed eiusmod tempor incidunt ut labore et
   dolore magna aliqua. Ut enim ad minim veniam.

:: ADVERTENCIA ::  Lorem ipsum dolor sit amet, consectetur
   adipiscing elit, sed eiusmod tempor incidunt ut labore et
   dolore magna aliqua. Ut enim ad minim veniam.

:: ADVERTENCIA ADVERTENCIA ADVERTENCIA ADVERTENCIA ::

   Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed eiusmod tempor
   incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.

:: ADVERTENCIA ADVERTENCIA ADVERTENCIA ADVERTENCIA

   ::

   Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed eiusmod tempor
   incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.
```

### ..comandos >

los comandos de configuracion del documento, son un subconjunto de la sistaxis
`..comando >` que cumple con la forma `..comando > argumento`, sin embargo
existen multiples formas de este comando

```
..comando-de-configuracion > argumento

..comando > argumento
  cuerpo

..comando > argumento en
  multiples lineas

  cuerpo

..comando >
  solo cuerpo
```

ademas estos pueden ser simentricos, es decir tener una marca de cierre

```
..comando > argumento
  cuerpo
< comando..

..comando > argumento en
  multiples lineas

  cuerpo
< comando..


..comando >
  solo cuerpo
< comando..
```

(resaltar que los comandos de configuracion, *no son simetricos*, es decir, no
necesita (ni debe) escribirse el cierre del comando)

Bien? Para el primer caso, el *argumento* abarca hasta la aparicion de la primer
linea en blanco o sin la indentacion apropiada.

en este tipo de comandos se encuentran los de configuracion del documento (que
ya vimos al inicio)

```
..title    > titulo del documento
  puede abarcar varias lineas, siempre con indentacion y sin lineas
  en blanco

  esto queda fuera del comando titulo, (de hecho termina el bloque de
  comandos de configuracion)

..author   > nasciiboy
..mail     > nasciiboy@gmail.com
..style    > worg/worg.css
..options  > highlight
```


para los comandos que solo tienen `cuerpo`

```
..emph >
  toda esta seccion tiene enfasis
< emph..

..emph >
  tambien esta

..bold >
  esta es bold

..center >
  y esta va centrada
< center..

..quote >
  Cuando hago esto, la gente piensa que es porque quiero alimentar mi
  ego, ¿verdad? Por supuesto, ¡no pido que se le llame “Stallmanix!”

  --Richard Matthew Stallman
```

Estan diseñados para resaltar o aplicar alguna configuracion a una parrafo o
bloque extenso del documento

En este tipo de comandos se encuentran: `center`, `bold`, `emph`, `italic`,
`quote`, `example`, `pre`, `diagram`, `art`, `cols`.

De proporcionar un `argumento`, el exportador sencillamente lo ignora.

`cols` tiene una sintaxis especial que nos premite crear columnas con cualquier
tipo de contenido, su sintaxis es

```
..cols >
  primer columna

  ::

  segunda columna

  ..src > codigo-fuente
    algo

  ::

  tercer columna

  ..img > ruta/a/imagen
```

donde el delimitador de las columnas es la secuencia `::`. Debe estar en
solitario en su propia linea, con dos espacios de indentacion.

Por su parte los comandos con *argumento* y *cuerpo*, se utilizan para bloques
de tipo <q>codigo fuente</q>

```
..src > c
  #include <stdio.h> # esto es codigo en C
< src..

..src > go

  package biskana

  import "github.com/nasciiboy/regexp3"

  // esto es codigo en go


..src > sh
  echo "hola que hace"
```

aqui el contenido despues del (y en la misma linea que) `>` especifica el
lenguaje. Por su parte, el cuerpo del bloque, es toda linea que cumpla con la
indentacion de dos espacios.

Por ultimo, los bloques con *argumentos* de multiples lineas y *cuerpo*

```
..figure > esto el titulo
  de una mini seccion

  este es el contenido de la mini seccion
```

donde el *contenido* empieza luego de la primer linea en blanco.


```
..img > ruta/a/mi/imagen.jpg

  descripcion, contenido o lo que sea

..video > ruta/a/mi/video.ogg

  descripcion, contenido o lo que sea
```

#### Tablas

```
| encabezado    | otro e  |
|===============|=========|
| elemento uno  | algo x  |
|---------------|---------|
| elemento dos  | algo    |
|               | mas...  |
|---------------|---------|
| elemento tres | otra    |
|               | cosa    |
|               | @b(mas) |
```

el encabezado se coloca a la cima, delimitado con `|===|==|`

cada fila se separa con `|----|---|`

una tabla puede tener (o no) cuerpo o encabezado, pero las filas deben tener
siempre el mismo numero de columnas, en el futuro se hara una tabla super
dopada, de momento esto es lo que hay, por cierto, las tablas tienes soporte
completo de comandos `@`


## Exportar y visualizar

Una ves instalado morg y contando con nuestro documento (de momento) podemos
hacer 2 cosas:

1. Exportar a html

2. Visualizar en la terminal (caraceristica recien añadida con muchas carencias,
   como la falta de visualizacion de tablas, imagenes, videos y otras cosas mas)


Para lo primero el comondo a utilizar es

```
morg export mi-documento.morg
```

la extencion `.morg` no es necesaria, pero es una buena costumbre marcar de que
va cada fichero

Para visualizar el documento

```
morg tui mi-documento.morg
```

esto muestra el documento, dentro de la terminal, con algunos colores. Mencionar
que una vez se ejecuta el comando, se utiliza la dimencion actual de la terminal
y esta no se refresca en caso de cambio

Para desplazarse dentro del documento

```
q                ==> quitar
Home             ==> Inicio
End              ==> Final
PgUp             ==> Desplazarse una pagina hacia arriba
PgDown           ==> Desplazarse una pagina hacia abajo
Flecha arriba    ==> Subir
Flecha abajo     ==> Bajar
Flecha derecha   ==> Desplazamiento a la derecha
Flecha izquierda ==> Desplazamiento a la izquierda
```

La libreria que se utiliza para manipular la terminal, esta muy, muy, muy verde
en ocaciones aparecen artefactos, creo que al utilizar caracteres unicode
