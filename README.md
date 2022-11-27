<div align="center" id="top"> 
  <img src="./.github/app.gif" alt="Gorm Repository Boilerplate" />

&#xa0;

  <!-- <a href="https://gormrepositoryboilerplate.netlify.app">Demo</a> -->
</div>

<h1 align="center">Gorm Repository Boilerplate</h1>

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/bangadam/gorm-repository-boilerplate?color=56BEB8">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/bangadam/gorm-repository-boilerplate?color=56BEB8">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/bangadam/gorm-repository-boilerplate?color=56BEB8">

  <img alt="License" src="https://img.shields.io/github/license/bangadam/gorm-repository-boilerplate?color=56BEB8">

  <img alt="Github issues" src="https://img.shields.io/github/issues/bangadam/gorm-repository-boilerplate?color=56BEB8" />

  <img alt="Github forks" src="https://img.shields.io/github/forks/bangadam/gorm-repository-boilerplate?color=56BEB8" />

  <img alt="Github stars" src="https://img.shields.io/github/stars/bangadam/gorm-repository-boilerplate?color=56BEB8" />
</p>

<!-- Status -->

<!-- <h4 align="center">
	🚧  Gorm Repository Boilerplate 🚀 Under construction...  🚧
</h4>

<hr> -->

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#sparkles-features">Features</a> &#xa0; | &#xa0;
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Usage</a> &#xa0; | &#xa0;
  <a href="#memo-license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/bangadam" target="_blank">Author</a>
</p>

<br>

## :dart: About

Gorm Repository Boilerplate is a boilerplate for gorm repository pattern. It is a providing basic functions to CRUD and query entities as well as transactions and common error handling. It is a good starting point for gorm repository pattern.

## :sparkles: Features

:heavy_check_mark: CRUD;\

## :rocket: Technologies

The following tools were used in this project:

- [Golang](https://golang.org/)
- [Gorm](https://gorm.io/)

## :white_check_mark: Requirements

Before starting :checkered_flag:, you need to have [Go](https://golang.org/) installed.

## :checkered_flag: Usage

```go
package base

import "github.com/bangadam/gorm-repository-boilerplate"

type BaseRepository interface {
	gormrepository.TransactionRepository
	FindByName(target interface{}, name string, preloads ...string) error
}

type repository struct {
	gormrepository.TransactionRepository
}

func NewRepository(db *gorm.DB, logger logging.Logger) BaseRepository {
	return &repository{
		TransactionRepository: gormrepository.NewGormRepository(db, logger, "Creator"),
	}
}

func (r *repository) FindByName(target interface{}, name string, preloads ...string) error {
	return r.TransactionRepository.FindOneByField(target, "name", name, preloads...)
}
```

## :memo: License

This project is under license from MIT. For more details, see the [LICENSE](LICENSE.md) file.

Made with :heart: by <a href="https://github.com/bangadam" target="_blank">{{YOUR_NAME}}</a>

&#xa0;

<a href="#top">Back to top</a>
