[guia para utilizar morg](howto.md)

# MORG

que es **morg**?

el nombre de otro sistema de documentacion de marcas ligeras, basado en otros
sistemas de documentacion (de marcas ligeras)

## por que?

tener libros impresos esta chulo, aunque poco practico e ineficiente es, si se
compara con los formatos de documentacion electronicos, por ello surgieron las
paginas man, los derivados de TeX, XML, la web y sobre ella la wikipedia. Sin
embargo seguimos forjando informacion, pensando en imprimir en papel, con la
consecuencia directa de que todos los formatos de documentacion surgidos despues
del transistor son o demasiado complejos (XML, TeX, Texinfo, LaTeX) o tienen
demasiadas carencias (markdow, ReStructuredText, mediaWiki, org-mode)

Tomando en cuenta que como especie nuestro presente y futuro depende de la
cantidad de conocimiento que poseemos, compartimos y utilizamos, es preocupante
que aun no exista un formato apropiado para crear, conservar, traducir y
distribuir nuestra literatura, por ello **es momento de forjar un sistema a la
altura**, apenas mas complicado que ascii, e igual de valido que cualquier
derivado de GML->SGML->XML->HTML o alguna variente de TeX

morg es una pequeña propuesta (aun en fase muy experimental (== inestable)) de como
podria ser ese supuesto nuevo formato de documentacion universal, apto tanto
para un simple blog personal, como poesia, novelas, o el mas complejo articulo
cientifico

pero crear este nuevo sistema de documentacion no es el fin de la mision, solo
establece el inicio de la ardua labor de preservar, expandir y difundir a todo
habitante, estudiante y academico, toda la informacion habida y por haber, de
forma libre, permanente y sin restricciones. Debemos crear una gran biblioteca
con un formato en texto plano, explorable incluso con un sencillo editor de
texto, o en sus formas mas elaboradas mediante una interfaz web o de linea
e comandos, pero sin caer en los vicios y defectos del actual html, es decir,
sin publicidad, drm, que este plenamente estandarizado y libre de censura

(debajo, un arbol conceptual de la estructura de la libreria)

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

es un error, permitir que la informacion desaparesca, mas aun, dejar su custodia
a entidades mesquinas que solo permiten el acceso a unos cuantos elegidos

## Como debe ser un sistema de documentacion ideal
### Inmediato

Debe estar disponible en todo momento. Las paginas man, fueron un gran acierto
de nuestros ancestros. El problema? <code class="command" >cat</code> o <code
class="command" >less</code> con una base de datos de documentos en texto
plano serian igualmente eficientes. No habria colores, pero a cambio
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

Entoces por que no utilizamos info (o su hermano latex) para escribir la
wikipedia?

- Hay que leer un manual (en ingles) de muchas paginas para utilizarlo como es
  debido

- Esta lleno de marcas y cosas misticas (pensadas para imprimir libros)

- Una ves finalizado el documento hay que "compilar" para exportar a otros
  formatos mas manejables, es decir, pasar del fuente <span class="file" >.texi</span> a info, html, pdf,
  ... y si no compila te comes los mocos!

pude que usted este influienciado por falsos mitos que rodean a los derivados de tex
(info|latex) (ver [A](http://www.danielallington.net/2016/09/the-latex-fetish/)
o [B](http://karl-voit.at/2017/08/26/latex-fetish/)), aun asi, nada impide hacer
realidad las nobles metas que se proponen estos sistemas de composicion de textos

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
que lejos de hacerlos independientes, los vuelven <q><em>facilitadores</em></q>
de estos ultimos

El formato que creemos ha de ser tan agradable a la vista que incluso no
requiera ninguna herramienta especial para su visualizacion y creacion, los mas
intrepidos haran alarde de valerse solo con `less`, `more`, `cat` o mariposas.

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

A diferencia de otros formatos de marcas ligeras, todos los encabezados generan
un identificador interno, al cual se puede hacer referencia dentro del
documento. Tal identificador, se forja apartir del texto, substituyento los
espacios en blanco por guiones (`-`)

En caso de que se desee tener un identificador distinto al texto, o si existen
dos o mas encabezados con el mismo nombre, se puede utilizar

```
** identificador <> contenido del encabezado
```

donde `<>` actua como separador entre el identificador personalizado y el texto
que aparece como nombre del encabezado, por ejemplo

```
* Seccion Uno

  Sed eiusmod tempor incidunt ut labore et dolore magna aliqua.

** Ejemplos seccion uno <> Ejemplos

   ...

* Seccion Dos

  Lorem ipsum dolor sit amet, consectetur adipiscing elit.

** Ejemplos seccion dos <> Ejemplos

   ...
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
* Seccion uno

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

se sigue la sintaxis de una lista y sin dejar lineas en blanco deben aparecen
dos puntos seguidos (`::`) con al menos un espacio en blanco del lado izquierdo,
la definicion puede aparecer a continuacion, en lineas distintas, pero debe
mantener la indentacion

```
- elemento ::

  Lorem ipsum dolor sit amet, consectetur adipiscing elit.

  Dolore magna aliqua. Ut enim ad minim veniam.
```

#### y si tengo una novela
##### Dialogos

```
> "Dialogo, Lorem ipsum dolor sit amet, consectetur adipiscing
  elit, sed eiusmod tempor incidunt ut labore et dolore magna
  aliqua. Ut enim ad minim veniam."
```

para indicar la pertenencia del dialogo a un personaje, podria optarse por la
siguiente sintaxis

```
> personaje A :: "Dialogo, Lorem ipsum dolor sit amet",

> personaje B :: aliqua. Ut enim ad minim veniam."
```

##### "separadores"

```
  Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed eiusmod tempor
  incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.

  > "Dialogo, Lorem", ipsum dolor sit amet, consectetur adipiscing
    elit, sed eiusmod tempor incidunt ut labore et dolore magna
    aliqua. Ut enim ad minim veniam.

  ....

  Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed eiusmod tempor
  incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.
```

un separador indica esos cambios <em>abruptos</em> de escena argumental,
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

:: PELIGRO-PELIGRO-PELIGRO-PELIGRO-PELIGRO-PELIGRO

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
- A :: acronym      - a :: abbr           - 0 ::            - : :: def
- B ::              - b :: bold           - 1 ::            - = ::
- C :: smallCaps    - c :: code           - 2 ::            - ? ::
- D ::              - d ::                - 3 ::            - @ :: escape
- E :: error        - e :: emph           - 4 ::            - ` ::
- F :: func         - f :: file           - 5 ::            - ' :: ‘samp’
- G ::              - g ::                - 6 ::            - " :: “quote”
- H ::              - h ::                - 7 ::            - # :: path
- I ::              - i :: italic         - 8 ::
- J ::              - j ::                - 9 ::
- K :: keyword      - k :: kbd            - ^ :: sup
- L ::              - l :: link           - _ :: sub
- M :: Math         - m :: math (label)   - \ ::
- N :: --> note     - n :: note -->       - | ::
- O :: option       - o ::                - * ::
- P ::              - p ::                - + ::
- Q ::              - q :: quote (label)  - - :: —exp—
- R :: result       - r :: ref            - . ::
- S ::              - s :: strike         - / ::
- T :: radiotarget  - t :: target         - % :: (parentesis)
- U ::              - u :: underline      - & :: symbol
- V :: var          - v :: verbatim       - $ :: command
- W ::              - w ::                - ~ ::
- X ::              - x ::                - ! :: warning
- Y ::              - y ::
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

al expandir `@)` se substituye por `)`, asi:


```
<b>1). Punto uno</b>
```

por lo general el unico aspecto donde seria incomodo utilizar el signo `@` seria
dentro de las direcciones de correo, donbre habria que colocar `@@`, a mi
parecer un precio razonable

##### comentarios

```
@ linea comentada
```

una linea cuyo primer caracter visible sea `@` separada por almenos un espacio
en blanco del texto, comenta la linea en cuestion

#### mas alla del ASCII

preferiblemente se utilizara un sistema de codificacion <q>moderno</q> como
UTF-8.

Opcionalmente (y para no vernos en la necesidad de buscar un caracter
complicado) se puede crear el <q>comando</q> `&`, por ejemplo

```
@&{nombreGenericoDeCaracterComplicado}
```

a modo de ilustrar esto si utilizamos `@&{leftarrow}`, en el texto a exportar
sera reemplazado por `⇐`

#### math

para las formulas matematicas <q>en linea</q> ya que desconosco bastante en este
tema, podriamos no reinventar la rueda y tomar las formulas Tex

```
@M{\formula\Matematica\Tex}
```

(he investigado un poco sobre el tema, quiza sea mejor opcion tomar
la sintaxis que utiliza Libre Office Math o un recien llegado como AsciiMath)

adicionalmente, existira el comando `m` (`@m[cosas]`) para indicar que una seccion
es/o supone ser una formula, de forma similar a como se utiliza un enfasis, pero
con la diferencia, de no tener que invocar al "renderizador" de formulas
matematicas, por ejemplo en lugar de utilizar `@M{H^2O_9}`, mediante caracteres
unicode podria substituirse con `@m{H²O₉}`, con el beneficio adicional de tener un
<em>texto fuente</em> com mayor legibilidad

#### enlaces

```
@l{ruta}
```

equivale a

```
<a href="ruta">ruta</a>
```

y

```
@l{ruta<>descripcion}
```

es equivalente a

```
<a href="ruta">descripcion</a>
```

Ahora que conocemos los comandos `@` y la forma de crear enlaces, podemos
profundizar en un tema antes mencionado: <b>los encabezados</b>, su forma de
referenciarlos y ejemplos de lo que se produciria en una exportacion a html

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

(mas lejos y retorcido aun) un enlace dentro de un encabezado con enfasis

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

##### el mecanismo de los comandos `@`

todas los <q>comandos</q> `@` tienen la estructura `@x(izquierda<>derecha)`. Donde
`izquierda` viene a ser contenido personalizado y opcional. Por su parte `derecha` es
el contenido <i>por defecto</i> del comando. Finalmente `<>` actua a modo de separador entre ambos

Cuando un comando, por ejemplo, un enlace requiere una `izquierda` y este no se ha
especificado, se genera apartir de `derecha`, extrayendo las marcas especiales
dejando unicamente el texto.

Cuando el comando no requiere `izquierda` y esta se proporciona, el comando
o lo ignora o se utiliza como identificador o etiqueta segun sea el caso (nota:
esto aun se encuentra en consideracion y podria requerir sintaxis adicional)

Cuando el comando esta dentro de otro comando, el comando interno *pasa* su
contenido en la `derecha` al comando externo, por ejemplo:

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
diseñados para ser incluidos dentro de parrafos de texto (de hecho para prevenir
errores, si estos no se cierran, se hace de forma automatica al llegar al final
de cada parrafo, o cuando aparece un nuevo elemeto del lenguaje de marcas

Por su parte los <em>comandos de bloque</em> se utilizan para afectar secciones
enteras, indicar acciones complejas o para contenido especializado, como pueden
ser complejas equaciones matematicas o porciones de codigo fuente, sin
preocuparse por el significado de la sintaxis regular

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

El contenido abarca hasta la aparicion de la primer linea en blanco o sin la
indentacion apropiada.

en este tipo de comandos se encuentran los de configuracion del documento

```
..title    > titulo del documento
  puede abarcar varias lineas, siempre con indentacion y sin lineas
  en blanco
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
..figure > esto es el titulo
  de una mini seccion

  este es el contenido de la mini seccion
```

donde el *contenido* empieza luego de la primer linea en blanco.

Como vez, todo depende del comando que se utilize... este es el aspecto mas
complejo del lenguaje de marcas ligeras que propongo y el como, que y cual
funcion desempeña cada comando, es arbitrario y sujeto a una especificacion
unilateral.

dentro de estos, cada seccion puede tener un significado particular, interpretar
o ignorar elementos como los comandos `@` y otras especificaciones, como
alteracion del comportamiento de un comando, secciones adicionales del bloque,
etc.

a continuacion veremos como seria la indicacion para modificar el comportamiento
de un comando de bloque, <q><b>los parametros</b></q>

##### parametros

Aunque los parametros pueden ser variados, deben ser pocos, uno, dos, maxime tres
por comando de bloque. Remarcar que el formato no es para crear espectaculos
visuales.

en la mayoria de herramientas de linea de comandos un parametro tiene la
sintaxis

    --algo="cosa"

o

    -algo cosa

bien?, para morg esto tiene la forma

    algo

o

    algo()

o

    algo( cosa )

o

    algo( cosa-a, cosa-b, ... )

una forma de entenderlo, es como una funcion que tiene valores establecidos por
defecto, y segun el numero de parametros que se envien el resto se colocan de
forma automatica, y en caso de que un parametro (o su numero) sea erroneo,
las secciones erroneas se subtituyen con su valor por defecto.

La diferencia radica en que no se trata de una funcion, sino del nombre de un
parametro y sus valores!

ahora, estos parametros aparecen dentro de los bloques, entre la declaracion del
bloque y el simbolo `>`, quedando asi

```
..bloque algo algo1() algoMas2( "cadena", 123, .21, `raw`, identificador )  >
  cosas
< bloque..
```

encontramos un ejemplo de uso dentro de los bloques de codigo fuente para
indicar un tema especifico (en la exportacion), o numerar el contenido

```
..code n(10) style( "monokai" ) > c
  #include <stdio.h> # esto es codigo en C

  int main(){

    printf("hola, que tal\n");
    return 0;
  }
< code..
```

como es una cosa muy chula, tambien puede utilizarse para especificar opciones
en la configuracion del documento

```
..options > fancyCode() toc
```

aqui se indicaria que deseamos resaltado de codigo sin especificar uno es
particular, solo su resaltado o etiquetado en una exportacion, el segundo
parametro `toc`, sirve para indicar que queremos que la exportacion incluya una
tabla de contenidos

##### ladrillos

esta es otra caracteristica mas (y espero sea la ultima) que complica un poco
mas nuestros bloques de codigo, sin embargo creo que su utilidad en ciertos
bloques, como el codigo fuente (quiza el mas exigente), justifica su existencia

Imagina que tienes un bloque de codigo fuente, y en algun momento lo deseas
"ejecutar" y colocar la salida de su ejecucion dentro del documento, pues hay es
donde aparecen los ladrillos

```
..code > sh
  echo "hola, mundo"
<>
  hola, mundo
< code..
```

exacto, (de nuevo) el `<>` hace de separador, en este caso entre el contenido
principal y los ladrillos... y en este podriamos indicar algurna cosa como por
ejemplo

```
..code > sh
  random
<> ejecucion 1
  1236
<> ejecucion 2
  47
< code..
```

y asi sucesivamente.

otro comando interesante seria `cols` para expresar que deseamos crear columnas,
con su contenido

```
..cols >

  este es el texto de la columna 1

<>

  este es el texto de la columna 2

<>

  contenido de la columna 3

...
```

#### srci

este es un comando de bloque complementario con el de codigo fuente donde se
simula un prompt de algun lenguaje y su salida, sin hechar mano de los ladrillos

```
..srci > lenguaje
  > (esto simularia ser (codigo)
  ^   (en lisp (o algo)
  ^            (por el estilo)))
  esta seria la supuesta salida
  que produce el codigo
```

La entrada o prompt de lenguaje inicia con `>` y cada linea inmediatamente
despues que inicie con `^`, tambien se asume como parte de esta

lo interesante de este bloque, es que con una sintaxis comun se puede
simular la supuesta estrada de casi cualquier lenguaje!

Para casos donde el lenguaje tenga mas de un promt, habria que utilizar un
parametro adicional y supongo que diversos bloques, por ejemplo para simular una
secion como usuario de a pie en bash y luego un logueo como root, seria

```
..srci prompt( "$ " ) > sh
  > rm -rf /boot
  permiso denegado
  > su
  password

aqui con algun texto se alerta de los peligros de ser root y ejecutar
comandos con poder ilimitado

..srci prompt( "# " ) > sh
  > whoami
  root
  > rm -rf /boot
  error: lo siento su princesa se encuentra en otro castillo
```

#### tablas

Sin duda un tema complejo, podria tenerse una tabla totalmente funcional con
formulas y demas, pero para inciar:

```
+----------------------+
|   mi tabla compleja  |
+=======+====+=========+
|  a    | b  |    c    |
+---+---+----+----+----+
| d | e | f  |  g |  h |
|   |   +----+----+----+
|   |   | i  |    j    |
|   +---+----+---------+
|   | k | l  |         |
+---+   +----+|||m||||||
|   |   | o  |         |
| n |   +----+---------+
|   |   | p  |    q    |
+---+---+----+---------+
```

el estilo (con esquinas `+` y contornos `|` o `-`) de tablas que utiliza
reStructuredText... con el adicional de poder dividir el encabezado
con un divisor que abarque toda la linea, substituyendo `-` por `=`.

Tambien se brindan "pies" de tabla con un divisor que abarque toda la linea,
substituyendo `-` por `~`.

#### porg

los ficheros `po` producidos con `gettext` se utilizan para traducir documentos
de un lenguaje a otro, siempre que gettext no muera en el intento... con morg podemos
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

### un poco de accion

De momento, a modo de <q>prueba de concepto</q> existe **morg**, un programa
(mal) desarrollado en golang, con el cual se puede exportar (con limitaciones)
algunos conceptos basicos del lenguaje al formato **html** asi como poder
visualizar el documento dentro de un terminal. Para mas informacion
vea [Como iniciar con morg](howto.md).

De forma simplificada, para optener el programa

``` sh
go get github.com/nasciiboy/morg
```

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

#### el codigo

(mis habilidades como programador son pocas y limitadas, sin embargo, creo en el
codigo como forma de expresion y transmision de conocimiento, si las siguientes
metas resultan irreales lo puedo comprender, aunque no aceptar...)

sin importar el lenguaje, el codigo debe ser elegante o almenos claro y
sencillo, evintando dependencias inecesarias que dificulten su adaptacion a
otros lenguajes o peor aun, el aprendisaje de otros programadores. En resumen se
buscara siempre ser una implementacion de referencia con toques didacticos en la
que cualquier indicio de aparicion de cruft sera señal de refleccion y futura
refactorizacion e incluso reescritura

conceptualmente se plano dividir el programa en varias secciones, `katana` el
encargado de parcear el documento, para entregar una estructura de datos
sencilla

`biskana` *el traductor* a otros lenguajes o dicho de otra manera el
*renderizador* (que hace realidad nuestras fantacias) de texto a texto

`nirvana` (cofff, en un principio iba a ser `hana` (flor) pero lo olvide!)
encargado de desplegar el documento de forma visual como TUI o como GUI

`*ana` (aun no implementado) la base de datos donde se consulta si tenemos el
documento de forma local o debe realizarce una peticion externa. Ademas debe
regresar el documento en el formato original

El primer componente que programe fue `biskana`, de hay, el
deseo de terminar los nombres en `ana`. Alguna propuesta interesante?

Como cohesionador de todo, el propio `morg` (aun no me convence el nombre), se
aceptan sujerencias!

Por cierto `katana` hace uso de un motor de expresiones regulares elaborado
desde cero llamado `regexp4` que por mera casualidad tambien programe. Lo
utilizo por estas razones

1. funciona!

2. no ha aparecido alguna exprecion que revase su limitada capacidad

3. es sencillo y facil de modificar (por mi, almenos), ademas cuando surge un
   error se donde buscar

4. puede portarse con relativa facilidad a cualquier lenguaje (creo), el
   desarrollo original fue hecho en C. Cuando digo en C me refiero a solo C, sin
   recurrir a ninguna libreria, ni siquiera la libreria estandar. Su port a go
   no fue demasiado traumatico, encima se vio veneficiado por la orientacion a
   objetos

#### porque Go

por nada en especial... cuando empece a escribir el codigo en C (en abril del 2016)
mis habilidades no daban para mucho, no es que ahora sea un jodido guru,
pero algo he aprendido, he? pero por que C? por velocidad y eficiencia, si vas a
hacer un proyecto tan ambicioso, mal seria que fuese lagueado y demorara en
exeso (mas de 5 segundos) para mostrar el contenido.

Aun si el proyecto se escribiese en un lenguaje interpretado en algun momento
deberia portarse a un lenguaje veloz (no mas de 6 veces inferior al rendimiento
de C)

podria haber elegido C++... pero me cruce con Go, que con el paso del tiempo ha
llegado a probacar fascinacion en mi, es solo 3 veces mas lento que C y tiene ideas muy
interesantes. Como unicas desventajas creo que su tipado fuerte es un fastidio, al igual que
el estilo de codificacion, llaves forsosas para instrucciones simples y el no
contar con punteros de verdad (al menos con un nivel de indireccion) obliga a
hacer algunos apaños.

Ademas de su eficiencia, a su fabor Go (con su rica libreria estandar) esta
pensado para ser verdaderamente multiplataforma, es sencillo en su orientacion a
objetos, claro y automatico con la gestion de dependencias y de lectura agradable

#### primer ejemplo

seria grosero no monstrar ni un poco, asi que
[aqui](https://github.com/nasciiboy/tgpl) encontran un ejemplo del formato
actual, con exportacion a html + una muestra de porg

Para ver el resultado en todo su esplendor, clona o baja una copia del repo y
visualiza en el navegador el html.

### zen

Extenso, esto es ya, un programador al inicio de su travesia soy y antes de
empezar a programar, modificar o agregar funciones deberia establece una
especificacion para el formato.

Mas tarde, como calentamiento hacer un exportador robusto y luego los demas
componentes. De camino integrarlo en algunos cms, darle soporte en nuestros
editores favoritos, al tiempo de otorgarle facilidas de autocompletado y
resaltado, establecer una estructura capaz de mantenerse por cuenta propia y/o
por la coperacion desinteresado de empresas y gobiernos, y solamente despues de
ello absorver todo el contenido que sea posible y dominar la galaxia

## como puedo ayudar

Pagameeeee un salario, (enserio, [mi correo](mailto:nasciiboy@gmail.com))

Contrata a un grupo de programadores motivados, que compartan este sueño (y
pagame un salario!)

traduce, difunde, discute y comenta

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
obligado) bajo esta misma licencia.

Mientras tanto programare lo que pueda, cuando pueda, simplemente por
ser un reto emocionante

Happy Hacking!
