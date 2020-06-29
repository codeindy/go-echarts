package charts

import (
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// Surface3D represents a 3D surface chart.
type Surface3D struct {
	Chart3D
}

func (Surface3D) Type() string { return types.ChartSurface3D }

// NewSurface3D creates a new 3d surface chart.
func NewSurface3D() *Surface3D {
	c := &Surface3D{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.initChart3D()
	return c
}

// AddXYAxis adds both the X axis and the Y axis.
func (c *Surface3D) AddXYAxis(xAxis, yAxis interface{}) *Surface3D {
	c.xData = xAxis
	c.yData = yAxis
	return c
}

// AddZAxis adds the Z axis.
func (c *Surface3D) AddZAxis(name string, zAxis interface{}, opts ...SeriesOpts) *Surface3D {
	c.addZAxis(types.ChartSurface3D, name, zAxis, opts...)
	return c
}
