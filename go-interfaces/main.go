package main

import (
	"fmt"
	"log"
)

type IWorker interface {
	Operation(message string) error
}

type Processor struct {
	Id     int
	Worker IWorker
}

func (p *Processor) DoWork(msg string) {
	err := p.Worker.Operation(msg)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("It works\nId: %d!\nMessage: %s", p.Id, msg)
}

func main() {
	var primary PrimaryWorker
	var secondary  SecondaryWorker
	primary.Operation("Primary Op")
	secondary.Operation("Secondary Op")

	pr1 := Processor{
		Id:     1,
		Worker: primary,
	}
	pr2 := Processor{
		Id:     2,
		Worker: secondary,
	}

	pr1.DoWork("Doing Processor 1")
	pr2.DoWork("Doing Processor 2")
}

type PrimaryWorker struct {
}

func (p PrimaryWorker) Operation(message string) error {
	if message != "" {
		return fmt.Errorf("error: processing primary worker")
	}
	log.Println("sucess: primary worker processed.")
	return nil
}

type SecondaryWorker struct {
}

func (p SecondaryWorker) Operation(message string) error {
	if message != "" {
		return fmt.Errorf("\nerror: processing secondary worker")
	}
	log.Println("success: secondary worker processed.")
	return nil
}
