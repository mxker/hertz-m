// Code generated by hertz generator. DO NOT EDIT.

package Person

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	person "hertz-m/biz/handler/person"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.POST("/list", append(_personlistMw(), person.PersonList)...)
	{
		_person_info := root.Group("/person_info", _person_infoMw()...)
		_person_info.POST("/:age", append(_personinfoMw(), person.PersonInfo)...)
	}
}
