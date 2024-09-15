package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cucumber/godog"
)


func iSendRequestTo(ctx context.Context, arg1, arg2 string) (context.Context, error) {
	return context.WithValue(ctx, "iSendRequestTo", fmt.Sprintf("arg1:%s,arg2:%s", arg1, arg2)), nil
}

func theResponseCodeShouldBe(ctx context.Context, arg1 int) (context.Context, error) {
	return context.WithValue(ctx, "theResponseCodeShouldBe", fmt.Sprintf("arg1:%d", arg1)), nil
}

func theResponseShouldMatchJson(ctx context.Context, arg1 *godog.DocString) error {
	log.Printf("ctx: %+v", ctx)
	log.Print(ctx.Value("iSendRequestTo"))
	return godog.ErrSkip
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)"$`, iSendRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, theResponseShouldMatchJson)
}

