package md5

import (
	"encoding/hex"
	. "github.com/Bitnick2002/goa/jasmine"
	"testing"
)

func Test_(t *testing.T) {
	Describe("md5 iteration", func() {
		It("should iterate sum ", func() {

		})
		Describe("When 'Message' iterate 1 times", func() {
			It("should be 4c2a8fe7eaf24721cc7a9f0175115bd4", func() {
				result := IteraionSum([]byte("Message"), 1)
				Expect(hex.EncodeToString(result)).ToBe("4c2a8fe7eaf24721cc7a9f0175115bd4")
			})
		})
		Describe("When 'Message' iterate 2 times", func() {
			It("should be 35d04ebbe271c1e202b5e77385f97346", func() {
				result := IteraionSum([]byte("Message"), 2)
				Expect(hex.EncodeToString(result)).ToBe("35d04ebbe271c1e202b5e77385f97346")
			})
		})
		Describe("When 'Message' iterate 3 times", func() {
			It("should be 9bc95ee7034560bb79d891c625ab48ce", func() {
				result := IteraionSum([]byte("Message"), 3)
				Expect(hex.EncodeToString(result)).ToBe("9bc95ee7034560bb79d891c625ab48ce")
			})
		})
	})
}
