package composite

import "testing"

func TestComposite(t *testing.T) {
	root := Composite{Name: "服装"}
	c1 := Composite{Name: "男装"}
	c2 := Composite{Name: "女装"}

	leaf1 := Leaf{Name: "衬衣"}
	leaf2 := Leaf{Name: "夹克"}
	leaf3 := Leaf{Name: "裙子"}
	leaf4 := Leaf{Name: "套装"}

	root.AddChild(&c1)
	root.AddChild(&c2)
	c1.AddChild(&leaf1)
	c1.AddChild(&leaf2)
	c2.AddChild(&leaf3)
	c2.AddChild(&leaf4)

	root.Print("")
}
