package main

import (
	"fmt"

	"github.com/rafael180496/libcore/database"
)

func main() {
	var (
		conexion database.StConect
	)
	path := "data.ini"
	err := conexion.ConfigINI(path)
	if err != nil {
		fmt.Printf("Error:%s", err.Error())
	}
	fmt.Printf("Conexion:%s", conexion.Conexion.ToString())
	fmt.Printf("Probando...")
	fmt.Printf("prueba:%v", conexion.Test())
}
