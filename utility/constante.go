package utility

type (
	/*Pc : es un tipo para las paletas de colores.*/
	Pc string
)

var (
	/*TypeContent : contents type para los formatos de correo */
	TypeContent = map[string]string{
		"html": "text/html",
		"text": "text/plain",
	}
	/*SMTPPORT : indica los puertos disponible para el smtp */
	SMTPPORT = map[string]string{
		"gmail1":   "587",
		"hotmail1": "25",
		"hotmail2": "465",
		"yahoo1":   "465",
		"yahoo2":   "587",
	}
	/*SMTPURL :  Host de correos */
	SMTPURL = map[string]string{
		"gmail":   "smtp.gmail.com",
		"hotmail": "smtp.live.com",
		"yahoo":   "smtp.mail.yahoo.com",
	}

	/*Msj : mensajes de error en las librerias*/
	Msj = StMsj{
		Store: map[string]string{
			"GE01": "Mensaje no encontrado.",
			//libreria encrip
			"EN01": "Error en generar phisher.",
			"EN02": "Error en generar bloque.",
			"EN03": "Error en leer el bloque.",
			"EN04": "Error en abrir el bloque.",
			//libreria convert
			"CO01": "Error en convertir float.",
			"CO02": "Error en convertir int64.",
			"CO03": "Error en convertir int32.",
			//libreria archivo,
			"AR01": "No existe el archivo.",
			"AR02": "Error al crear directorio.",
			"AR03": "Error al crear archivo.",
			"AR04": "Error al crear log.",
			"AR05": "No existe el directorio.",
			"AR06": "Problemas en leer directorio.",
			"AR07": "Error al eliminar un archivo.",
			//libreria database
			"CN01": "Error al convertir a json.",
			"CN02": "Error al convertir a string.",
			"CN03": "Error al convertir a string.",
			"CN04": "No tiene el insert.",
			"CN05": "No tiene los prefijos validos.",
			"CN06": "No se obtuvieron las columnas.",
			"CN07": "No se obtuvieron las filas .",
			"CN08": "Error al leer archivo config json.",
			"CN09": "No existe el arhivo config json.",
			"CN10": "No existe el arhivo config ini.",
			"CN11": "Error al leer archivo config ini.",
			"CN12": "El archivo config no es valido.",
			"CN13": "Tipo de DB no compatible.",
			"CN14": "Error al conectar ala DB.",
			"CN15": "Cantidad de fila es cero.",
			"CN16": "Error al cerrar ala DB.",
			"CN17": "Error en commit.",
			"CN18": "Error en al ejecutar querie.",
			"CN19": "Error en al obtener columna.",
			"CN20": "El db es sqllite necesita el archivo.db",
			//libreria json
			"JS01": "Error al generar un nuevo json.",
			"JS02": "Error al convertir un json.",
			"JS03": "Json null.",
			"JS04": "Error al capturar json.",
			//Libreria de Email
			"ER01": "Error en enviar el mensaje",
			"ER02": "No pasa las validaciones",
			/*libreria server */
			"SR01": "Tipo  de peticion incorrecta.",
			"SR02": "El Path es invalido.",
			"SR03": "No contiene una funcion de peticion.",
		},
	}
	/*EXT : extensiones de archivos */
	EXT = map[string]string{
		"JSON": ".json",
		"INI":  ".ini",
		"XML":  ".xml",
		"RUT":  ".rut",
		"SQL":  ".sql",
		"DB":   ".bd",
	}
)

const (
	/*colores disponibles */

	/*Green : verde */
	Green Pc = "g"
	/*Red : rojo */
	Red Pc = "r"
	/*Blue : azul */
	Blue Pc = "b"
	/*Cyan : celeste */
	Cyan Pc = "c"
	/*White : blanco */
	White Pc = "w"
	/*Black : negro */
	Black Pc = "bl"
	/*Yellow : amarillo*/
	Yellow Pc = "y"
	/*Magenta : magenta */
	Magenta Pc = "m"
	/*HBlack : negro fuerte */
	HBlack Pc = "hbl"
	/*HRed : rojo fuerte */
	HRed Pc = "hr"
	/*HGreen : verde fuerte */
	HGreen Pc = "hg"
	/*HYellow : amarrillo fuerte */
	HYellow Pc = "hy"
	/*HBlue : azul fuerte */
	HBlue Pc = "hb"
	/*HMagenta : magenta fuerte*/
	HMagenta Pc = "hm"
	/*HCyan : celeste fuerte */
	HCyan Pc = "hc"
	/*HWhite : blanco fuerte */
	HWhite Pc = "hw"
	/*FORMFE : Formato de fecha para los archivo YYYYMMDD*/
	FORMFE = "%d%02d%02d"
	/*FormatFechaPostgresql : fomato de fecha de la base  de datos de postgresql*/
	FormatFechaPostgresql = "2006-01-02"
)
