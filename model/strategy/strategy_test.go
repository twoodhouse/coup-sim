package strategy
import (
	// "fmt"
	"testing"
	"github.com/twoodhouse/coup-sim/model/strategies/noLieStrategy"
)

func TestStrategyInterface(t *testing.T) {
  noLieStrategy := noLieStrategy.New()
	noLieStrategy.GetAction()
	printStrategyName(&noLieStrategy)
}

func printStrategyName(strategy Interface) {
  strategy.GetAction()
}
