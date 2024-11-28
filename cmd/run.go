package cmd

import (
	"github.com/JohnGrimm/Stress-Test-fg/internal/stresstest"
)

func Execute(config stresstest.Config) {
	stresstest.ExecutarTestes(config)
}
