package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["argo_api/controllers:ActainicioController"] = append(beego.GlobalControllerRouter["argo_api/controllers:ActainicioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:ActainicioController"] = append(beego.GlobalControllerRouter["argo_api/controllers:ActainicioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:ActainicioController"] = append(beego.GlobalControllerRouter["argo_api/controllers:ActainicioController"],
		beego.ControllerComments{
			Method: "GetCV",
			Router: `/:contrato/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:ActainicioController"] = append(beego.GlobalControllerRouter["argo_api/controllers:ActainicioController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:ActainicioController"] = append(beego.GlobalControllerRouter["argo_api/controllers:ActainicioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:ActainicioController"] = append(beego.GlobalControllerRouter["argo_api/controllers:ActainicioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:NovedadController"] = append(beego.GlobalControllerRouter["argo_api/controllers:NovedadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:NovedadController"] = append(beego.GlobalControllerRouter["argo_api/controllers:NovedadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:NovedadController"] = append(beego.GlobalControllerRouter["argo_api/controllers:NovedadController"],
		beego.ControllerComments{
			Method: "GetCV",
			Router: `/:contrato/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:NovedadController"] = append(beego.GlobalControllerRouter["argo_api/controllers:NovedadController"],
		beego.ControllerComments{
			Method: "GetCVT",
			Router: `/:contrato/:vigencia/:tiponovedad`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:NovedadController"] = append(beego.GlobalControllerRouter["argo_api/controllers:NovedadController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:NovedadController"] = append(beego.GlobalControllerRouter["argo_api/controllers:NovedadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:NovedadController"] = append(beego.GlobalControllerRouter["argo_api/controllers:NovedadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:PlantillaDocumentoController"] = append(beego.GlobalControllerRouter["argo_api/controllers:PlantillaDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:PlantillaDocumentoController"] = append(beego.GlobalControllerRouter["argo_api/controllers:PlantillaDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:PlantillaDocumentoController"] = append(beego.GlobalControllerRouter["argo_api/controllers:PlantillaDocumentoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:PlantillaDocumentoController"] = append(beego.GlobalControllerRouter["argo_api/controllers:PlantillaDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:PlantillaDocumentoController"] = append(beego.GlobalControllerRouter["argo_api/controllers:PlantillaDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:TiponovedadController"] = append(beego.GlobalControllerRouter["argo_api/controllers:TiponovedadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:TiponovedadController"] = append(beego.GlobalControllerRouter["argo_api/controllers:TiponovedadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:TiponovedadController"] = append(beego.GlobalControllerRouter["argo_api/controllers:TiponovedadController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:TiponovedadController"] = append(beego.GlobalControllerRouter["argo_api/controllers:TiponovedadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["argo_api/controllers:TiponovedadController"] = append(beego.GlobalControllerRouter["argo_api/controllers:TiponovedadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

}
