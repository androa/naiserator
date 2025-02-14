package resourcecreator_test

import (
	"testing"

	nais "github.com/nais/naiserator/pkg/apis/nais.io/v1alpha1"
	networking_istio_io_v1alpha3 "github.com/nais/naiserator/pkg/apis/networking.istio.io/v1alpha3"
	rbac_istio_io_v1alpha1 "github.com/nais/naiserator/pkg/apis/rbac.istio.io/v1alpha1"
	"github.com/nais/naiserator/pkg/resourcecreator"
	"github.com/nais/naiserator/pkg/test/fixtures"
	"github.com/stretchr/testify/assert"
	"k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
	rbacv1 "k8s.io/api/rbac/v1"
)

type realObjects struct {
	deployment         *v1.Deployment
	service            *corev1.Service
	serviceAccount     *corev1.ServiceAccount
	hpa                *autoscalingv1.HorizontalPodAutoscaler
	ingress            *extensionsv1beta1.Ingress
	networkPolicy      *networkingv1.NetworkPolicy
	serviceRole        *rbac_istio_io_v1alpha1.ServiceRole
	serviceRoleBinding *rbac_istio_io_v1alpha1.ServiceRoleBinding
	virtualServices    []*networking_istio_io_v1alpha3.VirtualService
	role               *rbacv1.Role
	rolebinding        *rbacv1.RoleBinding
}

func getRealObjects(resources []runtime.Object) (o realObjects) {
	for _, r := range resources {
		switch v := r.(type) {
		case *v1.Deployment:
			o.deployment = v
		case *corev1.Service:
			o.service = v
		case *corev1.ServiceAccount:
			o.serviceAccount = v
		case *autoscalingv1.HorizontalPodAutoscaler:
			o.hpa = v
		case *extensionsv1beta1.Ingress:
			o.ingress = v
		case *networkingv1.NetworkPolicy:
			o.networkPolicy = v
		case *rbac_istio_io_v1alpha1.ServiceRole:
			o.serviceRole = v
		case *rbac_istio_io_v1alpha1.ServiceRoleBinding:
			o.serviceRoleBinding = v
		case *networking_istio_io_v1alpha3.VirtualService:
			o.virtualServices = append(o.virtualServices, v)
		case *rbacv1.Role:
			o.role = v
		case *rbacv1.RoleBinding:
			o.rolebinding = v
		}
	}
	return
}

// Test that a specified application spec results in the correct Kubernetes resources.
func TestCreate(t *testing.T) {
	t.Run("default application spec merges into empty struct", func(t *testing.T) {
		app := &nais.Application{}
		err := nais.ApplyDefaults(app)
		assert.NoError(t, err)
	})

	t.Run("application spec needs required parameters", func(t *testing.T) {
		app := fixtures.MinimalFailingApplication()
		opts := resourcecreator.NewResourceOptions()
		err := nais.ApplyDefaults(app)
		assert.NoError(t, err)

		resources, err := resourcecreator.Create(app, opts)
		assert.Error(t, err)
		assert.Nil(t, resources)
	})

	t.Run("team label and application name is propagated to created resources", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		opts := resourcecreator.NewResourceOptions()
		err := nais.ApplyDefaults(app)
		assert.NoError(t, err)

		resources, err := resourcecreator.Create(app, opts)
		assert.NoError(t, err)

		objects := getRealObjects(resources)
		assert.Equal(t, app.Name, objects.deployment.Name)
		assert.Equal(t, app.Namespace, objects.deployment.Namespace)
		assert.Equal(t, app.Name, objects.deployment.Labels["app"])
		assert.Equal(t, app.Labels["team"], objects.deployment.Labels["team"])
	})

	t.Run("all basic resource types are created from an application spec", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		opts := resourcecreator.NewResourceOptions()
		err := nais.ApplyDefaults(app)
		assert.NoError(t, err)

		resources, err := resourcecreator.Create(app, opts)
		assert.NoError(t, err)

		objects := getRealObjects(resources)
		assert.NotNil(t, objects.hpa)
		assert.NotNil(t, objects.service)
		assert.NotNil(t, objects.serviceAccount)
		assert.NotNil(t, objects.deployment)
		assert.Nil(t, objects.ingress)
	})

	t.Run("an ingress object is created if ingress paths are specified", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		app.Spec.Ingresses = []string{"https://foo.bar/baz"}
		opts := resourcecreator.NewResourceOptions()
		err := nais.ApplyDefaults(app)
		assert.NoError(t, err)

		resources, err := resourcecreator.Create(app, opts)
		assert.NoError(t, err)

		objects := getRealObjects(resources)
		assert.NotNil(t, objects.ingress)
	})

	t.Run("erroneous ingress uris create errors", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		app.Spec.Ingresses = []string{"gopher://lol"}
		opts := resourcecreator.NewResourceOptions()
		err := nais.ApplyDefaults(app)
		assert.NoError(t, err)

		resources, err := resourcecreator.Create(app, opts)
		assert.Error(t, err)
		assert.Nil(t, resources)
	})

	t.Run("istio resources are omitted when access policy creation is disabled", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		opts := resourcecreator.NewResourceOptions()
		err := nais.ApplyDefaults(app)
		assert.NoError(t, err)

		resources, err := resourcecreator.Create(app, opts)
		assert.NoError(t, err)

		objects := getRealObjects(resources)
		assert.Nil(t, objects.virtualServices)
		assert.Nil(t, objects.serviceRoleBinding)
		assert.Nil(t, objects.serviceRole)
		assert.Nil(t, objects.networkPolicy)
	})

	t.Run("istio resources are created when access policy creation is enabled", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		app.Spec.Ingresses = []string{"host.domain.tld"}
		opts := resourcecreator.NewResourceOptions()
		opts.AccessPolicy = true
		err := nais.ApplyDefaults(app)
		assert.NoError(t, err)

		resources, err := resourcecreator.Create(app, opts)
		assert.NoError(t, err)

		objects := getRealObjects(resources)
		assert.Len(t, objects.virtualServices, 1)
		assert.NotNil(t, objects.serviceRoleBinding)
		assert.NotNil(t, objects.serviceRole)
		assert.NotNil(t, objects.networkPolicy)
	})

	t.Run("servicerole and servicerolebinding resources are created when access policy creation is enabled", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		opts := resourcecreator.NewResourceOptions()
		opts.AccessPolicy = true
		app.Spec.Ingresses = []string{"https://foo.bar"}

		err := nais.ApplyDefaults(app)
		assert.NoError(t, err)

		resources, err := resourcecreator.Create(app, opts)
		assert.NoError(t, err)

		objects := getRealObjects(resources)
		assert.NotNil(t, objects.virtualServices)
		assert.NotNil(t, objects.serviceRoleBinding)
		assert.NotNil(t, objects.serviceRole)
		assert.NotNil(t, objects.networkPolicy)
	})

	t.Run("leader election rbac is created when LE is requested", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		app.Spec.LeaderElection = true
		err := nais.ApplyDefaults(app)
		assert.NoError(t, err)

		opts := resourcecreator.NewResourceOptions()
		resources, err := resourcecreator.Create(app, opts)
		assert.NoError(t, err)

		objects := getRealObjects(resources)
		assert.NotNil(t, objects.role)
		assert.Equal(t, app.Name, objects.role.Name)
		assert.NotNil(t, objects.rolebinding)
		assert.Equal(t, app.Name, objects.rolebinding.Name)
	})
}
