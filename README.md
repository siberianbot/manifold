# Manifold

Manifold - a build tool.

## Motivation

A long time ago, in a galaxy far away... Okay, I'm get to the point.

A few time ago I need a tool to build a wide range of dependent things in different languages. This project is dead, but
the main idea around Manifold is still alive: a simple and fast tool to build set of projects written in different
languages which could be dependent on other projects.

Manifold was inspired by solutions from Visual Studio, but I didn't find any alternative to cross-cutting concern like
mine.

## Usage

Usage is simple.

1. Put a `.manifold.yml` in project root;
2. Run `manifold build`.

Manifold build all dependencies of this project and project itself.

## Steps

Each project defines a bunch of steps, which are required to execute a build.

## Projects or Workspaces

Manifold operates with two different types of projects: project (how obviously!) and workspaces. Workspace is just a
bunch of projects or workspaces, but it only contains paths to nested projects. Project dependencies can be names of
dependent projects, or paths to dependent projects. Manifold deals with this dependencies in different ways:

- Dependency by project name: Manifold seeks for project from upper-level workspaces and projects to find project or
  workspace with such name, otherwise - it fails;
- Dependency by path: Manifold checks that project with such path was loaded earlier and uses it, otherwise - it loads
  project by path.

## Next goals

I have several improvements that I want to implement in the future, like...

- Support of building project which utilizes .NET, CMake, npm, etc. Right now only command line is supported;
- Parallel building.