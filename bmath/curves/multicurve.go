package curves

import (
	"github.com/wieku/danser-go/bmath"
	"sort"
)

const minPartWidth = 0.0001

type MultiCurve struct {
	sections   []float32
	lines      []Linear
	length     float32
	firstPoint bmath.Vector2f
}

func NewMultiCurve(typ string, points []bmath.Vector2f, desiredLength float64) *MultiCurve {
	lines := make([]Linear, 0)

	if len(points) < 3 {
		typ = "L"
	}

	switch typ {
	case "P":
		lines = append(lines, ApproximateCircularArc(points[0], points[1], points[2], 0.125)...)
		break
	case "L":
		for i := 0; i < len(points)-1; i++ {
			lines = append(lines, NewLinear(points[i], points[i+1]))
		}
		break
	case "B":
		lastIndex := 0
		for i, p := range points {
			if (i == len(points)-1 && p != points[i-1]) || (i < len(points)-1 && points[i+1] == p) {
				pts := points[lastIndex : i+1]

				if len(pts) > 2 {
					lines = append(lines, ApproximateBezier(pts)...)
				} else if len(pts) == 2 {
					lines = append(lines, NewLinear(pts[0], pts[1]))
				}

				lastIndex = i + 1
			}
		}
		break
	case "C":

		if points[0] != points[1] {
			points = append([]bmath.Vector2f{points[0]}, points...)
		}

		if points[len(points)-1] != points[len(points)-2] {
			points = append(points, points[len(points)-1])
		}

		for i := 0; i < len(points)-3; i++ {
			lines = append(lines, ApproximateCatmullRom(points[i:i+4], 50)...)
		}
		break
	}

	length := float32(0.0)

	for _, l := range lines {
		length += l.GetLength()
	}

	firstPoint := points[0]

	diff := float64(length) - desiredLength

	for len(lines) > 0 {
		line := lines[len(lines)-1]

		if float64(line.GetLength()) > diff+minPartWidth {
			pt := line.PointAt((line.GetLength() - float32(diff)) / line.GetLength())
			lines[len(lines)-1] = NewLinear(line.Point1, pt)
			break
		}

		diff -= float64(line.GetLength())
		lines = lines[:len(lines)-1]
	}

	length = 0.0

	for _, l := range lines {
		length += l.GetLength()
	}

	sections := make([]float32, len(lines)+1)
	sections[0] = 0.0
	prev := float32(0.0)

	for i := 0; i < len(lines); i++ {
		prev += lines[i].GetLength()
		sections[i+1] = prev
	}

	return &MultiCurve{sections, lines, length, firstPoint}
}

func (mCurve *MultiCurve) PointAt(t float32) bmath.Vector2f {
	if len(mCurve.lines) == 0 {
		return mCurve.firstPoint
	}

	desiredWidth := mCurve.length * bmath.ClampF32(t, 0.0, 1.0)

	withoutFirst := mCurve.sections[1:]
	index := sort.Search(len(withoutFirst), func(i int) bool {
		return withoutFirst[i] >= desiredWidth
	})

	//log.Println(len(mCurve.lines), desiredWidth, mCurve.length, index)

	return mCurve.lines[index].PointAt((desiredWidth - mCurve.sections[index]) / (mCurve.sections[index+1] - mCurve.sections[index]))
}

func (mCurve *MultiCurve) GetLength() float32 {
	return mCurve.length
}

func (mCurve *MultiCurve) GetStartAngle() float32 {
	if len(mCurve.lines) > 0 {
		return mCurve.lines[0].GetStartAngle()
	}
	return 0.0
}

func (mCurve *MultiCurve) GetEndAngle() float32 {
	if len(mCurve.lines) > 0 {
		return mCurve.lines[len(mCurve.lines)-1].GetEndAngle()
	}
	return 0.0
}

func (mCurve *MultiCurve) GetLines() []Linear {
	return mCurve.lines
}
