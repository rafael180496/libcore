# LIBCORE

## **Creaciones de librerias**

### **TEST:**

Carpeta donde contiene ejemplos de las librerias creadas.

### **Utility:**

Carpeta contiene funciones variadas de todo el proyecto.

### **Database:**

Carpeta contiene funciones variadas sobre conexion a base de datos.

### **Server:**

Carpeta contiene funciones variadas sobre servicio con ECHO.

### **Go module**

Es la forma de obtener dependicias y de liberarse
del **GOPATH** esto crea un entorno cerrado al proyecto.

#### **Documentacion**

* [Introduccion]([https://medium.com/mindorks/create-projects-independent-of-gopath-using-go-modules-802260cdfb51])

* [GoLibModule]([https://blog.golang.org/using-go-modules])

  **Inicializar**

```bash
    go mod init github.com/{your_username}/{repo_name}
    or go test -v #para inicializar las dependencias
    go build
    ./gomod
```

### **Librerias Externas:**

**Toda libreria externa se debe de intalar para que
funcione correctamente el proyecto.**

* [fatih/color]([https://github.com/fatih/color])

* [rana/ora]([https://gopkg.in/rana/ora.v4])

* [lib/pq]([https://github.com/lib/pq])

* [go-sql-driver/mysql]([https://github.com/go-sql-driver/mysql])

* [denisenkom/go-mssqldb]([https://github.com/denisenkom/go-mssqldb])

* [jmoiron/sqlx]([github.com/jmoiron/sqlx])

* [go-ini/ini]([https://github.com/go-ini/ini])

* [labstack/echo]([https://github.com/labstack/echo])

* [satori/go.uuid]([https://github.com/satori/go.uuid])

## **Nomenclatura:**

* Todo paquete comienza con **cr** para entender que es del proyecto core.
* Las constante son todas en mayuscula y deben de estan en un archivo aparte con las variables globales.
* Las funciones deben de tener una descripcion resumida pero directas del funcionamiento.
* Las funciones o estructuras publicas comienzan con mayusculas.
* Todo atributo de una estrucutra publica comienza con mayusculas.
* El uso de puntero se identificara con un **ptr** al comienzo de la variable ejemplo **PtrConexion**.
* Los nombres de las variables no pueden llevar separaciones con  **_** si no para separar llevara otra mayuscula.
* Las constante deben de tener una descripcion asi como tambien las estructuras.
* Las estructura deberan de comenzar con las 3 primeras iniciales de la libreria.
* Las estructura comienzan con **st**.
* Los nombre de archivos que contienen constante o variables globales terminan con **const**.
* Toda importancia estatica debe llevar una descripcion.
* Todo mensaje se creara en la variable global **Msj** con su codigo.
* Al usar canales usar las variables con el prefijo **Chan** al final esto indica que es una variable de canal.

## **Ejemplo de descripciones:**

**Funciones:**

~~~~go
/*Cierre : cierra las conexiones de base de datos intanciadas*/
func (p *StConexion) Cierre() error {}
~~~~

**Estructura:**

~~~~go
/*StCadConexion : Estructura para generar la cadena de  conexiones de base de datos */
type StCadConexion struct {
    Nombre `json:"nombre"`
}
~~~~

**Constante:**

```go
/*Mong : conexion tipo mongodb */
const Mong = "MONGO"
```

**Importanciones Estaticas:**

```go
 /*Conexion a mysql*/
import(
    _ "github.com/go-sql-driver/mysql"
    )
```

**Variables Globales e punteros:**

```go
var (
    /*EXT : extensiones de archivos */
        EXT = map[string]string{
        "JSON": ".json",
        "INI":  ".ini",
        "XML":  ".xml",
        }
        ptrNombre *string
    )
```

**Mensaje:**

```go
    MensajesGrnl = map[string]string{
        "ERR1":"No existe el archivo"
        }
```

**Canales:**

```go
    EjemploChan := make(chan int)
```
