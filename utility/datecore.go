package utility

import (
	"strings"
	"time"
)

/*DatePostgreSQL : formato Time especial con parseo para base de datos postgresql con manipulaciones en json*/
type DatePostgreSQL struct {
	time.Time
}

/*UnmarshalJSON : fomateo especial a json para los tipo Time en golang*/
func (sd *DatePostgreSQL) UnmarshalJSON(input []byte) error {
	newTime, err := time.Parse(FormatFechaPostgresql, strings.Trim(string(input), `"`))
	if err != nil {
		return err
	}
	sd.Time = newTime
	return nil
}
