package commands

import (
	"context"
	"fmt"

	"github.com/containers/image/v5/copy"
	"github.com/containers/image/v5/signature"
	"github.com/containers/image/v5/transports/alltransports"
	"github.com/containers/image/v5/types"
)

// PullImage pulls an image from a registry and stores it as an OCI archive
func PullImage(image string) error {
	ctx := context.Background()

	// Allow all images (disable signature verification for now)
	policy, err := signature.DefaultPolicy(nil)
	if err != nil {
		return fmt.Errorf("failed to get default policy: %v", err)
	}

	policyCtx, err := signature.NewPolicyContext(policy)
	if err != nil {
		return fmt.Errorf("failed to create policy context: %v", err)
	}
	defer policyCtx.Destroy()

	// Define source (Docker Hub)
	sourceRef, err := alltransports.ParseImageName("docker://docker.io/library/" + image)
	if err != nil {
		return fmt.Errorf("failed to parse source image: %v", err)
	}

	// Define destination (OCI layout on disk)
	destRef, err := alltransports.ParseImageName("oci:/home/arcadian/Downloads/" + image + ":latest")
	if err != nil {
		return fmt.Errorf("failed to parse destination: %v", err)
	}

	// Create an empty system context
	systemCtx := &types.SystemContext{}

	// Copy the image
	_, err = copy.Image(ctx, policyCtx, destRef, sourceRef, &copy.Options{
		SourceCtx:      systemCtx,
		DestinationCtx: systemCtx,
	})
	if err != nil {
		return fmt.Errorf("failed to pull image: %v", err)
	}

	fmt.Println("Successfully pulled image:", image)
	return nil
}

// func main() {
// 	err := PullImage("alpine")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}
// }
