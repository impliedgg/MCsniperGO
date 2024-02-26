package claimer

import (
	"github.com/Kqzz/MCsniperGO/pkg/mc"
	"github.com/valyala/fasthttp"
)

func (claim *Claim) SendRequest(account *mc.MCaccount, client *fasthttp.Client) {
	var statusCode int
	var failType mc.FailType
	var err error
	switch account.Type {
	case mc.MsPr:
		statusCode, failType, err = account.CreateProfile(claim.Username, client)
	case mc.MsGp:
		statusCode, failType, err = account.CreateProfile(claim.Username, client)
	case mc.Ms:
		statusCode, failType, err = account.ChangeUsername(claim.Username, client)
	}

	claim.Claimer.respchan <- ClaimResponse{statusCode, failType, err}
}
