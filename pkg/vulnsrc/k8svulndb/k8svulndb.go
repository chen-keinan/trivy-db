package k8svulndb

import (
	"path/filepath"

	"github.com/aquasecurity/trivy-db/pkg/types"
	"github.com/aquasecurity/trivy-db/pkg/vulnsrc/osv"
	"github.com/aquasecurity/trivy-db/pkg/vulnsrc/vulnerability"
)

var (
	k8sDir = filepath.Join("vuln-list-k8s", "upstream")
)

func NewVulnSrc() osv.OSV {
	sources := map[types.Ecosystem]types.DataSource{
		vulnerability.Kubernetes: {
			ID:   vulnerability.K8sVulnDB,
			Name: "Official Kubernetes CVE Feed",
			URL:  "https://kubernetes.io/docs/reference/issues-security/official-cve-feed/index.json",
		},
	}
	return osv.New(k8sDir, vulnerability.K8sVulnDB, sources, &defaultTransformer{})
}

type defaultTransformer struct{}

func (t *defaultTransformer) TransformAdvisories(advs []osv.Advisory, entry osv.Entry) ([]osv.Advisory, error) {
	return advs, nil
}
