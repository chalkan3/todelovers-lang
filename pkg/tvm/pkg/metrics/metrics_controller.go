package metrics

type MetricsController interface {
	GorotinesManger() GorotineManager
}

type metricsController struct {
	gm GorotineManager
}

func (mc *metricsController) GorotinesManger() GorotineManager {
	return mc.gm
}

func NewMetricsController() MetricsController {
	return &metricsController{
		gm: NewGorotineManager(),
	}
}
