package abstractfactory

type AbstractFactory interface {
	createProductA() string
	createProductB() string
}

type PCFactory struct {
}

func (factory *PCFactory) createProductA() string {
	return "PC productA"
}

func (factory *PCFactory) createProductB() string {
	return "PC productB"
}

type PhoneFactory struct {
}

func (factory *PhoneFactory) createProductA() string {
	return "Phone productA"
}

func (factory *PhoneFactory) createProductB() string {
	return "Phone productB"
}
