/*

Copyright (C) 2020-2021  Daniele Rondina <geaaru@sabayonlinux.org>
Credits goes also to Gogs authors, some code portions and re-implemented design
are also coming from the Gogs project, which is using the go-macaron framework
and was really source of ispiration. Kudos to them!

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

*/
package specs

type LxdCEnvironment struct {
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	File    string `json:"-" yaml:"-"`

	TemplateEngine LxdCTemplateEngine `json:"template_engine,omitempty" yaml:"template_engine,omitempty"`

	Projects []LxdCProject `json:"projects" yaml:"projects"`

	Commands             []LxdCCommand `json:"commands,omitempty" yaml:"commands,omitempty"`
	IncludeCommandsFiles []string      `json:"include_commands_files,omitempty" yaml:"include_commands_files,omitempty"`

	Profiles             []LxdCProfile `json:"profiles,omitempty" yaml:"profiles,omitempty"`
	IncludeProfilesFiles []string      `json:"include_profiles_files,omitempty" yaml:"include_profiles_files,omitempty"`
	Networks             []LxdCNetwork `json:"networks,omitempty" yaml:"networks,omitempty"`
	IncludeNetworkFiles  []string      `json:"include_networks_files,omitempty" yaml:"include_networks_files,omitempty"`
	Storages             []LxdCStorage `json:"storages,omitempty" yaml:"storages,omitempty"`
	IncludeStorageFiles  []string      `json:"include_storage_files,omitempty" yaml:"include_storage_files,omitempty"`
}

type LxdCProfile struct {
	Name        string                       `json:"name" yaml:"name"`
	Description string                       `json:"description" yaml:"description"`
	Config      map[string]string            `json:"config" yaml:"config"`
	Devices     map[string]map[string]string `json:"devices" yaml:"devices"`
}

type LxdCNetwork struct {
	Name        string            `json:"name" yaml:"name"`
	Type        string            `json:"type" yaml:"type"`
	Description string            `json:"description" yaml:"description"`
	Config      map[string]string `json:"config" yaml:"config"`
}

type LxdCStorage struct {
	Name          string            `json:"name" yaml:"name"`
	Driver        string            `json:"driver" yaml:"driver"`
	Description   string            `json:"description,omitempty" yaml:"description,omitempty"`
	Documentation string            `json:"documentation,omitempty" yaml:"documentation,omitempty"`
	Config        map[string]string `json:"config" yaml:"config"`
}

type LxdCHook struct {
	Event      string   `json:"event" yaml:"event"`
	Node       string   `json:"node" yaml:"node"`
	Commands   []string `json:"commands,omitempty" yaml:"commands,omitempty"`
	Out2Var    string   `json:"out2var,omitempty" yaml:"out2var,omitempty"`
	Err2Var    string   `json:"err2var,omitempty" yaml:"err2var,omitempty"`
	Entrypoint []string `json:"entrypoint,omitempty" yaml:"entrypoint,omitempty"`
	Flags      []string `json:"flags,omitempty" yaml:"flags,omitempty"`
	Disable    bool     `json:"disable,omitempty" yaml:"disable,omitempty"`
}

type LxdCHooks struct {
	Hooks []LxdCHook `json:"hooks,omitempty" yaml:"hooks,omitempty"`
}

type LxdCTemplateEngine struct {
	Engine string   `json:"engine" yaml:"engine"`
	Opts   []string `json:"opts,omitempty" yaml:"opts,omitempty"`
}

type LxdCProject struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	IncludeGroupFiles []string `json:"include_groups_files,omitempty" yaml:"include_groups_files,omitempty"`
	IncludeEnvFiles   []string `json:"include_env_files,omitempty" yaml:"include_env_files,omitempty"`
	IncludeHooksFiles []string `json:"include_hooks_files,omitempty" yaml:"include_hooks_files,omitempty"`

	Environments []LxdCEnvVars `json:"vars,omitempty" yaml:"vars,omitempty"`

	Groups      []LxdCGroup `json:"groups" yaml:"groups"`
	NodesPrefix string      `json:"nodes_prefix,omitempty" yaml:"nodes_prefix,omitempty"`

	Hooks           []LxdCHook           `json:"hooks" yaml:"hooks"`
	ConfigTemplates []LxdCConfigTemplate `json:"config_templates,omitempty" yaml:"config_templates,omitempty"`
}

type LxdCProjectSanitized struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	IncludeGroupFiles []string `json:"include_groups_files,omitempty" yaml:"include_groups_files,omitempty"`
	IncludeEnvFiles   []string `json:"include_env_files,omitempty" yaml:"include_env_files,omitempty"`
	IncludeHooksFiles []string `json:"include_hooks_files,omitempty" yaml:"include_hooks_files,omitempty"`

	Groups      []LxdCGroup `json:"groups" yaml:"groups"`
	NodesPrefix string      `json:"nodes_prefix,omitempty" yaml:"nodes_prefix,omitempty"`

	Hooks           []LxdCHook           `json:"hooks" yaml:"hooks"`
	ConfigTemplates []LxdCConfigTemplate `json:"config_templates,omitempty" yaml:"config_templates,omitempty"`
}

type LxdCGroup struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Connection  string `json:"connection,omitempty" yaml:"connection,omitempty"`

	CommonProfiles []string          `json:"common_profiles,omitempty" yaml:"common_profiles,omitempty"`
	Config         map[string]string `json:"config,omitempty" yaml:"config,omitempty"`

	Ephemeral bool `json:"ephemeral,omitempty" yaml:"ephemeral,omitempty"`

	Nodes       []LxdCNode `json:"nodes" yaml:"nodes"`
	NodesPrefix string     `json:"nodes_prefix,omitempty" yaml:"nodes_prefix,omitempty"`

	Hooks             []LxdCHook           `json:"hooks" yaml:"hooks"`
	IncludeHooksFiles []string             `json:"include_hooks_files,omitempty" yaml:"include_hooks_files,omitempty"`
	ConfigTemplates   []LxdCConfigTemplate `json:"config_templates,omitempty" yaml:"config_templates,omitempty"`
}

type LxdCEnvVars struct {
	EnvVars map[string]interface{} `json:"envs,omitempty" yaml:"envs,omitempty"`
}

type LxdCNode struct {
	Name              string `json:"name" yaml:"name"`
	NamePrefix        string `json:"name_prefix,omitempty" yaml:"name_prefix,omitempty"`
	ImageSource       string `json:"image_source" yaml:"image_source"`
	ImageRemoteServer string `json:"image_remote_server,omitempty" yaml:"image_remote_server,omitempty"`

	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Config map[string]string `json:"config,omitempty" yaml:"config,omitempty"`

	SourceDir string `json:"source_dir,omitempty" yaml:"source_dir,omitempty"`

	Entrypoint []string `json:"entrypoint,omitempty" yaml:"entrypoint,omitempty"`

	ConfigTemplates []LxdCConfigTemplate `json:"config_templates,omitempty" yaml:"config_templates,omitempty"`
	SyncResources   []LxdCSyncResource   `json:"sync_resources,omitempty" yaml:"sync_resources,omitempty"`
	Profiles        []string             `json:"profiles,omitempty" yaml:"profiles,omitempty"`

	Hooks             []LxdCHook `json:"hooks" yaml:"hooks"`
	IncludeHooksFiles []string   `json:"include_hooks_files,omitempty" yaml:"include_hooks_files,omitempty"`
}

type LxdCConfigTemplate struct {
	Source      string `json:"source" yaml:"source"`
	Destination string `json:"dst" yaml:"dst"`
}

type LxdCSyncResource struct {
	Source      string `json:"source" yaml:"source"`
	Destination string `json:"dst" yaml:"dst"`
}

type LxdCCommand struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Project     string `json:"project" yaml:"project"`
	ApplyAlias  bool   `json:"apply_alias,omitempty" yaml:"apply_alias,omitempty"`

	SkipSync bool `json:"skip_sync,omitempty" yaml:"skip_sync,omitempty"`
	Destroy  bool `json:"destroy,omitempty" yaml:"destroy,omitempty"`

	NodesPrefix string `json:"nodes_prefix,omitempty" yaml:"nodes_prefix,omitempty"`

	EnableFlags  []string `json:"enable_flags,omitempty" yaml:"enable_flags,omitempty"`
	DisableFlags []string `json:"disable_flags,omitempty" yaml:"disable_flags,omitempty"`

	EnableGroups  []string `json:"enable_groups,omitempty" yaml:"enable_groups,omitempty"`
	DisableGroups []string `json:"disable_groups,omitempty" yaml:"disable_groups,omitempty"`

	Envs     LxdCEnvVars `json:"envs,omitempty" yaml:"envs,omitempty"`
	VarFiles []string    `json:"vars_files,omitempty" yaml:"vars_files,omitempty"`

	IncludeHooksFiles []string `json:"include_hooks_files,omitempty" yaml:"include_hooks_files,omitempty"`
}
