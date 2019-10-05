package heap

// 代表java对象实例
type Object struct {
	class *Class
	fields Slots  //对象的field通过 Field.slotId 索引到此次表项中的实例数据
}


func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

// getters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.fields
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

