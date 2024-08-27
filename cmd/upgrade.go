package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/kube"
)

type applyCmd struct {
	release                  string
	chart                    string
	chartVersion             string
	chartRepo                string
	detailedExitCode         bool
	devel                    bool
	disableValidation        bool
	disableOpenAPIValidation bool
	enableDNS                bool
	namespace                string // namespace to assume the release to be installed into. Defaults to the current kube config namespace.
	valueFiles               valueFiles
	values                   []string
	stringValues             []string
	stringLiteralValues      []string
	jsonValues               []string
	fileValues               []string
	reuseValues              bool
	resetValues              bool
	allowUnreleased          bool
	noHooks                  bool
	includeTests             bool
	postRenderer             string
	postRendererArgs         []string
	insecureSkipTLSVerify    bool
	install                  bool
	normalizeManifests       bool
	threeWayMerge            bool
	extraAPIs                []string
	kubeVersion              string
	useUpgradeDryRun         bool
	Options
}

func (a *applyCmd) runHelm() error {
	var err error

	args := []string{"pull", a.chart}

	cmd := exec.Command(os.Getenv("HELM_BIN"), args...)

	return nil
}

func newChartCommand() *cobra.Command {
	apply := applyCmd{
		namespace: os.Getenv("HELM_NAMESPACE"),
	}

	cmd := &cobra.Command{
		Use:     "upgrade [flags] [RELEASE] [CHART]",
		Short:   "Upgrade a helm chart and apply a specific values file",
		Long:    `TODO`,
		Example: `Todo`,
		Args: func(cmd *cobra.Command, args []string) error {
			return checkArgsLength(len(args), "release name", "chart path")
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return apply.runHelm()
		},
	}
}
