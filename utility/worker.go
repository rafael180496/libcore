package utility

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type (

	/*FormatWorker : contiene los formatos de los worker para poder obtenerlos de forma de un json*/
	FormatWorker []struct {
		Tp  string `json:"tp"`
		Key string `json:"key"`
		Fr  string `json:"fr"`
	}
	/*MasterWorker : contiene varios worker y se administran de forma individual con maps*/
	MasterWorker struct {
		Jobs   map[string]Job
		Config FormatWorker
	}
	/*Worker : Orquestador de procesos paralelo*/
	Worker struct {
		start     bool
		finalid   bool
		fexec     time.Time
		ticker    *time.Ticker
		hr        HrTime
		job       Job
		valid     chan bool
		err       chan error
		activo    bool
		indticker bool
	}
	/*Job : alias para un proceso*/
	Job func(chan bool, chan error)
)

func (p *MasterWorker) ReloadWork(key string) (*Worker, error) {
	var TaskMap Worker
	var err error
	indresp := false
	Block{
		Try: func() {
			for _, v := range p.Config {
				if v.Key == key {
					work, err := NewWork(v.Fr, v.Tp, p.Jobs[v.Key])
					if err != nil {
						break
					}
					TaskMap = work
					indresp = true
					break
				}
			}
		},
		Catch: func(e Exception) {
			err = fmt.Errorf("error en procesar el workers")
		},
	}.Do()
	if !indresp {
		return nil, fmt.Errorf("el work no existe")
	}
	return &TaskMap, err
}

/*LoadWorkers : carga todos los worker y lo envie en un map*/
func (p *MasterWorker) LoadWorkers() (map[string]*Worker, error) {
	TaskMap := make(map[string]*Worker)
	var err error
	Block{
		Try: func() {
			for _, v := range p.Config {
				work, err := NewWork(v.Fr, v.Tp, p.Jobs[v.Key])
				if err != nil {
					break
				}
				TaskMap[v.Key] = &work
			}
		},
		Catch: func(e Exception) {
			err = fmt.Errorf("error en procesar el map workers")
		},
	}.Do()
	return TaskMap, err
}

/*LoadConfig : carga el archivo de configuracion de los worker como archivo json*/
func (p *MasterWorker) LoadConfig(PathJSON string) error {
	var (
		err     error
		config  FormatWorker
		ptrArch *os.File
	)
	if !FileExt(PathJSON, "JSON") {
		return Msj.GetError("CN09")
	}
	PathJSON, err = TrimFile(PathJSON)
	if err != nil {
		return Msj.GetError("CN08")
	}
	ptrArch, err = os.Open(PathJSON)
	if err != nil {
		return Msj.GetError("CN08")
	}
	defer ptrArch.Close()
	decJSON := json.NewDecoder(ptrArch)
	err = decJSON.Decode(&config)
	if err != nil {
		return Msj.GetError("CN08")
	}
	return nil
}

/*Finally : Finalizacion del proceso donde el indicativo reset es para reintentar secuencia*/
func (p *Worker) Finally() {
	p.start = false
}

/*SetActivo : modifica para que se desactive la tarea*/
func (p *Worker) SetActivo(act bool) {
	p.activo = act
}

/*StartGen : ejecuta una tarea en paralelo en forma general*/
func (p *Worker) StartGen() {
	if p.indticker {
		p.Start()
	} else {
		p.StartTime()
	}
}

/*Start : Ejecuta el proceso paralelo*/
func (p *Worker) Start() {
	if !p.start && p.activo {
		p.start = true
		p.valid = make(chan bool)
		p.err = make(chan error)
		go p.job(p.valid, p.err)
	}
}

/*StartTime : ejecuta un proceso paralelo mediante una hora especifica*/
func (p *Worker) StartTime() {
	if !DateEquals(p.fexec, time.Now()) && p.finalid {
		p.finalid = false
	}
	if p.hr.EqualNow() && !p.finalid {
		p.fexec = time.Now()
		p.finalid = true
		p.Start()
	}
}

/*Tick : regresa el secuencial de tiempo*/
func (p *Worker) Tick() <-chan time.Time {
	return p.ticker.C
}

/*Valid : envia si el proceso termino*/
func (p *Worker) Valid() <-chan bool {
	return p.valid
}

/*Err :  envia el error en un canal de ejecucion*/
func (p *Worker) Err() <-chan error {
	return p.err
}

/*NewWork : crea una tarea con formatos y tipo de de secuencia
1 - Ticker
2 - Time
Formatos:
05 hr
54 min
12:16 time
*/
func NewWork(format, tp string, job Job) (Worker, error) {
	switch tp {
	case "hr", "min":
		tiker, err := NewTicker(format, tp)
		if err != nil {
			return Worker{}, err
		}
		wrol := NewWorkTicker(tiker, job)
		return wrol, nil
	case "time":
		hrtime, err := NewHrTimeStr(format)
		if err != nil {
			return Worker{}, err
		}
		wrol, err := NewWorkTime(hrtime.hr, hrtime.min, job)
		if err != nil {
			return Worker{}, err
		}
		return wrol, nil
	default:
		return Worker{}, fmt.Errorf("error en tipo")
	}
}

/*NewWorkTicker : crea una tarea mediante un ticker para ponerlo en un servicio*/
func NewWorkTicker(p *time.Ticker, job Job) Worker {
	return Worker{
		finalid:   false,
		job:       job,
		ticker:    p,
		start:     false,
		activo:    true,
		indticker: true,
	}
}

/*NewWorkTime : crea una tarea mediante una hora exacta para ponerlo en un servicio*/
func NewWorkTime(hr, min int, job Job) (Worker, error) {
	hrTime, err := NewHrTime(hr, min)
	if err != nil {
		return Worker{}, err
	}
	return Worker{
		finalid:   false,
		job:       job,
		hr:        hrTime,
		start:     false,
		activo:    true,
		indticker: false,
	}, nil
}
