package ulid

import (
	"crypto/rand"
	"sync"
	"time"

	pkgulid "github.com/oklog/ulid/v2"
)

var (
	entropy    = pkgulid.Monotonic(rand.Reader, 0)
	entropyMtx sync.Mutex
	nowfunc    = time.Now
)

func MustNew() string {
	entropyMtx.Lock()
	defer entropyMtx.Unlock()
	newULID, err := pkgulid.New(pkgulid.Timestamp(nowfunc()), entropy)
	if err != nil {
		panic(err)
	}
	return newULID.String()
}
