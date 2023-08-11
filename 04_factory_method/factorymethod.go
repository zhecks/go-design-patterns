package factorymethod

type ExportOperateApi interface {
	export() string
	father() string
}

type OperatorFactory interface {
	create() ExportOperateApi
}

type BaseExportOperate struct {
}

func (eo *BaseExportOperate) father() string {
	return "father"
}

type ExportTxtOperate struct {
	*BaseExportOperate
}

func (eo *ExportTxtOperate) export() string {
	return "txt export operate"
}

type ExportTxtFactory struct {
}

func (etf *ExportTxtFactory) create() ExportOperateApi {
	return &ExportTxtOperate{
		&BaseExportOperate{},
	}
}

type ExportTomlOperate struct {
	*BaseExportOperate
}

func (eo *ExportTomlOperate) export() string {
	return "toml export operate"
}

type ExportTomlFactory struct {
}

func (etf *ExportTomlFactory) create() ExportOperateApi {
	return &ExportTomlOperate{
		&BaseExportOperate{},
	}
}
