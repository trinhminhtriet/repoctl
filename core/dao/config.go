package dao

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gookit/color"
	"github.com/trinhminhtriet/repoctl/core"
	"gopkg.in/yaml.v3"
)

var (
	DEFAULT_SHELL         = "bash -c"
	DEFAULT_SHELL_PROGRAM = "bash"
	ACCEPTABLE_FILE_NAMES = []string{"repoctl.yaml", "repoctl.yml", ".repoctl.yaml", ".repoctl.yml"}

	DEFAULT_THEME = Theme{
		Name:   "default",
		Stream: DefaultStream,
		Table:  DefaultTable,
		Tree:   DefaultTree,
		TUI:    DefaultTUI,
		Block:  DefaultBlock,
		Color:  core.Ptr(true),
	}

	DEFAULT_TARGET = Target{
		Name: "default",

		All: false,
		Cwd: false,

		Projects: []string{},
		Paths:    []string{},
		Tags:     []string{},

		TagsExpr: "",
	}

	DEFAULT_SPEC = Spec{
		Name:   "default",
		Output: "stream",

		Parallel: false,
		Forks:    4,

		IgnoreErrors:      false,
		IgnoreNonExisting: false,

		OmitEmptyRows:    false,
		OmitEmptyColumns: false,

		ClearOutput: true,
	}
)

type Config struct {
	// Internal
	EnvList        []string  `yaml:"-"`
	ImportData     []Import  `yaml:"-"`
	ThemeList      []Theme   `yaml:"-"`
	SpecList       []Spec    `yaml:"-"`
	TargetList     []Target  `yaml:"-"`
	ProjectList    []Project `yaml:"-"`
	TaskList       []Task    `yaml:"-"`
	Path           string    `yaml:"-"`
	Dir            string    `yaml:"-"`
	UserConfigFile *string   `yaml:"-"`
	ConfigPaths    []string  `yaml:"-"`
	Color          bool      `yaml:"-"`

	Shell         string `yaml:"shell"`
	SyncRemotes   *bool  `yaml:"sync_remotes"`
	SyncGitignore *bool  `yaml:"sync_gitignore"`
	ReloadTUI     *bool  `yaml:"reload_tui_on_change"`

	// Intermediate
	Env      yaml.Node `yaml:"env"`
	Import   yaml.Node `yaml:"import"`
	Themes   yaml.Node `yaml:"themes"`
	Specs    yaml.Node `yaml:"specs"`
	Targets  yaml.Node `yaml:"targets"`
	Projects yaml.Node `yaml:"projects"`
	Tasks    yaml.Node `yaml:"tasks"`
}

func (c *Config) GetContext() string {
	return c.Path
}

func (c *Config) GetContextLine() int {
	return -1
}

// Returns the config env list as a string splice in the form [key=value, key1=$(echo 123)]
func (c Config) GetEnvList() []string {
	var envs []string
	count := len(c.Env.Content)
	for i := 0; i < count; i += 2 {
		env := fmt.Sprintf("%v=%v", c.Env.Content[i].Value, c.Env.Content[i+1].Value)
		envs = append(envs, env)
	}

	return envs
}

func getUserConfigFile(userConfigPath string) *string {
	// Flag
	if userConfigPath != "" {
		if _, err := os.Stat(userConfigPath); err == nil {
			return &userConfigPath
		}
	}

	// Env
	val, present := os.LookupEnv("MANI_USER_CONFIG")
	if present {
		return &val
	}

	// Default
	defaultUserConfigDir, _ := os.UserConfigDir()
	defaultUserConfigPath := filepath.Join(defaultUserConfigDir, "repoctl", "config.yaml")
	if _, err := os.Stat(defaultUserConfigPath); err == nil {
		return &defaultUserConfigPath
	}

	return nil
}

// Function to read Repoctrl configs.
func ReadConfig(configFilepath string, userConfigPath string, colorFlag bool) (Config, error) {
	color := CheckUserColor(colorFlag)
	var configPath string

	userConfigFile := getUserConfigFile(userConfigPath)

	// Try to find config file in current directory and all parents
	if configFilepath != "" {
		filename, err := filepath.Abs(configFilepath)
		if err != nil {
			return Config{}, err
		}

		configPath = filename
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return Config{}, err
		}

		// Check first cwd and all parent directories, then if not found,
		// check if env variable MANI_CONFIG is set, and if not found
		// return no config found
		filename, err := core.FindFileInParentDirs(wd, ACCEPTABLE_FILE_NAMES)
		if err != nil {
			val, present := os.LookupEnv("MANI_CONFIG")
			if present {
				filename = val
			} else {
				return Config{}, err
			}
		}

		filename, err = core.ResolveTildePath(filename)
		if err != nil {
			return Config{}, err
		}

		filename, err = filepath.Abs(filename)
		if err != nil {
			return Config{}, err
		}

		configPath = filename
	}

	dat, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	// Found config, now try to read it
	var config Config

	config.Path = configPath
	config.Dir = filepath.Dir(configPath)
	config.UserConfigFile = userConfigFile
	config.Color = color

	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		re := ResourceErrors[Config]{Resource: &config, Errors: []error{err}}
		return config, FormatErrors(re.Resource, re.Errors)
	}

	// Set default shell command
	if config.Shell == "" {
		config.Shell = DEFAULT_SHELL
	} else {
		config.Shell = core.FormatShell(config.Shell)
	}

	// Set Sync Gitignore
	if config.SyncGitignore == nil {
		syncGitignore := true
		config.SyncGitignore = &syncGitignore
	}

	// Set Reload TUI
	if config.ReloadTUI == nil {
		reloadTUI := false
		config.ReloadTUI = &reloadTUI
	}

	// Set Sync Remote
	if config.SyncRemotes == nil {
		syncRemotes := false
		config.SyncRemotes = &syncRemotes
	}

	configResources, err := config.importConfigs()
	if err != nil {
		return config, err
	}

	config.TaskList = configResources.Tasks
	config.ProjectList = configResources.Projects
	config.ThemeList = configResources.Themes
	config.SpecList = configResources.Specs
	config.TargetList = configResources.Targets
	config.EnvList = configResources.Envs

	config.CheckConfigNoColor()

	for _, configPath := range configResources.Imports {
		config.ConfigPaths = append(config.ConfigPaths, configPath.Path)
	}

	// Set default theme if it's not set already
	_, err = config.GetTheme(DEFAULT_THEME.Name)
	if err != nil {
		config.ThemeList = append(config.ThemeList, DEFAULT_THEME)
	}

	// Set default spec if it's not set already
	_, err = config.GetSpec(DEFAULT_SPEC.Name)
	if err != nil {
		config.SpecList = append(config.SpecList, DEFAULT_SPEC)
	}

	// Set default target if it's not set already
	_, err = config.GetTarget(DEFAULT_TARGET.Name)
	if err != nil {
		config.TargetList = append(config.TargetList, DEFAULT_TARGET)
	}

	// Parse all tasks
	taskErrors := make([]ResourceErrors[Task], len(configResources.Tasks))
	for i := range configResources.Tasks {
		taskErrors[i].Resource = &configResources.Tasks[i]
		configResources.Tasks[i].ParseTask(config, &taskErrors[i])
	}

	var configErr = ""
	for _, taskError := range taskErrors {
		if len(taskError.Errors) > 0 {
			configErr = fmt.Sprintf("%s%s", configErr, FormatErrors(taskError.Resource, taskError.Errors))
		}
	}

	if configErr != "" {
		return config, &core.ConfigErr{Msg: configErr}
	}

	return config, nil
}

// Open repoctl config in editor
func (c Config) EditConfig() error {
	return openEditor(c.Path, -1)
}

func openEditor(path string, lineNr int) error {
	editor := os.Getenv("EDITOR")
	var args []string

	if lineNr > 0 {
		switch editor {
		case "nvim":
			args = []string{fmt.Sprintf("+%v", lineNr), path}
		case "vim":
			args = []string{fmt.Sprintf("+%v", lineNr), path}
		case "vi":
			args = []string{fmt.Sprintf("+%v", lineNr), path}
		case "emacs":
			args = []string{fmt.Sprintf("+%v", lineNr), path}
		case "nano":
			args = []string{fmt.Sprintf("+%v", lineNr), path}
		case "code": // visual studio code
			args = []string{"--goto", fmt.Sprintf("%s:%v", path, lineNr)}
		case "idea": // Intellij
			args = []string{"--line", fmt.Sprintf("%v", lineNr), path}
		case "subl": // Sublime
			args = []string{fmt.Sprintf("%s:%v", path, lineNr)}
		case "atom":
			args = []string{fmt.Sprintf("%s:%v", path, lineNr)}
		case "notepad-plus-plus":
			args = []string{"-n", fmt.Sprintf("%v", lineNr), path}
		default:
			args = []string{path}
		}
	} else {
		args = []string{path}
	}

	cmd := exec.Command(editor, args...)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// Open repoctl config in editor and optionally go to line matching the task name
func (c Config) EditTask(name string) error {
	configPath := c.Path
	if name != "" {
		task, err := c.GetTask(name)
		if err != nil {
			return err
		}
		configPath = task.context
	}

	dat, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	type ConfigTmp struct {
		Tasks yaml.Node
	}

	var configTmp ConfigTmp
	err = yaml.Unmarshal([]byte(dat), &configTmp)
	if err != nil {
		return err
	}

	lineNr := 0
	if name == "" {
		lineNr = configTmp.Tasks.Line - 1
	} else {
		for _, task := range configTmp.Tasks.Content {
			if task.Value == name {
				lineNr = task.Line
				break
			}
		}
	}

	return openEditor(configPath, lineNr)
}

// Open repoctl config in editor and optionally go to line matching the project name
func (c Config) EditProject(name string) error {
	configPath := c.Path
	if name != "" {
		project, err := c.GetProject(name)
		if err != nil {
			return err
		}
		configPath = project.context
	}

	dat, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	type ConfigTmp struct {
		Projects yaml.Node
	}

	var configTmp ConfigTmp
	err = yaml.Unmarshal([]byte(dat), &configTmp)
	if err != nil {
		return err
	}

	lineNr := 0
	if name == "" {
		lineNr = configTmp.Projects.Line - 1
	} else {
		for _, project := range configTmp.Projects.Content {
			if project.Value == name {
				lineNr = project.Line
				break
			}
		}
	}

	return openEditor(configPath, lineNr)
}

func InitMani(args []string, initFlags core.InitFlags) ([]Project, error) {
	// Choose to initialize repoctl in a different directory
	// 1. absolute or
	// 2. relative or
	// 3. working directory
	var configDir string
	if len(args) > 0 && filepath.IsAbs(args[0]) {
		// absolute path
		configDir = args[0]
	} else if len(args) > 0 {
		// relative path
		wd, err := os.Getwd()
		if err != nil {
			return []Project{}, err
		}
		configDir = filepath.Join(wd, args[0])
	} else {
		// working directory
		wd, err := os.Getwd()
		if err != nil {
			return []Project{}, err
		}
		configDir = wd
	}

	err := os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		return []Project{}, err
	}

	configPath := filepath.Join(configDir, "repoctl.yaml")
	if _, err := os.Stat(configPath); err == nil {
		return []Project{}, &core.AlreadyManiDirectory{Dir: configDir}
	}

	url, err := core.GetWdRemoteUrl(configDir)
	if err != nil {
		return []Project{}, err
	}

	rootName := filepath.Base(configDir)
	rootPath := "."
	rootUrl := url
	rootProject := Project{Name: rootName, Path: rootPath, Url: rootUrl}
	projects := []Project{rootProject}
	if initFlags.AutoDiscovery {
		prs, err := FindVCSystems(configDir)
		if err != nil {
			return []Project{}, err
		}
		RenameDuplicates(prs)

		projects = append(projects, prs...)
	}

	funcMap := template.FuncMap{
		"projectItem": func(name string, path string, url string) string {
			var txt = name + ":"

			if name != path {
				txt = txt + "\n    path: " + path
			}

			if url != "" {
				txt = txt + "\n    url: " + url
			}

			return txt
		},
	}

	tmpl, err := template.New("init").Funcs(funcMap).Parse(`projects:
  {{- range .}}
  {{ (projectItem .Name .Path .Url) }}
  {{ end }}
tasks:
  hello:
    desc: Print Hello World
    cmd: echo "Hello World"
`,
	)
	if err != nil {
		return []Project{}, err
	}

	// Create repoctl.yaml
	f, err := os.Create(configPath)
	if err != nil {
		return []Project{}, err
	}

	err = tmpl.Execute(f, projects)
	if err != nil {
		return []Project{}, err
	}

	err = f.Close()
	if err != nil {
		return []Project{}, err
	}

	// Update gitignore file if VCS set to git
	hasUrl := false
	for _, project := range projects {
		if project.Url != "" {
			hasUrl = true
			break
		}
	}

	if hasUrl && initFlags.SyncGitignore {
		// Add gitignore file
		gitignoreFilepath := filepath.Join(configDir, ".gitignore")
		if _, err := os.Stat(gitignoreFilepath); os.IsNotExist(err) {
			err := os.WriteFile(gitignoreFilepath, []byte(""), 0644)
			if err != nil {
				return []Project{}, err
			}
		}

		var projectNames []string
		for _, project := range projects {
			if project.Url == "" {
				continue
			}

			if project.Path == "." {
				continue
			}

			projectNames = append(projectNames, project.Path)
		}

		// Add projects to gitignore file
		err = UpdateProjectsToGitignore(projectNames, gitignoreFilepath)
		if err != nil {
			return []Project{}, err
		}
	}

	fmt.Println("\nInitialized repoctl repository in", configDir)
	fmt.Println("- Created repoctl.yaml")

	if hasUrl && initFlags.SyncGitignore {
		fmt.Println("- Created .gitignore")
	}

	return projects, nil
}

func RenameDuplicates(projects []Project) {
	projectNamesCount := make(map[string]int)
	// Find duplicate names
	for _, p := range projects {
		projectNamesCount[p.Name] += 1
	}

	// Rename duplicate projects
	for i, p := range projects {
		if projectNamesCount[p.Name] > 1 {
			projects[i].Name = p.Path
		}
	}
}

func CheckUserColor(colorFlag bool) bool {
	_, present := os.LookupEnv("NO_COLOR")
	if present || !colorFlag {
		color.Disable()
		return false
	}

	return true
}

func (c *Config) CheckConfigNoColor() {
	for _, env := range c.EnvList {
		name := strings.Split(env, "=")[0]
		if name == "NO_COLOR" {
			color.Disable()
		}
	}
}
