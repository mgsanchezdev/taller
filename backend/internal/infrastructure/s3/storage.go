package s3

// Storage provides S3-compatible storage for photos, videos, and PDFs.
// TODO: implement with minio or AWS SDK.
// Path structure: vehicles/{plate}/{entry_id|repair_id}/{photos|videos|docs}/
type Storage interface {
	// Upload(key string, body io.Reader, contentType string) (url string, err error)
	// GetURL(key string) string
}
