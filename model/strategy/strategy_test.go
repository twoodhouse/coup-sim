package strategy
import (
	"testing"
	"github.com/twoodhouse/coup-sim/model/strategies/noLieStrategy"
	// "github.com/twoodhouse/coup-sim/model/log"
)

func TestStrategyInterface(t *testing.T) {
  noLieStrategy := noLieStrategy.New()
	testInterface(&noLieStrategy)
}

func testInterface(strategy Interface) {
	// log := log.New()
	// strategy.GetAction(log)
}
