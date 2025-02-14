package resourcecreator

const (
	NaisAppNameEnv   = "NAIS_APP_NAME"
	NaisNamespaceEnv = "NAIS_NAMESPACE"
	NaisAppImageEnv  = "NAIS_APP_IMAGE"

	IstioNetworkingAPIVersion               = "networking.istio.io/v1alpha3"         // API version of the Networking resources
	IstioRBACAPIVersion                     = "rbac.istio.io/v1alpha1"               // API version of the RBAC resources
	IstioIngressGatewayLabelValue           = "ingressgateway"                       // Label value denoting the ingress gateway pod selector
	IstioIngressGatewayServiceAccount       = "istio-ingressgateway-service-account" // Service account name that Istio ingress gateway is running as
	IstioNamespace                          = "istio-system"                         // Which namespace Istio is installed in
	IstioPrometheusServiceAccount           = "istio-prometheus-service-account"     // Service account name that Prometheus is running as
	IstioServiceEntryLocationExternal       = "MESH_EXTERNAL"                        // Service entries external to the cluster
	IstioServiceEntryResolutionDNS          = "DNS"                                  // Service entry lookup type
	IstioGatewayPrefix                      = "istio-system/ingress-gateway-%s"
	IstioVirtualServiceTotalWeight    int32 = 100 // The total weight of all routes must equal 100
)
