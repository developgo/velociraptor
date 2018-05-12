package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"www.velocidex.com/golang/velociraptor/config"
	"www.velocidex.com/golang/velociraptor/context"
	"www.velocidex.com/golang/velociraptor/crypto"
	"www.velocidex.com/golang/velociraptor/executor"
	"www.velocidex.com/golang/velociraptor/http_comms"
)

func RunClient() {
	kingpin.Parse()

	ctx := context.Background()
	config, err := config.LoadConfig(*config_path)
	if err != nil {
		kingpin.FatalIfError(err, "Unable to load config file")
	}
	ctx.Config = config
	manager, err := crypto.NewClientCryptoManager(
		&ctx, []byte(config.Client_private_key))
	if err != nil {
		kingpin.FatalIfError(err, "Unable to parse config file")
	}

	exe, err := executor.NewClientExecutor(config)
	if err != nil {
		kingpin.FatalIfError(err, "Can not create executor.")
	}

	comm, err := http_comms.NewHTTPCommunicator(
		ctx,
		manager,
		exe,
		config.Client_server_urls,
	)
	if err != nil {
		kingpin.FatalIfError(err, "Can not create HTTPCommunicator.")
	}

	comm.Run()
}
