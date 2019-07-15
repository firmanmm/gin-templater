package gintemplater

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Templater struct {
	logger    *log.Logger
	engine    *gin.Engine
	builder   *templaterBuilder
	watcher   *templaterWatcher
	hotReload bool
	outputDir string
}

func (t *Templater) Run() {
	t.builder.initBuild()
	if t.hotReload {
		t.watcher.Run()
		t.logger.Println("Templater started")
	}
}

func (t *Templater) Stop() {
	if t.hotReload {
		t.watcher.Stop()
		t.logger.Println("Templater stopped")
	}
}

func (t *Templater) reload() {
	t.engine.LoadHTMLGlob(t.outputDir + "/*")
	t.logger.Println("Templater reloaded")
}

func NewTemplater(engine *gin.Engine, conf *Config) *Templater {
	instance := new(Templater)
	instance.engine = engine
	instance.outputDir = conf.OutputDir
	instance.hotReload = conf.AutoReload
	instance.logger = log.New(os.Stdout, "[GIN-TEMPLATER] ", log.Ltime)
	instance.builder = newTemplaterBuilder(conf.InputDir, conf.OutputDir, instance.logger)
	instance.watcher = newTemplaterWatcher(conf.OutputDir, instance.logger)
	instance.watcher.addListener(instance.builder.generate)
	return instance
}
