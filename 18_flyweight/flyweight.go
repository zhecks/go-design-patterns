package flyweight

import "fmt"

type ImageFlyweightFactory struct {
	images map[string]*ImageFlyweight
}

var imageFactory *ImageFlyweightFactory

func GetInstance() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			images: map[string]*ImageFlyweight{},
		}
	}
	return imageFactory
}

func (factory *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image, ok := factory.images[filename]
	if !ok {
		image = NewImageFlyweight(filename)
		factory.images[filename] = image
	}
	return image
}

type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

func (i *ImageFlyweight) Data() string {
	return i.data
}
