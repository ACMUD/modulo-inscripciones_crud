package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:CorreoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:CorreoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:CorreoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:CorreoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:CorreoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:CorreoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:CorreoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:CorreoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:CorreoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:CorreoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:DiaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:DiaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:DiaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:DiaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:DiaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:DiaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:DiaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:DiaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:DiaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:DiaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoUsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoUsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoUsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoUsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:EstadoUsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FechasController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FechasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FechasController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FechasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FechasController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FechasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FechasController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FechasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FechasController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FechasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FranjaHorariaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FranjaHorariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FranjaHorariaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FranjaHorariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FranjaHorariaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FranjaHorariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FranjaHorariaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FranjaHorariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FranjaHorariaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:FranjaHorariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:GrupoInteresController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:GrupoInteresController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:GrupoInteresController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:GrupoInteresController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:GrupoInteresController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:GrupoInteresController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:GrupoInteresController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:GrupoInteresController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:GrupoInteresController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:GrupoInteresController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioPersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioPersonaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioPersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioPersonaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioPersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioPersonaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioPersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioPersonaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioPersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:HorarioPersonaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:IdentificacionController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:IdentificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:IdentificacionController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:IdentificacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:IdentificacionController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:IdentificacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:IdentificacionController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:IdentificacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:IdentificacionController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:IdentificacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PeriodoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PeriodoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PeriodoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PeriodoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PeriodoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PeriodoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PeriodoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PeriodoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PeriodoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PeriodoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolPersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolPersonaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolPersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolPersonaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolPersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolPersonaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolPersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolPersonaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolPersonaController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:RolPersonaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SalonController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SalonController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SalonController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SalonController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SalonController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SalonController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SalonController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SalonController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SalonController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SalonController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SedeController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SedeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SedeController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SedeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SedeController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SedeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SedeController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SedeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SedeController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:SedeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoCorreoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoCorreoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoCorreoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoCorreoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoCorreoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoCorreoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoCorreoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoCorreoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoCorreoController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoCorreoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoIdentificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoIdentificacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoIdentificacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoIdentificacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/ACMUD/modulo-inscripciones_crud/controllers:TipoIdentificacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
