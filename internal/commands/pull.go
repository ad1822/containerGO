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
	"github.com/fatih/color"
)

func PullImage(image string) error {
	ctx := context.Background()
	utils.Logger(color.FgBlue, fmt.Sprintln("⬇️ Pulling image:", image))

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

	dir := utils.GetContainerBaseDir("Images")
	destPath := filepath.Join(dir, fmt.Sprintf("%s:latest", image))
	destRef, err := alltransports.ParseImageName(fmt.Sprintf("oci:%s", destPath))
	if err != nil {
		return fmt.Errorf("failed to parse destination: %w", err)
	}

	systemCtx := &types.SystemContext{
		// Uncomment and add credentials if pulling private images
		// DockerAuthConfig: &types.DockerAuthConfig{
		// 	Username: "your-username",
		// 	Password: "your-password",
		// },
	}

	imageName, err := copy.Image(ctx, policyCtx, destRef, sourceRef, &copy.Options{
		SourceCtx:      systemCtx,
		DestinationCtx: systemCtx,
	})

	if err != nil {
		return fmt.Errorf("failed to pull image: %w", err)
	} else {
		// println(string(imageName))
		utils.Logger(color.FgGreen, string(imageName))
	}

	utils.Logger(color.FgBlue, fmt.Sprintln("✅ Successfully pulled image:", image))

	return nil
}
