package ulid

import (
	"crypto/rand"
	"fmt"
	"sync"
	"time"
	"unsafe"

	pkgulid "github.com/oklog/ulid/v2"
)

var (
	entropy    = pkgulid.Monotonic(rand.Reader, 0)
	entropyMtx sync.Mutex
	nowfunc    = time.Now
)

type ULID []byte

func MustNew() ULID {
	entropyMtx.Lock()
	defer entropyMtx.Unlock()
	newULID, err := pkgulid.New(pkgulid.Timestamp(nowfunc()), entropy)
	if err != nil {
		panic(err)
	}
	return newULID[:]
}

func Parse(idStr string) (ULID, error) {
	ulid, err := pkgulid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ulid: %w", err)
	}
	return ulid[:], nil
}

// String implements Stringer interface
func (id ULID) String() string {
	var newID pkgulid.ULID
	_ = newID.UnmarshalBinary(id)
	return newID.String()
}

// MarshalJSON implements encoding.MarshalJSON interface
func (id ULID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + id.String() + `"`), nil
}

// UnmarshalJSON implements encoding.UnmarshalJSON interface
func (id *ULID) UnmarshalJSON(in []byte) error {
	// Remove double quotes at the beginning and the end
	in = in[1 : len(in)-1]
	parsed, err := Parse(*(*string)(unsafe.Pointer(&in)))
	if err != nil {
		return err
	}

	*id = parsed
	return nil
}

// TODO: Implement basic interfaces
// encoding.BinaryMarshaler, encoding.BinaryUnmarshaler, encoding.TextMarshaler, encoding.TextUnmarshaler, sql.Scanner, sql/driver.Valuer
