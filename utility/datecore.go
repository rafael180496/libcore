package utility

import (
	"encoding/xml"
	"strings"
	"time"
)

/*Date : formato Time especial con parseo los xml y json de manera que el formato puede variar*/
type Date struct {
	time.Time
}

/*Equals : Valida dos date a -> b

1 -> mayor a

0 -> igual b

-1 -> menor a
*/
func (p *Date) Equals(data Date) int {
	switch {
	case p.Year() == data.Year() && p.Month() == data.Month() && p.Day() == data.Day():
		return 0
	case p.Year() > data.Year(), p.Year() == data.Year() && p.Month() > data.Month(), p.Year() == data.Year() && p.Month() == data.Month() && p.Day() > data.Day():
		return 1
	default:
		return -1
	}
}

/*UnmarshalJSON : fomateo especial a json para los tipo Time en golang*/
func (p *Date) UnmarshalJSON(input []byte) error {
	newTime, err := StringToDate(strings.Trim(string(input), `"`))
	if err != nil {
		return err
	}
	p.Time = newTime
	return nil
}

/*UnmarshalXML : fomateo especial a xml para los tipo Time en golang*/
func (p *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := StringToDate(v)
	if err != nil {
		return err
	}
	*p = Date{parse}
	return nil
}
