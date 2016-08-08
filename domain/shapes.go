package domain

type View []*Block

type Shape struct {
	views     []View
	viewIndex int
	visible   bool
}

func (s *Shape) Init() {
	s.viewIndex = 0
	s.visible = false
}

func (s *Shape) Rotate() {
	s.viewIndex += 1
	if s.viewIndex > len(s.views)-1 {
		s.viewIndex = 0
	}
}

func (s *Shape) RotateBack() {
	s.viewIndex -= 1
	if s.viewIndex < 0 {
		s.viewIndex = len(s.views) - 1
	}
}

func (s *Shape) GetBlocks() []*Block {
	return s.views[s.viewIndex]
}

func SquareShape(colour BlockColour) *Shape {

	return &Shape{
		viewIndex: 0,
		visible:   false,
		views: []View{
			View{
				&Block{X: 0, Y: 0, Colour: colour},
				&Block{X: 1, Y: 0, Colour: colour},
				&Block{X: 0, Y: 1, Colour: colour},
				&Block{X: 1, Y: 1, Colour: colour},
			},
		},
	}
}

func BarShape(colour BlockColour) *Shape {

	return &Shape{
		viewIndex: 0,
		visible:   false,
		views: []View{
			View{
				&Block{X: 0, Y: 0, Colour: colour},
				&Block{X: 1, Y: 0, Colour: colour},
				&Block{X: 2, Y: 0, Colour: colour},
				&Block{X: 3, Y: 0, Colour: colour},
			},
			View{
				&Block{X: 2, Y: 0, Colour: colour},
				&Block{X: 2, Y: 1, Colour: colour},
				&Block{X: 2, Y: 2, Colour: colour},
				&Block{X: 2, Y: 3, Colour: colour},
			},
		},
	}
}

func LeftLShape(colour BlockColour) *Shape {

	return &Shape{
		viewIndex: 0,
		visible:   false,
		views: []View{
			View{
				&Block{X: 0, Y: 0, Colour: colour},
				&Block{X: 1, Y: 0, Colour: colour},
				&Block{X: 2, Y: 0, Colour: colour},
				&Block{X: 2, Y: 1, Colour: colour},
			},
			View{
				&Block{X: 2, Y: 0, Colour: colour},
				&Block{X: 2, Y: 1, Colour: colour},
				&Block{X: 2, Y: 2, Colour: colour},
				&Block{X: 1, Y: 2, Colour: colour},
			},
			View{
				&Block{X: 0, Y: 2, Colour: colour},
				&Block{X: 1, Y: 2, Colour: colour},
				&Block{X: 2, Y: 2, Colour: colour},
				&Block{X: 0, Y: 1, Colour: colour},
			},
			View{
				&Block{X: 0, Y: 0, Colour: colour},
				&Block{X: 0, Y: 1, Colour: colour},
				&Block{X: 0, Y: 2, Colour: colour},
				&Block{X: 1, Y: 0, Colour: colour},
			},
		},
	}
}

func RightLShape(colour BlockColour) *Shape {

	return &Shape{
		viewIndex: 0,
		visible:   false,
		views: []View{
			View{
				&Block{X: 0, Y: 0, Colour: colour},
				&Block{X: 1, Y: 0, Colour: colour},
				&Block{X: 2, Y: 0, Colour: colour},
				&Block{X: 0, Y: 1, Colour: colour},
			},
			View{
				&Block{X: 2, Y: 0, Colour: colour},
				&Block{X: 2, Y: 1, Colour: colour},
				&Block{X: 2, Y: 2, Colour: colour},
				&Block{X: 1, Y: 0, Colour: colour},
			},
			View{
				&Block{X: 0, Y: 2, Colour: colour},
				&Block{X: 1, Y: 2, Colour: colour},
				&Block{X: 2, Y: 2, Colour: colour},
				&Block{X: 2, Y: 1, Colour: colour},
			},
			View{
				&Block{X: 0, Y: 0, Colour: colour},
				&Block{X: 0, Y: 1, Colour: colour},
				&Block{X: 0, Y: 2, Colour: colour},
				&Block{X: 1, Y: 2, Colour: colour},
			},
		},
	}
}

func LeftStepShape(colour BlockColour) *Shape {

	return &Shape{
		viewIndex: 0,
		visible:   false,
		views: []View{
			View{
				&Block{X: 0, Y: 1, Colour: colour},
				&Block{X: 1, Y: 1, Colour: colour},
				&Block{X: 1, Y: 0, Colour: colour},
				&Block{X: 2, Y: 0, Colour: colour},
			},
			View{
				&Block{X: 0, Y: 0, Colour: colour},
				&Block{X: 0, Y: 1, Colour: colour},
				&Block{X: 1, Y: 1, Colour: colour},
				&Block{X: 1, Y: 2, Colour: colour},
			},
		},
	}
}

func RightStepShape(colour BlockColour) *Shape {

	return &Shape{
		viewIndex: 0,
		visible:   false,
		views: []View{
			View{
				&Block{X: 0, Y: 0, Colour: colour},
				&Block{X: 1, Y: 0, Colour: colour},
				&Block{X: 1, Y: 1, Colour: colour},
				&Block{X: 2, Y: 1, Colour: colour},
			},
			View{
				&Block{X: 0, Y: 2, Colour: colour},
				&Block{X: 0, Y: 1, Colour: colour},
				&Block{X: 1, Y: 1, Colour: colour},
				&Block{X: 1, Y: 0, Colour: colour},
			},
		},
	}
}
