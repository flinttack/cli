package ecosystem

import (
	"github.com/flinttack/flinttack/internal/ecosystem/bun"
	"github.com/flinttack/flinttack/internal/ecosystem/bundler"
	"github.com/flinttack/flinttack/internal/ecosystem/cabal"
	"github.com/flinttack/flinttack/internal/ecosystem/cargo"
	"github.com/flinttack/flinttack/internal/ecosystem/cocoapods"
	"github.com/flinttack/flinttack/internal/ecosystem/composer"
	"github.com/flinttack/flinttack/internal/ecosystem/conan"
	"github.com/flinttack/flinttack/internal/ecosystem/deno"
	"github.com/flinttack/flinttack/internal/ecosystem/golang"
	"github.com/flinttack/flinttack/internal/ecosystem/gradle"
	"github.com/flinttack/flinttack/internal/ecosystem/julia"
	"github.com/flinttack/flinttack/internal/ecosystem/mix"
	"github.com/flinttack/flinttack/internal/ecosystem/npm"
	"github.com/flinttack/flinttack/internal/ecosystem/nuget"
	"github.com/flinttack/flinttack/internal/ecosystem/pdm"
	"github.com/flinttack/flinttack/internal/ecosystem/pipenv"
	"github.com/flinttack/flinttack/internal/ecosystem/pnpm"
	"github.com/flinttack/flinttack/internal/ecosystem/poetry"
	"github.com/flinttack/flinttack/internal/ecosystem/pub"
	"github.com/flinttack/flinttack/internal/ecosystem/pypi"
	"github.com/flinttack/flinttack/internal/ecosystem/renv"
	"github.com/flinttack/flinttack/internal/ecosystem/swiftpm"
	"github.com/flinttack/flinttack/internal/ecosystem/yarn"
)

func All() []Ecosystem {
	return []Ecosystem{
		pnpm.New(),
		yarn.New(),
		bun.New(),
		npm.New(),
		deno.New(),
		cargo.New(),
		golang.New(),
		pypi.New(),
		poetry.New(),
		pdm.New(),
		pipenv.New(),
		composer.New(),
		bundler.New(),
		nuget.New(),
		swiftpm.New(),
		cocoapods.New(),
		conan.New(),
		renv.New(),
		julia.New(),
		cabal.New(),
		pub.New(),
		mix.New(),
		gradle.New(),
	}
}

func Prefer(found []Ecosystem) []Ecosystem {
	found = preferFamily(found, []string{"pnpm", "yarn", "bun", "npm"})
	found = preferFamily(found, []string{"pypi", "poetry", "pdm", "pipenv"})
	return found
}

func preferFamily(found []Ecosystem, order []string) []Ecosystem {
	present := map[string]bool{}
	member := map[string]bool{}
	for _, name := range order {
		member[name] = true
	}
	for _, e := range found {
		present[e.Name()] = true
	}
	chosen := ""
	for _, name := range order {
		if present[name] {
			chosen = name
			break
		}
	}
	if chosen == "" {
		return found
	}
	out := make([]Ecosystem, 0, len(found))
	for _, e := range found {
		if member[e.Name()] && e.Name() != chosen {
			continue
		}
		out = append(out, e)
	}
	return out
}
