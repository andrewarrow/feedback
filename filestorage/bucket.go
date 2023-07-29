package filestorage

import (
	"context"

	"google.golang.org/api/option"
)

type Client struct {
}
type Bucket struct {
}
type Object struct {
}
type Writer struct {
	ContentType string
}

func NewClient(ctx context.Context, option *option.ClientOption) *Client {
	c := Client{}
	return &c
}

func (c *Client) Bucket(s string) *Bucket {
	b := Bucket{}
	return &b
}

func (b *Bucket) Object(s string) *Object {
	o := Object{}
	return &o
}

func (o *Object) NewWriter(ctx context.Context) *Writer {
	w := Writer{}
	return &w
}

func (w *Writer) Close() {
}

func (w *Writer) Write(b []byte) (int, error) {
	return 0, nil
}
