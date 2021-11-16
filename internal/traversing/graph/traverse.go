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

func Traverse(path string) (*NodeCollection, *TraverseError) {
	return processRoot(path)
}

func processRoot(path string) (*NodeCollection, *TraverseError) {
	ctx, err := newContext(path)
	nodeCollection := newNodeCollection()

	if err != nil {
		return nil, newTraverseError(path, err)
	}

	node, traverseErr := produceNode(ctx)

	if traverseErr != nil {
		return nil, traverseErr
	}

	nodeCollection.Nodes = append(nodeCollection.Nodes, node)
	nodeCollection.Root = node

	traverseErr = processNode(node, nodeCollection)

	if traverseErr != nil {
		return nil, traverseErr
	}

	return nodeCollection, nil
}

func processNode(node *Node, nodeCollection *NodeCollection) *TraverseError {
	newNodes := make([]*Node, 0)

	for _, dependencyInfo := range node.Dependencies {
		err := (*TraverseError)(nil)
		dependencyNode := (*Node)(nil)

		switch dependencyInfo.Kind() {
		case dependents.DependentPathInfoKind:
			isNew := false

			path := dependencyInfo.(dependents.DependentPathInfo).Path
			isNew, dependencyNode, err = processByPath(path, nodeCollection)

			if err == nil && isNew {
				newNodes = append(newNodes, dependencyNode)
				nodeCollection.Nodes = append(nodeCollection.Nodes, dependencyNode)
			}

			break

		case dependents.DependentProjectInfoKind:
			project := dependencyInfo.(dependents.DependentProjectInfo).Project
			dependencyNode, err = processByProject(project, nodeCollection)
			break

		default:
			panic(fmt.Sprintf("unknown dependency kind %v", dependencyInfo.Kind()))
		}

		if err != nil {
			return err
		}

		link := newNodeLink(node, dependencyNode)
		nodeCollection.Links = append(nodeCollection.Links, link)
	}

	for _, newNode := range newNodes {
		err := processNode(newNode, nodeCollection)

		if err != nil {
			return err
		}
	}

	return nil
}

func processByPath(path string, nodeCollection *NodeCollection) (bool, *Node, *TraverseError) {
	ctx, err := newContext(path)

	if err != nil {
		return false, nil, newTraverseError(path, err)
	}

	dependencyNode := (*Node)(nil)

	for _, node := range nodeCollection.Nodes {
		if node.BuildInfo.Path() == ctx.path {
			dependencyNode = node
			break
		}
	}

	if dependencyNode != nil {
		return false, dependencyNode, nil
	}

	dependencyNode, traverseErr := produceNode(ctx)

	if traverseErr != nil {
		return false, nil, traverseErr
	}

	return true, dependencyNode, nil
}

func processByProject(project string, nodeCollection *NodeCollection) (*Node, *TraverseError) {
	dependencyNode := (*Node)(nil)

	for _, node := range nodeCollection.Nodes {
		if node.BuildInfo.Name() == project {
			dependencyNode = node
			break
		}
	}

	if dependencyNode == nil {
		// TODO: handle error properly
		return nil, newTraverseError("", errors.New(fmt.Sprintf("project not found: %s", project)))
	}

	return dependencyNode, nil
}

func produceNode(ctx *contextImpl) (*Node, *TraverseError) {
	// TODO: log.Println(fmt.Sprintf("%s...", ctx.path))

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

	return newNode(buildInfo, dependencies), nil
}
