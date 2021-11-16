package graph

import (
	"errors"
	"fmt"
	"log"
	"manifold/internal/document_definition"
	"manifold/internal/traversing/build_info"
	"manifold/internal/traversing/dependents"
	"os"
	"strings"
)

func Traverse(path string) (*Node, *TraverseError) {
	return traverseNested(path, nil)
}

func traverseNested(path string, ctx *contextImpl) (*Node, *TraverseError) {
	ctx, err := newContext(path, ctx)

	if err != nil {
		return nil, newTraverseError(path, err)
	}

	return traverse(ctx)
}

func traverse(ctx *contextImpl) (*Node, *TraverseError) {
	log.Println(fmt.Sprintf("%s...", ctx.path))

	file, openErr := os.Open(ctx.path)

	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	if openErr != nil {
		return nil, newTraverseError(ctx.path, openErr)
	}

	document, readErr := document_definition.Read(file)

	if readErr != nil {
		return nil, newTraverseError(ctx.path, readErr)
	}

	originalDir, _ := os.Getwd()

	_ = os.Chdir(ctx.dir)
	defer os.Chdir(originalDir)

	buildInfo, dependencies := build_info.FromDocumentDefinition(document, ctx)

	for _, warning := range ctx.warnings {
		log.Println(fmt.Sprintf("warning: %s", warning))
	}

	if !ctx.IsValid() {
		lines := []string{fmt.Sprintf("%d error(s) occurred:", len(ctx.errors))}
		lines = append(lines, ctx.errors...)

		err := errors.New(strings.Join(lines, "\n\t"))

		return nil, newTraverseError(ctx.path, err)
	}

	node := Node{
		BuildInfo:    buildInfo,
		Dependencies: make([]*Node, 0),
	}

	for _, dependency := range dependencies {
		var dependencyNode *Node
		var dependencyError *TraverseError

		switch dependency.Kind() {
		case dependents.DependentPathInfoKind:
			pathDependency := dependency.(dependents.DependentPathInfo)
			// TODO: check that dependency by path is already loaded
			dependencyNode, dependencyError = traverseNested(pathDependency.Path, ctx)
			break

		case dependents.DependentProjectInfoKind:
			// TODO: traverse, but back...
			panic("not implemented")

		default:
			panic(fmt.Sprintf("unknown dependency kind %v", dependency.Kind()))
		}

		if dependencyError != nil {
			return nil, dependencyError
		}

		node.Dependencies = append(node.Dependencies, dependencyNode)
	}

	return &node, nil
}
