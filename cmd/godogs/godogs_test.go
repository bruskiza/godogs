package main

import (
  "context"
  "errors"
  "fmt"
  "testing"

  "github.com/cucumber/godog"
)


func thereAreGodogs(ctx context.Context, available int) (context.Context, error) {
  return context.WithValue(ctx, "dawgs", available), nil
}

func iEat(ctx context.Context, num int) (context.Context, error) {
  available, ok := ctx.Value("dawgs").(int)
  if !ok {
    return ctx, errors.New("there are no godogs available")
  }

  if available < num {
    return ctx, fmt.Errorf("you cannot eat %d godogs, there are %d available", num, available)
  }

  available -= num

  return context.WithValue(ctx, "dawgs", available), nil
}

func thereShouldBeRemaining(ctx context.Context, remaining int) error {
  available, ok := ctx.Value("dawgs").(int)
  if !ok {
    return errors.New("there are no godogs available")
  }

  if available != remaining {
    return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, available)
  }

  return nil
}

func TestFeatures(t *testing.T) {
  suite := godog.TestSuite{
    ScenarioInitializer: InitializeScenario,
    Options: &godog.Options{
      Format:   "pretty",
      Paths:    []string{"features"},
      TestingT: t, // Testing instance that will run subtests.
    },
  }

  if suite.Run() != 0 {
    t.Fatal("non-zero status returned, failed to run feature tests")
  }
}

func InitializeScenario(sc *godog.ScenarioContext) {
  sc.Step(`^there are (\d+) godogs$`, thereAreGodogs)
  sc.Step(`^I eat (\d+)$`, iEat)
  sc.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
}

