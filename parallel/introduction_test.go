package parallel_test

import (
	"testing"

	"github.com/ponta0929/goplactice/parallel"
)

func TestDataConflict(t *testing.T) {
	parallel.DataConflict()
}

func TestDataConflictBadSol(t *testing.T) {
	parallel.DataConflictBadSol()
}

func TestDataConflictSol(t *testing.T) {
	parallel.DataConflictSol()
}

func TestDataConflictSolSync(t *testing.T) {
	parallel.DataConflictSolSync()
}

func TestDeadRock(t *testing.T) {
	//実行するとエラー
	if false {
		parallel.DeadRock()
	}
}

func TestLiveRock(t *testing.T) {
	parallel.LiveRock()
}

func TestResourceDepletion(t *testing.T) {
	parallel.ResourceDepletion()
}
