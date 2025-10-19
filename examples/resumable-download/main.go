// Copyright 2024 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This example demonstrates how to use resumable downloads to fetch
// specific byte ranges from container registry layers.
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <digest-ref> <start-byte> <end-byte>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s gcr.io/my-repo/my-image@sha256:abc123... 0 1023\n", os.Args[0])
		os.Exit(1)
	}

	// Parse the digest reference
	ref, err := name.NewDigest(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to parse digest: %v", err)
	}

	// Parse start and end byte offsets
	var start, end int64
	if _, err := fmt.Sscanf(os.Args[2], "%d", &start); err != nil {
		log.Fatalf("Failed to parse start byte: %v", err)
	}
	if _, err := fmt.Sscanf(os.Args[3], "%d", &end); err != nil {
		log.Fatalf("Failed to parse end byte: %v", err)
	}

	// Fetch the byte range
	log.Printf("Fetching bytes %d-%d from %s...", start, end, ref.Name())
	rc, err := remote.LayerRange(ref, start, end)
	if err != nil {
		log.Fatalf("Failed to fetch byte range: %v", err)
	}
	defer rc.Close()

	// Copy the range to stdout
	n, err := io.Copy(os.Stdout, rc)
	if err != nil {
		log.Fatalf("Failed to read bytes: %v", err)
	}

	log.Printf("\nSuccessfully read %d bytes", n)
}
