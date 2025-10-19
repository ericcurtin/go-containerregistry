# Resumable Download Example

This example demonstrates how to use the resumable download feature to fetch specific byte ranges from container registry layers using HTTP range requests.

## Usage

```bash
go run main.go <digest-ref> <start-byte> <end-byte>
```

## Example

```bash
# Fetch the first 1024 bytes from a layer
go run main.go gcr.io/my-repo/my-image@sha256:abc123... 0 1023

# Resume a download starting from byte 1024
go run main.go gcr.io/my-repo/my-image@sha256:abc123... 1024 2047
```

## Use Cases

- **Resumable Downloads**: If a download is interrupted, you can resume from where it left off
- **Partial Content Access**: Access only the portion of a layer you need
- **Progressive Loading**: Load content incrementally for better user experience
- **Bandwidth Optimization**: Download only the required portions of large layers

## Notes

- Range requests require a digest reference (not a tag)
- The byte offsets are inclusive (start and end bytes are both included)
- Hash verification is not performed on partial content
- Not all registries support range requests (though most modern ones do)
