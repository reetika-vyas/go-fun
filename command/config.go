package commander

import (
	"os"
	"io/ioutil"
	"fmt"
)

/* PSSH */
const CLUSTER_PATH = "/tmp/clusters"
const OUTPUT_PATH = "/tmp/output"
const ERROR_PATH = "/tmp/error"
const CONSOLE_FILE = CLUSTER_PATH + "/console.txt"

const RELEASE_FILE = "/Users/amanpreet.singh/Documents/release.txt"
const DEFAULT_PERM = os.FileMode(0644)     //Owner RW,Group R,Other R
const DIR_DEFAULT_PERM = os.FileMode(0755) //Owner RWX,Group RX,Other RX
const DEFAULT_PARALELISM = 50

const DEBUG_FILE string = "/tmp/kohandebug"

/* Date/Time */
const LAYOUT = "Jan 2, 2006 at 3:04pm (MST)"

func DebugControl(flag bool) {
	if flag {
		PrintSkyBlue("Enabling Debug Mode")
		ioutil.WriteFile(DEBUG_FILE, []byte{}, DEFAULT_PERM)
	} else {
		PrintRed("Disabling Debug Mode")
		os.Remove(DEBUG_FILE)
	}
	PrintYellow(fmt.Sprintf("Debug Mode: %v", IsDebugMode()))
}

func IsDebugMode() bool {
	return PathExists(DEBUG_FILE)
}
