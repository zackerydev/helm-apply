package cmd

import "github.com/spf13/cobra"

const rootCmdLongUsage = `
The Helm Diff Plugin

* Shows a diff explaining what a helm upgrade would change:
    This fetches the currently deployed version of a release
  and compares it to a local chart plus values. This can be
  used to visualize what changes a helm upgrade will perform.

* Shows a diff explaining what had changed between two revisions:
    This fetches previously deployed versions of a release
  and compares them. This can be used to visualize what changes
  were made during revision change.

* Shows a diff explaining what a helm rollback would change:
    This fetches the currently deployed version of a release
  and compares it to the previously deployed version of the release, that you
  want to rollback. This can be used to visualize what changes a
  helm rollback will perform.
`

func New() *cobra.Command {
	// chartCommand := newChartCommand()

	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply a specific values file inside of a helm chart",
		Long: `
    Helm Apply

    This plugin allows you to apply a specific values file inside of a helm chart.
    `,
		Args: cobra.ExactArgs(1),
		// Args: chartCommand.Args,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Applying values file")
		},
	}

	cmd.SetHelpCommand(&cobra.Command{})
	return cmd
}
