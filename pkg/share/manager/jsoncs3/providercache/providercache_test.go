// Copyright 2018-2022 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package providercache_test

import (
	"context"
	"io/ioutil"
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	collaboration "github.com/cs3org/go-cs3apis/cs3/sharing/collaboration/v1beta1"
	"github.com/cs3org/reva/v2/pkg/share/manager/jsoncs3/providercache"
	"github.com/cs3org/reva/v2/pkg/storage/utils/metadata"
)

var _ = Describe("Cache", func() {
	var (
		c       providercache.Cache
		storage metadata.Storage

		storageID = "storageid"
		spaceID   = "spaceid"
		shareID   = "storageid$spaceid!share1"
		share1    = &collaboration.Share{
			Id: &collaboration.ShareId{
				OpaqueId: "share1",
			},
		}
		ctx    context.Context
		tmpdir string
	)

	BeforeEach(func() {
		ctx = context.Background()

		var err error
		tmpdir, err = ioutil.TempDir("", "providercache-test")
		Expect(err).ToNot(HaveOccurred())

		err = os.MkdirAll(tmpdir, 0755)
		Expect(err).ToNot(HaveOccurred())

		storage, err = metadata.NewDiskStorage(tmpdir)
		Expect(err).ToNot(HaveOccurred())

		c = providercache.New(storage)
		Expect(c).ToNot(BeNil())
	})

	AfterEach(func() {
		if tmpdir != "" {
			os.RemoveAll(tmpdir)
		}
	})

	Describe("Add", func() {
		It("adds a share", func() {
			s := c.Get(storageID, spaceID, shareID)
			Expect(s).To(BeNil())

			Expect(c.Add(ctx, storageID, spaceID, shareID, share1)).To(Succeed())

			s = c.Get(storageID, spaceID, shareID)
			Expect(s).ToNot(BeNil())
			Expect(s).To(Equal(share1))
		})

		It("sets the mtime", func() {
			Expect(c.Add(ctx, storageID, spaceID, shareID, share1)).To(Succeed())
			Expect(c.Providers[storageID].Spaces[spaceID].Mtime).ToNot(Equal(time.Time{}))
		})

		It("updates the mtime", func() {
			Expect(c.Add(ctx, storageID, spaceID, shareID, share1)).To(Succeed())
			old := c.Providers[storageID].Spaces[spaceID].Mtime
			Expect(c.Add(ctx, storageID, spaceID, shareID, share1)).To(Succeed())
			Expect(c.Providers[storageID].Spaces[spaceID].Mtime).ToNot(Equal(old))
		})
	})

	Context("with an existing entry", func() {
		BeforeEach(func() {
			Expect(c.Add(ctx, storageID, spaceID, shareID, share1)).To(Succeed())
		})

		Describe("Get", func() {
			It("returns the entry", func() {
				s := c.Get(storageID, spaceID, shareID)
				Expect(s).ToNot(BeNil())
			})
		})

		Describe("Remove", func() {
			It("removes the entry", func() {
				s := c.Get(storageID, spaceID, shareID)
				Expect(s).ToNot(BeNil())
				Expect(s).To(Equal(share1))

				Expect(c.Remove(ctx, storageID, spaceID, shareID)).To(Succeed())

				s = c.Get(storageID, spaceID, shareID)
				Expect(s).To(BeNil())
			})

			It("updates the mtime", func() {
				Expect(c.Add(ctx, storageID, spaceID, shareID, share1)).To(Succeed())
				old := c.Providers[storageID].Spaces[spaceID].Mtime
				Expect(c.Remove(ctx, storageID, spaceID, shareID)).To(Succeed())
				Expect(c.Providers[storageID].Spaces[spaceID].Mtime).ToNot(Equal(old))
			})
		})

		Describe("Persist", func() {
			It("handles non-existent storages", func() {
				Expect(c.Persist(ctx, "foo", "bar")).To(Succeed())
			})
			It("handles non-existent spaces", func() {
				Expect(c.Persist(ctx, storageID, "bar")).To(Succeed())
			})

			It("persists", func() {
				Expect(c.Persist(ctx, storageID, spaceID)).To(Succeed())
			})

			It("updates the mtime", func() {
				oldMtime := c.Providers[storageID].Spaces[spaceID].Mtime

				Expect(c.Persist(ctx, storageID, spaceID)).To(Succeed())
				Expect(c.Providers[storageID].Spaces[spaceID].Mtime).ToNot(Equal(oldMtime))
			})

		})

		Describe("PersistWithTime", func() {
			It("does not persist if the mtime on disk is more recent", func() {
				Expect(c.PersistWithTime(ctx, storageID, spaceID, time.Now().Add(-3*time.Hour))).ToNot(Succeed())
			})
		})

		Describe("Sync", func() {
			BeforeEach(func() {
				Expect(c.Persist(ctx, storageID, spaceID)).To(Succeed())
				// reset in-memory cache
				c = providercache.New(storage)
			})

			It("downloads if needed", func() {
				s := c.Get(storageID, spaceID, shareID)
				Expect(s).To(BeNil())

				Expect(c.Sync(ctx, storageID, spaceID)).To(Succeed())

				s = c.Get(storageID, spaceID, shareID)
				Expect(s).ToNot(BeNil())
			})

			It("does not download if not needed", func() {
				s := c.Get(storageID, spaceID, shareID)
				Expect(s).To(BeNil())

				c.Providers[storageID] = &providercache.Spaces{
					Spaces: map[string]*providercache.Shares{
						spaceID: {
							Mtime: time.Now(),
						},
					},
				}
				Expect(c.Sync(ctx, storageID, spaceID)).To(Succeed()) // Sync from disk won't happen because in-memory mtime is later than on disk

				s = c.Get(storageID, spaceID, shareID)
				Expect(s).To(BeNil())
			})
		})
	})
})