package helpers

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// BuildRegionParams build strings from regions slice
func BuildRegionParams(regions ...string) string {
	var res []string
	for _, region := range regions {
		res = append(res, fmt.Sprintf(`"%s"`, region))
	}
	return fmt.Sprintf("[%s]", strings.Join(res, ", "))
}

// SetPostTFCleanUp schedule final terraform clean up
func SetPostTFCleanUp(t *testing.T, opts *terraform.Options) {
	if _, ok := os.LookupEnv("POLKADOT_TEST_NO_POST_TF_CLEANUP"); !ok {
		t.Log("Setting terrafrom deferred cleanup...")
		t.Cleanup(func() {
			terraform.Destroy(t, opts)
		})
	} else {
		t.Log("Skipping terrafrom deferred cleanup...")
	}
}

// SetInitialTFCleanUp schedule initial terraform clean up
func SetInitialTFCleanUp(t *testing.T, opts *terraform.Options) {
	if _, ok := os.LookupEnv("POLKADOT_TEST_INITIAL_TF_CLEANUP"); ok {
		t.Log("Starting terrafrom cleanup...")
		terraform.Destroy(t, opts)
	} else {
		t.Log("Skipping terrafrom cleanup...")
	}
}
