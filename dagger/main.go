package main

import (
	"context"
	"dagger/acme-ochestrator/internal/dagger"
	"fmt"
)

type Acmeochestrator struct{}

var apps = map[string]*dagger.AcmeplatformApp{
	"acmeapp": dag.Acmeapp().AsAcmeplatformApp(),
}

func (m *Acmeochestrator) DeployApp(ctx context.Context, name string) (string, error) {
	a, ok := apps[name]

	if !ok {
		return "", fmt.Errorf("app %s not found", name)
	}

	err := dag.Acmeplatform().Deploy(ctx, a)
	return "", err
}
