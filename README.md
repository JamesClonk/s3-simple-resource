# s3-simple-resource
Concourse resource for simple file upload/download from S3-compatible provider, using signature V2 signing.

## Source Configuration

* `access_key`: *Required.* The access key to use.

* `secret_key`: *Required.* The secret key to use.

* `bucket`: *Required.* The name of the bucket.

* `host`: *Required.* Host endpoint of S3-compatible provider.

* `insecure`: *Optional.* Switch to insecure HTTP.

## Behavior

### `in`: Download a file from the bucket.

Given a file specified by `source`, download it from the S3 bucket, save it under `target`.

#### Parameters

* `source`: *Required.* File to download.

* `target`: *Required.* Name/destination of new file.


### `out`: Upload a file to the bucket.

Given a file specified by `source`, upload it to the S3 bucket, save it under `target`.

#### Parameters

* `source`: *Required.* File to upload.

* `target`: *Required.* Name/destination of new file in the bucket.
