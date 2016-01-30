/*
   Copyright 2016 Fabio Berchtold

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package s3resource

import (
	"log"

	"github.com/minio/minio-go"
)

type Client interface {
	UploadFile(bucket, source, target string) error
	DownloadFile(bucket, source, target string) error
}

type client struct {
	client minio.CloudStorageClient
}

func NewClient(accessKey, secretKey, host string, insecure bool) (Client, error) {
	s3Client, err := minio.NewV2(host, accessKey, secretKey, insecure)
	if err != nil {
		return nil, err
	}

	return &client{
		client: s3Client,
	}, nil
}

func (c *client) UploadFile(bucket, source, target string) error {
	n, err := c.client.FPutObject(bucket, target, source, "application/octet-stream")
	if err != nil {
		return err
	}
	log.Printf("uploaded [%s] to [%s/%s] successfully\n", source, n, bucket, target)
	return nil
}

func (c *client) DownloadFile(bucket, source, target string) error {
	err := c.client.FGetObject(bucket, source, target)
	if err != nil {
		return err
	}
	log.Printf("downloaded [%s/%s] to [%s] successfully\n", source, bucket, target)
	return nil
}
