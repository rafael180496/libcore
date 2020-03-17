package test

import (
	"testing"

	utl "github.com/rafael180496/libcore/utility"
)

/*TestCpfile : Copia un archivo orig a un archivo dest*/
func TestCpfile(t *testing.T) {
	pathorig := "config/pr1/prueba.txt"
	utl.CpFile(pathorig, "config/pr2/")
}

/*TestCpDir : Copia un directorio orig a un archivo dest*/
func TestCpDir(t *testing.T) {
	pathorig := "config/pr1/"
	utl.CpDir(pathorig, "config/pr2/")
}

/*TestRmFile : test para eliminar un archivo*/
func TestRmFile(t *testing.T) {
	pathorig := "config/pr2/prueba.txt"
	utl.RmFile(pathorig)
}

/*TestRmDir : Elimna un directorio completo*/
func TestRmDir(t *testing.T) {
	pathorig := "config/pr2/"
	utl.RmDir(pathorig)
}

/*TestListDir : lista la informacion de un directorio*/
func TestListDir(t *testing.T) {
	pathorig := "config/pr1"
	info, _ := utl.ListDir(pathorig)
	utl.PrintPc("%v", info[0].Name())
}
