package datasets

import "kwil/pkg/log"

type DatasetConnectionOpts func(*Dataset)

func WithPath(path string) DatasetConnectionOpts {
	return func(d *Dataset) {
		d.path = path
	}
}

func WithLogger(logger log.Logger) DatasetConnectionOpts {
	return func(d *Dataset) {
		d.log = logger
	}
}
