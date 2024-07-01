package bucket

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type BucketType int

const (
	AwsS3 BucketType = iota
)

type BucketInterface interface {
	Upload(io.Reader, string) error
	Download(string, string) (*os.File, error)
	Delete(string) error
}

type Bucket struct {
	p BucketInterface
}

func New(bt BucketType, cfg any) (b *Bucket, err error) {
	rt := reflect.TypeOf(cfg)

	switch bt {
	case AwsS3:
		if rt.Name() != "AwsConfig" {
			return nil, fmt.Errorf("config is not compatible with bucket type: %s", rt.Name())
		}

		b.p = newAwsSession(cfg.(AwsConfig))
	default:
		return nil, fmt.Errorf("config is not compatible with bucket type: %s", rt.Name())
	}
	return
}

func (b *Bucket) Upload(r io.Reader, name string) error {
	return b.p.Upload(r, name)
}

func (b *Bucket) Download(name string, path string) (*os.File, error) {
	return b.p.Download(name, path)
}

func (b *Bucket) Delete(name string) error {
	return b.p.Delete(name)
}
