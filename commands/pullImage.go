package commands

import (
	"context"
	"fmt"

	"github.com/containers/image/v5/copy"
	"github.com/containers/image/v5/signature"
	"github.com/containers/image/v5/transports/alltransports"
	"github.com/containers/image/v5/types"
)

func PullImage(image string) error {
	ctx := context.Background()

	policy, err := signature.DefaultPolicy(nil)
	if err != nil {
		return fmt.Errorf("failed to get default policy: %v", err)
	}

	policyCtx, err := signature.NewPolicyContext(policy)
	if err != nil {
		return fmt.Errorf("failed to create policy context: %v", err)
	}
	defer policyCtx.Destroy()

	sourceRef, err := alltransports.ParseImageName("docker://docker.io/library/" + image)
	if err != nil {
		return fmt.Errorf("failed to parse source image: %v", err)
	}

	destRef, err := alltransports.ParseImageName("oci:/home/arcadian/Downloads/" + image + ":latest")
	if err != nil {
		return fmt.Errorf("failed to parse destination: %v", err)
	}

	systemCtx := &types.SystemContext{}

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
