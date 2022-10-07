package table

import (
	"github.com/yaoapp/yao/widgets/action"
	"github.com/yaoapp/yao/widgets/component"
	"github.com/yaoapp/yao/widgets/field"
	"github.com/yaoapp/yao/widgets/hook"
)

// DSL the table DSL
type DSL struct {
	Root        string              `json:"-"`
	ID          string              `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Action      *ActionDSL          `json:"action"`
	Layout      *LayoutDSL          `json:"layout"`
	Fields      *FieldsDSL          `json:"fields"`
	ComputesIn  field.ComputeFields `json:"-"`
	ComputesOut field.ComputeFields `json:"-"`
	CProps      field.CloudProps    `json:"-"`
}

// ActionDSL the table action DSL
type ActionDSL struct {
	Bind              *BindActionDSL  `json:"bind,omitempty"`
	Setting           *action.Process `json:"setting,omitempty"`
	Component         *action.Process `json:"component,omitempty"`
	Search            *action.Process `json:"search,omitempty"`
	Get               *action.Process `json:"get,omitempty"`
	Find              *action.Process `json:"find,omitempty"`
	Save              *action.Process `json:"save,omitempty"`
	Create            *action.Process `json:"create,omitempty"`
	Insert            *action.Process `json:"insert,omitempty"`
	Delete            *action.Process `json:"delete,omitempty"`
	DeleteIn          *action.Process `json:"delete-in,omitempty"`
	DeleteWhere       *action.Process `json:"delete-where,omitempty"`
	Update            *action.Process `json:"update,omitempty"`
	UpdateIn          *action.Process `json:"update-in,omitempty"`
	UpdateWhere       *action.Process `json:"update-where,omitempty"`
	BeforeFind        *hook.Before    `json:"before:find,omitempty"`
	AfterFind         *hook.After     `json:"after:find,omitempty"`
	BeforeSearch      *hook.Before    `json:"before:search,omitempty"`
	AfterSearch       *hook.After     `json:"after:search,omitempty"`
	BeforeGet         *hook.Before    `json:"before:get,omitempty"`
	AfterGet          *hook.After     `json:"after:get,omitempty"`
	BeforeSave        *hook.Before    `json:"before:save,omitempty"`
	AfterSave         *hook.After     `json:"after:save,omitempty"`
	BeforeCreate      *hook.Before    `json:"before:create,omitempty"`
	AfterCreate       *hook.After     `json:"after:create,omitempty"`
	BeforeInsert      *hook.Before    `json:"before:insert,omitempty"`
	AfterInsert       *hook.After     `json:"after:insert,omitempty"`
	BeforeDelete      *hook.Before    `json:"before:delete,omitempty"`
	AfterDelete       *hook.After     `json:"after:delete,omitempty"`
	BeforeDeleteIn    *hook.Before    `json:"before:delete-in,omitempty"`
	AfterDeleteIn     *hook.After     `json:"after:delete-in,omitempty"`
	BeforeDeleteWhere *hook.Before    `json:"before:delete-where,omitempty"`
	AfterDeleteWhere  *hook.After     `json:"after:delete-where,omitempty"`
	BeforeUpdate      *hook.Before    `json:"before:update,omitempty"`
	AfterUpdate       *hook.After     `json:"after:update,omitempty"`
	BeforeUpdateIn    *hook.Before    `json:"before:update-in,omitempty"`
	AfterUpdateIn     *hook.After     `json:"after:update-in,omitempty"`
	BeforeUpdateWhere *hook.Before    `json:"before:update-where,omitempty"`
	AfterUpdateWhere  *hook.After     `json:"after:update-where,omitempty"`
}

// BindActionDSL action.bind
type BindActionDSL struct {
	Model  string                 `json:"model,omitempty"`  // bind model
	Store  string                 `json:"store,omitempty"`  // bind store
	Table  string                 `json:"table,omitempty"`  // bind table
	Form   string                 `json:"form,omitempty"`   // bind form
	Option map[string]interface{} `json:"option,omitempty"` // bind option
}

// LayoutDSL the table layout
type LayoutDSL struct {
	Primary string           `json:"primary,omitempty"`
	Header  *HeaderLayoutDSL `json:"header,omitempty"`
	Filter  *FilterLayoutDSL `json:"filter,omitempty"`
	Table   *ViewLayoutDSL   `json:"table,omitempty"`
}

// HeaderLayoutDSL layout.header
type HeaderLayoutDSL struct {
	Preset  *PresetHeaderDSL      `json:"preset,omitempty"`
	Actions []component.ActionDSL `json:"actions"`
}

// PresetHeaderDSL layout.header.preset
type PresetHeaderDSL struct {
	Batch  *BatchPresetDSL  `json:"batch,omitempty"`
	Import *ImportPresetDSL `json:"import,omitempty"`
}

// BatchPresetDSL layout.header.preset.batch
type BatchPresetDSL struct {
	Columns   []component.InstanceDSL `json:"columns,omitempty"`
	Deletable bool                    `json:"deletable,omitempty"`
}

// ImportPresetDSL layout.header.preset.import
type ImportPresetDSL struct {
	Name      string               `json:"name,omitempty"`
	Operation []OperationImportDSL `json:"operation,omitempty"`
}

// OperationImportDSL  layout.header.preset.import.operation[*]
type OperationImportDSL struct {
	Title string `json:"title,omitempty"`
	Link  string `json:"link,omitempty"`
}

// FilterLayoutDSL layout.filter
type FilterLayoutDSL struct {
	BtnAddText string              `json:"btnAddText,omitempty"`
	Columns    component.Instances `json:"columns,omitempty"`
}

// ViewLayoutDSL layout.table
type ViewLayoutDSL struct {
	Props     component.PropsDSL  `json:"props,omitempty"`
	Columns   component.Instances `json:"columns,omitempty"`
	Operation OperationTableDSL   `json:"operation,omitempty"`
}

// OperationTableDSL layout.table.operation
type OperationTableDSL struct {
	Fold    bool              `json:"fold,omitempty"`
	Actions component.Actions `json:"actions"`
}

// FieldsDSL the table fields DSL
type FieldsDSL struct {
	Filter field.Filters `json:"filter,omitempty"`
	Table  field.Columns `json:"table,omitempty"`
}