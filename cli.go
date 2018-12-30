package main

import (
	"flag"
	"fmt"
	"os"
)

const Usage = `
add --data Data "添加区块数据"
print            "查看数据"
`

type CLI struct {
	bc *BlockChain
}

func (cli *CLI) Run() {
	if len(os.Args) < 2 {
		fmt.Printf(Usage)
		os.Exit(1)
	}
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	//pringCmd := flag.NewFlagSet("pring", flag.ExitOnError)
	addBCmdPara := addCmd.String("data", "", "blog info")
	switch os.Args[1] {
	case "add":
		err :=addCmd.Parse(os.Args[2:])
		CheckErr("add",err)
		if addCmd.Parsed() {
			if *addBCmdPara=="" {
				cli.add()
			}
		}
	case "print":
		err :=addCmd.Parse(os.Args[2:])
		CheckErr("print",err)
		if addCmd.Parsed() {
				cli.print()
		}
	}
}

func (cli *CLI) print() {
	fmt.Println("test")
}

func (cli *CLI) add() {
	fmt.Println("gdsg")
}
