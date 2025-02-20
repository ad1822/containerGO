package commands

import (
	"archive/tar"
	"compress/gzip"
	"containerGO/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/containers/image/v5/oci/layout"
	"github.com/containers/image/v5/types"
	"github.com/fatih/color"
	"github.com/opencontainers/go-digest"
)

type Manifest struct {
	Layers []struct {
		Digest string `json:"digest"`
	} `json:"layers"`
}

func ExtractRootFS(imagePath, extractTo string) error {
	ref, err := layout.ParseReference(imagePath)
	if err != nil {
		return fmt.Errorf("failed to parse OCI image: %v", err)
	}

	imgSrc, err := ref.NewImageSource(context.Background(), &types.SystemContext{})
	if err != nil {
		return fmt.Errorf("failed to get image source: %v", err)
	}
	defer imgSrc.Close()

	manifestBytes, _, err := imgSrc.GetManifest(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to get manifest: %v", err)
	}

	var manifest Manifest
	if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
		return fmt.Errorf("failed to parse manifest: %v", err)
	}

	// Ensure extraction directory exists
	if err := os.MkdirAll(extractTo, 0755); err != nil {
		return fmt.Errorf("failed to create extraction directory: %v", err)
	}

	// Extract layers in correct order (Base → Latest)
	for i := 0; i < len(manifest.Layers); i++ {
		layer := manifest.Layers[i]
		// fmt.Println("Extracting layer:", layer.Digest)
		utils.Logger(color.FgCyan, fmt.Sprintf("Extracting layer : %s", layer.Digest))

		layerDigest := digest.Digest(layer.Digest)
		layerReader, _, err := imgSrc.GetBlob(context.Background(), types.BlobInfo{Digest: layerDigest}, nil)
		if err != nil {
			return fmt.Errorf("failed to get layer: %v", err)
		}
		defer layerReader.Close()

		err = extractCompressedTar(layerReader, extractTo)
		if err != nil {
			return fmt.Errorf("failed to extract layer: %v", err)
		}
	}

	utils.Logger(color.FgBlue, fmt.Sprintf("✅ Root Filesystem extracted to : %s", extractTo))
	// fmt.Println("Root filesystem extracted to:", extractTo)
	return nil
}

func extractCompressedTar(reader io.Reader, dest string) error {
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %v", err)
	}
	defer gzipReader.Close()

	utils.Logger(color.FgCyan, "Decompressing layer before extraction")
	return extractTar(gzipReader, dest)
}

func extractTar(reader io.Reader, dest string) error {
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar header: %v", err)
		}

		target := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}

		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return fmt.Errorf("failed to create parent directory: %v", err)
			}
			outFile, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				return fmt.Errorf("failed to create file: %v", err)
			}
			defer outFile.Close()

			if _, err := io.Copy(outFile, tarReader); err != nil {
				return fmt.Errorf("failed to write file: %v", err)
			}

		case tar.TypeSymlink:
			if _, err := os.Lstat(target); err == nil {
				fmt.Printf("Warning: removing existing file/directory to create symlink: %s\n", target)
				if err := os.RemoveAll(target); err != nil {
					return fmt.Errorf("failed to remove existing file/directory: %v", err)
				}
			}
			if err := os.Symlink(header.Linkname, target); err != nil {
				return fmt.Errorf("failed to create symlink: %v", err)
			}
		}
	}
	return nil
}
