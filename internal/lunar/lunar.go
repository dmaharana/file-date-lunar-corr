package lunar

import (
	"math"
	"time"

	"github.com/dyindude/moonphase"
)

// Phase represents one of the 8 lunar phases
type Phase string

const (
	New             Phase = "New"
	WaxingCrescent  Phase = "Waxing Crescent"
	FirstQuarter    Phase = "First Quarter"
	WaxingGibbous   Phase = "Waxing Gibbous"
	Full            Phase = "Full"
	WaningGibbous   Phase = "Waning Gibbous"
	LastQuarter     Phase = "Last Quarter"
	WaningCrescent  Phase = "Waning Crescent"
)

var PhaseNames = []Phase{
	New,
	WaxingCrescent,
	FirstQuarter,
	WaxingGibbous,
	Full,
	WaningGibbous,
	LastQuarter,
	WaningCrescent,
}

// GetPhase returns the lunar phase name for a given date
func GetPhase(dt time.Time) Phase {
	m := moonphase.New(dt)
	p := m.Phase() // 0.0 - 1.0 (0=New, 0.5=Full)

	// Shift by half a bin to center the phases
	// New Moon is centered around 0.0
	// 8 bins: each bin is 1/8 = 0.125 wide
	// Half a bin is 0.0625
	
	idx := int(math.Floor((p + 0.0625) * 8)) % 8
	return PhaseNames[idx]
}

// GetPhaseFloat returns the raw phase value (0.0 - 1.0)
func GetPhaseFloat(dt time.Time) float64 {
	m := moonphase.New(dt)
	return m.Phase()
}
