package service

import (
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
	"html/template"
	"os"
)

// @title    CreateTemp
// @description   函数的详细描述
// @auth                     （2020/04/05  20:22 ）
// @return    err             error
func CreateTemp(a model.AutoCodeStruct) (err error) {
	basePath := "./resource/template"
	modelTmpl, err := template.ParseFiles(basePath + "/te/model.go.tpl")
	if err != nil {
		return err
	}
	apiTmpl, err := template.ParseFiles(basePath + "/te/api.go.tpl")
	if err != nil {
		return err
	}
	routerTmpl, err := template.ParseFiles(basePath + "/te/router.go.tpl")
	if err != nil {
		return err
	}
	feapiTmpl, err := template.ParseFiles(basePath + "/fe/api.js.tpl")
	if err != nil {
		return err
	}
	readmeTmpl, err := template.ParseFiles(basePath + "/readme.txt.tpl")
	if err != nil {
		return err
	}
	//自动化总目录
	_autoCode := "./autoCode/"
	//自动化后台代码目录
	_te := "./autoCode/te/"
	_dir := _te + a.PackageName
	_modeldir := _te + a.PackageName + "/model"
	_apidir := _te + a.PackageName + "/api"
	_routerdir := _te + a.PackageName + "/router"
	//自动化前台代码目录
	_fe := "./autoCode/fe/"
	_fe_dir := _fe + a.PackageName
	_fe_apidir := _fe + a.PackageName + "/api"
	err = utils.CreateDir(_autoCode, _te, _dir, _modeldir, _apidir, _routerdir, _fe, _fe_dir, _fe_apidir)
	if err != nil {
		return err
	}
	model, err := os.OpenFile(_te+a.PackageName+"/model/model.go", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	api, err := os.OpenFile(_te+a.PackageName+"/api/api.go", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	router, err := os.OpenFile(_te+a.PackageName+"/router/router.go", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	feapi, err := os.OpenFile(_fe+a.PackageName+"/api/api.js", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	readme, err := os.OpenFile(_autoCode+"readme.txt", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	// 生成代码
	{
		err = modelTmpl.Execute(model, a)
		if err != nil {
			return err
		}
		err = apiTmpl.Execute(api, a)
		if err != nil {
			return err
		}
		err = routerTmpl.Execute(router, a)
		if err != nil {
			return err
		}
		err = feapiTmpl.Execute(feapi, a)
		if err != nil {
			return err
		}
		err = readmeTmpl.Execute(readme, a)
		if err != nil {
			return err
		}
	}
	_ = model.Close()
	_ = api.Close()
	_ = router.Close()
	_ = feapi.Close()
	_ = readme.Close()
	fileList := []string{
		_te + a.PackageName + "/model/model.go",
		_te + a.PackageName + "/api/api.go",
		_te + a.PackageName + "/router/router.go",
		_fe + a.PackageName + "/api/api.js",
		_autoCode + "readme.txt",
	}
	err = utils.ZipFiles("./ginvueadmin.zip", fileList, ".", ".")
	if err != nil {
		return err
	}
	err = os.RemoveAll(_autoCode)
	if err != nil {
		return err
	}
	return nil
}
