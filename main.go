package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:    "wallet",
				Aliases: []string{"w"},
				Usage:   "options for wallet",
				Subcommands: []*cli.Command{
					{
						Name:    "create",
						Aliases: []string{"c"},
						Usage:   "create new wallet",
						Action:  ActionWalletCreate,
					},
					{
						Name:    "balance",
						Aliases: []string{"b"},
						Usage:   "show wallet balance",
						Action:  ActionWalletBalance,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "address",
								Usage:    "wallet address (hex)",
								Required: true,
								Aliases:  []string{"a"},
							},
							&cli.StringFlag{
								Name:     "node_url",
								Usage:    "ethereum json-rpc node url",
								Required: true,
								Aliases:  []string{"n"},
							},
						},
					},
				},
			},
			{
				Name:    "contract",
				Aliases: []string{"c"},
				Usage:   "options for contract",
				Subcommands: []*cli.Command{
					{
						Name:    "deploy",
						Aliases: []string{"d"},
						Usage:   "deploy smart contract on-chain",
						Action:  ActionContractDeploy,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "private_key",
								Usage:    "private key of a wallet address that will deploy the contract",
								Required: true,
								Aliases:  []string{"k"},
							},
							&cli.StringFlag{
								Name:     "address",
								Usage:    "wallet address",
								Required: true,
								Aliases:  []string{"a"},
							},
							&cli.StringFlag{
								Name:     "node_url",
								Usage:    "ethereum json-rpc node url",
								Required: true,
								Aliases:  []string{"n"},
							},
						},
					},
					{
						Name:    "status",
						Aliases: []string{"s"},
						Usage:   "status of smart contract on-chain",
						Action:  ActionContractStatus,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "address",
								Usage:    "address of a deployed contract",
								Required: true,
								Aliases:  []string{"a"},
							},
							&cli.StringFlag{
								Name:     "node_url",
								Usage:    "ethereum json-rpc node url",
								Required: true,
								Aliases:  []string{"n"},
							},
							&cli.StringFlag{
								Name:     "private_key",
								Usage:    "private key of a wallet address that will deploy the contract",
								Required: true,
								Aliases:  []string{"k"},
							},
						},
					},
				},
			},
			{
				Name:    "file",
				Aliases: []string{"m"},
				Usage:   "options for file",
				Subcommands: []*cli.Command{
					{
						Name:    "add",
						Aliases: []string{"a"},
						Usage:   "add proof of file existence on-chain",
						Action:  ActionFileAdd,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "private_key",
								Usage:    "private key of a wallet address that will deploy the contract",
								Required: true,
								Aliases:  []string{"k"},
							},
							&cli.StringFlag{
								Name:      "file",
								Usage:     "file proof of which will be stored on chain",
								Required:  true,
								Aliases:   []string{"f"},
								TakesFile: true,
							},
							&cli.StringFlag{
								Name:     "node_url",
								Usage:    "ethereum json-rpc node url",
								Required: true,
								Aliases:  []string{"n"},
							},
							&cli.StringFlag{
								Name:     "contract_address",
								Usage:    "address of a deployed contract",
								Required: true,
								Aliases:  []string{"a"},
							},
							&cli.StringFlag{
								Name:     "contract_owner",
								Usage:    "address of a contract owner",
								Required: true,
								Aliases:  []string{"o"},
							},
						},
					},
					{
						Name:    "verify",
						Aliases: []string{"v"},
						Usage:   "verify proof of file on-chain",
						Action:  ActionFileVerify,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:      "file",
								Usage:     "file of which proof will be stored on chain",
								Required:  true,
								Aliases:   []string{"f"},
								TakesFile: true,
							},
							&cli.StringFlag{
								Name:     "contract_address",
								Usage:    "address of a deployed contract",
								Required: true,
								Aliases:  []string{"a"},
							},
							&cli.StringFlag{
								Name:     "node_url",
								Usage:    "ethereum json-rpc node url",
								Required: true,
								Aliases:  []string{"n"},
							},
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
