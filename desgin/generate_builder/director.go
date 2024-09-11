package main

// 主管的作用是组织建造过程。(可选)
type director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func (d *director) dismantleHouse() House {
	// 拆除步骤 1 2 3
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
