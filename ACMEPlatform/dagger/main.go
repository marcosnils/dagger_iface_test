package main

import (
	"context"
	"dagger/acme-platform/internal/dagger"
	"fmt"
)

type App interface {
	DaggerObject
	Build() *dagger.Container
	Test(context.Context) (string, error)
	Healthcheck(ctx context.Context, endpoint string) error
	Secrets(ctx context.Context, env string) ([]*dagger.Secret, error)
}

type Acmeplatform struct{}

func (m *Acmeplatform) Deploy(ctx context.Context, app App) error {
	_, err := app.Test(ctx)
	if err != nil {
		return err
	}

	_, _ = app.Secrets(ctx, "prod")

	// do something with secrets

	app.Build().Sync(ctx)

	// deploy container to $TARGET_ENV

	err = app.Healthcheck(ctx, "prod")
	if err != nil {
		return err
	}

	return nil
}

func (m *Acmeplatform) Scan(ctx context.Context, c *dagger.Container) error {
	fmt.Println("Scanning container for vulnerabilities")
	return nil
}
