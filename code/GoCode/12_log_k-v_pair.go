package main

import (
    "github.com/go-logr/logr"
    "github.com/go-logr/zapr"
    "go.uber.org/zap"
)

type MyStruct struct {
    Log logr.Logger
}

func (m *MyStruct) logMssg(k, v string) {
    m.Log.WithValues(k, v).Info("Adding key value pair")
}

func main() {
    logger, err := zap.NewDevelopment()
    if err != nil {
        panic(err)
    }

    myStruct := &MyStruct{
        Log: zapr.NewLogger(logger),
    }

    myStruct.logMssg("Key", "Value")
}