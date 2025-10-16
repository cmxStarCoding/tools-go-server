package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func InitCmd() {
	app := &cli.App{
		Name:  "mycli",                             // 设置命令行应用的名称
		Usage: "A simple command-line application", // 命令行应用的描述
		Commands: []*cli.Command{
			{
				Name:    "greet",                 // 子命令名称
				Aliases: []string{"hello", "hi"}, // 子命令别名
				Usage:   "Greet someone",         // 子命令描述
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",          // 命令行参数名称
						Aliases:  []string{"n"},   // 参数的别名
						Usage:    "Name to greet", // 参数的说明
						Required: true,            // 是否是必需的
					},
				},
				Action: func(c *cli.Context) error {
					// 获取用户输入的 name 参数
					name := c.String("name")
					fmt.Printf("Hello, %s!\n", name)
					return nil
				},
			},
			{
				Name:  "sum",             // 另一个子命令
				Usage: "Sum two numbers", // 子命令描述
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "a", // 第一个数字
						Usage:    "First number",
						Required: true, // 必须提供
					},
					&cli.IntFlag{
						Name:     "b", // 第二个数字
						Usage:    "Second number",
						Required: true, // 必须提供
					},
				},
				Action: func(c *cli.Context) error {
					// 获取 a 和 b 的值
					a := c.Int("a")
					b := c.Int("b")
					fmt.Printf("The sum of %d and %d is %d\n", a, b, a+b)
					return nil
				},
			},
		},
	}

	// 启动 CLI 应用
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error running the application:", err)
	}

}
