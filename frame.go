package insensate

import (
	"github.com/fractalcat/emogo"
	_ "code.google.com/p/goprotobuf/proto"
)

func NewEpocSensor(label string, value int, quality int) *EpocSensor {
	s := new(EpocSensor)
	l, v, q := label, int32(value), int32(quality)
	s.Label = &l
	s.Value = &v
	s.Quality = &q
	return s
}

func (e *EpocFrame) addSensor(label string, value int, quality int) {
	s := NewEpocSensor(label, value, quality)
	e.Sensors = append(e.Sensors, s)
}

func NewEpocFrame(e *emogo.EmokitFrame) *EpocFrame {
	f := new(EpocFrame)
	if e.BatteryFrame() {
		bat := uint32(e.Battery())
		f.Battery = &bat
	}
	counter := uint32(e.Counter())
	f.Counter = &counter
	x, y := e.Gyro()
	x_, y_ := int32(x), int32(y)
	f.GyroX = &x_
	f.GyroY = &y_
	f.Sensors = make([]*EpocSensor, 0)
	f.addSensor("F3", e.F3.Value, e.F3.Quality)
	f.addSensor("FC6", e.FC6.Value, e.FC6.Quality)
	f.addSensor("P7", e.P7.Value, e.P7.Quality)
	f.addSensor("T8", e.T8.Value, e.T8.Quality)
	f.addSensor("F7", e.F7.Value, e.F7.Quality)
	f.addSensor("F8", e.F8.Value, e.F8.Quality)
	f.addSensor("T7", e.T7.Value, e.T7.Quality)
	f.addSensor("P8", e.P8.Value, e.P8.Quality)
	f.addSensor("AF4", e.AF4.Value, e.AF4.Quality)
	f.addSensor("F4", e.F4.Value, e.F4.Quality)
	f.addSensor("AF3", e.AF3.Value, e.AF3.Quality)
	f.addSensor("O2", e.O2.Value, e.O2.Quality)
	f.addSensor("O1", e.O1.Value, e.O1.Quality)
	f.addSensor("FC5", e.FC5.Value, e.FC5.Quality)
	return f
}
