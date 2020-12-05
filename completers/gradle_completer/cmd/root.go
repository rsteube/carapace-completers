package cmd

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gradle",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().Bool("build-cache", false, "Enables the Gradle build cache. Gradle will try to reuse outputs from previous builds.")
	rootCmd.Flags().StringP("build-file", "b", "", "Specify the build file.")
	rootCmd.Flags().Bool("configure-on-demand", false, "Configure necessary projects only. Gradle will attempt to reduce configuration time for large multi-project builds. [incubating]")
	rootCmd.Flags().String("console", "", "Specifies which type of console output to generate. Values are 'plain', 'auto' (default), 'rich' or 'verbose'.")
	rootCmd.Flags().Bool("continue", false, "Continue task execution after a task failure.")
	rootCmd.Flags().BoolP("continuous", "t", false, "Enables continuous build. Gradle does not exit and will re-execute tasks when task file inputs change.")
	rootCmd.Flags().Bool("daemon", false, "Uses the Gradle Daemon to run the build. Starts the Daemon if not running.")
	rootCmd.Flags().BoolP("debug", "d", false, "Log in debug mode (includes normal stacktrace).")
	rootCmd.Flags().StringP("dependency-verification", "F", "", "Configures the dependency verification mode (strict, lenient or off) [incubating]")
	rootCmd.Flags().BoolP("dry-run", "m", false, "Run the builds with all task actions disabled.")
	rootCmd.Flags().StringP("exclude-task", "x", "", "Specify a task to be excluded from execution.")
	rootCmd.Flags().Bool("export-keys", false, "Exports the public keys used for dependency verification. [incubating]")
	rootCmd.Flags().Bool("foreground", false, "Starts the Gradle Daemon in the foreground.")
	rootCmd.Flags().BoolP("full-stacktrace", "S", false, "Print out the full (very verbose) stacktrace for all exceptions.")
	rootCmd.Flags().StringP("gradle-user-home", "g", "", "Specifies the gradle user home directory.")
	rootCmd.Flags().Bool("include-build", false, "Include the specified build in the composite.")
	rootCmd.Flags().BoolP("info", "i", false, "Set log level to info.")
	rootCmd.Flags().StringP("init-script", "I", "", "Specify an initialization script.")
	rootCmd.Flags().String("max-workers", "", "Configure the number of concurrent workers Gradle is allowed to use.")
	rootCmd.Flags().Bool("no-build-cache", false, "Disables the Gradle build cache.")
	rootCmd.Flags().Bool("no-configure-on-demand", false, "Disables the use of configuration on demand. [incubating]")
	rootCmd.Flags().Bool("no-daemon", false, "Do not use the Gradle daemon to run the build. Useful occasionally if you have configured Gradle to always run with the daemon by default.")
	rootCmd.Flags().Bool("no-parallel", false, "Disables parallel execution to build projects.")
	rootCmd.Flags().BoolP("no-rebuild", "a", false, "Do not rebuild project dependencies.")
	rootCmd.Flags().Bool("no-scan", false, "Disables the creation of a build scan. For more information about build scans, please visit https://gradle.com/build-scans.")
	rootCmd.Flags().Bool("offline", false, "Execute the build without accessing network resources.")
	rootCmd.Flags().Bool("parallel", false, "Build projects in parallel. Gradle will attempt to determine the optimal number of executor threads to use.")
	rootCmd.Flags().String("priority", "", "Specifies the scheduling priority for the Gradle daemon and all processes launched by it. Values are 'normal' (default) or 'low' [incubating]")
	rootCmd.Flags().Bool("profile", false, "Profile build execution time and generates a report in the <build_dir>/reports/profile directory.")
	rootCmd.Flags().String("project-cache-dir", "", "Specify the project-specific cache directory. Defaults to .gradle in the root project directory.")
	rootCmd.Flags().StringP("project-dir", "p", "", "Specifies the start directory for Gradle. Defaults to current directory.")
	rootCmd.Flags().StringP("project-prop", "P", "", "Set project property for the build script (e.g. -Pmyprop=myvalue).")
	rootCmd.Flags().BoolP("quiet", "q", false, "Log errors only.")
	rootCmd.Flags().Bool("refresh-dependencies", false, "Refresh the state of dependencies.")
	rootCmd.Flags().Bool("refresh-keys", false, "Refresh the public keys used for dependency verification. [incubating]")
	rootCmd.Flags().Bool("rerun-tasks", false, "Ignore previously cached task results.")
	rootCmd.Flags().Bool("scan", false, "Creates a build scan. Gradle will emit a warning if the build scan plugin has not been applied. (https://gradle.com/build-scans)")
	rootCmd.Flags().StringP("settings-file", "c", "", "Specify the settings file.")
	rootCmd.Flags().BoolP("stacktrace", "s", false, "Print out the stacktrace for all exceptions.")
	rootCmd.Flags().Bool("status", false, "Shows status of running and recently stopped Gradle Daemon(s).")
	rootCmd.Flags().Bool("stop", false, "Stops the Gradle Daemon if it is running.")
	rootCmd.Flags().StringP("system-prop", "D", "", "Set system property of the JVM (e.g. -Dmyprop=myvalue).")
	rootCmd.Flags().Bool("update-locks", false, "Perform a partial update of the dependency lock, letting passed in module notations change version. [incubating]")
	rootCmd.Flags().BoolP("version", "v", false, "Print version info.")
	rootCmd.Flags().BoolP("warn", "w", false, "Set log level to warn.")
	rootCmd.Flags().String("warning-mode", "", "Specifies which mode of warnings to generate. Values are 'all', 'fail', 'summary'(default) or 'none'")
	rootCmd.Flags().Bool("write-locks", false, "Persists dependency resolution for locked configurations, ignoring existing locking information if it exists [incubating]")
	rootCmd.Flags().BoolP("write-verification-metadata", "M", false, "Generates checksums for dependencies used in the project (comma-separated list) [incubating]")
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"build-file":              carapace.ActionFiles(""),
		"console":                 carapace.ActionValues("plain", "auto", "rich", "verbose"),
		"dependency-verification": carapace.ActionValues("strict", "lenient", "off"),
		"exclude-task":            ActionTasks(),
		"priority":                carapace.ActionValues("normal", "low"),
		"project-cache-dir":       carapace.ActionDirectories(),
		"project-dir":             carapace.ActionDirectories(),
		"settings-file":           carapace.ActionFiles("settings.gradle"),
		"warning-mode":            carapace.ActionValues("all", "fail", "summary", "none"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		ActionTasks(),
	)
}

func locateBuildConfig() (target string, err error) {
	folder := "."
	if rootCmd.Flag("project-dir").Changed {
		folder = rootCmd.Flag("project-dir").Value.String()
	}

	if target, err = util.FindReverse(folder, "build.gradle"); err != nil {
		target, err = util.FindReverse(folder, "build.gradle.kts")
	}
	return
}

func ActionTasks() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var content []byte
		if target, err := locateBuildConfig(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			if content, err = ioutil.ReadFile(target); err != nil {
				return carapace.ActionMessage(err.Error())
			}
		}

		checksum := fmt.Sprintf("%x", sha1.Sum(content))

		if tasks, err := tasks(checksum); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValuesDescribed(tasks...)
		}
	})
}

func tasks(checksum string) (tasks []string, err error) {
    // TODO create/use carapace.TempDir() func
	filename := fmt.Sprintf("%v/carapace/gradle/%v", os.TempDir(), checksum)
	if tasks, err = loadTasks(filename); err != nil {
		if tasks, err = parseTasks(); err == nil {
			err = saveTasks(filename, tasks)
		}
	}
	return
}

func saveTasks(filename string, tasks []string) (err error) {
	var m []byte
	if m, err = json.Marshal(tasks); err == nil {
		if err = os.MkdirAll(filepath.Dir(filename), 0777); err == nil {
			err = ioutil.WriteFile(filename, m, 0644)
		}
	}
	return
}

func loadTasks(filename string) (tasks []string, err error) {
	var content []byte
	if content, err = ioutil.ReadFile(filename); err == nil {
		err = json.Unmarshal(content, &tasks)
	}
	return
}

func parseTasks() ([]string, error) {
	output, err := exec.Command("gradle", "tasks", "--all").Output()
	if err != nil {
		return nil, err
	}

	patternTaskGroups := regexp.MustCompile(`[^ ]+ tasks\n-----+\n(?P<tasks>(.+\n)*)`)
	patternTask := regexp.MustCompile(`^(?P<task>[^ ]+)( - (?P<description>.*))?`)

	result := make([]string, 0)
	taskGroups := patternTaskGroups.FindAllStringSubmatch(string(output), -1)
	for _, taskGroup := range taskGroups {
		for _, line := range strings.Split(taskGroup[1], "\n") {
			if patternTask.MatchString(line) {
				task := patternTask.FindStringSubmatch(line)
				result = append(result, task[1], task[3])
			}
		}
	}
	return result, nil
}
