package composite

import "fmt"

type Component interface {
	AddChild(child Component)
	RemoveChild(child Component)
	GetChildren(index int) *Component
	Print(preStr string)
}

type Composite struct {
	Name            string
	ChildComponents []Component
}

func (c *Composite) AddChild(child Component) {
	if c.ChildComponents == nil {
		c.ChildComponents = make([]Component, 0)
	}
	c.ChildComponents = append(c.ChildComponents, child)
}

func (c *Composite) RemoveChild(child Component) {
	for i, component := range c.ChildComponents {
		if component == child {
			c.ChildComponents = append(c.ChildComponents[:i], c.ChildComponents[i+1:]...)
			break
		}
	}
}

func (c *Composite) GetChildren(index int) *Component {
	if index >= 0 && index < len(c.ChildComponents) {
		return &c.ChildComponents[index]
	}
	return nil
}

func (c *Composite) Print(preStr string) {
	fmt.Println(preStr + c.Name)
	for _, component := range c.ChildComponents {
		component.Print(preStr + " ")
	}
}

type Leaf struct {
	Name string
}

func (leaf *Leaf) AddChild(_ Component) {
	panic(fmt.Errorf("leaf can't add child"))
}

func (leaf *Leaf) RemoveChild(_ Component) {
	panic(fmt.Errorf("leaf can't remove child"))
}

func (leaf *Leaf) GetChildren(_ int) *Component {
	panic(fmt.Errorf("leaf can't get children"))
}

func (leaf *Leaf) Print(preStr string) {
	fmt.Println(preStr + leaf.Name)
}
