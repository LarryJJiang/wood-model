package gredis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"time"
	"wood/pkg/app"
	bizcode "wood/pkg/bizerror"
	"wood/pkg/logging"
)

// ErrForciblyClose
var (
	// windows
	// wsarecv: An existing connection was forcibly closed by the remote host.
	ErrForciblyClose = "connection was forcibly closed by the remote host"
)

// Publish .
func Publish(args ...interface{}) (err error) {
	conn := RedisConn.Get()
	defer conn.Close()

	// publish
	_, err = conn.Do("PUBLISH", args...)
	if err != nil {
		return
	}
	return
}

// Subscribe .
func Subscribe(ctx context.Context, channels []string, fn func(redis.Message) error, fnErrExit bool) (isReady, isGracefulExit bool, err error) {
	psc := &redis.PubSubConn{Conn: RedisConn.Get()}
	defer psc.Close()

	// channel
	var channelArgs = redis.Args{}.AddFlat(channels)
	if err := psc.Subscribe(channelArgs...); err != nil {
		return isReady, isGracefulExit, err
	}
	defer func() { _ = psc.Unsubscribe(channelArgs...) }()

	// done
	done := make(chan error, 1)
	defer close(done)

	// process
	go func() {
		for {
			switch res := psc.Receive().(type) {
			case error:
				done <- res
				return
			case redis.Subscription:
				if res.Count == 0 {
					done <- nil
					return
				}
			case redis.Message:
				if fnErr := fn(res); fnErr != nil && fnErrExit {
					done <- fnErr
					return
				}
			case redis.Pong:

			}
		}
	}()
	isReady = true

	// ticker
	ticker := time.NewTicker(healthCheckPeriod)
	defer ticker.Stop()

	// loop
loop:
	for {
		select {
		case <-ticker.C:
			// Send ping to test health of connection and server. If
			// corresponding pong is not received, then receive on the
			// connection will timeout and the receive goroutine will exit.
			if err = psc.Ping(""); err != nil {
				break loop
			}
		case <-ctx.Done():
			isGracefulExit = true
			break loop
		case err := <-done:
			if err == nil {
				isGracefulExit = true
			}
			// Return error from the receive goroutine.
			return isReady, isGracefulExit, err
		}
	}
	return
}

// pub sub
const (
	CustomerTeamChangeKey     = "local_channel_for_customer_team_change" // key
	CustomerTeamChangeMessage = "1"                                      // msg
)

// CustomerTeamChangePublish .
func CustomerTeamChangePublish() (err error) {
	// publish
	err = Publish(CustomerTeamChangeKey, CustomerTeamChangeMessage)
	if err != nil {
		logging.Error(err)
		err = app.NewResponseErr(bizcode.ErrorRedis, nil)
		return
	}
	return
}
