package symbol

import (
	"fmt"
	"reflect"

	. "github.com/kocircuit/kocircuit/lang/circuit/eval"
	. "github.com/kocircuit/kocircuit/lang/circuit/model"
	. "github.com/kocircuit/kocircuit/lang/go/kit/tree"
)

type BlobSymbol struct {
	Value reflect.Value `ko:"name=value"` // []byte
}

func (blob *BlobSymbol) Bytes() []byte {
	return blob.Value.Bytes()
}

func (blob *BlobSymbol) String() string {
	return fmt.Sprintf("Blob<%d>", blob.Value.Len())
}

func (blob *BlobSymbol) Equal(sym Symbol) bool {
	if other, ok := sym.(*BlobSymbol); ok {
		left, right := blob.Bytes(), other.Bytes()
		if len(left) != len(right) {
			return false
		} else {
			for i := range left {
				if left[i] != right[i] {
					return false
				}
			}
			return true
		}
	} else {
		return false
	}
}

func (blob *BlobSymbol) Splay() Tree {
	return NoQuote{blob.String()}
}

func (blob *BlobSymbol) Hash() string {
	return BytesID(blob.Bytes()).String()
}

func (blob *BlobSymbol) Type() Type {
	return BlobType{}
}

func (blob *BlobSymbol) Select(span *Span, path Path) (Shape, Effect, error) {
	if len(path) == 0 {
		return blob, nil, nil
	} else {
		return nil, nil, span.Errorf(nil, "blob value %v cannot be selected into", blob)
	}
}

func (blob *BlobSymbol) Augment(span *Span, _ Knot) (Shape, Effect, error) {
	return nil, nil, span.Errorf(nil, "blob value %v cannot be augmented", blob)
}

func (blob *BlobSymbol) Invoke(span *Span) (Shape, Effect, error) {
	return nil, nil, span.Errorf(nil, "blob value %v cannot be invoked", blob)
}

type BlobType struct{}

func (blob BlobType) IsType() {}

func (blob BlobType) String() string {
	return "Blob"
}

func (blob BlobType) Splay() Tree {
	return NoQuote{"Blob"}
}