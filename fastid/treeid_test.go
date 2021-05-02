package treeid

import (
	"testing"
)


func TestArrayTree_Alloc(t *testing.T) {
	tree := NewArrayTree(3)

	for n := 0; n <= 3; n++ {
		id, err := tree.Alloc()
		if err != nil {
			t.Errorf("unexpected err: %v", err)
		}
		t.Logf("allocated id = %d \n", id)
		if id != n {
			t.Errorf("id returned (%d) does not match exepected (%d) \n", id, n)
		}
	}
	_, err := tree.Alloc()
	t.Logf("expected error: %v", err)

	t.Log("freeing ids 1 and 3")
	err = tree.Free(1)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	err = tree.Free(3)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	id, err := tree.Alloc()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	t.Logf("got id = %d", id)
	id, err = tree.Alloc()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	t.Logf("got id = %d", id)

}
