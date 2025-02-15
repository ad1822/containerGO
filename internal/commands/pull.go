package commands

import (
	"containerGO/internal/utils"
	"context"
	"fmt"
	"path/filepath"

	"github.com/containers/image/v5/copy"
	"github.com/containers/image/v5/signature"
	"github.com/containers/image/v5/transports/alltransports"
	"github.com/containers/image/v5/types"
)

func PullImage(image string) error {
	ctx := context.Background()
	fmt.Println("Pulling image:", image)

	// Load default policy for image pulling
	policy, err := signature.DefaultPolicy(nil)
	if err != nil {
		return fmt.Errorf("failed to get default policy: %w", err)
	}

	policyCtx, err := signature.NewPolicyContext(policy)
	if err != nil {
		return fmt.Errorf("failed to create policy context: %w", err)
	}
	defer policyCtx.Destroy()

	// Handle image format to avoid registry duplication
	sourceRef, err := alltransports.ParseImageName(fmt.Sprintf("docker://%s", image))
	if err != nil {
		return fmt.Errorf("failed to parse source image: %w", err)
	}

	// Ensure the destination format is correct
	dir := utils.GetContainerBaseDir("Images")
	destPath := filepath.Join(dir, fmt.Sprintf("%s:latest", image))
	destRef, err := alltransports.ParseImageName(fmt.Sprintf("oci:%s", destPath))
	if err != nil {
		return fmt.Errorf("failed to parse destination: %w", err)
	}

	// Create a system context for authentication if needed
	systemCtx := &types.SystemContext{
		// Uncomment and add credentials if pulling private images
		// DockerAuthConfig: &types.DockerAuthConfig{
		// 	Username: "your-username",
		// 	Password: "your-password",
		// },
	}

	// Copy the image from the source to the destination
	_, err = copy.Image(ctx, policyCtx, destRef, sourceRef, &copy.Options{
		SourceCtx:      systemCtx,
		DestinationCtx: systemCtx,
	})
	if err != nil {
		return fmt.Errorf("failed to pull image: %w", err)
	}

	fmt.Println("Successfully pulled image:", image)
	return nil
}
