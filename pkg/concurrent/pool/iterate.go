package pool

import "context"

// IterateWithStep iterate data with step
func IterateWithStep(ctx context.Context, data []string, step int, process func(data []string) error) error {
	for i := 0; i < len(data); i += step {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		nextI := i + step
		if nextI > len(data) {
			nextI = len(data)
		}

		err := process(data[i:nextI])
		if err != nil {
			return err
		}
	}
	return nil
}
