package git

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ResolveCommit(repo, version string) (string, error) {
	switch {
	case len(version) == 40 && func() bool { _, err := hex.DecodeString(version); return err == nil }():
		return version, nil
	case strings.HasPrefix(version, "refs/"):
		url := repo + "/info/refs?service=git-upload-pack"
		resp, err := http.Get(url)
		if err != nil {
			return "", fmt.Errorf("%v %v", url, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("%v HTTP %v", url, resp.StatusCode)
		}

		refsBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("%v %v", url, err)
		}

		for _, line := range strings.Split(string(refsBytes), "\n") {
			if len(line) < 46 || line[44] != ' ' {
				continue
			}
			if line[45:] == version {
				return line[4:44], nil
			}
		}

		return "", fmt.Errorf("did not find a ref named %q in .git/info/refs", version)
	default:
		return "", fmt.Errorf("version %q doesn't look like either a commit hash or a ref name", version)
	}
}
