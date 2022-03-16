package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type zipfiles []string

func newFiles(fls []string, p *[]string) *zipfiles {
	*p = fls
	return (*zipfiles)(p)
}

func (f *zipfiles) Set(fls string) error {
	*f = zipfiles(strings.Split(fls, ","))
	return nil
}

func (f *zipfiles) Get() interface{} {
	return []string(*f)
}

func (f *zipfiles) String() string {
	return strings.Join([]string(*f), ",")
}

var (
	files      []string
	outputName string
	isCompress bool
)

func init() {
	flag.Var(newFiles([]string{}, &files), "files", "The files need zip")
	flag.StringVar(&outputName, "output", "myzip.zip", "The output name")
	flag.BoolVar(&isCompress, "isCompress", false, "Is compress file")
}

type myzipOptions struct {
	isCompress bool
}

type MyzipOptions interface {
	apply(*myzipOptions)
}

type tempFunc func(*myzipOptions)

type funcMyzipOption struct {
	f tempFunc
}

func (fdo *funcMyzipOption) apply(opt *myzipOptions) {
	fdo.f(opt)
}

func newFuncMyzipOption(f tempFunc) *funcMyzipOption {
	return &funcMyzipOption{f: f}
}

func withIsCompress(isCompress bool) MyzipOptions {
	return newFuncMyzipOption(func(mo *myzipOptions) {
		mo.isCompress = isCompress
	})
}

func defaultMyzipOption() myzipOptions {
	return myzipOptions{
		isCompress: true,
	}
}

type MyzipArgv struct {
	files      []string
	outputName string
	options    myzipOptions
}

func addFileToZip(rootPath string, file string, zipWriter *zip.Writer, isCompress bool) error {
	// 打开要写入的文件
	fileToZip, err := os.Open(filepath.Join(rootPath, file))
	if err != nil {
		return err
	}
	defer fileToZip.Close()
	// 获取文件信息
	fileInfo, err := fileToZip.Stat()
	if err != nil {
		return err
	}
	// 设置压缩头部信息
	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	// 设置文件名
	header.Name = strings.TrimPrefix(file, string(filepath.Separator))
	if fileInfo.IsDir() {
		newFiles, err := fileToZip.ReadDir(-1)
		if err != nil {
			return err
		}
		for _, newFile := range newFiles {
			addFileToZip(rootPath, filepath.Join(file, newFile.Name()), zipWriter, isCompress)
		}
		return nil
	}

	// 设置压缩模式
	if isCompress {
		header.Method = zip.Deflate
	} else {
		// 设置压缩模式
		header.Method = zip.Store
	}
	// 写入头部信息
	fileWriter, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	// 拷贝文件
	written, err := io.Copy(fileWriter, fileToZip)
	if err != nil {
		return err
	}
	fmt.Printf("add file %s to zip success, %d \n", file, written)
	return nil
}

func ZipFiles(files []string, outputName string, options ...MyzipOptions) error {
	// 设置默认参数
	argv := &MyzipArgv{
		files:      files,
		outputName: outputName,
		options:    defaultMyzipOption(),
	}
	for _, option := range options {
		option.apply(&argv.options)
	}

	// 获取运行目录
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	// 创建输出文件
	outputFile, err := os.Create(filepath.Join(rootPath, outputName))
	if err != nil {
		return err
	}
	defer outputFile.Close()
	zipWriter := zip.NewWriter(outputFile)
	defer zipWriter.Close()
	// 压缩文件
	for _, file := range files {
		err := addFileToZip(rootPath, file, zipWriter, argv.options.isCompress)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	options := []MyzipOptions{
		withIsCompress(isCompress),
	}
	err := ZipFiles(files, outputName, options...)
	if err != nil {
		panic(err)
	}
}
