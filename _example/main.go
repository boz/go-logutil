package main

import (
	"context"
	"errors"
	stdlog "log"
	"os"

	logutil "github.com/boz/go-logutil"
)

func main() {
	log := logutil.Default()
	testit(log)

	log = log.WithComponent("server")
	testit(log)

	ctx := context.Background()
	log = logutil.FromContextOrDefault(ctx)
	testit(log)

	ctx = logutil.NewContext(ctx, log.WithComponent("ctx"))
	log, _ = logutil.FromContext(ctx)
	testit(log)

	log = logutil.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags|stdlog.Lshortfile), os.Stdout)
	testit(log)

	log.Fatalf("bye")
}

func testit(log logutil.Log) {
	log.Infof("infof %v", "a")
	log.Debugf("debugf %v", "b")
	log.Warnf("warnf %v", "c")
	log.Errorf("error %v", "d")

	log.ErrWarn(errors.New("bad"), "errwarn %v", "e")
	log.Err(errors.New("bad"), "err %v", "f")
}
