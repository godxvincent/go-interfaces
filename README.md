# Learning Go Interfaces

Este repo contiene ejemplos básicos de como funcionan las interfaces y un ejemplo básico de como implementar el patron builder en go.

## Tabla de contenido

- [Notas sobre interfaces y structs](##Notas-sobre-interfaces-y-structs)
- [Ejemplos códificados en este repo](#Ejemplos-códificados-en-este-repo)





------------

## Notas sobre interfaces y structs
1. Para que un struct pueda ser retornado en lugar de una interfaz dicha structura `(clase)` debe implementar como en java todos los metodos
2. Ahora bien el punto uno también aplica en el caso de  querer pasarle a una funcion un parametro cuyo parametro esta definido del tipo de una `(interface)`, para eso la  `(clase)` debe implementar todos los metodos.

    > Note: As a reminder, a type is said to implement an interface when it defines method sets available on the interface type. [Exploring structs and interfaces in Go](https://blog.logrocket.com/exploring-structs-and-interfaces-in-go/ "Exploring structs and interfaces in Go")
3. Sin embargo, cuando se define una estructura `(clase)` que tiene un atributo del tipo de esa interface la `(clase)` solo debe implementar el metodo que requiera y así podrá pasar dicho `(objeto)` de esa `(clase)` a la función. 
4. Cuando una en una estructura se define un campo del tipo de otra estructura y durante el programa se hace llamado al objeto de esa ultima estructura de forma directa se dice que los campos fueron promovidos
    ```go
    type Direccion struct {
        direccionLarga string
        barrio         string
        localidad      string
    }

    type Persona struct {
        nombre string
        edad   int
        Direccion
    }

    func main() {
        persona := &Persona{}
        fmt.Println(persona.barrio) // Esto es posible porque el campo dirección ha sido promovido
    }
    ```
5. Es posible hacer que una clase tenga un campo del tipo interface, al hacer esto la estructura de manera implicita implementa los metodos de la interface lo que ayuda que dicha clase solo tenga que implementar *interceptar* los metodos que requiere sobre escribir. [embedding-in-go-part-3-interfaces-in-structs](https://eli.thegreenplace.net/2020/embedding-in-go-part-3-interfaces-in-structs/)


## Ejemplos códificados en este repo

1. **Definición de estructuras (Clases):** Archivo `models/usuario.go` lineas `3 a 9`
    * Esta estructura define sus atributos en minuscula lo que hace no sean exportables sus atributos. Equivalente a `private` en java
    * Debido al punto anterior se definen metodos get para setear los valores lineas `11 a 22`
    * No se implementa algo similar a *setters* ya que se implementa el builder en el mismo archivo

2. **Implementación patron builder:** `models/usuario.go` lineas `24 en adelante` y `main.go`
    *  Para hacer uso del builder del usuario se crea una interface de nombre `UserBuilderInterface` junto con los métodos que la `struct` debe implementar lineas `24 a 31`
    * Se crea un `struct` de nombre `userBuilder` no exportable que implementará lo indicado por la interface lineas `33 a 39`
    * Se crean dos funciones exportables que crean objetos concretos del tipo `UserBuilder` en lugar de los tipos`userBuilderInterface` esto gracias a que el `struct` implementa la interface lineas `41 a 47`
    * Se implementan los metodos definidos por la interface lineas `49 en adelante`
    * La forma de usar el builder y la `struct` se pueden ver en el archivo `main.go` lineas `16 a 28`

3. **Estructuras con tipos:** `main.go`
    * Se pueden definir `struct` solo con los tipos y apuntadores. Sin embargo, estos no pueden repetirse en la misma estructura lineas `9 a 12`
    * Ejemplo de su implementación lineas `29 a 31`
    * Nota adicional al definir un tipo con apuntador no se puede pasar el valor directo hay que definir la variable y luego si envíar la referencia.

4. **Funciones e Interfaces:** 
    
    Para esta parte hay dos implementaciones. La primera son funciones que reciben como parametros una `interface` y el argumento de entrada es una `struct` que implementa dicha interface. 

    * La interface se encuentra definida en el archivo `models/interfaces.go` y se llama `SaludosI`
    * La `struct` que la implementa se encuentra definida en el archivo `models/estructuras.go` y se llama `Saludos`.
    * La función `ProbarInterfaceSaludos` define un parametro del tipo de la interface y esta función se encuentra en `models/interfaces.go`
    * El llamado de dicha función y creación del objeto que corresponde a la `struct` se encuentra en el `main.go` en las lineas `31 a 35`

    La segunda implementación consiste en la misma función que espera una interfaz, pero esta vez recibe un `struct` que solo implementa de manera explicita uno de los metodos de las interfaces. Esto se debe a que al tener embebida la interface como un atributo go interpreta esto como una clase de herencia o implementación explicita de las interfaces. Esto permite que el desarrollador pueda solo **sobreescribir** una de las funciones que requiera intervenir.

    * La `struct` que tiene la interface embebida se llama `SaludosInterfaceEmbebida` y esta en el archivo `models/estructuras.go`
    * Esta `struct` solo implementa el metodo `Generico()` y se encuentra en el mismo archivo `models/estructuras.go`
    * En el `main.go` se puede ver ver en las líneas `42 a 44` que se crea un objeto de esta `struct` y se invoca la función sin arrojar error.

5. **Interface generica:** 
    * Se define una función básica que recibe una interface vacia para validar el hecho de que se puede pasar cualquier tipo.
