package gintemplater

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

//Templater instance to perform hot reload on Gin Web Framework
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
	t.reload()
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

//NewTemplater will return a new and ready to use Templater instance
func NewTemplater(engine *gin.Engine, conf *Config) *Templater {
	instance := new(Templater)
	instance.engine = engine
	instance.outputDir = conf.OutputDir
	instance.hotReload = conf.AutoReload
	instance.logger = log.New(os.Stdout, "[GIN-TEMPLATER] ", log.Ltime)
	instance.builder = newTemplaterBuilder(conf.InputDir, conf.OutputDir, instance.logger)
	instance.watcher = newTemplaterWatcher(conf.InputDir, instance.logger)
	if conf.AutoReload {
		rebuildEv := func(data string) {
			instance.reload()
		}
		instance.watcher.addListener(instance.builder.generate)
		instance.watcher.addListener(rebuildEv)
	}
	return instance
}
