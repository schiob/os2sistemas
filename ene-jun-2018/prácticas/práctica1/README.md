# Práctica 1
Preparándonos para usar go, git y github, escriban el siguiente programa.

## hola.go
Escribe un programa que reciba por stdin una palabra `x` e imprima  `hola` desde un hilo y la palabra `x` desde otro hilo, debe de imprimir cada palabra 10 veces.

### Input
Una sola linea con una palabra.

### Output
Imprime 20 lineas, 10 con `hola` y 10 con la palabra `x`, como cada hilo trabaja por su cuenta el orden de cada palabra será 'aleatorio'.


### Ejemplos

#### Entrada
```
chio
```
#### Salida

```
hola
hola
chio
hola
chio
chio
chio
hola
hola
chio
chio
hola
chio
hola
hola
chio
chio
hola
hola
chio
```

## Referencias
 - [tour de go - goroutines](https://tour.golang.org/concurrency/1)
 - [go by example - goroutines](https://gobyexample.com/goroutines)
 - [artículo explicando concurrencia y goroutines en go](https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html)
 - [golang-book - concurrency](https://www.golang-book.com/books/intro/10)

## Consejos
Investiguen mucho, y hagan pruebas.

Si les sale un error tranquilos, googlen. También está el grupo para preguntas.

![don`t panic](https://i.imgur.com/jClfB2I.gif "Dont panic")
