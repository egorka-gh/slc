package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/egorka-gh/sm/slc"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func main() {
	var csvDir string
	var postDir string

	flag.StringVar(&csvDir, "indir", "", "Input dir with csv from slc")
	flag.StringVar(&postDir, "postdir", "", "Output dir for xml files")

	flag.Parse()

	if csvDir == "" || postDir == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if csvDir == postDir {
		fmt.Println("Директории должны быть разными")
		os.Exit(1)
	}

	//check in dir
	if err := checkDir(csvDir); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//check out dir
	if err := checkDir(postDir); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//recreate done folder
	doneDir := filepath.Join(csvDir, "done")
	if err := os.MkdirAll(doneDir, 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//scan in dir
	dirScan, err := os.Open(csvDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	csvFiles, err := dirScan.Readdir(0)
	dirScan.Close()
	slc := &slc.Client{
		BornIn:   "slcltsVVSeqKh08Z2LmCbg==",
		Client:   534,
		IDprefix: "SLC",
		IWopcode: 4,
		IWstate:  1,
		LocFrom:  33,
	}
	for index := range csvFiles {
		fi := csvFiles[index]
		if fi.Mode().IsDir() {
			continue
		}
		fmt.Println(fi.Name())
		err := processCSV(filepath.Join(csvDir, fi.Name()), postDir, slc)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//move file to done folder
		err = os.Rename(filepath.Join(csvDir, fi.Name()), filepath.Join(doneDir, fi.Name()))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}

func processCSV(csvFile, outDir string, slc *slc.Client) error {
	//open file
	fl, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer fl.Close()
	//win1251 -> utf8
	r := transform.NewReader(fl, charmap.Windows1251.NewDecoder())
	p, err := slc.ParseIW(r)
	if err != nil {
		return err
	}
	//create output file
	out, err := os.Create(filepath.Join(outDir, p.Name+".xml"))
	if err != nil {
		return err
	}
	defer out.Close()

	return p.Encode(out)
}

func checkDir(dir string) error {
	fi, err := os.Stat(dir)
	if err != nil {
		return err
	}
	if !fi.Mode().IsDir() {
		return fmt.Errorf("Не верная папка '%s'", dir)
	}
	return nil
}
