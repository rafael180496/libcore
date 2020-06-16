# **Intrucciones:**

## **Intalacion de libreria [rana/ora]([https://gopkg.in/rana/ora.v4])**

### **LINUX**

**Creamos una carpeta nueva y le asignamos permisos:**

```bash
sudo mkdir /usr/local/share/pkgconfig
sudo chmod 777 /usr/local/share/pkgconfig
```

**Accedemos a la carpeta que hemos creados:**

```bash
cd /usr/local/share/pkgconfig
```

**Creamos un archivo con el nombre oci8.pc con el siguiente contenido:**

```c
prefix=/usr
version=19.3
build=client64

libdir=${prefix}/lib/oracle/${version}/${build}/lib
includedir=${prefix}/include/oracle/${version}/${build}

Name: oci8
Description: Oracle database engine
Version: ${version}
Libs: -L${libdir} -lclntsh
Libs.private:
Cflags: -I${includedir}
```

### **Validar que se tiene instalado completamente el Oracle Instant Client en su version 12.1 / 19.3:**

```bash
oracle-instantclient19.3-basic-19.3.0.0.0-1.x86_64
oracle-instantclient19.3-devel-19.3.0.0.0-1.x86_64
oracle-instantclient19.3-precomp-19.3.0.0.0-1.x86_64
oracle-instantclient19.3-sqlplus-19.3.0.0.0-1.x86_64
oracle-instantclient19.3-tools-19.3.0.0.0-1.x86_64
```

### **Validar que se posee las variables de entornos configuradas:**

```bash
sudo nano ~/.profile

export PKG_CONFIG_PATH=/usr/local/share/pkgconfig
export ORACLE_HOME=/usr/lib/oracle/19.3/client64
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$ORACLE_HOME/lib:$ORACLE_HOME

export OCI_HOME=/usr/lib/oracle/19.3/client64/lib
export OCI_LIB_DIR=$ORACLE_HOME/lib
export OCI_INC_DIR=/usr/include/oracle/19.3/client64
export OCI_INCLUDE_DIR=/usr/include/oracle/19.3/client64
export OCI_VERSION=12

export PATH=$PATH:$ORACLE_HOME/bin
export GOPATH=/home/gvalladares/Repo/goProyectos

export C_INCLUDE_PATH=/usr/include:/usr/include/oracle/19.3/client64:/usr/local/include
```

### **Si se modifico el profile, se debera de recargar las variables de entorno:**

```bash
source ~/.profile
```

### **En el proyecto que requerimo el paquete procederemos a descargarla e instalarla:**

```bash
go get -u -v gopkg.in/rana/ora.v4
```

### **Windows:**

#### **Requisitos:**

* Oracle Client 64-bit  (Cliente12.1 o superior de preferencia)
* Instalacion de Go en 64-bit
* 64-bit MinGW-64 instalado:[mingw-w64]([https://sourceforge.net/projects/mingw-w64/]).
* Versión de GIT para Windows

### **Configuración de pkg-config**

Crear la carpeta  **C:\pkg-config** y descomprime en su interior los siguientes archivos:

```bash
http://ftp.gnome.org/pub/gnome/binaries/win64/dependencies/gettext-runtime_0.18.1.1-2_win64.zip

http://ftp.gnome.org/pub/gnome/binaries/win64/dependencies/pkg-config_0.23-2_win64.zip

http://ftp.gnome.org/pub/gnome/binaries/win64/glib/2.26/glib_2.26.1-1_win64.zip
```

Crear una carpeta siguiente **c:\pkg-config\PKG_CONFIG_PATH**
y crea un archivo con el nombre **oci8.pc** en el interior de la carpeta anteriormente creada:

```c
oracle=c:/app/client/user/product/12.1.0/client_1
prefix=/devel/target/XXXXXXXXXXXXXXXXXXXXXXXXXX
exec_prefix=${prefix}
libdir=${oracle}/oci/lib/msvc
includedir=${oracle}/oci/include
glib_genmarshal=glib-genmarshal
gobject_query=gobject-query
glib_mkenums=glib-mkenums
Name: oci8
Version: 12.1
Description: oci8 library
Libs: -L${libdir} -loci
Cflags: -I${includedir}
```

### **Variables de entorno**

Se agrega en la variable de entorno PATH, la ruta del directorio BIN de la instalación **MinGW-64** y de la ruta de pkg-config: **c:\pkg-config\bin**

Se necesita crear una variable de entorno con el nombre de **PKG_CONFIG_PATH** con el valor siguiente **C:\pkg-config\PKG_CONFIG_PATH.**

#### **Instalación de la librería para Oracle**

```bash
go get gopkg.in/rana/ora.v4

go install -v gopkg.in/rana/ora.v4
```

#### **Codigo de prueba:**

```go
package main

import (
        "database/sql"
        _ "gopkg.in/rana/ora.v4"
        "fmt"
)

func main() {
        db, err := sql.Open("ora", "user/pass@servidor/DB")
        if err !=nil {
                fmt.Println("Connection error", err)
        }
        defer db.Close()

        qry := "SELECT user FROM dual"
        fmt.Println("Running query:", qry)
        rows, err := db.Query(qry)

        if err != nil {
                fmt.Println("Error:", err)
        }

        defer rows.Close()
        var user string
        for rows.Next() {
                if err = rows.Scan(&user); err != nil {
                        fmt.Println("Error:", err)
                        break
                }
                fmt.Println(user)
        }
}
```

### **Validación del entorno**

```bash
C:\>systeminfo | findstr /B /C:"OS Name" /C:"OS Version"
OS Name:                   Microsoft Windows 10 Pro
OS Version:                10.0.15063 N/A Build 15063

C:\>gcc --version
gcc (x86_64-posix-seh-rev2, Built by MinGW-W64 project) 7.1.0
Copyright (C) 2017 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

C:\>go version
go version go1.8.3 windows/amd64

C:\>tnsping
TNS Ping Utility for 64-bit Windows: Version 12.1.0.2.0 - Production on 24-AUG-2017 09:13:00
Copyright (c) 1997, 2014, Oracle.  All rights reserved.
TNS-03502: Insufficient arguments.  Usage:  tnsping <address> [<count>]

C:\>git --version
git version 2.14.1.windows.1
```
