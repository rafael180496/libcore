package database

import (
	"fmt"
	"strings"

	utl "github.com/rafael180496/libcore/utility"
)

type (
	/*DataTable : maneja un crud completo y genera script automaticos*/
	DataTable struct {
		table string
		rows  []StData
		index []string
	}
)

/*ValidRow : valida si tiene filas llenadas*/
func (p *DataTable) ValidRow() bool {
	return utl.ReturnIf(len(p.rows) > 0, true, false).(bool)
}

/*GetCols : Obtiene el nombre de las columnas de cada fila*/
func (p *DataTable) GetCols() ([]string, error) {
	var cols []string
	item, err := p.GetRow(1)
	if err != nil {
		return cols, err
	}
	cols = item.KeyColum()
	if len(cols) <= 0 {
		return cols, utl.Msj.GetError("DT04")
	}
	return cols, nil
}

/*GetIndex : Obtiene el nombre de los indices*/
func (p *DataTable) GetIndex() []string {
	return p.index
}

/*GetTable : Obtiene el nombre de la tabla*/
func (p *DataTable) GetTable() string {
	return p.table
}

/*SetTable : Modifica el nombre de la tabla*/
func (p *DataTable) SetTable(table string) {
	p.table = strings.ToUpper(utl.Trim(table))
}

/*LenRows : Obtienen la cantidad de fila*/
func (p *DataTable) LenRows() int {
	return len(p.rows)
}

/*LenIndex : Obtienen la cantidad de llaves primarias para update o delete*/
func (p *DataTable) LenIndex() int {
	return len(p.index)
}

/*AddIndexs : agrega varias claves primarias*/
func (p *DataTable) AddIndexs(cols ...string) error {
	for _, col := range cols {
		err := p.AddIndex(col)
		if err != nil {
			return err
		}
	}
	return nil
}

/*AddRows : Agrega varias fila */
func (p *DataTable) AddRows(rows ...StData) {
	for _, item := range rows {
		p.AddRow(item)
	}
}

/*AddRow : Agrega una fila */
func (p *DataTable) AddRow(row StData) {
	p.rows = append(p.rows, upperMap(row))
}

/*AddIndex : agrega una llave para los delete o update*/
func (p *DataTable) AddIndex(col string) error {
	item, err := p.GetRow(1)
	if err != nil {
		return err
	}
	if !item.ValidColum(col) {
		return utl.Msj.GetError("DT02")
	}
	p.index = append(p.index, col)
	return nil
}

/*GetRow : Obtiene una fila X por el numero de fila*/
func (p *DataTable) GetRow(row int) (StData, error) {
	if p.LenRows() < row {
		return StData{}, utl.Msj.GetError("DT01")
	}
	return p.rows[(row - 1)], nil
}

/*GetRows : Obtiene todas las fila*/
func (p *DataTable) GetRows(row int) []StData {
	return p.rows
}

/*GenInserts : genera insert masivos para modificaciones de base de datos*/
func (p *DataTable) GenInserts() ([]StQuery, error) {
	var queries []StQuery
	clone := *p
	sqltemp, err := sqldinamic(clone, "i")
	if err != nil {
		return queries, err
	}
	fmt.Println(sqltemp)
	return queries, nil
}

/*sqldinamic : genera los sql temporales para los crud*/
func sqldinamic(data DataTable, acc string) (string, error) {
	table := utl.Trim(data.GetTable())
	if table == "" {
		return "", utl.Msj.GetError("DT03")
	}
	cols, err := data.GetCols()
	if err != nil {
		return "", err
	}
	if !validDuplid(cols) {
		return "", utl.Msj.GetError("DT05")
	}
	switch acc {
	case "i":
		sqltmp := sqlinsert(table, cols)
		return sqltmp, nil
	default:
		return "", nil

	}
}
func upperMap(data StData) StData {
	var datanew StData
	for k := range data {
		kNew := strings.ToUpper(k)
		datanew[kNew] = data[k]
	}
	return datanew
}
func sqlinsert(table string, cols []string) string {
	sqltmp := fmt.Sprintf("INSERT INTO %s (", table)
	for i, item := range cols {
		if i == (len(cols) - 1) {
			sqltmp = fmt.Sprintf("%s%s) VALUES(", sqltmp, item)
		} else {
			sqltmp = fmt.Sprintf("%s%s,", sqltmp, item)
		}
	}
	for i, item := range cols {
		if i == (len(cols) - 1) {
			sqltmp = fmt.Sprintf("%s:%s)", sqltmp, item)
		} else {
			sqltmp = fmt.Sprintf("%s:%s,", sqltmp, item)
		}
	}
	return sqltmp
}

/*validDuplid : valida si las columnas estan duplicadas*/
func validDuplid(cols []string) bool {
	for i, col := range cols {
		for j, colaux := range cols {
			if i != j && colaux == col {
				return false
			}
		}
	}
	return true
}
