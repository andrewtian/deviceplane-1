package engine

import (
	"context"
	"errors"

	"github.com/deviceplane/deviceplane/pkg/spec"
)

var (
	ErrInstanceNotFound = errors.New("instance not found")
)

type Engine interface {
	CreateContainer(context.Context, string, spec.Service) (string, error)
	StartContainer(context.Context, string) error
	ListContainers(context.Context, map[string]struct{}, map[string]string, bool) ([]Instance, error)
	StopContainer(context.Context, string) error
	RemoveContainer(context.Context, string) error

	PullImage(context.Context, string) error
}

type Instance struct {
	ID      string
	Labels  map[string]string
	Running bool
}
