// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package analyzers

import (
	"istio.io/istio/pkg/config/analysis"
	"istio.io/istio/pkg/config/analysis/analyzers/annotations"
	"istio.io/istio/pkg/config/analysis/analyzers/authz"
	"istio.io/istio/pkg/config/analysis/analyzers/deployment"
	"istio.io/istio/pkg/config/analysis/analyzers/deprecation"
	"istio.io/istio/pkg/config/analysis/analyzers/destinationrule"
	"istio.io/istio/pkg/config/analysis/analyzers/envoyfilter"
	"istio.io/istio/pkg/config/analysis/analyzers/externalcontrolplane"
	"istio.io/istio/pkg/config/analysis/analyzers/gateway"
	"istio.io/istio/pkg/config/analysis/analyzers/injection"
	"istio.io/istio/pkg/config/analysis/analyzers/k8sgateway"
	"istio.io/istio/pkg/config/analysis/analyzers/multicluster"
	"istio.io/istio/pkg/config/analysis/analyzers/schema"
	"istio.io/istio/pkg/config/analysis/analyzers/service"
	"istio.io/istio/pkg/config/analysis/analyzers/serviceentry"
	"istio.io/istio/pkg/config/analysis/analyzers/sidecar"
	"istio.io/istio/pkg/config/analysis/analyzers/telemetry"
	"istio.io/istio/pkg/config/analysis/analyzers/virtualservice"
	"istio.io/istio/pkg/config/analysis/analyzers/webhook"
)

// All returns all analyzers
func All() []analysis.Analyzer {
	analyzers := []analysis.Analyzer{
		// Please keep this list sorted alphabetically by pkg.name for convenience
		&annotations.K8sAnalyzer{},
		&authz.AuthorizationPoliciesAnalyzer{},
		&deployment.ServiceAssociationAnalyzer{},
		&deployment.ApplicationUIDAnalyzer{},
		&deprecation.FieldAnalyzer{},
		&externalcontrolplane.ExternalControlPlaneAnalyzer{},
		&gateway.IngressGatewayPortAnalyzer{},
		&gateway.CertificateAnalyzer{},
		&gateway.SecretAnalyzer{},
		&gateway.ConflictingGatewayAnalyzer{},
		&injection.Analyzer{},
		&injection.ImageAnalyzer{},
		&injection.ImageAutoAnalyzer{},
		&k8sgateway.SelectorAnalyzer{},
		&multicluster.MeshNetworksAnalyzer{},
		&service.PortNameAnalyzer{},
		&sidecar.SelectorAnalyzer{},
		&virtualservice.ConflictingMeshGatewayHostsAnalyzer{},
		&virtualservice.DestinationHostAnalyzer{},
		&virtualservice.DestinationRuleAnalyzer{},
		&virtualservice.GatewayAnalyzer{},
		&virtualservice.JWTClaimRouteAnalyzer{},
		&virtualservice.RegexAnalyzer{},
		&destinationrule.CaCertificateAnalyzer{},
		&serviceentry.ProtocolAddressesAnalyzer{},
		&webhook.Analyzer{},
		&envoyfilter.EnvoyPatchAnalyzer{},
		&telemetry.ProdiverAnalyzer{},
		&telemetry.SelectorAnalyzer{},
		&telemetry.DefaultSelectorAnalyzer{},
		&telemetry.LightstepAnalyzer{},
	}

	analyzers = append(analyzers, schema.AllValidationAnalyzers()...)

	return analyzers
}

// AllCombined returns all analyzers combined as one
func AllCombined() analysis.CombinedAnalyzer {
	return analysis.Combine("all", All()...)
}
