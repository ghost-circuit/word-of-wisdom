# fxslog
a small shim package providing a slog logger for uber's fx library

## Usage

The WithLogger function returns an fx.Option that configures fx to use a 
*slog.Logger for logging. Using this option requires that you provide a 
*slog.Logger.


```golang
func main() {
	flag.Parse()

	fx.New(
		//fx.WithLogger(
		//	func(logger *slog.Logger) fxevent.Logger {
		//		return &fxslog.SlogLogger{Logger: logger.With("component", "uber/fx")}
		//	},
		//),
        
        // This is a thin wrapper over the snippet from above.
        fxslog.WithLogger(),
        
        fx.Provide(
            // We still need to provide a logger. The fxslog.WithLogger 
            // returns a fx.Option that consumes *slog.Logger
            func() *slog.Logger{
                return slog.Default()
            },
        )
		...
	).Run()
}

```