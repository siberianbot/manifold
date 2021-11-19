package test

import (
	"fmt"
	"manifold/internal/steps"
)

type FakeContext struct {
	Errors       []string
	Warnings     []string
	File         string
	StepProvider steps.StepProvider
}

func (f *FakeContext) GetStepProvider() steps.StepProvider {
	return f.StepProvider
}

func (f *FakeContext) AddError(message string, params ...interface{}) {
	f.Errors = append(f.Errors, fmt.Sprintf(message, params...))
}

func (f *FakeContext) AddWarning(message string, params ...interface{}) {
	f.Warnings = append(f.Warnings, fmt.Sprintf(message, params...))
}

func (f FakeContext) IsValid() bool {
	return len(f.Errors) == 0
}

func (f FakeContext) CurrentFile() string {
	return f.File
}

func NewFakeContext() FakeContext {
	return FakeContext{
		Errors:       make([]string, 0),
		Warnings:     make([]string, 0),
		File:         "",
		StepProvider: steps.NewDefaultStepProvider(),
	}
}
