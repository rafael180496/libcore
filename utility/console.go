package utility

import (
	"fmt"
	"net"
	"runtime"
	"strings"

	color "github.com/fatih/color"
)

/*MaxCPUtask : maxima multi hilos que puede obtener   */
func MaxCPUtask() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

/*NumProcSet : setea los numeros de proceso validos */
func NumProcSet(n int) int {
	return runtime.GOMAXPROCS(n)
}

/*GetCPU : envia la cantidad de CPU disponible en el procesador */
func GetCPU() int {
	n := runtime.NumCPU()
	if n <= 0 {
		return 1
	}
	return n
}

/*MsjPc : envia un string de color personalizado de las
constantes disponible en la libreria */
func MsjPc(c Pc, format string, arg ...interface{}) string {
	var (
		menj strings.Builder
	)
	if !IsLinux() {
		fmt.Fprintf(&menj, format, arg...)
	} else {
		d := color.New(sendColor(c), color.Bold)
		d.Fprintf(&menj, format, arg...)
	}
	return menj.String()
}

/*PrintPc : muestra un printf con color personalizado para consolas
basadas en linux*/
func PrintPc(c Pc, format string, arg ...interface{}) {
	if !IsLinux() {
		fmt.Printf(format, arg...)
	} else {
		d := color.New(sendColor(c), color.Bold)
		d.Printf(format, arg...)

	}
}

/*MsjBlue : Enviar un string de color celeste funciona para consolas
basadas en linux */
func MsjBlue(format string, arg ...interface{}) string {
	return MsjPc(Blue, format, arg...)
}

/*MsjRed : Enviar un string de color rojo para consolas
basadas en linux*/
func MsjRed(format string, arg ...interface{}) string {
	return MsjPc(Red, format, arg...)
}

/*MsjGreen : Enviar un string de color verde para consolas
basadas en linux*/
func MsjGreen(format string, arg ...interface{}) string {
	return MsjPc(Green, format, arg...)
}

/*PrintGreen : muestra un printf con color verde para consolas
basadas en linux*/
func PrintGreen(format string, arg ...interface{}) {
	PrintPc(Green, format, arg...)
}

/*PrintRed : muestra un printf con color rojo para consolas
basadas en linux*/
func PrintRed(format string, arg ...interface{}) {
	PrintPc(Red, format, arg...)
}

/*sendColor : envia el color correcto en atributo */
func sendColor(item Pc) color.Attribute {
	switch item {
	/*Green : verde */
	case Green:
		return color.FgGreen
		/*Red : rojo */
	case Red:
		return color.FgRed
		/*Blue : azul */
	case Blue:
		return color.FgBlue
		/*Cyan : celeste */
	case Cyan:
		return color.FgCyan
		/*White : blanco */
	case White:
		return color.FgWhite
		/*Black : negro */
	case Black:
		return color.FgBlack
		/*Yellow : amarillo*/
	case Yellow:
		return color.FgYellow
		/*Magenta : magenta */
	case Magenta:
		return color.FgMagenta
		/*HiBlack : negro fuerte */
	case HBlack:
		return color.FgHiBlack
		/*HRed : rojo fuerte */
	case HRed:
		return color.FgHiRed
		/*HGreen : verde fuerte */
	case HGreen:
		return color.FgHiGreen
		/*HYellow : amarrillo fuerte */
	case HYellow:
		return color.FgHiYellow
		/*HBlue : azul fuerte */
	case HBlue:
		return color.FgHiBlue
		/*HMagenta : magenta fuerte*/
	case HMagenta:
		return color.FgHiMagenta
		/*HCyan : celeste fuerte */
	case HCyan:
		return color.FgHiCyan
		/*HWhite : blanco fuerte */
	case HWhite:
		return color.FgHiWhite
	default:
		return color.FgWhite

	}

}

/*GetLocalIPV4 : te envia la ip local que contiene la maquina que estas ejecutando el programa */
func GetLocalIPV4() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

/*IsLinux : Valida si estas en un sistema operativo linux */
func IsLinux() bool {
	switch runtime.GOOS {
	case "windows":
		return false

	case "linux", "darwin":
		return true
	default:
		return false
	}
}
