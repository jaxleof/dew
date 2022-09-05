package cmd

import (
	"log"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	Open.AddCommand(OpenGyms)
	Open.AddCommand(OpenMashup)
	Open.AddCommand(OpenStatus)
	Open.AddCommand(OpenRandom)
	rootCmd.AddCommand(Open)
}

var Open = &cobra.Command{
	Use:   "open",
	Short: "a shortcut of opening codeforces website",
	Run: func(cmd *cobra.Command, args []string) {
		OpenWebsite("https://codeforces.com")
	},
}

var OpenGyms = &cobra.Command{
	Use:   "gym",
	Short: "a shortcut of opening codeforces gyms",
	Run: func(cmd *cobra.Command, args []string) {
		OpenWebsite("https://codeforces.com/gyms")
	},
}

var OpenMashup = &cobra.Command{
	Use:   "mashup",
	Short: "a shortcut of opening codeforces mashup",
	Run: func(cmd *cobra.Command, args []string) {
		OpenWebsite("https://codeforces.com/mashups")
	},
}

var OpenStatus = &cobra.Command{
	Use:   "status",
	Short: "a shortcut of opening codeforces status",
	Run: func(cmd *cobra.Command, args []string) {
		OpenWebsite("https://codeforces.com/problemset/status?my=on")
	},
}

var OpenRandom = &cobra.Command{
	Use:   "random",
	Short: "a shortcut of opening codeforces status",
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigFile("./codeforces/config.yaml")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}
		var name = viper.GetString("random")
		OpenWebsite("https://codeforces.com/problemset/problem/" + name[:len(name)-1] + "/" + name[len(name)-1:])
	},
}

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func OpenWebsite(path string) {
	run, ok := commands[runtime.GOOS]
	if !ok {
		log.Fatalf("don't know how to open things on %s platform", runtime.GOOS)
	}
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", "start", path)
	} else {
		cmd = exec.Command(run, path)
	}
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func OpenRandomFunc(info problemInfo) {
	OpenWebsite("https://codeforces.com/problemset/problem/" + strconv.Itoa(info.ContestId) + "/" + info.Index)
}
