package scene

import "fmt"

// SuspendScene does nothing
type SuspendScene struct {
}

func (s *SuspendScene) Initialize() {
}

func (s *SuspendScene) Destroy() {
	fmt.Printf("TEMP: before suspend destroy\n")
	ReportMemoryUsage()
}

func (s *SuspendScene) Drive() {
}
