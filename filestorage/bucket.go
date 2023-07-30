package filestorage

import (
	"context"
	"io/ioutil"
	"os"

	"google.golang.org/api/option"
)

type Client struct {
	BucketPath string
}
type Bucket struct {
	BucketPath string
}
type Object struct {
	Filename   string
	BucketPath string
}
type Writer struct {
	ContentType string
	Filename    string
	BucketPath  string
}
type Reader struct {
	Filename   string
	BucketPath string
}

func NewClient(ctx context.Context, option option.ClientOption) (*Client, error) {
	c := Client{}
	return &c, nil
}

func (c *Client) Bucket(s string) *Bucket {
	b := Bucket{}
	b.BucketPath = c.BucketPath
	return &b
}

func (b *Bucket) Object(s string) *Object {
	o := Object{}
	o.Filename = s
	o.BucketPath = b.BucketPath
	return &o
}

func (o *Object) NewWriter(ctx context.Context) *Writer {
	w := Writer{}
	w.Filename = o.Filename
	w.BucketPath = o.BucketPath
	return &w
}

func (o *Object) NewReader(ctx context.Context) (*Reader, error) {
	r := Reader{}
	r.Filename = o.Filename
	r.BucketPath = o.BucketPath
	return &r, nil
}

func (o *Object) Delete(ctx context.Context) {
	os.Remove(o.BucketPath + "/" + o.Filename)
}

func (w *Writer) Close() {
}

func (w *Writer) Write(b []byte) (int, error) {
	ioutil.WriteFile(w.BucketPath+"/"+w.Filename, b, 0644)
	return len(b), nil
}
