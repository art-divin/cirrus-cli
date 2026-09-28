package options

type TartOptions struct {
	LazyPull bool
	NoUnmount bool
    RunArgs  []string
    // NoCleanup keeps the cloned VM after the task instead of deleting it (set by --tart-no-cleanup).
    NoCleanup bool
}
