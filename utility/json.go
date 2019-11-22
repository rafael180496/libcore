package utility

import (
	"bytes"
	"encoding/json"
)

type (

	/*JSON : Contiene datos JSON de muchos objetos.*/
	JSON []byte
)

/*NewJSON : Crea un JSON con cualquier data*/
func NewJSON(v interface{}) (JSON, error) {
	d, err := json.Marshal(v)
	if err != nil {
		return d, SendErrorCod("JS01")
	}
	return d, nil
}

/*ParseJSON : Captura el JSON con cualquier data*/
func ParseJSON(d JSON, v interface{}) error {
	err := json.Unmarshal(d, v)
	if err != nil {
		return SendErrorCod("JS02")
	}
	return nil
}

/*JSONtoObj : convierte objetos JSON en map.*/
func JSONtoObj(d JSON) ([]map[string]interface{}, error) {
	var objs []map[string]interface{}
	if bytes.Equal(d, []byte("null")) {
		return objs, SendErrorCod("JS03")
	}
	var v interface{}
	err := ParseJSON(d, &v)
	if err != nil {
		return nil, err
	}
	switch vv := v.(type) {
	case []interface{}:
		for _, o := range vv {
			objs = append(objs, o.(map[string]interface{}))
		}
	case map[string]interface{}:
		objs = []map[string]interface{}{vv}
	case []map[string]interface{}:
		objs = vv
	default:
		return nil, SendErrorCod("JS04")
	}

	return objs, nil
}

/*ObjtoJSON : convierte maps en objetos JSON .*/
func ObjtoJSON(Encabezado []string, filas [][]interface{}) (JSON, error) {
	var b bytes.Buffer
	b.Write([]byte("["))
	for i, fila := range filas {
		if i > 0 {
			b.Write([]byte(","))
		}
		b.Write([]byte("{"))
		for j, v := range fila {
			if j > 0 {
				b.Write([]byte(","))
			}
			d, err := NewJSON(v)
			if err != nil {
				return nil, err
			}
			EncabezadoStr := "null"
			if len(Encabezado) > 0 && len(Encabezado) > j {
				EncabezadoStr = Encabezado[j]
			}
			b.Write([]byte(`"` + EncabezadoStr + `":` + string(d)))
		}
		b.Write([]byte("}"))
	}
	b.Write([]byte("]"))
	return JSON(b.Bytes()), nil
}
