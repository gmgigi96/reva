package utils

import (
	"encoding/base32"
	"fmt"
	"strings"

	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
)

func DecodeSpaceID(raw string) (storageID, path string, ok bool) {
	// The input is expected to be in the form of <storage_id>$<base32(<path>)
	s := strings.SplitN(raw, "$", 2)
	if len(s) != 2 {
		return
	}

	storageID = s[0]
	encodedPath := s[1]
	p, err := base32.StdEncoding.DecodeString(encodedPath)
	if err != nil {
		return
	}

	path = string(p)
	ok = true
	return
}

func DecodeResourceID(raw string) (storageID, path, itemID string, ok bool) {
	// The input is expected to be in the form of <storage_id>$<base32(<path>)!<item_id>
	s := strings.SplitN(raw, "!", 2)
	if len(s) != 2 {
		return
	}
	itemID = s[1]
	storageID, path, ok = DecodeSpaceID(s[0])
	return
}

func EncodeResourceID(r *provider.ResourceId) string {
	return fmt.Sprintf("%s$%s!%s", r.StorageId, r.SpaceId, r.OpaqueId)
}

func EncodeSpaceID(path string) string {
	return base32.StdEncoding.EncodeToString([]byte(path))
}
