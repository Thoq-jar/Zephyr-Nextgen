package main

import (
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"
)

func banner() {
	banner := `
+----------------------------------+
+ ▗▄▄▄▄▖▗▄▄▄▖▗▄▄▖ ▗▖ ▗▖▗▖  ▗▖▗▄▄▖  +
+    ▗▞▘▐▌   ▐▌ ▐▌▐▌ ▐▌ ▝▚▞▘ ▐▌ ▐▌ +
+  ▗▞▘  ▐▛▀▀▘▐▛▀▘ ▐▛▀▜▌  ▐▌  ▐▛▀▚▖ +
+ ▐▙▄▄▄▖▐▙▄▄▖▐▌   ▐▌ ▐▌  ▐▌  ▐▌ ▐▌ +
+----------------------------------+

You are using: Zephyr Utils.

`
	fmt.Print(banner)
}

func run(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("command execution failed: %w\nOutput: %s", err, output)
	}
	return string(output), nil
}

func buildProject() {
	output, err := run("go", "mod", "tidy")
	if err != nil {
		fmt.Printf("%v\nOutput: %v\n", err, output)
		return
	}

	files, err := filepath.Glob("src/*")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if len(files) == 0 {
		fmt.Println("No files found in src/*")
		return
	}

	args := append([]string{"build", "-o", "Zephyr"}, files...)
	output, err = run("go", args...)
	if err != nil {
		fmt.Printf("%v\nOutput: %v\n", err, output)
		return
	}
}

func runProject() {
	buildProject()
	output, err := run("./Zephyr")
	if err != nil {
		fmt.Printf("%v\nOutput: %v\n", err, output)
		return
	}
	fmt.Printf("%v\n", output)
}

func setupProject() {
	output, err := run("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@latest")
	if err != nil {
		fmt.Printf("%v\nOutput: %v\n", err, output)
		return
	}
	fmt.Printf("%v\n", output)
}

func lintProject() {
	output, err := run("golangci-lint", "run")
	if err != nil {
		fmt.Printf("%v\nOutput: %v\n", err, output)
		return
	}
	if output == "" {
		fmt.Println("All good!")
	} else {
		fmt.Printf("%v\n", output)
	}
}

func cleanProject() {
	output, err := run("go", "clean")
	if err != nil {
		fmt.Printf("%v\n: %v\n", err, output)
		return
	}
	cleanTwo, err := run("rm", "-rf", "Zephyr")
	if err != nil {
		fmt.Printf("%v\n: %v\n", err, output)
		return
	}
	fmt.Printf("%v\n", cleanTwo)
}

func main() {
	buildFlag := flag.Bool("build", false, "Build the project")
	runFlag := flag.Bool("run", false, "Run the project")
	setupFlag := flag.Bool("setup", false, "Setup the project")
	lintFlag := flag.Bool("lint", false, "Lint the project")
	cleanFlag := flag.Bool("clean", false, "Clean the project")
	flag.Parse()

	if !*buildFlag && !*runFlag && !*setupFlag && !*lintFlag && !*cleanFlag {
		fmt.Println("Usage: --build to build the project, " +
			"--run to build and run the project, " +
			"--setup to configure the project or " +
			"--lint to lint the project")
		return
	}

	banner()

	if *buildFlag {
		fmt.Println("Building...")
		buildProject()
		fmt.Println("Done!")
	}

	if *runFlag {
		fmt.Println("Launching...")
		runProject()
		fmt.Println("Done!")
	}

	if *setupFlag {
		fmt.Println("Setting up...")
		setupProject()
		fmt.Println("Done!")
	}

	if *lintFlag {
		fmt.Println("Linting...")
		lintProject()
	}

	if *cleanFlag {
		fmt.Println("Cleaning...")
		cleanProject()
		fmt.Println("Done!")
	}
}
