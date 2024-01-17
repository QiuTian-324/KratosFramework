package serve

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/spf13/cobra"
	"go-Hermes/app/user/internal/conf"
	"net/http"
	"os"
)

var (
	// Name is the name of the compiled software.
	Name = "users.service"
	// Version is the version of the compiled software.
	Version = "v1"
	// Flagconf is the config flag.
	Flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&Flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(rr), // consul 的引入
	)
}

// runProfDebug 性能调试
func runProfDebug(debugAddr string) {
	go http.ListenAndServe(debugAddr, nil)
}

func StartServe(cmd *cobra.Command, args []string) {
	var err error
	var c config.Config
	if len(args) > 0 {
		configKey := args[0]
		c, err = LoadConfig(configKey)
		if err != nil {
			panic(err)
		}
		defer c.Close()
	} else {
		c, err = LoadConfig(Flagconf)
		if err != nil {
			panic(err)
		}
		defer c.Close()
	}

	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Registry, logger)
	if err != nil {
		panic(err)
	}

	defer cleanup()

	pprofDebugAddr, _ := cmd.Flags().GetString("pprof-debug-addr")
	if pprofDebugAddr != "" {
		runProfDebug(pprofDebugAddr)
	}

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func LoadConfig(path string) (config.Config, error) {
	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
	defer func(c config.Config) {
		err := c.Close()
		if err != nil {
			return
		}
	}(c)
	if err := c.Load(); err != nil {
		return nil, err
	}
	return c, nil
}
