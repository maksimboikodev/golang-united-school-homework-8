package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	argOp       = "operation"
	argId       = "id"
	argItem     = "item"
	argFilename = "fileName"

	opList     = "list"
	opAdd      = "add"
	opFindById = "findById"
	opRemove   = "remove"
)

var (
	errorNotSpecifiedOperation = errors.New("-operation flag has to be specified")
	errorNotSpecifiedFileName  = errors.New("-fileName flag has to be specified")
	errorNotSpecifiedItem      = errors.New("-item flag has to be specified")
	errorNotSpecifiedId        = errors.New("-id flag has to be specified")
)

func parseArgs() Arguments {
	flagId := flag.String("id", "1", "users id")
	flagOperation := flag.String("operation", "list", "supported operations: add, list, findById, remove")
	flagItem := flag.String("item", "{\"id\": \"1\", \"email\": \"email@test.com\", \"age\": 23}", "JSON example: {\"id\": \"1\", \"email\": \"email@test.com\", \"age\": 23}")
	flagFileName := flag.String("fileName", "test.json", "JSON file with data")

	flag.Parse()

	return Arguments{
		"id":        *flagId,
		"operation": *flagOperation,
		"item":      *flagItem,
		"fileName":  *flagFileName,
	}
}

func Perform(args Arguments, writer io.Writer) error {
	if args["operation"] == "" {
		return errorNotSpecifiedOperation
	}

	if args["fileName"] == "" {
		return errorNotSpecifiedFileName
	}

	switch args["operation"] {
	case "list":
		return listOperation(args, writer)
	case "findById":
		return findByIdOperation(args, writer)
	case "add":
		return addOperation(args, writer)
	case "remove":
		return removeOperation(args, writer)
	default:
		return fmt.Errorf("Operation %s not allowed!", args["operation"])
	}
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
