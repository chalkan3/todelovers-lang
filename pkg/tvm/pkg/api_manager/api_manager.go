package apimanager

import (
	"log"
	cpapi "mary_guica/pkg/tvm/internal/api/control_plane"
	eapi "mary_guica/pkg/tvm/internal/api/events"
	control "mary_guica/pkg/tvm/pkg/control_plane"
)

type APIManager interface {
	Launch()
	KeepAlive()
}
type apiManager struct {
	eventsAPI       eapi.API
	controlPlaneAPI cpapi.API
}

func NewAPIManager(cp control.ControlPlane) APIManager {
	return &apiManager{
		eventsAPI:       eapi.NewAPI(),
		controlPlaneAPI: cpapi.NewAPI(cp),
	}
}

func (a *apiManager) Launch() {
	log.Println("[INFO][VM][API] - Lauching events API")
	a.eventsAPI.Serve()
	log.Println("[INFO][VM][API] - Lauching control plane API")
	a.controlPlaneAPI.Serve()

}

func (a *apiManager) KeepAlive() {

}
