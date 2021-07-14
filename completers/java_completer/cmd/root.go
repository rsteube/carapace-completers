package cmd

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "java",
	Short: "Launches a Java application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-D") &&
			len(arg) > 2 &&
			rootCmd.Flag(arg[1:]) == nil {
			rootCmd.Flags().Bool(arg[1:], false, "set a system property") // fake flag to prevent errors
		}
	}

	carapace.Override(carapace.Opts{
		LongShorthand:   true,
		OptArgDelimiter: ":",
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("d32", false, "use a 32-bit data model if available")
	rootCmd.Flags().Bool("d64", false, "use a 64-bit data model if available")
	rootCmd.Flags().Bool("server", false, "to select the \"server\" VM")
	rootCmd.Flags().String("cp", "", "class search path of directories and zip/jar files")
	rootCmd.Flags().Bool("classpath", false, "class search path of directories and zip/jar files")
	rootCmd.Flags().Bool("D", false, "set a system property") // duh - no space between flag and value allowed so add it as bool and add any entered value as fake flag
	rootCmd.Flags().String("verbose", "", "enable verbose output")
	rootCmd.Flags().Bool("version", false, "print product version and exit")
	rootCmd.Flags().Bool("showversion", false, "print product version and continue")
	rootCmd.Flags().Bool("?", false, "print help message")
	rootCmd.Flags().String("jar", "", "jar file to execute")
	rootCmd.Flags().String("javaagent", "", "load Java programming language agent, see java.lang.instrument")
	rootCmd.Flags().Bool("X", false, "print help on non-standard options")
	rootCmd.Flags().String("ea", "", "enable assertions with specified granularity")
	rootCmd.Flags().String("enableassertions", "", "enable assertions with specified granularity")
	rootCmd.Flags().String("da", "", "disable assertions with specified granularity")
	rootCmd.Flags().String("disableassertions", "", "disable assertions with specified granularity")
	rootCmd.Flags().Bool("esa", false, "enable system assertions")
	rootCmd.Flags().Bool("enablesystemassertions", false, "enable system assertions")
	rootCmd.Flags().Bool("dsa", false, "disable system assertions")
	rootCmd.Flags().Bool("disablesystemassertions", false, "disable system assertions")
	rootCmd.Flags().String("agentlib", "", "load native agent library")
	rootCmd.Flags().String("agentpath", "", "load native agent library by full pathname")
	rootCmd.Flags().String("splash", "", "show splash screen with specified image")
	rootCmd.Flags().Bool("help", false, "print help message")

	rootCmd.Flag("agentlib").NoOptDefVal = " "
	rootCmd.Flag("agentpath").NoOptDefVal = " "
	rootCmd.Flag("da").NoOptDefVal = " "
	rootCmd.Flag("disableassertions").NoOptDefVal = " "
	rootCmd.Flag("javaagent").NoOptDefVal = " "
	rootCmd.Flag("splash").NoOptDefVal = " "
	rootCmd.Flag("verbose").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"agentpath": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
		"cp": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
		"classpath": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
		"jar": carapace.ActionFiles(".jar"),
		"javaagent": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles(".jar")
			default:
				return carapace.ActionValues()
			}
		}),
		"verbose": carapace.ActionValues("class", "gc", "jni"),
	})
}
