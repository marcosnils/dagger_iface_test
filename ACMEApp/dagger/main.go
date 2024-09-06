package main

import (
	"context"
	"dagger/acme-app/internal/dagger"
	"fmt"
)

type Acmeapp struct{}

func (m *Acmeapp) Build(ctx context.Context) *dagger.Container {
	c := dag.Container()
	dag.Acmeplatform().Scan(ctx, c)
	return c
}

func (m *Acmeapp) Test(ctx context.Context) (string, error) {
	return "", nil
}

func (m *Acmeapp) Secrets(ctx context.Context, env string) ([]*dagger.Secret, error) {
	return []*dagger.Secret{
		dag.SetSecret("db_password", "verysecure"),
		dag.SetSecret("api_key", "moresecre"),
	}, nil
}

func (m *Acmeapp) Healthcheck(ctx context.Context, endpoint string) error {
	return fmt.Errorf("healthcheck failed")
}
