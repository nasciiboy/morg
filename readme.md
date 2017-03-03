# MORG

que es **morg**?

otro sistema de documentacion de marcas ligeras, basado en otros sistemas de
documentacion (de marcas ligeras), que intenta dominar el mundo, transformar a
la humanidad, terminar con el trabajo, la propiedad intelectual y forrar a su
creador

## por que?

Como toda creacion humana la informatica ha respondido a las necesidades de los
presentes deacuerdo a los conocimientos y habilidades disponibles. Acaso al
inicio alguien considero la necesidad de codificar informacion mas alla del
conjunto de caracteres USAmericano? claro que no, los humanos no somos tan
inteligentes, y avansamos a base de apaños, resolviendo unas pocas dificultades
de vez en vez, lamentablemente solo mirando en retrospectiva esto se torna
evidente, y al hacerlo quedamos obligados a cuestionar nuestra realidad.

En esta ocacion veamos en perspectiva al mas importante recurso de la humanidad
<q>la informacion</q> y los medios a nuestro alcance para crearla e interactuar
con ella.

<blockquote>
<em>seccion pendiente</em> llena de palabras con muchas letras, donde se
describe el trayecto de los sistemas para documentar informacion
</blockquote>

Con lo anterior llegamos a la siguente conclucion, <em>tener libros impresos
esta chulo, aunque poco practico e ineficiente es, si se compara con otos
formatos de documentacion</em>, por ello surgieron las paginas man, los
derivados de TeX, XML, la web y sobre ella la wikipedia. Sin embargo seguimos
forjando informacion, pensando en imprimir en papel, con la consecuencia directa de
que todos los formatos de documentacion surgidos despues del transistor son un
dolor en el culo, y es momento de forjar un sistema a la altura, apenas mas
complicado que ascii, e igual de valido que cualquier hijo vastardo de
GML->SGML->XML->HTML

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
*groff* o alguna de sus variantes, si has intentado crear una pagina man, o
incluso has sido tan intrepido como para documentar tus cosas en man, habras
decistido al poco tiempo, no hay nada mas feo he inteligible que una pagina de
manual en groff. Por ello GNU lanzo **info** que sin duda es mas util que man,
ademas TexInfo es menos feo que groff.

Entoces por que no utilizamos info para escribir la wikipedia?

- Hay que leer un manual (en ingles) de muchas paginas para utilizarlo como es
  debido

- Esta lleno de marcas y cosas misticas (pensadas para imprimir libros)

- Una ves finalizado el documento hay que "compilar" para exportar a otros
  formato mas manejables, es decir, pasar del fuente <span class="file" >.texi</span> a info, html, pdf,
  ... y si no compila te comes los mocos!

### Practico

El formato pdf se utiliza mucho, ha de ser bueno, si no por que
habria tantos libros escaneados?

Si no puedes realizar una busqueda de culquier palabra dentro del documento *no*
puede ser bueno, si la forma de acceder al fuente para modificar algun error no
esta a tu alcance *es* infame y si has de recorrerlo por paginas es *perverso*

### Modificable (por humanos)

Si aspiras a ser un *heroe del teclado* y por "curiosidad" se te ha ocurrido
mirar el codigo html de cualquer pagina web, habras llegado a la conclucion de
que el mejor lugar para guardar un mensaje que nadie ha de ver jamas, esta dentro
de una etiqueta html anidada sobre cientos de etiquetas html en una linea unica
sin ningun salto de linea.

Un formato que acepta tales aberraciones deberia ser prohibido o almenos
intervenido por un consejo de sabios, para evitar tal desgracia.

### WYSIWYMAG

What You See Is What You Mean And Get (Lo que ves es lo que quieres decir y
obtener)

La estructura del documento ha de ser minimamente agradable a la vista y
proporcionar la herramientas necesarias para utilizarlo en la creacion de
cualquier timpo de documento, desde un post a un libro o publicaciones
cientificas de cualquier indole, teniendo siempre en cosideracion que el
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
herramientas para la creacion de blogs como en el caso de wordPress

Por alguna razon desconocida los sistemas de marcado ligero son comodos, sin
embargo, si algun error se les ha de atribuir (que probablemente sea la razon de
que existan tantos) es, **no valerse por si mismos**, al mas minimo
inconveniente se recurre a trozos de codigo html o latex, resultando en
horrendos engendros <q><em>facilitadores</em></q> de estos ultimos.

El formato que creemos ha de ser tan agradable a la vista que incluso no
requiera ninguna herramienta especial para su visualizacion y creacion, los mas
intrepidos haran alarde de valerse solo con `less`, `more`, `cat` o mariposas.

Ya, esta bonito, sencillo y autosuficiente! algo mas? que tal un **Proyecto de
<s>Complementacion</s> Documentacion Humana**, por que no crear un super
repositorio (distribuido/federado) que contenga (al menos) toda obra escrita, ya
sea un blog, un libro e incluso la wikipedia, todo con el mismo formato,
separado en secciones elegibles y descargables a medida del disco duro


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

un error es, permitir que la informacion desaparesca, mas aun, dejarla en custodia
de entes que solo permiten el acceso a unos cuantos elegidos.

A un lado el copyrigth, los bytes no tienen dueños!

Sobre la implementacion... solo hay ideas, informacion mas adelante

### sintaxis
#### Estructura e indentacion

Un buen sistema de documentacion priorisa la estructura sobre el aspecto.

La estructura minima, consiste en separar el documento en secciones,
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
  muy muy muy muy muy muymuy muy muy muy
  muy muy extenso
```

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

#### y si tengo una novela

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
antecede con `@` (nota: para que aparesca `@` hay que colocar `@@`)

```
@b(1@). punto uno)
```

al expandir la etiqueta `@)` se substituye por `)`, asi:

```
<b>1). Punto uno</b>
```

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
@&{leftarrow}
```

#### math

para las formulas matematicas (inline) y ya que desconosco bastante en este
tema, podriamos no reinventar la rueda y tomar las formulas Tex

```
@m{\formula\Matematica\Tex}
```

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

ahora, todos los encabezados generan un indentificador interno apartir de su
nombre, por ejemplo

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

para hacer una referencia interna a un encabezado hariamos asi

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

todas los <q>comandos</q> tienen esta estructura `@x(custom<>contenido)`. Donde
`custom` es un parametro personalizado y opcional. Por su parte `contenido` es
el contenido del comando

Cuando un comando, por ejemplo, los enlaces requieren un `custom` y este no se ha
especificado, se genera apartir del `contenido`, extrayendo las marcas expeciales
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

un *radio target* conbierte en un enlace a cualquier palabra que encaja con la
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

```
.. comando > contenido
```

o en su forma "simetrica"

```
.. comando >
  contenido
< comando ..
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

Bien? Hay varios tipos de comandos y varias formas de optener su contenido. Por
un lado tenemo comandos donde el cuerpo se define en una sola linea o mas
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

  esto queda fuera del comando titulo

..author   > nasciiboy
..mail     > nasciiboy@gmail.com
..style    > worg/worg.css
..options  > highlight
```

Adicionalmente, los comandos de configuracion se colocan al inicio del documento
y terminan cuando aparece el primer parrafo o encabezado. Los comandos de
configuracion no deben tener espacios en blanco al inicio, ni etiqueta de
cierre, es decir `< title..` no significa nada para el comando `title`.

tambien tenemos comandos que solo tiene cuerpo

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

esta diseñada para resaltar o aplicar alguna configuracion a una parrafo o
bloque extenso del documento

luego tenemos comandos con *argumentos* y *cuerpo*, como pueden ser los bloques
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

por ultimo, teneos bloque con *argumentos* de multiples lineas y *cuerpo*

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

ya vimos que son los argumentos y el cuerpo, los `parametros` son modificadores
para el comando, como podrian ser:

- interpretar algun tipo de enfasis dentro de un
  bloque de codigo

- agregar un identificador

- establecer la orientacion visual de los elementos

- ...

Aunque los parametros pueden ser variados, deben ser pocos, uno, dos maxime tres
por comando de bloque. Remarcar que el formato no es para crear espectaculos
visuales.

los bloques propuestos

```
block o custom // como bloques personalizables
img
video
figure
quote
verse
emph
center
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

a resaltar `options`, que sirve para especificar como ha de comportarse el
traductor/visualisador del formato, de este, al igual que los parametros aun no
defino su estructura (las mas comunes son: `cosa`, `bandera:valor`,
`bandera=valor`)

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

donde un una `@` al mismo nivel de indentacion del inicio del nombre del
encabezado seguido por un espacio en blanco establece un subencabezado

poder concatenar varias definiciones

```
- A ::
  B ::
  C :: exadecimal
```

o

```
- A :: B :: C :: exadecimal
```

variables de substitucion

```
@V{variable definida en alguna parte}
```

(la siguiente, es una idea que abandone en pro de parametros en los bloque, pero igual
resulta interesante)

En algunos bloques poder establecer la orientacion del contenido, o agregar una
sintaxis especifica, como:

```
^^ elemento :: descripcion
>> elemento :: descripcion
<< elemento :: descripcion
__ elemento :: descripcion
```

La posicion del elemento se establece de la combinacion de dos caracteres, a
elegir `<`, `>`, `^` y `_`. El primer caracter indica la posicion del elemento
(`^`) superior, (`_`) inferior, (`>`) derecha o (`<`) izquierda.

La segunda combninacion establece la alineacion del contenido.
(`::`) justificado, (`<:`) a la izquierda, (`:>`) a la
derecha o (`>:<`) centrado.

```
<< elemento :: descripcion
```

genera


```
|------------------+-----------------------|
|                  | descripcion           |
|                  |                       |
|     elemento     |                       |
|                  |                       |
|                  |                       |
|------------------+-----------------------|
```

```
<^elemento >:< descripcion
```

genera

```
|------------------+-----------------------|
|    elemento      |      descripcion      |
|                  |                       |
|                  |                       |
|                  |                       |
|                  |                       |
|------------------+-----------------------|
```

```
^> elemento >:< descripcion
```

genera

```
|-----------------------------------------|
|                                         |
|                                         |
|                           elemento      |
|                                         |
|                                         |
|-----------------------------------------|
|              descripcion                |
|                                         |
|                                         |
|                                         |
|                                         |
|-----------------------------------------|
```

en org-mode, existe una forma para colocar el resultado de ejecutar un bloque de
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

otra duda surge y aunque antes se expuso como posibilidad es la forma de
concatenar los comandos `@`. Deberian poderse concateran de la forma
`@abcd(cosas)`? o solo deberia haber ordenes sencillas? si pueden concatenarse
cuales? y en que orden de ejecucion?

para terminar dos opciones para establecer un encabezado/bloque con un
idetificador

```
** @t(identificador-personalizado<>contenido del encabezado)

..figure > @t(identificador-personalizado<>argumento)
```

si el encabezado, inicia con un comando `t` especifica el identificador y texto a
utilizar

es decir, con

```
** @t(identificador<>titulo del encabezado)
```

y

```
** @t(identificador) titulo del encabezado
```

generar

```
<h2 id="identificador" >titulo del encabezado</h2>
```

aunque considero mas congruente utilizar

```
** identificador <> contenido del encabezado

..figure > identificador <> argumento
```

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

<em>por que esto es sencillo?</em> estas trabajando con el contenido original, se puede
contrastar la traduccion directamente y no requiere compilaciones ni trucos
complejos.

Para generar la traduccion, solo es necesario borrar las lineas con la marca
especial, por que la estructura del documento siempre esta precente, es decir,
siempre tenemos el producto final, solo eliminamos lo inecesario!

Si agregamos un programa que haga todo automagicamente, con una pre-traduccion,
el tabajo sera pan comido!

incluso y fantaceando, podria haber ficheros para reemplazar rss (rorg) y que el
navegador interprete directamente morg. Las fantacias no cuestan nada.

### biskana

`biskana` es lo que hay hasta el momento, una primitiva libreria escrita en
lenguaje de programacion go para exportar el formato *morg* a html, funcionan
muchos de los comandos `@` de forma limitada, aunque el sistema
`@x(identificador<>contenido)` extrae los datos perfectamente, ademas si un
comando no cuenta con llave de cierre, el resaltado abarca hasta el final del
parrafo y se cierra automaticamente

tambien el sistema para extraer el contenido de un bloque de comando funciona de
manera correcta, claro esta, solo para los casos definidos

para los bloques de codigo funte se brinda la posibilidad de resaltado de
sintaxis agregando un enlace de java-script a la configuracion.

Existe un wrapper en go de pygments con el cual podemos resaltar la sintaxis
mediente un simple css, obviamente requiere tener instalado pygments y esperar
un tiempo consideramlemente mayor para generar el resultado

aun no hay codigo para manejar las tablas de datos (proximamente)

#### Como usar

Primero tenemos que tener instalado *go* en el
sistema, [esta](https://golang.org/doc/install) es la guia de instalacion
oficial (la verdad no recuerdo como lo instale)

`biskana` es solo un componete del proyecto, de momento el unico escrito, para
su uso se envolvio en el comando `morg`
([aqui](https://github.com/nasciiboy/morg) el codigo en todo su explendor), el
cual hara muchas mas cosas, mientras tanto se limita a exportar a html los
ficheros que le indiquemos.

Asi optenemos el codigo y generamos el ejecutable

``` sh
go get -v github.com/nasciiboy/morg
```

(desconosco la funcion de `-v`)

si ya tenemos intalado morg, para actualizar

``` sh
go get -u -v github.com/nasciiboy/morg
```

para utilizar el comando

``` sh
morg fichero-uno.morg fichero-dos.morg ...
```

el resultado es un fichero en el directorio actual de tabajo con terminacion `.html`

para activar el resaltado con [highlight.js](https://highlightjs.org/)
(java-script) tienes que colocar la siguiente opcion de configuracion dentro del
fichero morg

``` morg
..options > highlight
```

coloque un estilo por defecto en el codigo fuente, para modificarlo lo mas
sencillo es hacerlo en el fichero resultado (a mano), ademas para esta opcion
debes tener la capeta `highlight` en la misma dereccion del resultado

para optener tu carpeta `highlight` ve a [esta](https://highlightjs.org/download/) pagina y especifica que lenguajes
deseas, descarga, descomprime y coloca en la misma carpeta del resultado

Si por el contrario si (tienes instalado y) quieres marcar de resaltado de
codigo utilizando `pygments` debes utilazar esta opcion de configuracion

``` morg
..options > pygments
```

luego especificar mediante una hoja de estilos css el estilo de
resaltado. Puedes opteren estilos css para
pygments [aqui](https://github.com/jwarby/jekyll-pygments-themes/)

la hoja de estilo se especifica asi

``` morg
..style > direccion/a/mi/style.css
```

no aseguro que el programa no te explote en la cara, de momento ha servido para
dos pequeñas pruebas

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

#### el codigo

sin importar el lenguaje, el codigo debe ser elegante o almenos claro y
sencillo, evintando dependencias inecesarias que dificulten su adaptacion a
otros lenguajes o peor aun el aprendisaje de otros programadores. En resumen se
buscara siempre ser una implementacion de referencia con toques didacticos en la
que cualquier indicio de aparicion de cruft sera señal de refleccion y futura
refactorizacion e incluso reescritura

`biskana` no es el programa definitivo, lo pienso como un componente *el
traductor* a otros lenguajes o dicho de otra manera el *renderizador* (que hace
realidad nuesras fantacias) de texto a texto, cuando el codigo alcance cierto
grado de madures sera dividido en otro companente `katana` este sera el parser
del formato que genere una estructura o formato de aun mas facil traduccion,
clasificando cada seccion del documento.

Momentaneamente supongo que la mejor manera en que de deben
clasificarse/generarse las secciones es como entes independientes sin estar
influenciados por el nivel de anidamento por ejemplo

```
* seccion 1
** seccion 2
*** seccion 3
```

no se traduce como

```
<div><h1>seccion 1</h1>
  <div><h2>seccion 2</h3>
    <div><h3>seccion 3</h3>
    </div>
  </div>
</div>
```

en su lugar se traduce como

```
<h1>seccion 1</h1>
<div class="h1">
</div>

<h2>seccion 2</h2>
<div class="h2">
</div>

<h3>seccion 3</h3>
<div class="h3">
</div>
```

cada seccion es "relativamente" independiente de otra, pues contiene toda la
informacion para visualirarse por si misma, unicamente su contenido tiene un
nivel de anidamiento

aunque de momento no es asi, lo mismo planeo para las listas

```
- elemento 1
- elemento 2

  + subelemento 2.1

- elemento 3
```

no se traduce como

```
<ul>
  <li>elemento 1</li>
  <li>elemento 2</li>
    <ul>
      <li>subelemento 2.1</li>
    </ul>
  <li>elemento 3</li>
</ul>
```

si no como

```
<p class="li-1" >elemento 1</p>
<p class="li-1" >elemento 2</p>
<p class="li-2" >subelemento 2.1</p>
<p class="li-1" >elemento 3</p>
```

este enfoque requiere de un monton de nuevas etiquetas, en el caso de exportar a
html, por el otro lado en un programa, esto no tiene ningun inconveniente,
incluso sera mas sencillo la visualizacion

Varios modulos, necesario sera implementar, uno, encargado de la busqueda (local
y web) del un documento, `katana` para parsear el documento, `biskana` como
exportador, `*ana` (`*`? algun buen nombre que termine en ana? se me ocurre
`nirvana` y `hana`) como visualizador/navegador del documento y `morg` (nombre
tentativo) como cohesianador de todo

son solo ideas, que tal?

por cierto `biskana` hace uso de un motor de expresiones regulares elaborado
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

#### primer ejemplo

seria grosero no monstrar ni un poco, asi que aqui estan dos ejemplos con el
`biskana` (`morg`) y el formato actual (exportacion a html)

El primero es un libro (aun en proceso de escritura) sobre el motor de
expresiones regulares Recursive Regexp
Raptor [aqui](https://github.com/nasciiboy/raptor-book) el repo, el resultado
es
[este](https://github.com/nasciiboy/raptor-book/blob/master/raptor-book.html).

Para ver el resultado en todo su esplendor, clona o baja una copia del repo y
visualiza en el navegador el html.

El segundo, es una colaboracion que estaba haciendo para traducir un manual de
emacs. Por motivos varios no he terminado y necesita una serio correccion.

<em>Por que poner algo tan vago?</em> en el se muestra el concepto de utilizar morg para
traducir manuales de forma
sencilla. [aqui](https://github.com/nasciiboy/emacs-lisp-intro-es) el repo, el
resultado (no muy bueno por no actualizar el formato)
es
[este](https://github.com/nasciiboy/emacs-lisp-intro-es/blob/master/emacs-lisp-intro_es.html) y
el fichero protagonista, el
*.porg*
[aqui](https://github.com/nasciiboy/emacs-lisp-intro-es/blob/master/emacs-lisp-intro_es.porg)

como se hace el *po*-org? primero se pasa el material a traducir a *morg*, luego
se duplica cada seccion, poniendo el duplicado justo debajo y comentando el
original con alguna marca especial, para el caso utilice `#` y se va
traducciondo cada seccion.

El proceso de exportar la traduccion es igual de sencillo, copias el *porg* y le
cambias de nombre, luego le aplicas una regexp de substitucion con el programa
de preferencia que elimine todas las lineas que inicien con la marca especiol y
voala, el documento traducido!

### zen

Extenso esto es ya, momento y pausa, el codigo no es elegante, apenas soy un
programador al inicio de su travesia y recien llegado a golang. Antes de empezar
a programar, modificar o agregar funciones debe establece una especificacion
para el formato y de ser posible un consenso global.

Mas tarde, como calentamiento hacer un exportador robusto y luego sus
componentes. De camino integrarlo en algunos cms (como hugo), Darle soporte en
nuestros editores faboritos, al tiempo de otorgarle facilidas de autocompletado y
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

Estas en una organizacion afin a este ideal, acupas un cargo de gobierno?
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
