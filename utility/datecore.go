package utility

import (
	"encoding/xml"
	"strings"
	"time"
)

type (
	/*Date : formato Time especial con parseo los xml y json de manera que el formato puede variar*/
	Date struct {
		time.Time
	}
	/*HrTime : valida las horas en un time*/
	HrTime struct {
		hr, min int
	}
)

/*Valid : valida la estructura*/
func (p *HrTime) Valid() error {
	if p.hr > 24 || p.hr < 0 {
		return StrErr("La hora es invalida colocar en formato de 24HR")
	}
	if p.min > 59 || p.min < 0 {
		return StrErr("los Minutos es invalida colocar en formato de 59Min")
	}
	return nil
}

/*SetHr : setea un tipo time */
func (p *HrTime) SetHr(hr, min int) error {
	val := HrTime{
		hr:  hr,
		min: min,
	}
	if err := val.Valid(); err != nil {
		return err
	}
	*p = val
	return nil
}

/*GetHr : extrae un tipo time */
func (p *HrTime) GetHr() (hr, min int) {
	return p.hr, p.min
}

/*SetHrTime : setea un tipo time */
func (p *HrTime) SetHrTime(hr, min int) error {
	val := HrTime{
		hr:  hr,
		min: min,
	}
	if err := val.Valid(); err != nil {
		return err
	}
	*p = val
	return nil
}

/*NewHrTime : crea un time */
func NewHrTime(hr, min int) (HrTime, error) {
	val := HrTime{
		hr:  hr,
		min: min,
	}
	if err := val.Valid(); err != nil {
		return HrTime{}, err
	}
	return val, nil
}

/*NewHrTimeDate : crea un time  con un tipo date*/
func NewHrTimeDate(data Date) (HrTime, error) {
	val := HrTime{
		hr:  data.Hour(),
		min: data.Minute(),
	}
	if err := val.Valid(); err != nil {
		return HrTime{}, err
	}
	return val, nil
}

/*EqualNow : Valida si un hrtime es identico a un date de la hora actual*/
func (p *HrTime) EqualNow() bool {
	date := time.Now()
	return ReturnIf(date.Hour() == p.hr && date.Minute() == p.min, true, false).(bool)
}

/*EqualDate : Valida si un hrtime es identico a un date*/
func (p *HrTime) EqualDate(date Date) bool {
	return ReturnIf(date.Hour() == p.hr && date.Minute() == p.min, true, false).(bool)
}

/*HrTime : crea una instancia de hrtime*/
func (p *Date) HrTime() HrTime {
	return HrTime{
		hr:  p.Hour(),
		min: p.Minute(),
	}
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
