package utility

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type (
	/*StArchMa : estructura maestra que crea directorios o archivos masivos.*/
	StArchMa []StArch
	/*StArch : estructura para registra los archivos y generarlos.*/
	StArch struct {
		Path   string
		IndDir bool
	}
	/*StLog : Estructura para crear log personalizados por medio de la fecha
	directorio/fecha.log
	*/
	StLog struct {
		//Directorio contenido
		Dir string
		//Nombre del log mas fecha
		Name string
		//Fecha
		Fe time.Time
		//Prefijo
		Prefix string
	}
)

/*Create : crear el archivo o directorio vacio si existe no lo crea */
func (p *StArchMa) Create() error {
	for _, item := range *p {
		err := item.Create()
		if err != nil {
			return nil
		}
	}
	return nil
}

/*Create : crear el archivo o directorio vacio si existe no lo crea */
func (p *StArch) Create() error {
	if !FileExist(p.Path, p.IndDir) {
		var err error
		if p.IndDir {
			err = DirNew(p.Path)
		} else {
			_, err = FileNew(p.Path)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

/*Printf : Ingresa un texto en los logs asignado. */
func (p *StLog) Printf(format string, args ...interface{}) error {
	err := p.Init()
	if err != nil {
		return err
	}
	log.Printf(format, args...)
	return nil
}

/*Init : Inicializa el log para comenzarlo a usar */
func (p *StLog) Init() error {
	NameArch := fmt.Sprintf("%s/%s%s.log", Trim(p.Dir), Trim(p.Name), Trim(TimetoStr(p.Fe)))
	if !FileExist(p.Dir, true) {
		return Msj.GetError("AR05")
	}
	log.SetPrefix(p.Prefix)
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	file, err := FileNew(NameArch)
	if err != nil {
		return err
	}
	log.SetOutput(file)
	return nil
}

/*NewStLog : Crea una nueva intancia de StLog */
func NewStLog(dir, name, prefix string, fe time.Time) (StLog, error) {
	LogNew := StLog{
		Dir:    dir,
		Name:   name,
		Prefix: prefix,
		Fe:     fe,
	}
	err := LogNew.Init()
	if err != nil {
		return LogNew, err
	}

	return LogNew, nil
}

/*FileExt : Valida las extenciones de archivos.*/
func FileExt(Path string, ext string) bool {
	if !FileExist(Path, false) {
		return false
	}
	return ReturnIf(strings.Index(Path, EXT[ext]) > 0, true, false).(bool)
}

/*FileRename : Renombra a un archivo como tambien lo puede mover a otro directorio de manera nativa.*/
func FileRename(PathOrigen, PathNuevo string) error {
	err := os.Rename(PathOrigen, PathNuevo)
	if err != nil {
		return err
	}
	return nil
}

/*FileExist : Valida si el archivo del path existe antes de procesarlo.
Valida tambien si existe un directorio con el inddir en true
*/
func FileExist(Path string, inddir bool) bool {
	info, err := os.Stat(Path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	if inddir {
		return info.IsDir()
	}
	return true
}

/*FileNew : crea un archivo X*/
func FileNew(p string) (*os.File, error) {
	f, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, Msj.GetError("AR03")
	}
	return f, nil
}

/*TrimFile : Renombra a una carpeta o archivo quitandole todos los espacio regresando
el path del nuevo archivo.*/
func TrimFile(Path string) (string, error) {
	if !FileExist(Path, false) {
		return "", Msj.GetError("AR01")
	}
	PathOrig := Path
	Path = Trim(strings.Replace(Path, "\r", "", -1))
	err := FileRename(PathOrig, Path)
	if err != nil {
		return "", err
	}
	return Path, nil
}

/*DirNew : Crea una carpeta vacia en el sistema*/
func DirNew(Path string) error {
	err := os.MkdirAll(PlecaAdd(Path), os.ModePerm)
	if err != nil {
		return Msj.GetError("AR02")
	}
	return nil
}

/*CpFile : copia un archivo Origen a un directorio destino*/
func CpFile(PathOrig, PathDest string) error {
	fileNew := new(os.File)
	PathDest = PlecaAdd(PathDest)
	if !FileExist(PathOrig, false) {
		return Msj.GetError("AR01")
	}
	if !FileExist(PathDest, true) {
		return Msj.GetError("AR05")
	}

	pathNew, err := TrimFile(PathOrig)
	if err != nil {
		return err
	}
	fileOrig, err := os.Open(pathNew)
	if err != nil {
		return err
	}
	infoFile, err := os.Stat(fileOrig.Name())
	if err != nil {
		return err
	}
	pathFinal := PathDest + infoFile.Name()
	fileNew, err = FileNew(pathFinal)
	if err != nil {
		return err
	}
	_, err = io.Copy(fileNew, fileOrig)
	if err != nil {
		return err
	}
	return nil
}

/*CpDir : copia una carpeta entera a una carpeta destino*/
func CpDir(PathOrig, PathDest string) error {
	PathOrig = PlecaAdd(PathOrig)
	PathDest = PlecaAdd(PathDest)
	if !FileExist(PathOrig, true) {
		return Msj.GetError("AR05")
	}
	if !FileExist(PathDest, true) {
		err := DirNew(PathDest)
		if err != nil {
			return err
		}
	}
	archs, err := ListDir(PathOrig)
	if err != nil {
		return err
	}
	for _, arch := range archs {
		if arch.IsDir() {
			err = CpDir(PathOrig+arch.Name(), PathDest+arch.Name())
			if err != nil {
				return err
			}
		} else {
			err = CpFile(PathOrig+arch.Name(), PathDest)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

/*RmDir : Elimina un directorio entero*/
func RmDir(src string) error {
	src = PlecaAdd(src)
	if !FileExist(src, true) {
		return Msj.GetError("AR05")
	}

	archs, err := ListDir(src)
	if err != nil {
		return err
	}
	for _, arch := range archs {
		if arch.IsDir() {
			err = RmDir(src + arch.Name())
			if err != nil {
				return err
			}
		} else {
			err = RmFile(src + arch.Name())
			if err != nil {
				return err
			}
		}
	}
	err = RmFile(src)
	if err != nil {
		return err
	}

	return nil
}

/*RmFile : elimina un archivo exacto*/
func RmFile(file string) error {
	if !FileExist(file, false) {
		return Msj.GetError("AR01")
	}
	err := os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}

/*ListDir : lista la infomacion que contiene una carpeta*/
func ListDir(src string) ([]os.FileInfo, error) {
	src = PlecaAdd(src)
	if !FileExist(src, true) {
		return nil, Msj.GetError("AR05")
	}
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return nil, Msj.GetError("AR06")
	}
	return files, nil
}

/*PlecaAdd : coloca la pleca de un directorio "/" */
func PlecaAdd(src string) string {
	if src[len(src)-1] != '/' {
		src = src + "/"
	}
	return src
}
