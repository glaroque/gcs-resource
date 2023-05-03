package check

import gcsresource "github.com/frodenas/gcs-resource"

type CheckRequest struct {
	Source  gcsresource.Source  `json:"source"`
	Version gcsresource.Version `json:"version"`
}

type CheckResponse []gcsresource.Version
