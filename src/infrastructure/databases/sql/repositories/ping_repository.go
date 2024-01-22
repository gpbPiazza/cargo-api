package repositories

import "context"

func Ping(ctx context.Context) error {
	// ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()

	// if err := globalDB.PingContext(ctx); err != nil {
	// 	return err
	// }

	return nil
}
