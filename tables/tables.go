package tables

import (
	"os"
	"strings"
)

type Table struct {
	// database parameters

	Name    string
	Columns []Column

	// file parameters

	struct_name    string
	path           string
	filename_full  string
	filename_short string
	package_name   string
	file           *os.File
	imports        []string
}

const (
	// keywords

	package_str       string = "package "
	import_str        string = "import "
	type_str          string = "type "
	struct_str        string = " struct "
	func_str          string = "func "
	open_braces       string = "{"
	close_braces      string = "}"
	open_parenthesis  string = "("
	close_parenthesis string = ")"
	pointer_str       string = " *"

	// escapped chars

	new_line string = "\n"
	tab      string = "\t"
	quote    string = "\""
	space    string = " "

	//defaults

	default_path = "database/"
)

func New(name string) *Table {
	name = strings.ToLower(name)
	struct_name := name
	if struct_name[len(name)-1:] == "s" {
		struct_name = name[:len(name)-1]
	}
	t := Table{
		Name:           name,
		path:           default_path,
		package_name:   name,
		filename_short: name + ".go",
		struct_name:    capitalize(struct_name),
		imports:        []string{"test/database"},
	}
	return &t
}

func (t *Table) ChangePath(new_path string) {
	if new_path[len(new_path)-1:] != "/" {
		new_path += "/"
	}
	t.path = new_path
	t.filename_full = t.path + t.filename_short
}

func (t *Table) ChangeFilename(new_name string) {
	t.filename_short = new_name
	t.filename_full = t.path + t.filename_short
}

func (t *Table) writeln(strs ...string) (err error) {
	for _, str := range strs {
		_, err = t.file.WriteString(str)
		if err != nil {
			return err
		}
	}
	_, err = t.file.WriteString(new_line)
	return err
}

func (t *Table) checkImports() (err error) {
	for _, c := range t.Columns {
		i := c.Type.Import()
		if i == "" {
			continue
		}

		t.addImport(i)
	}
	return err
}

func (t *Table) addImport(i string) {
	for _, imp := range t.imports {
		if i == imp {
			return
		}
	}
	t.imports = append(t.imports, i)
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func quoteString(str string) string {
	return quote + str + quote
}
func (t *Table) Create() error {
	if err := os.MkdirAll(t.path, os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(t.filename_full)
	if err != nil {
		return err
	}

	t.file = file
	return nil
}

func (t *Table) WriteHead() (err error) {
	_, err = t.file.WriteString(package_str + t.Name + new_line)
	t.writeln()
	return err
}

func (t *Table) WriteImports() (err error) {
	t.checkImports()
	if len(t.imports) == 0 {
		return nil
	}

	err = t.writeln(import_str + open_parenthesis)
	if err != nil {
		return err
	}

	for _, imp := range t.imports {
		err = t.writeln(tab + quoteString(imp))
		if err != nil {
			return err
		}
	}

	err = t.writeln(close_parenthesis)
	if err != nil {
		return err
	}
	t.writeln()

	return nil
}

func (t *Table) WriteType() (err error) {
	err = t.writeln(type_str + t.struct_name + struct_str + open_braces)
	if err != nil {
		return err
	}

	for _, c := range t.Columns {
		err = t.writeln(tab + c.ParseField())
		if err != nil {
			return err
		}
	}

	err = t.writeln(close_braces)
	if err != nil {
		return err
	}
	t.writeln()

	return nil
}

func (t *Table) WriteMethods() (err error) {
	err = t.writeInsert()
	if err != nil {
		return err
	}
	err = t.writeUpdate()
	if err != nil {
		return err
	}
	err = t.writeDelete()
	return err
}

func (t *Table) writeInsert() (err error) {
	method := "Insert"
	variable := string(strings.ToLower(t.struct_name)[0])
	reference := open_parenthesis + variable + pointer_str + t.struct_name + close_parenthesis
	method_head := func_str + reference + " " + method + open_parenthesis + close_parenthesis + " error " + open_braces
	err = t.writeln(method_head)
	if err != nil {
		return err
	}
	err = t.writeln(tab + "tx := database.DB.Save(" + variable + ")")
	if err != nil {
		return err
	}
	err = t.writeln(tab + "if tx.Error != nil {")
	if err != nil {
		return err
	}
	err = t.writeln(tab + tab + "return tx.Error")
	if err != nil {
		return err
	}
	err = t.writeln(tab + close_braces)
	if err != nil {
		return err
	}
	err = t.writeln(tab + "return nil")
	if err != nil {
		return err
	}
	err = t.writeln(close_braces)
	if err != nil {
		return err
	}
	t.writeln()
	return err
}

func (t *Table) writeUpdate() (err error) {
	method := "Update"
	variable := string(strings.ToLower(t.struct_name)[0])
	reference := open_parenthesis + variable + pointer_str + t.struct_name + close_parenthesis
	method_head := func_str + reference + " " + method + open_parenthesis + close_parenthesis + " error " + open_braces
	err = t.writeln(method_head)
	if err != nil {
		return err
	}
	err = t.writeln(tab + "tx := database.DB.Save(" + variable + ")")
	if err != nil {
		return err
	}
	err = t.writeln(tab + "if tx.Error != nil {")
	if err != nil {
		return err
	}
	err = t.writeln(tab + tab + "return tx.Error")
	if err != nil {
		return err
	}
	err = t.writeln(tab + close_braces)
	if err != nil {
		return err
	}
	err = t.writeln(tab + "return nil")
	if err != nil {
		return err
	}
	err = t.writeln(close_braces)
	if err != nil {
		return err
	}
	t.writeln()
	return err
}

func (t *Table) writeDelete() (err error) {
	method := "Delete"
	variable := string(strings.ToLower(t.struct_name)[0])
	reference := open_parenthesis + variable + pointer_str + t.struct_name + close_parenthesis
	method_head := func_str + reference + " " + method + open_parenthesis + close_parenthesis + " error " + open_braces
	err = t.writeln(method_head)
	if err != nil {
		return err
	}
	err = t.writeln(tab + "tx := database.DB.Delete(" + variable + ")")
	if err != nil {
		return err
	}
	err = t.writeln(tab + "if tx.Error != nil {")
	if err != nil {
		return err
	}
	err = t.writeln(tab + tab + "return tx.Error")
	if err != nil {
		return err
	}
	err = t.writeln(tab + close_braces)
	if err != nil {
		return err
	}
	err = t.writeln(tab + "return nil")
	if err != nil {
		return err
	}
	err = t.writeln(close_braces)
	if err != nil {
		return err
	}
	t.writeln()
	return err
}

func (t *Table) Run() (err error) {
	err = t.Create()
	if err != nil {
		return err
	}
	err = t.WriteHead()
	if err != nil {
		return err
	}
	err = t.WriteImports()
	if err != nil {
		return err
	}
	err = t.WriteType()
	if err != nil {
		return err
	}
	err = t.WriteMethods()
	if err != nil {
		return err
	}

	return err
}
