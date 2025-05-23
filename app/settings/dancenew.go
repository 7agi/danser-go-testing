package settings

var CursorDance = initCursorDance()

func initCursorDance() *cursorDance {
	return &cursorDance{
		Movers: []*mover{
			DefaultsFactory.InitMover(),
		},
		Spinners: []*spinner{
			DefaultsFactory.InitSpinner(),
		},
		ComboTag:           false,
		Battle:             false,
		DoSpinnersTogether: true,
		TAGSliderDance:     false,
		MoverSettings: &moverSettings{
			Bezier: []*bezier{
				DefaultsFactory.InitBezier(),
			},
			Flower: []*flower{
				DefaultsFactory.InitFlower(),
			},
			HalfCircle: []*circular{
				DefaultsFactory.InitCircular(),
			},
			Spline: []*spline{
				DefaultsFactory.InitSpline(),
			},
			Momentum: []*momentum{
				DefaultsFactory.InitMomentum(),
			},
			ExGon: []*exgon{
				DefaultsFactory.InitExGon(),
			},
			Linear: []*linear{
				DefaultsFactory.InitLinear(),
			},
			Pippi: []*pippi{
				DefaultsFactory.InitPippi(),
			},
		},
	}
}

type mover struct {
	Mover             string `combo:"spline,bezier,circular,linear,axis,aggressive,flower,momentum,exgon,pippi"`
	SliderDance       bool
	RandomSliderDance bool
}

func (d *defaultsFactory) InitMover() *mover {
	return &mover{
		Mover:             "spline",
		SliderDance:       false,
		RandomSliderDance: false,
	}
}

type spinner struct {
	Mover         string  `combo:"heart,triangle,square,cube,circle"`
	centerOffset  string  `vector:"true" left:"CenterOffsetX" right:"CenterOffsetY"`
	CenterOffsetX float64 `min:"-1000" max:"1000"`
	CenterOffsetY float64 `min:"-1000" max:"1000"`
	Radius        float64 `max:"200" format:"%.0fo!px"`
}

func (d *defaultsFactory) InitSpinner() *spinner {
	return &spinner{
		Mover:  "circle",
		Radius: 100,
	}
}

type cursorDance struct {
	Movers             []*mover   `new:"InitMover" wiki:"Help|https://github.com/Wieku/danser-go/wiki/Movers#available-movers"`
	Spinners           []*spinner `new:"InitSpinner" wiki:"Help|https://github.com/Wieku/danser-go/wiki/Movers#available-spinner-movers"`
	ComboTag           bool       `liveedit:"false"`
	Battle             bool       `liveedit:"false"`
	DoSpinnersTogether bool       `liveedit:"false"`
	TAGSliderDance     bool       `label:"TAG slider dance" liveedit:"false"`
	MoverSettings      *moverSettings
}

type moverSettings struct {
	Bezier     []*bezier   `new:"InitBezier"`
	Flower     []*flower   `new:"InitFlower"`
	HalfCircle []*circular `new:"InitCircular"`
	Spline     []*spline   `new:"InitSpline"`
	Momentum   []*momentum `new:"InitMomentum"`
	ExGon      []*exgon    `new:"InitExGon"`
	Linear     []*linear   `new:"InitLinear"`
	Pippi      []*pippi    `new:"InitPippi"`
}
