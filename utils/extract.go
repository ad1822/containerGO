package utils

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/containers/image/v5/oci/layout"
	"github.com/containers/image/v5/types"
	"github.com/opencontainers/go-digest"
)

// Manifest represents the OCI image manifest
type Manifest struct {
	Layers []struct {
		Digest string `json:"digest"`
	} `json:"layers"`
}

// ExtractRootFS extracts the root filesystem from the pulled image
func ExtractRootFS(imagePath, extractTo string) error {
	ref, err := layout.ParseReference(imagePath)
	fmt.Println("Ref : ", ref)
	if err != nil {
		return fmt.Errorf("failed to parse OCI image: %v", err)
	}

	imgSrc, err := ref.NewImageSource(context.Background(), &types.SystemContext{})
	fmt.Println("Image Source", imgSrc)
	if err != nil {
		return fmt.Errorf("failed to get image source: %v", err)
	}
	defer imgSrc.Close()

	// Get the raw manifest JSON
	manifestBytes, _, err := imgSrc.GetManifest(context.Background(), nil)
	fmt.Println("Manifest Bytes : ", manifestBytes)
	if err != nil {
		return fmt.Errorf("failed to get manifest: %v", err)
	}

	// Unmarshal the JSON into the Manifest struct
	var manifest Manifest
	if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
		return fmt.Errorf("failed to parse manifest: %v", err)
	}

	// Ensure the extraction directory exists
	if err := os.MkdirAll(extractTo, 0755); err != nil {
		return fmt.Errorf("failed to create extraction directory: %v", err)
	}

	// Extract layers (root filesystem)
	for _, layer := range manifest.Layers {
		fmt.Println("Extracting layer:", layer.Digest)

		// Convert string to digest.Digest
		layerDigest := digest.Digest(layer.Digest)

		// Get layer content
		layerReader, _, err := imgSrc.GetBlob(context.Background(), types.BlobInfo{Digest: layerDigest}, nil)
		if err != nil {
			return fmt.Errorf("failed to get layer: %v", err)
		}

		// Decompress and extract the tarball
		err = extractCompressedTar(layerReader, extractTo)
		if err != nil {
			return fmt.Errorf("failed to extract layer: %v", err)
		}
	}

	fmt.Println("Root filesystem extracted to:", extractTo)
	return nil
}

// extractCompressedTar decompresses a gzipped tar archive before extracting
func extractCompressedTar(reader io.Reader, dest string) error {
	// Try to decompress the layer
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		// If it's not gzip, assume it's a raw tar
		fmt.Println("Layer is not compressed, extracting directly")
		return extractTar(reader, dest)
	}
	defer gzipReader.Close()

	fmt.Println("Decompressing layer before extraction")
	return extractTar(gzipReader, dest)
}

// extractTar extracts a tar archive to the specified directory
func extractTar(reader io.Reader, dest string) error {
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return err
			}
			outFile, err := os.Create(target)
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}
	return nil
}
