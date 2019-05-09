// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/ACMUD/modulo-inscripciones_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/estado",
			beego.NSInclude(
				&controllers.EstadoController{},
			),
		),

		beego.NSNamespace("/fechas",
			beego.NSInclude(
				&controllers.FechasController{},
			),
		),

		beego.NSNamespace("/correo",
			beego.NSInclude(
				&controllers.CorreoController{},
			),
		),

		beego.NSNamespace("/estado_usuario",
			beego.NSInclude(
				&controllers.EstadoUsuarioController{},
			),
		),

		beego.NSNamespace("/periodo",
			beego.NSInclude(
				&controllers.PeriodoController{},
			),
		),

		beego.NSNamespace("/grupo_interes",
			beego.NSInclude(
				&controllers.GrupoInteresController{},
			),
		),

		beego.NSNamespace("/sede",
			beego.NSInclude(
				&controllers.SedeController{},
			),
		),

		beego.NSNamespace("/persona",
			beego.NSInclude(
				&controllers.PersonaController{},
			),
		),

		beego.NSNamespace("/rol",
			beego.NSInclude(
				&controllers.RolController{},
			),
		),

		beego.NSNamespace("/horario_persona",
			beego.NSInclude(
				&controllers.HorarioPersonaController{},
			),
		),

		beego.NSNamespace("/tipo_correo",
			beego.NSInclude(
				&controllers.TipoCorreoController{},
			),
		),

		beego.NSNamespace("/identificacion",
			beego.NSInclude(
				&controllers.IdentificacionController{},
			),
		),

		beego.NSNamespace("/salon",
			beego.NSInclude(
				&controllers.SalonController{},
			),
		),

		beego.NSNamespace("/dia",
			beego.NSInclude(
				&controllers.DiaController{},
			),
		),

		beego.NSNamespace("/horario",
			beego.NSInclude(
				&controllers.HorarioController{},
			),
		),

		beego.NSNamespace("/franja_horaria",
			beego.NSInclude(
				&controllers.FranjaHorariaController{},
			),
		),

		beego.NSNamespace("/rol_persona",
			beego.NSInclude(
				&controllers.RolPersonaController{},
			),
		),

		beego.NSNamespace("/tipo_identificacion",
			beego.NSInclude(
				&controllers.TipoIdentificacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
