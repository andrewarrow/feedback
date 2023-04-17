package router

import (
	"testing"
)

func TestDecorcate(t *testing.T) {
	list := []MSA{}
	object1 := MSA{"foo": "boo", "foo_id": 456}
	item1 := MSA{"thing_id": 123, "something": "else", "object": object1}
	list = append(list, item1)
	object2 := MSA{"foo": "boo", "foo_id": 789}
	item2 := MSA{"thing_id": 999, "something": "else", "object": object2}
	list = append(list, item2)

	idMap := gatherDecorateIds(list, 0)

	if len(idMap) != 2 {
		t.Fatalf("%d", len(idMap))
	}

}
