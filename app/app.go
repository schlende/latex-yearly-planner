package app

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/compose"
	"github.com/kudrykv/latex-yearly-planner/app/config"
	"github.com/kudrykv/latex-yearly-planner/app/tpls"
	"github.com/urfave/cli/v2"
)

const (
	fConfig = "config"
	pConfig = "preview"
)

func New() *cli.App {
	return &cli.App{
		Name: "plannergen",

		Writer:    os.Stdout,
		ErrWriter: os.Stderr,

		Flags: []cli.Flag{
			&cli.PathFlag{Name: fConfig, Required: true},
			&cli.BoolFlag{Name: pConfig, Required: false},
		},

		Action: action,
	}
}

func action(c *cli.Context) error {
	var (
		fn  Composer
		ok  bool
		cfg config.Config
		err error
	)

	preview := c.Bool(pConfig)

	pathConfigs := strings.Split(c.Path(fConfig), ",")
	if cfg, err = config.New(pathConfigs...); err != nil {
		return fmt.Errorf("config new: %w", err)
	}

	wr := &bytes.Buffer{}

	t := tpls.New()

	if err = t.Document(wr, cfg); err != nil {
		return fmt.Errorf("tex document: %w", err)
	}

	if err = ioutil.WriteFile("out/"+RootFilename(pathConfigs[len(pathConfigs)-1]), wr.Bytes(), 0600); err != nil {
		return fmt.Errorf("ioutil write file: %w", err)
	}

	for _, file := range cfg.Pages {
		wr.Reset()

		var mom []page.Modules
		for _, block := range file.RenderBlocks {
			if fn, ok = ComposerMap[block.FuncName]; !ok {
				return fmt.Errorf("unknown func " + block.FuncName)
			}

			modules, err := fn(cfg, block.Tpls)

			// Only one page per unique module if preview flag is enabled
			if preview {
				modules = filterUniqueModules(modules)
			}

			if err != nil {
				return fmt.Errorf("%s: %w", block.FuncName, err)
			}

			mom = append(mom, modules)
		}

		if len(mom) == 0 {
			return fmt.Errorf("modules of modules must have some modules")
		}

		allLen := len(mom[0])
		for _, mods := range mom {
			if len(mods) != allLen {
				return errors.New("some modules are not aligned")
			}
		}

		for i := 0; i < allLen; i++ {
			for j, mod := range mom {
				if err = t.Execute(wr, mod[i].Tpl, mod[i]); err != nil {
					return fmt.Errorf("execute %s on %s: %w", file.RenderBlocks[j].FuncName, mod[i].Tpl, err)
				}
			}
		}

		if err = ioutil.WriteFile("out/"+file.Name+".tex", wr.Bytes(), 0600); err != nil {
			return fmt.Errorf("ioutil write file: %w", err)
		}
	}

	return nil
}

func RootFilename(pathconfig string) string {
	if idx := strings.LastIndex(pathconfig, "/"); idx >= 0 {
		pathconfig = pathconfig[idx+1:]
	}

	if strings.HasSuffix(pathconfig, ".yml") {
		pathconfig = pathconfig[:len(pathconfig)-len(".yml")]
	}

	if strings.HasSuffix(pathconfig, ".yaml") {
		pathconfig = pathconfig[:len(pathconfig)-len(".yaml")]
	}

	return pathconfig + ".tex"
}

type Composer func(cfg config.Config, tpls []string) (page.Modules, error)

var ComposerMap = map[string]Composer{
	"title":          compose.Title,
	"goals":          compose.Goals,
	"annual":         compose.Annual,
	"quarterly":      compose.Quarterly,
	"monthly":        compose.Monthly,
	"weekly":         compose.Weekly,
	"daily":          compose.Daily,
	"daily_reflect":  compose.DailyReflect,
	"daily_notes":    compose.DailyNotes,
	"daily_planning": compose.DailyPlanning,
	"notes_indexed":  compose.NotesIndexed,
}

func filterUniqueModules(array []page.Module) []page.Module {
	filtered := make([]page.Module, 0)
	found := map[string]bool{}

	for _, val := range array {
		if _, present := found[val.Tpl]; !present {
			filtered = append(filtered, val)
			found[val.Tpl] = true
		}
	}

	return filtered
}
