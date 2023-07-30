
The real 

`cloud.google.com/go/storage` and `google.golang.org/api/option`
packages are what production code will eventually use. But, during dev time
it's nice to just write files to the hard drive of your
[one free tier google compute instance](https://many.pw/hosting).
You get 30 GB for free! And storing in a real bucket will cost some money.

```
func getStorageClient() (*filestorage.Client, string) {
	keyPath := ""
	client, err := filestorage.NewClient(context.Background(),
		option.WithCredentialsFile(keyPath))
	if err != nil {
		return nil, ""
	}
	client.BucketPath = "/bucket"
	bucket := "unique-name"
	return client, bucket
}
```

This is the same make client code you'll need but it's using this fake `filestorage`
package which does everything the real client does but uses your machine's hard drive.

Write:

```
	client, bucket := getStorageClient()
	w := client.Bucket(bucket).Object(filename).NewWriter(context.Background())
	w.ContentType = "application/octet-stream"
	w.Write(data)
	w.Close()
```

Delete:

```
	client, bucket := getStorageClient(c)
	ctx := context.Background()
	client.Bucket(bucket).Object(object).Delete(ctx)
```

How much space is left:

```
func getAvailableDiskSpace(path string) (uint64, error) {
	var stat syscall.Statfs_t

	err := syscall.Statfs(path, &stat)
	if err != nil {
		return 0, err
	}

	availableSpace := stat.Bavail * uint64(stat.Bsize)
	return availableSpace, nil
}
```
