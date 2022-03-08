package cilium

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/templater"
	"github.com/aws/eks-anywhere/pkg/types"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/constants"
)

const namespace = constants.KubeSystemNamespace

//go:embed networkPolicy.yaml
var networkPolicyAllowAll string

type Cilium struct {
	*Upgrader
}

func NewCilium(client Client, helm Helm) *Cilium {
	return &Cilium{
		Upgrader: NewUpgrader(client, helm),
	}
}

func (c *Cilium) GenerateManifest(ctx context.Context, managementCluster, workloadCluster *types.Cluster, clusterSpec *cluster.Spec, provider providers.Provider) ([]byte, error) {
	ciliumManifest, err := c.templater.GenerateManifest(ctx, clusterSpec)
	//ciliumManifest, err := networking.LoadManifest(clusterSpec, clusterSpec.VersionsBundle.Cilium.Manifest)
	fmt.Printf("[RAJ] cilium manifest: %v\n", string(ciliumManifest))
	values := map[string]interface{}{
		"managementCluster":  !managementCluster.ExistingManagement,
		"providerNamespaces": provider.GetDeployments(),
	}
	// if always mode
	bytes, err := templater.Execute(networkPolicyAllowAll, values)
	if err != nil {
		return nil, err
	}
	content := templater.AppendYamlResources(ciliumManifest, bytes)
	//return ciliumManifest, err
	return content, err
}
