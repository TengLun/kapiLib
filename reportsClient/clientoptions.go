package reportsclient

import "net/http"

// SetClientHTTPClient allows a custom http client to be used by the client
// struct
func SetClientHTTPClient(customClient *http.Client) func(*Client) error {
	return func(c *Client) error {
		c.machine = customClient
		return nil
	}
}

// SetEmailList configures which emails a client will send reports to
func SetEmailList(emails []string) func(*Client) error {
	return func(c *Client) error {
		c.EmailList = emails
		return nil
	}
}

// SetDeliverToS3Bucket configures report delivery to a custom s3 bucket
func SetDeliverToS3Bucket(region, bucketName, accessKey, secretKey string) func(*Client) error {
	return func(c *Client) error {
		c.DeliveryMethod = []string{"S3bucket", region, bucketName, accessKey, secretKey}
		return nil
	}
}

// SetTimezone sets the report delivery timezone
func SetTimezone(tz string) func(*Client) error {
	return func(c *Client) error {
		c.TimeZone = tz
		return nil
	}
}
