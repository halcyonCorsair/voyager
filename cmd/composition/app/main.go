package app

import (
	"context"
	"flag"
	"math/rand"
	"os"
	"time"

	"github.com/atlassian/ctrl"
	ctrlApp "github.com/atlassian/ctrl/app"
	"github.com/atlassian/voyager/cmd"
	"github.com/atlassian/voyager/pkg/composition"
	"github.com/atlassian/voyager/pkg/util/crash"
)

func Main() {
	rand.Seed(time.Now().UnixNano())
	cmd.RunInterruptably(runWithContext)
}

func runWithContext(ctx context.Context) error {
	crash.InstallAPIMachineryLoggers()
	controllers := []ctrl.Constructor{
		&composition.ControllerConstructor{},
	}
	a, err := ctrlApp.NewFromFlags("composition", controllers, flag.CommandLine, os.Args[1:])
	if err != nil {
		return err
	}
	return a.Run(ctx)
}
