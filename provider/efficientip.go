/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"
	// "crypto/tls"
	// "fmt"
	// "net/http"
	// "os"
	// "strings"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

const (
)

type EIPConfig struct {
	DomainFilter DomainFilter
	ZoneIDFilter ZoneIDFilter
	DryRun       bool
}

type EIPProvider struct {
	domainFilter DomainFilter
	zoneIDFilter ZoneIDFilter
	dryRun       bool
}

func NewEIPProvider(config EIPConfig) (*EIPProvider, error) {
	log.Debug("eip provider")

	provider := &EIPProvider{
		domainFilter: config.DomainFilter,
		zoneIDFilter: config.ZoneIDFilter,
		dryRun: config.DryRun,
	}

	return provider, nil
}

func (p *EIPProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	log.Infof("eip applychanges")

	log.Debugf(" changes %v", changes.Create)
	log.Debugf(" update %v", changes.UpdateNew)
	log.Debugf(" delete %v", changes.Delete)

	if len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete) == 0 {
		return nil
	}

	return nil
}

func (p *EIPProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	log.Infof("eip records")

	var endpoints []*endpoint.Endpoint

	return endpoints, nil
}
