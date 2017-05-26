[guia para utilizar morg](howto.md)

# MORG

que es **morg**?

el nombre (provisional) de otro sistema de documentacion de marcas ligeras,
basado en otros sistemas de documentacion (de marcas ligeras), que intenta
dominar el mundo, transformar a la humanidad, terminar con el trabajo, la
propiedad intelectual y forrar a su creador

## por que?

Como toda creacion humana la informatica ha respondido a las necesidades de los
presentes deacuerdo a los conocimientos y habilidades disponibles.

Acaso al inicio alguien considero la necesidad de codificar informacion mas alla
del conjunto de caracteres USAmericano (ASCII)? Claro que no, los humanos no somos tan
inteligentes, y avanzamos a base de apaños, resolviendo unas pocas dificultades
de vez en vez. Lamentablemente solo mirando en retrospectiva esto se torna
evidente, y al hacerlo quedamos obligados a cuestionar nuestra realidad.

En esta ocacion veamos en perspectiva al mas importante recurso de la humanidad
<q>la informacion</q> y los medios a nuestro alcance para crearla e interactuar
con ella.

<blockquote>
<em>seccion pendiente</em> llena de palabras con muchas letras, donde se
describe el trayecto de los sistemas para documentar informacion
</blockquote>

Con lo anterior llegamos a la siguente conclucion, <em>tener libros impresos
esta chulo, aunque poco practico e ineficiente es, si se compara con otros
formatos de documentacion</em>, por ello surgieron las paginas man, los
derivados de TeX, XML, la web y sobre ella la wikipedia. Sin embargo seguimos
forjando informacion, pensando en imprimir en papel, con la consecuencia directa de
que todos los formatos de documentacion surgidos despues del transistor son un
dolor en el culo.

Es momento de forjar un sistema a la altura, apenas mas complicado que ascii, e
igual de valido que cualquier hijo vastardo de GML->SGML->XML->HTML o alguna
variente de TeX.

## Como debe ser un sistema de documentacion ideal
### **Inmediato**

Debe estar disponible en todo momento. Las paginas man, fueron un gran acierto
de nuestros ancestros. El problema? <code class="command" >cat</code> o <code
class="command" >less</code> con una base de datos de documentos en texto
plano serian igualmente eficientes. No tendriamos colores, pero a cambio
podriamos agregar nuevas paginas y/o secciones de forma mas elegante y rapida.

### Sencillo

Si la wikipedia existe no es por algun genio del marketing vende motos, o por
un loco programador hasta arriba de flow, no, no, no, la razon es *roff*,
*groff* o alguna de sus variantes.

Si has intentado crear una pagina man, o incluso has sido tan intrepido como
para documentar tus cosas en man, habras decistido al poco tiempo, no hay nada
mas feo, he initeligible que una pagina de manual en groff. Por ello GNU lanzo
**info** que sin duda es mas util que man, ademas TexInfo es menos feo que
groff.

Entoces por que no utilizamos info (o para el caso latex) para escribir la
wikipedia?

- Hay que leer un manual (en ingles) de muchas paginas para utilizarlo como es
  debido

- Esta lleno de marcas y cosas misticas (pensadas para imprimir libros)

- Una ves finalizado el documento hay que "compilar" para exportar a otros
  formatos mas manejables, es decir, pasar del fuente <span class="file" >.texi</span> a info, html, pdf,
  ... y si no compila te comes los mocos!

### Practico

El formato pdf se utiliza mucho, ha de ser bueno, si no, por que
habria tantos libros escaneados?

Si no puedes realizar una busqueda de culquier palabra dentro del documento *no*
puede ser bueno, si la forma de acceder al fuente para modificar algun error no
esta a tu alcance *es* infame y si has de recorrerlo por paginas es *perverso*

### Modificable (por humanos)

Si aspiras a ser un *heroe del teclado* y por "curiosidad" se te ha ocurrido
mirar el codigo html de cualquer pagina web, habras llegado a la conclucion de
que el mejor lugar para guardar un mensaje, que nadie ha de ver jamas, esta dentro
de una etiqueta html, anidada sobre cientos de etiquetas html, en una linea unica
sin ningun salto de linea.

Un formato que acepta tales aberraciones deberia ser prohibido o almenos
intervenido por un consejo de sabios, para evitar tal desgracia.

### WYSIWYMAG

What You See Is What You Mean And Get (Lo que ves es lo que quieres decir y
obtener)

La estructura del documento ha de ser minimamente agradable a la vista y
proporcionar la herramientas necesarias para utilizarlo en la creacion de
cualquier tipo de documento, desde un post a un libro o publicaciones
cientificas de cualquier indole, teniendo siempre en consideracion que el
proposito y fin ultimo es la documentacion, no convertirse en la base para crear
interfaces visuales.

Animaciones neon, anuncios publicitarios, botones "sociales", typografias con
sombras, colores que afectan la vista (y el buen gusto), no son el objetivo del
formato, de eso ya seguiran encargandose los formatos existentes

## Propuesta

html es tan feo que la wikipedia utiliza mediawiki (uno de los tantos lenguajes
de marcas ligeros que existen). Por su parte, sitios como github directamente
pasan de html, fomentando el uso de markdown, org, ReStructured Text, texto
plano, etc. Algo similar ocurre con plataformas de gestion de contenido y
herramientas para la creacion de blogs como en el caso de WordPress

Por alguna razon desconocida los sistemas de marcado ligero son comodos, sin
embargo, el mayor error y provable razon de que existan tantos lenguajes de
marcas ligeras es **no valerse por si mismos**, al mas minimo inconveniente se
recurre a trozos de codigo html o latex, resultando en horrendos engendros
<q><em>facilitadores</em></q> de estos.

El formato que creemos ha de ser tan agradable a la vista que incluso no
requiera ninguna herramienta especial para su visualizacion y creacion, los mas
intrepidos haran alarde de valerse solo con `less`, `more`, `cat` o mariposas.

Ya esta tio, es bonito, sencillo y autosuficiente! algo mas? que tal un **Proyecto de
<s>Complementacion</s> Documentacion Humana**, por que no crear un super
repositorio (distribuido/federado) que contenga (al menos) toda obra escrita, ya
sea un blog, un libro e incluso la wikipedia, todo con el mismo formato,
separado en secciones elegibles y descargables a medida del disco duro, algo asi


```
origen
   .
   ├── recursos
   │   ├── imgenes
   │   ├── videos
   │   ├── programas
   │   ├── codigo
   │   └── musica
   ├── man
   │   ├── 1
   │   ├── 2
   │   ...
   │   └── N
   ├── blog
   │   ├── gnusr
   │   ├── nasciiboy
   │   ... cualquiera
   │   └── emacsChan
   ├── books
   │   ├── scifi
   │   ├── math
   │   ...
   │   └── ajedrez

   ...

   └── wikipedia
```

que, que ventajas traera esto?

- librarse de DRMs, mierda script, cookies, sistemas de rastreo, publicidad

- crear una <q>biblioteca</q> permanente, no mas errores 404

- olvidar el exesivo tiempo de carga y recursos que consume un navegador, pues
  seria substituido por tuis o guis enfocadas unica y exclusivamente para
  interactuar con el formato... aunque claro, para quienes no puedan o quieran
  despegarse de navegador, seguiran teniendo una version web


es un error, permitir que la informacion desaparesca, mas aun, dejarla en custodia
de entes que solo permiten el acceso a unos cuantos elegidos.

Sobre la implementacion... solo hay ideas.

### sintaxis
#### Estructura e indentacion

Un buen sistema de documentacion priorisa la estructura sobre el aspecto.

La estructura minima, consiste en separar el documento en secciones o
encabezados gerarquizados

*una marca un nivel*: un encabezado inicia con el signo `*` seguido por (un)
espacio(s) y el nombre de la seccion.

El numero de `*` indica el nivel del encabezado, su equivalente  en html seria

- `*` == `h1`
- `**` == `h2`
- `***` == `h3`
- `****` == `h4`
- `*****` == `h5`


```
* nivel Uno

  Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
  eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut
  enim ad minim veniam.

** nivel dos

   Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
   eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut
   enim ad minim veniam.

   Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
   eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut
   enim ad minim veniam.

*** nivel tres

    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
    eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut
    enim ad minim veniam.
```

El contenido de cada encabezado inicia tras dejar una linea de espacio en blanco
y ha de indentarse (opcional y a gusto) con un numero de espacios igual al
numero de `*`, mas un espacio.

Para mantener una estetica agradable los titulares extensos
pueden colocarse de la forma

```
* encabezado muy muy muy muy muy muy muy
  muy muy muy muy muy muy muy muy muy muy
  muy muy extenso
```

A deferencia de otros formatos de marcas ligeras, todos los encabezados generan
un identificador interno, al cual se puede hacer referencia dentro del
documento. Tal identificador, se forja apartir del texto, substituyento los
espacios en blanco por guiones (`-`)

En caso de que se desee tener un identificador distinto al texto, o si existen
dos o mas encabezados con el mismo nombre, se puede utilizar

```
** identificador <> contenido del encabezado
```

donde `<>` actua como separador entre el identificador personalizado y el texto
que aparece como nombre del encabezado

#### Listas

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
elemnentos del formato


```
* nivel uno

  1. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
     eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut
     enim ad minim veniam.

     a) Lorem ipsum dolor sit amet.

        - Lorem ipsum dolor sit amet, consectetur adipiscing elit.

  2. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
     eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut
     enim ad minim veniam.
```

#### definiciones

```
- elemento :: definicion Lorem ipsum dolor sit amet, consectetur
  adipiscing elit, sed eiusmod tempor incidunt ut labore et
  dolore magna aliqua. Ut enim ad minim veniam.

+ elemento :: definicion Lorem ipsum dolor sit amet, consectetur
  adipiscing elit, sed eiusmod tempor incidunt ut labore et
  dolore magna aliqua. Ut enim ad minim veniam.
```

#### y si tengo una novela
##### Dialogos

```
> "Dialogo, Lorem ipsum dolor sit amet, consectetur adipiscing
  elit, sed eiusmod tempor incidunt ut labore et dolore magna
  aliqua. Ut enim ad minim veniam."
```

los dialogos tienen la mismas normas que una lista.

mencionar que anque existen varias formas de solvertar el asunto para indicar
a que personaje pertenece el dialogo, podria optarse por la siguiente sintaxis

```
> personaje A :: "Dialogo, Lorem ipsum dolor sit amet",

> personaje B :: aliqua. Ut enim ad minim veniam."
```

##### "separadores"

```
  Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed eiusmod tempor
  incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.

  > "Dialogo, Lorem ipsum dolor sit amet, consectetur adipiscing
    elit, sed eiusmod tempor incidunt ut labore et dolore magna
    aliqua. Ut enim ad minim veniam."

  ....

  Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed eiusmod tempor
  incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.
```

un separador representan esos cambios <em>abruptos</em> de escena argumental,
que en papel son representados con una pagina en blanco o varios espacios
vacios, sin cambiar de capitulo

para indicar el separador se colocan cuatro puntos en una linea, sin ningun
caracter distinto a ecepcion de espacios en blanco, este separador debe tener
almenos una linea en blanco sobre y debajo del mismo

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

#### resaltado

nadie quiere tener etiquetas a lo html

```
<etiqueta>
  <etiqueta>
    contenido
  </fin_etiqueta>
</fin_etiqueta>
```

los lenguajes de marcas ligeras lo manejan de forma un poco mas agradable

```
(org)       *bold*
(markdown)  **bold**
(mediawiki) '''bold'''

algun otro  <^bold^>
```

no obstante con esta aproximacion pronto se crean ambiguedades, ademas de
limitarse a 3 o 4 formas de etiquetar el contenido antes de recurrir a marcas
exoticas o recaer en etiquetas html.

A espera de una mejor alternativa, podria recurrirse al estilo de marcas de
texinfo… con un leve retoque al formato.

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
marcas que los caracteres ascii algo estaremos haciendo mal.

Algunas propuestas:

```
- A :: acronym      - a :: abbr       - 0 ::            - : :: def
- B ::              - b :: bold       - 1 ::            - ; ::
- C :: smallCaps    - c :: code       - 2 ::            - = ::
- D ::              - d :: data       - 3 ::            - ? ::
- E :: error        - e :: emph       - 4 ::            - @ :: escape
- F :: func         - f :: file       - 5 ::            - ` ::
- G ::              - g ::            - 6 ::            - ' :: ‘samp’
- H ::              - h ::            - 7 ::            - " :: “quote”
- I ::              - i :: italic     - 8 ::            - # :: path
- J ::              - j ::            - 9 ::
- K :: keyword      - k :: kbd        - ^ :: sup
- L :: label        - l :: link       - _ :: sub
- M ::              - m :: math       - \ ::
- N :: defnote      - n :: note       - | ::
- O :: option       - o ::            - * ::
- P ::              - p ::            - + ::
- Q ::              - q :: quote      - , ::
- R :: result       - r :: ref        - - :: —exp—
- S ::              - s :: strike     - . ::
- T :: radiotarget  - t :: target     - / ::
- U ::              - u :: underline  - % :: (parentesis)
- V :: var          - v :: verbatim   - & :: symbol
- W :: warnig       - w ::            - $ :: command
- X ::              - x ::            - ~ ::
- Y ::              - y ::            - ! ::
- Z ::              - z ::
```

que cada caracter solo tenga un significado permite concatenar acciones como en

```
@uisb(underlineItalicStrikeBold)
```

su equivalente html seria

```
<u><i><strike><b>underlineItalicStrikeBold</b></strike></i></u>
```

<em>y los `({[<>]})`?</em> Mas opciones, mas diversion.

Segun sea el contexto `{}` o `()` podrian requerir el *escape* de algun
caracter. Para minimizar la inclucion de signos extraños, los delimitadores se
aplican deacuerdo a la necesidad y gusto del <q>creador</q>.

cuando no haya *escapatoria*, para anular el significado de un caracter se
antecede con `@`, por ejemplo:

```
@b(1@). punto uno)
```

al expandir la etiqueta `@)` se substituye por `)`, asi:


```
<b>1). Punto uno</b>
```

(nota: para que aparesca `@` hay que colocar `@@`)

##### comentarios

```
@ linea comentada
```

una `@` al inicio de linea con almenos uno o mas espacios en blanco comenta la
linea en cuestion

#### mas alla del ASCII

preferiblemente se utilizara un sistema de codificacion <q>moderno</q> como
UTF-8.

Opcionalmente (y para no vernos en la necesidad de buscar un caracter
complicado) se puede echar mano del <q>comando</q> `&`, por ejemplo

```
@&{nombreGenericoDeCaracterComplicado}
```

a modo de ilustrar esto si utilizamos `@&{leftarrow}`, en el texto a exportar
sera reemplazado por `⇐`

#### math

para las formulas matematicas (inline) y ya que desconosco bastante en este
tema, podriamos no reinventar la rueda y tomar las formulas Tex

```
@m{\formula\Matematica\Tex}
```

ultimamente he investigado un poco sobre el tema, quiza sea mejor opcion tomar
la sintaxis que utiliza Libre Office Math o un recien llegado como AsciiMath

#### enlaces

```
@l{ruta}
```

equivalente a

```
<a href="ruta">ruta</a>
```

y

```
@l{ruta<>descripcion}
```

equivalente a

```
<a href="ruta">descripcion</a>
```

Como antes se menciono, todos los encabezados generan un indentificador interno
apartir de su nombre, por ejemplo

```
* encabezado
```

se traduce en html como

```
<h1 id="encabezado" >encabezado</h1>
```

y por ejemplo

```
* @b(encabezado) con @e(enfasis)
```

se traduce en html como

```
<h1 id="encabezado-con-enfasis" ><b>encabezado</b> con <em>enfasis</em></h1>
```

y un enlace dentro de un encabezado con enfasis

```
* @l(http://fsf.org/<>@b(link-encabezado)) con @e(enfasis)
```

se traduce en html como

```
<h1 id="link-encabezado-con-enfasis" ><a href="http://fsf.org/" ><b>link-encabezado</b></a> con <em>enfasis</em></h1>
```

para hacer una referencia interna a un encabezado, hariamos asi

```
@l(#encabezado)
```

que se traduce en

```
<a href="#encabezado" >encabezado</a>
```

o

```
@l(#encabezado<>lo que sea)
```

que se traduce en

```
<a href="#encabezado" >lo que sea</a>
```

(nota: en el futuro se planea presindir del signo `#` haciendo primero una
busqueda en todas las referencias internas del documento. En caso de no encotrar
coincidencias dejar la referencia tal cual)

##### como funciona esta magia?

todas los <q>comandos</q> `@` tienen la estructura `@x(custom<>contenido)`. Donde
`custom` es un parametro personalizado y opcional. Por su parte `contenido` es
el contenido del comando. Finalmente `<>` actua a modo de separador entre ambos

Cuando un comando, por ejemplo, un enlace requiere un `custom` y este no se ha
especificado, se genera apartir del `contenido`, extrayendo las marcas especiales
dejando unicamente el texto.

Cuando el comando no requiere de un `custom` y este es proporcionado, el comando
o lo ignora o se utiliza como identificador o etiqueta segun sea el caso (nota:
esto esta a consideracion y podria requerir sintaxis adicional)

Cuando el comando esta dentro de otro comando, el comando interno *pasa* su
contenido al comando externo, por ejemplo:

```
@l(#encabezado<>lo que sea con @e(enfasis con @b(algo<>bold)))
```

genera

```
<a href="#encabezado" >lo que sea con <em>enfasis con <b>bold</b></em></a>
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

tambien se pueden crear enlaces internos mediante el comando `t`

```
@t{target}
@t{target<>descripcion}
```

y radio objetivos con

```
@T{radioTarget}
@T{radioTarget,descripcion}
```

un *radio target* convierte en un enlace, a cualquier palabra que encaje con la
descripcion del radio opjetivo (ignorando entre mayusculas y minusculas), el
objetivo de todos los enlaces, es la declaracion misma del objetivo. Los bloques
de codigo quedan exentos de este comportamiento

#### notas

```
@n{enlace-a-nota}
@n{enlace-a-nota<>descripcion}
@n{nota en linea<>descripcion}
@N{objetivo-descripcion}
```

(nota: a esto aun le falta mas trabajo)

#### comandos de bloque

Los comandos `@` son para el contenido <q>en linea</q>, es decir, estan
diseñados para ser incluidos dentro de parrafos de texto.

Por su parte los <q>comandos de bloque<q> se utilizan para afectar secciones
enteras, indicar acciones complejas o para contenido especializado, como pueden
ser complejas equaciones matematicas o porciones de codigo fuente, sin
preocuparse por <q>el significado de la sintaxis regular</q>

La estructura basica de un comando de bloque es:


```
.. comando > contenido
```

o en su forma "simetrica"

```
..comando >
  contenido
< comando..
```

el contenido tiene que estar indentado con dos espacios por ejemplo

```
..comando >
  contenido

  mas contenido
```

o

```
..comando >
  contenido

  mas contenido

  ..otro-comando >
    contenido

    mas contenido
```

Bien? Pues hay varios tipos de comandos y varias formas de optener su
contenido.

Por un lado tenemos comandos donde el cuerpo se define en una sola linea o mas
indentadas.

```
..comando > contenido contenido contenido
  contenido
```

El contenido o *argumentos* abarcan hasta la aparicion de la primer linea en
blanco o sin la indentacion apropiada.

en este tipo de comandos se encuentran los de configuracion del documento

```
..title    > titulo del documento
  puede abarcar varias lineas, siempre con indentacion y sin lineas
  en blanco

  esto queda fuera del comando titulo y de hecho "rompe" la seccion de
  comandos de configuracion

..author   > nasciiboy
..mail     > nasciiboy@gmail.com
..style    > worg/worg.css
..options  > highlight
```

Adicionalmente, los comandos de configuracion se colocan al inicio del documento
y terminan en cuanto aparece el primer elemento que no sea un comando de
configuracion. Los comandos de configuracion no deben tener espacios en blanco
al inicio de la linea, ni etiqueta de cierre, es decir `< title..` no significa
nada para el comando `title`.

Tambien tenemos comandos que solo tiene cuerpo

```
..emph >
  toda esta seccion tiene enfasis
< emph..

..emph >
  tambien esta

..bold >
  esta es bold

..center >
  y esta va centrado
< center..

..quote >
  Cuando hago esto, la gente piensa que es porque quiero alimentar mi
  ego, ¿verdad? Por supuesto, ¡no pido que se le llame “Stallmanix!”

  --Richard Matthew Stallman
```

estan diseñados para resaltar o aplicar alguna configuracion a una parrafo o
bloque extenso del documento

Luego tenemos comandos con *argumentos* y *cuerpo*, como pueden ser los bloques
de codigo

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
lenguaje, por su parte, el cuerpo del bloque es toda linea que cumpla con la
indentacion

por ultimo, estan los bloques con *argumentos* de multiples lineas y *cuerpo*

```
..figure > esto el titulo
  de una mini seccion

  este es el contenido de la mini seccion
```

donde el *contenido* empieza luego de la primer linea en blanco.

Como vez, todo depende del comando que se utilize, en terminos generales
un comando tiene esta estructura

```
..comando parametros-especiales > argumentos

  //
  cuerpo
  //
```

tambien hay lugar para imagenes y video

```
..img parametros-especiales > direccion/a/mi/imagen.jpg

  descripcion, contenido o lo que sea

..video parametros-especiales > direccion/a/mi/video.mkv

  descripcion, contenido o lo que sea
```

Ya vimos que son los argumentos y el cuerpo, los `parametros` son modificadores
para el comando, como podrian ser:

- interpretar algun tipo de enfasis dentro de un bloque de codigo

- agregar un identificador

- establecer la orientacion visual de los elementos

- ...

Aunque los parametros pueden ser variados, deben ser pocos, uno, dos, maxime tres
por comando de bloque. Remarcar que el formato no es para crear espectaculos
visuales.

existen varias corrientes sobre la forma en que se pueden presentar los
<em>parametros</em>, por ejemplo: `cosa`, `bandera:valor`, `bandera=valor` y
`-bandera valor`. Al respecto, propongo que los parametros tengan la sintaxis de
una funcion, como la de la mayoria de lenguajes de programacion, es decir
`bandera( valor )`, donde bandera, `especifica` el parametro que deseamos
modificar/activar, y delimitado por parentecis el `valor` o valores a
establecer, este podria ser una serie de valores separados por coma, donde en
caso de precindir de ellos y aplicar el valor por defecto simpremente se dejaria
los parentecis en blanco.

Bueno, esto tiene un inconveniente, puesto que un comando de bloque, no deja de
ser una <q>funcion</q> y ahora tendriamos mas funciones, sumado al echo de que
para el poco instruido pudiese ser complicado... por otro lado no son funciones
tal cual, son solo una sintaxis distinta para los parametros, tambien evitan
cualquier anbiguedad respecto a la pertenencia de los valores y el alcance de
estos, sumado al echo de poder realizar un propocesamiento de los valores como
en un lenguaje de programacion, en fin, algo asi seria el aspecto para por
ejemplo indicar que deseamos numerar las lineas en un bloque de codigo fuente en
lenguaje C a partir con un indice que inicia en 42:

```
..src n(42) > c
  #include <stdio.h> # esto es codigo en C

  int main(){

    printf("hola, que tal\n");
    return 0;
  }
```

algunos bloques propuestos

```
block o custom // como bloques personalizables
img
video
figure
quote
verse
emph
center
tab
bold
italic
src
example/pre
cols
math
diagram
```

para la configuracion del documento

```
title
subtitle
author
translator
mail
licence
style
date
tags
description
id
options
lang/language
```

#### Tablas

Sin duda un tema complejo, podria tenerse una tabla totalmente funcional con
formulas y demas, pero para inciar:

```
| encabezado    | otro e  |
|===============|=========|
| elemento uno  | algo x  |
|---------------|---------|
| elemento dos  | algo a  |
|               |---------|
|               | algo b  |
|---------------|---------|
| d o s  c e l d a s      |
```

el encabezado se coloca a la cima, delimitado con `|===|==|`

cada elemento se separa con `|----|---|`

unir celdas es complicado podria tomarse en consideracion el numero exacto de
caracteres para obtener esta informacion, o colocar un signo <q>invisible</q> de
alineacion dentro la la tabla como `^`

#### a considerar

En algunos documentos se agrega un subtitulo en lugar de crear una subseccion
para esto pobria ofreserse algo como:

```
* encabezado
  @ subencabezado
```

donde una `@` al mismo nivel de indentacion del inicio del nombre del
encabezado seguido por un espacio en blanco establece un subencabezado

Otro asunto es poder concatenar varias definiciones

```
- A ::
  B ::
  C :: exadecimal
```

o de forma

```
- A :: B :: C :: exadecimal
```

Tambien proporcionar variables de substitucion

```
@V{variable definida en alguna parte}
```

En org-mode, existe una forma para colocar el resultado de ejecutar un bloque de
codigo fuente. Es asi:

```
#+BEGIN_SRC elisp
  (message "hola lisp")
#+END_SRC

#+RESULTS:
: hola lisp
```

en morg podria ser

```
..src > elisp
  (message "hola lisp")

>> hola lisp
```

o

```
..src > c
  #include <stdio.h>

  int main(){
    printf( "Hola mundo\n" );
    return 0;
  }

..result >
  "Hola mundo"
< src..
```

Otra duda surge y aunque antes se expuso como posibilidad, deberian poder
concatenarse los comandos `@` de la forma `@abcd(cosas)`? o solo existir ordenes
sencillas? si pueden concatenarse cuales? y en que orden de ejecucion?

#### porg

los ficheros `po` producidos con `gettext` se utilizan para traducir documentos
de un lenguaje a otro, si gettext no muere en el intento... con morg podemos
hacer algo mucho mas sencillo para traducir documentos

imaginemos que tenemos este texto

```
* nivel uno

  1. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
     eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut
     enim ad minim veniam.

     a) Lorem ipsum dolor sit amet.

        - Lorem ipsum dolor sit amet, consectetur adipiscing elit.

  2. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
     eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut
     enim ad minim veniam.
```

a si se ve como documento `porg`

```
#* nivel uno
* nivel uno

#   1. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
#      eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut
#      enim ad minim veniam.
  1. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
     eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut
     enim ad minim veniam.

#      a) Lorem ipsum dolor sit amet.
     a) Lorem ipsum dolor sit amet.

#         - Lorem ipsum dolor sit amet, consectetur adipiscing elit.
        - Lorem ipsum dolor sit amet, consectetur adipiscing elit.

#   2. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
#      eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut
#      enim ad minim veniam.
  2. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
     eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut
     enim ad minim veniam.
```

se toma el contenido fuente, se duplica cada seccion, se coloca justo debajo
del original y se marca con algun signo especial el contenido original.

<em>por que esto es sencillo?</em> estas trabajando con el contenido original, lo
cual permite contrastar la traduccion directamente y no requiere compilaciones
ni trucos complejos.

Para generar la traduccion, solo es necesario borrar las lineas con la marca
especial, por que la estructura del documento siempre esta precente, es decir,
siempre tenemos el producto final, solo eliminamos lo inecesario!

Si agregamos un programa que haga todo automagicamente, con una pre-traduccion,
el tabajo sera pan comido!

incluso y fantaceando, podria haber ficheros para reemplazar rss (rorg) y que el
navegador interprete directamente morg. Las fantacias no cuestan nada.

### Un poco de accion

De momento, a modo de <q>prueba de concepto</q> existe **morg**, un programa
(mal) desarrollado en golang, con el cual se puede exportar (con limitaciones)
algunos conceptos basicos del lenguaje al formato **html** asi como poder
visualizar el documento dentro de un terminal. Para mas informacion
vea [Como iniciar con morg](howto.md).

De forma simplificada, para optener el programa

``` sh
go get -v github.com/nasciiboy/morg
```

(desconosco la funcion de `-v`)

si ya tenemos instalado morg, para actualizar

``` sh
go get -u -v github.com/nasciiboy/morg
```

Para exportar un documento a html

``` sh
morg toHtml mi-fichero.morg
```

Para visualizar un documento

``` sh
morg tui mi-fichero.morg
```

funcionan varios de los comandos `@`, aunque el sistema
`@x(identificador<>contenido)` extrae los datos perfectamente, ademas si un
comando no cuenta con llave de cierre, el resaltado abarca hasta el final del
parrafo y se cierra automaticamente

El sistema para extraer el contenido de un comando de bloque tambien funciona de
manera correcta, claro esta, solo para los casos definidos

#### el codigo

sin importar el lenguaje, el codigo debe ser elegante o almenos claro y
sencillo, evintando dependencias inecesarias que dificulten su adaptacion a
otros lenguajes o peor aun, el aprendisaje de otros programadores. En resumen se
buscara siempre ser una implementacion de referencia con toques didacticos en la
que cualquier indicio de aparicion de cruft sera señal de refleccion y futura
refactorizacion e incluso reescritura

bueno, bueno, hasta ahora el codigo disponible <q>evoluciona</q> a base de
prueba y errar, tengo poquisima experiencia/formacion, encima soy un recien
llegado a go, en fin

conceptualmente se planea dividir el programa en varias secciones, `katana` el
encargado de parcear el documento, para entregar una estructura de datos
sencilla

`biskana` *el traductor* a otros lenguajes o dicho de otra manera el
*renderizador* (que hace realidad nuestras fantacias) de texto a texto

`nirvana` (cofff, en un principio iba a ser `hana` (flor) pero lo olvide!)
encargado de desplegar el documento de forma visual como TUI o como GUI

`*ana` (aun no implementado) la base de datos donde se consulta si tenemos el
documento de forma local o debe realizarce una peticion externa. Ademas debe
regresar el documento en el formato original

El primer componente que surgio fue `biskana`, me gusto la broma(?), de hay, el
deseo de terminar los nombres en `ana`. Alguna propuesta interesante?

Como cohesionador de todo, el propio `morg` (aun no me convence el nombre)

Por cierto `katana` hace uso de un motor de expresiones regulares elaborado
desde cero llamado `regexp3` que por mera casualidad programe, este motor esta
en desarrollo y en algun momento sera substituido por `regexp4`. Ambos utilizan
la misma sintaxis y aunque esta limitado solo a exrpesiones regulares, capturas
y pocas cosas mas. Lo utilizo por estas razones

1. funciona!

2. no ha aparecido alguna exprecion que revase su limitada capacidad

3. es sencillo y facil de modificar (por mi, almenos), ademas cuando surge un
   error se donde buscar

4. puede port-arse con relativa facilidad a cualquier lenguaje (creo), el
   desarrollo original fue hecho en C. Cuando digo en C me refiero a solo C, sin
   recurrir a ninguna libreria, ni siquiera la libreria estandar. Su port a go
   no fue demasiado traumatico, encima se vio veneficiado por la orientacion a
   objetos


#### porque Go

por nada en especial... cuando empece a escribir el codigo hace cosa de un año
en C mis habilidades no daban para mucho, no es que ahora sea un jodido guru,
pero algo he aprendido, he? pero por que C? por velocidad y eficiencia, si vas a
hacer un proyecto tan ambicioso, mal seria que fuese lagueado y demorara aun
unos pocos segundos en mostrar el contenido.

Aun si el proyecto se escribiese en un lenguaje interpretado en algun momento
deberia portarse a C o similar.

podria haber elegido C++... pero me cruce con Go, no es un lenguage que proboque
facinacion en mi, pero solo es 3 veces mas lento que C y tiene ideas muy
interesantes. Por añadir creo que su tipado fuerte es un fastidio, al igual que
el estilo de codificacion, llaves forsosas para instrucciones simples y el no
contar con punteros de verdad (al menos con un nivel de indireccion) obliga a
hacer apaños impresionantes. Desde luego, como recien llegado a go y tras una
ligera prueba del lenguaje este comentario puede ser una cagada, disculpad mi
ignorancia.

Por cierto, pygments, en un tremendo consumidor de tiempo, desconosco si es por
estar escrito en python, o por la naturaleza compleja de su labor. Habra que
reescribir una version mas veloz.

#### primer ejemplo

seria grosero no monstrar ni un poco, asi que aqui estan dos ejemplos con
`morg` y el formato actual (exportacion a html)

El primero es un libro (aun en proceso de escritura) sobre el motor de
expresiones regulares Recursive Regexp
Raptor [aqui](https://github.com/nasciiboy/raptor-book) el repo, el resultado
es
[este](https://github.com/nasciiboy/raptor-book/blob/master/raptor-book.html).

Para ver el resultado en todo su esplendor, clona o baja una copia del repo y
visualiza en el navegador el html.

El segundo, es una colaboracion que estaba haciendo para traducir un manual de
emacs. Por motivos varios no he terminado y necesita una seria correccion.

<em>Por que poner algo tan vago?</em> en el se muestra el concepto de utilizar morg para
traducir manuales de forma
sencilla. [aqui](https://github.com/nasciiboy/emacs-lisp-intro-es) el repo, el
resultado (no muy bueno por no actualizar el formato)
es
[este](https://github.com/nasciiboy/emacs-lisp-intro-es/blob/master/emacs-lisp-intro_es.html) y
el fichero protagonista, el
*.porg*
[aqui](https://github.com/nasciiboy/emacs-lisp-intro-es/blob/master/emacs-lisp-intro_es.porg)


### zen

Extenso, esto es ya, momento y pausa, el codigo no es elegante, apenas soy un
programador al inicio de su travesia y recien llegado a golang. Antes de empezar
a programar, modificar o agregar funciones debe establece una especificacion
para el formato y de ser posible un consenso global.

Mas tarde, como calentamiento hacer un exportador robusto y luego sus
componentes. De camino integrarlo en algunos cms (como hugo), Darle soporte en
nuestros editores favoritos, al tiempo de otorgarle facilidas de autocompletado y
resaltado.

Luego establecer una estructura capaz de mantenerse por cuenta propia y/o por
la coperacion desinteresado de empresas y gobiernos, absorver todo el contenido
que sea posible y dominar la galaxia

## como puedo ayudar

Pagameeeee un salario, enserio! (que medio te es ma sencillo?, contactame por
correo)

Contrata a un grupo de programadores motivados, que compartan este sueño, y
pagame un salario!

Programa, traduce, discute y comenta

Estas en una organizacion afin a este ideal, ocupas un cargo de gobierno?
comenta y difunde

## TODO

- definir el formato

- crear exportador robusto. De inicio a html, luego a otros formatos

- programar el resto de componentes

- estructura operativa, para poner en marcha el/los repositorios que albergaran
  blogs, libros, wikipedia y lo que se deje, para formar un repositorio global
  de conocimiento libre, permanente e imparable

- difusion, expansion y dominacion


Quiero terminar, aclarando lo siguiente: difundo estas ideas y codigo bajo la
licencia GNU AGPL v3, si tienes las habilidades para hacer lo planteado
programando todo por cuenta propia, te pido que lo hagas (aunque no estas
obligado) bajo esta licencia.

Incluso sin financiacion, programare lo que pueda, cuando pueda, simplemente por
ser un reto emocionante, empesando por incorporar el formato a hugo (en unos
meses o mas)


Happy Hacking!
