package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raggaer/pigo"
    "github.com/Cloakaac/cloak/util"
	"net/http"
    "time"
)

const (
    sandbox = "sandbox"
    live = "live"
)

var (
    baseURL string
    paypalToken *util.PaypalToken
)

type ShopController struct {
	*pigo.Controller
}

func init() {
    if pigo.Config.Key("paypal").String("mode") == sandbox {
        baseURL = "https://api.sandbox.paypal.com"
    } else {
        baseURL = "https://api.paypal.com"
    }
}

// Paypal shows the paypal buypoints page
func (base *ShopController) Paypal(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    base.Data["Errors"] = base.Session.GetFlashes("errors")
    base.Data["Success"] = base.Session.GetFlashes("success")
    base.Data["Points"] = pigo.Config.Key("paypal").Key("payment").Float("points")
    base.Data["Min"] = pigo.Config.Key("paypal").Key("payment").Float("min")
    base.Data["Max"] = pigo.Config.Key("paypal").Key("payment").Float("max")
    base.Data["Promo"] = pigo.Config.Key("paypal").Float("promo")
    base.Template = "paypal.html"
}

// PaypalPay process a paypal buypoints request
func (base *ShopController) PaypalPay(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    timeNow := time.Now().Unix()
    if paypalToken == nil || (timeNow + paypalToken.ExpiresIn) < timeNow {
        token, err := util.GetPaypalToken(baseURL, pigo.Config.Key("paypal").String("public"), pigo.Config.Key("paypal").String("private"))
        if err != nil {
            base.Error = err.Error()
            return
        }
        paypalToken = token
    }
}