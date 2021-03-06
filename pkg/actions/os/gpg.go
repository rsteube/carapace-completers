package os

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionGpgKeyIds completes GPG key ids
//   ABCDEF1234567890 (some GPG key)
//   ABCDEF1234567891 (another GPG key)
func ActionGpgKeyIds() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("sh", "-c", "gpg --list-keys --with-colons | awk -F: '/^pub:|^uid:/ { print $5 $10}'")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		})
	})
}
