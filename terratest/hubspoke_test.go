package hubspoketest

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

type moduleConfigs struct {
	TerraFiles []string
	SetFlags   struct{
		VarDefs string
		MainFile string
		BackendConfigFile string
		EnvVars string
		PlanFile string
	}
}

func (c *moduleConfigs) Configure() {

}
var mainConfs moduleConfigs


func TestHubSpoke(t *testing.T) {
	readDir()
	fmt.Println(mainConfs.TerraFiles)

	for _, v := range mainConfs.TerraFiles {
		fmt.Printf("The terraform files in this moduel directory are: ---- %v ----\n", v)
		if v == "main.tf" {
			mainConfs.SetFlags.MainFile = v
		} else if v == "variables.tf" {
			mainConfs.SetFlags.VarDefs = v
		}
	}
	fmt.Println(mainConfs.SetFlags.MainFile)
	fmt.Println(mainConfs.SetFlags.VarDefs)
}

func TestValidateSyntax(t *testing.T) {

}



//func readDir reads the directory for main.tf and variables.tf files
func readDir() {

	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		getconf, err := filepath.Glob("*.tf")

		if getconf != nil {
			mainConfs.TerraFiles = getconf
		}

		if err != nil {
			panic(err)
		}

		return nil

	})
	if err != nil {
		log.Fatalf("Unable to parse directory for conf folder, see error: %v", err)
	}

}
