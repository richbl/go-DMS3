package dms3build

// structPlatform contains build environment/platform details
type structPlatform struct {
	envName     string
	DirName     string
	compileTags string
}

// structComponent contains component details
type structComponent struct {
	srcName        string // component source filename
	exeName        string // compiled filename
	dirName        string // location for compiled component in release folder
	configFilename string // relevant config (TOML) file
	compile        bool   // whether component should be compiled
}

// platform types
const (
	LinuxArm7 PlatformType = iota
	LinuxAMD64
)

// PlatformType represents all available platform types
type PlatformType int

// BuildEnv contains platform build details
var BuildEnv = []structPlatform{
	{
		envName:     "linuxArm7",
		DirName:     "linux_arm7",
		compileTags: "GOOS=linux GOARCH=arm GOARM=7",
	},
	{
		envName:     "linuxAMD64",
		DirName:     "linux_amd64",
		compileTags: "GOOS=linux GOARCH=amd64",
	},
}

var components = []structComponent{
	{
		srcName:        "go_dms3client.go",
		exeName:        "go_dms3client",
		dirName:        "dms3client",
		configFilename: "dms3client.toml",
		compile:        true,
	},
	{
		srcName:        "go_dms3server.go",
		exeName:        "go_dms3server",
		dirName:        "dms3server",
		configFilename: "dms3server.toml",
		compile:        true,
	},
	{
		srcName:        "go_dms3mail.go",
		exeName:        "go_dms3mail",
		dirName:        "dms3mail",
		configFilename: "dms3mail.toml",
		compile:        true,
	},
	{
		srcName:        "install_dms3.go",
		exeName:        "install_dms3",
		dirName:        "dms3build",
		configFilename: "dms3build.toml",
		compile:        true,
	},
	{
		srcName:        "dms3build/remote_installers/dms3client_remote_installer.go",
		exeName:        "dms3client_remote_installer",
		dirName:        "dms3build",
		configFilename: "",
		compile:        true,
	},
	{
		srcName:        "dms3build/remote_installers/dms3server_remote_installer.go",
		exeName:        "dms3server_remote_installer",
		dirName:        "dms3build",
		configFilename: "",
		compile:        true,
	},
	{
		srcName:        "",
		exeName:        "dms3libs",
		dirName:        "dms3libs",
		configFilename: "dms3libs.toml",
		compile:        false,
	},
}
