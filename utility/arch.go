package utility

import (
	"log"
	"os"
	"strings"
	"time"
)

type (
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
	FechaText := TimetoStr(p.Fe)
	NameArch := strings.Replace(p.Dir+"/"+p.Name+FechaText+".log", " ", "", -1)

	if !FileExist(p.Dir, true) {
		return SendErrorCod("AR05")
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
	result := strings.Index(Path, EXT[ext])
	if result > 0 {
		return true
	}
	return false
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
		return nil, SendErrorCod("AR03")
	}
	return f, nil
}

/*TrimFile : Renombra a una carpeta o archivo quitandole todos los espacio regresando
el path del nuevo archivo.*/
func TrimFile(Path string) (string, error) {
	if !FileExist(Path, false) {
		return "", SendErrorCod("AR01")
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
	if Path[len(Path)-1] != '/' {
		Path = Path + "/"
	}
	err := os.MkdirAll(Path, os.ModePerm)
	if err != nil {
		return SendErrorCod("AR02")
	}
	return nil
}
