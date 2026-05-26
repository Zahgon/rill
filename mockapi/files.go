package mockapi

import (
	"context"
)

// DownloadFile simulates a file download. It returns the whole content as []byte.
func DownloadFile(ctx context.Context, url string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
