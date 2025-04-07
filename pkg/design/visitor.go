package desgin

import "fmt"

// ----------- 元素接口 -----------
type Element interface {
	Accept(Visitor)
}

// ----------- 文件结构 -----------
type File struct {
	Name string
	Size int // 文件大小，单位：KB
}

func (f *File) Accept(visitor Visitor) {
	visitor.VisitFile(f)
}

// ----------- 文件夹结构 -----------
type Folder struct {
	Name     string
	Children []Element
}

func (f *Folder) Accept(visitor Visitor) {
	visitor.VisitFolder(f)
}

// ----------- 访问者接口 -----------
type Visitor interface {
	VisitFile(*File)
	VisitFolder(*Folder)
}

// ----------- 具体访问者：大小计算器 -----------
type SizeCalculator struct {
	TotalSize int
}

func (s *SizeCalculator) VisitFile(file *File) {
	s.TotalSize += file.Size
}

func (s *SizeCalculator) VisitFolder(folder *Folder) {
	for _, child := range folder.Children {
		child.Accept(s)
	}
}

// ----------- 具体访问者：结构展示器 -----------
type StructureDisplayer struct {
	Indent string
}

func (d *StructureDisplayer) VisitFile(file *File) {
	fmt.Printf("%sFile: %s (%d KB)\n", d.Indent, file.Name, file.Size)
}

func (d *StructureDisplayer) VisitFolder(folder *Folder) {
	fmt.Printf("%sFolder: %s\n", d.Indent, folder.Name)
	previousIndent := d.Indent
	d.Indent += "  "
	for _, child := range folder.Children {
		child.Accept(d)
	}
	d.Indent = previousIndent
}

// ----------- 主函数 -----------
func main() {
	// 创建文件和文件夹
	file1 := &File{Name: "file1.txt", Size: 120}
	file2 := &File{Name: "file2.txt", Size: 80}
	subFolder := &Folder{
		Name:     "subfolder",
		Children: []Element{file2},
	}
	rootFolder := &Folder{
		Name:     "root",
		Children: []Element{file1, subFolder},
	}

	// 使用大小计算访问者
	sizeCalculator := &SizeCalculator{}
	rootFolder.Accept(sizeCalculator)
	fmt.Printf("总大小: %d KB\n", sizeCalculator.TotalSize)

	// 使用结构展示访问者
	fmt.Println("\n文件结构:")
	structureDisplayer := &StructureDisplayer{}
	rootFolder.Accept(structureDisplayer)
}
