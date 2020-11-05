package utility

import "time"

type (
	/*Worker : Orquestador de procesos paralelo*/
	Worker struct {
		start   bool
		finalid bool
		fexec   time.Time
		ticker  *time.Ticker
		hr      HrTime
		job     Job
		valid   chan bool
		err     chan error
	}
	/*Job : alias para un proceso*/
	Job func(chan bool, chan error)
)

/*Finally : Finalizcion del proceso donde el indicativo reset es para reintentar secuencia*/
func (p *Worker) Finally() {
	p.start = false
	close(p.err)
	close(p.valid)
}

/*Start : Ejecuta el proceso paralelo*/
func (p *Worker) Start() {
	if !p.start {
		p.start = true
		p.valid = make(chan bool)
		p.err = make(chan error)
		go p.job(p.valid, p.err)
	}
}

/*StartTime : ejecuta un proceso paralelo mediante una hora especifica*/
func (p *Worker) StartTime() {
	if !DateEquals(p.fexec, time.Now()) {
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

/*NewWorkTicker : crea una tarea mediante un ticker para ponerlo en un servicio*/
func NewWorkTicker(p *time.Ticker, job Job) Worker {
	return Worker{
		finalid: false,
		job:     job,
		ticker:  p,
		start:   false,
	}
}

/*NewWorkTime : crea una tarea mediante una hora exacta para ponerlo en un servicio*/
func NewWorkTime(hr, min int, job Job) (Worker, error) {
	hrTime, err := NewHrTime(hr, min)
	if err != nil {
		return Worker{}, err
	}
	return Worker{
		finalid: false,
		job:     job,
		hr:      hrTime,
		start:   false,
	}, nil
}
