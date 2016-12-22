
/package/kubectl/generate.go

`Generator` is an interface for things that can generate API
objects from input parameters. The `Generate` method creates
an API object given a set of parameters. The `ParamNames` method
returns the list of parameters that this generator uses.

    type Generator interface {
      Generate(params map[string]interface{}) (runtime.Object, error)
      ParamNames() []GeneratorParam
    }

`StructuredGenerator` is an interface for things that can generate API objects
without using parameters. The `StructuredGenerator` method creates an API
object using pre-configured parameters.

    type StructuredGenerator interface {
      StructuredGenerate() (runtime.Object, error)
    }

The `Generator` and `StructuredGenerator` interfaces are declared
in `/package/kubectl/generate.go`.

What types implement the Generator interface?
What types implement the StructuredGenerator interface?

There are many types that implement these interfaces. Here's
one of them.

## SecretForDockerRegistryGeneratorV1

The `SecretForDockerRegistryGeneratorV1` type implements the
`Generator` and `StructuredGenerator` interfaces.

Here's the type declaration in `/pkg/kubectl/secret_for_docker_registry.go`:

    type SecretForDockerRegistryGeneratorV1 struct {
      Name string
      Username string
      Email string
      Password string
      Server string
    }

This line of code verifies that the `SecretForDockerRegistryGeneratorV1` type
implements the `Generator` interface:

    var _ Generator = &SecretForDockerRegistryGeneratorV1{}

TODO: Understand better how this verification works.

Here are the method definitions for the `Genererator` interface:

    func (s SecretForDockerRegistryGeneratorV1) Generate(genericParams map[string]interface{}) (runtime.Object, error) {
      err := ValidateParams(s.ParamNames(), genericParams)
      ...
    }

    func (s SecretForDockerRegistryGeneratorV1) ParamNames() []GeneratorParam {
      return []GeneratorParam{
        {"name", true},
        {"docker-username", true},
        {"docker-email", true},
        {"docker-password", true},
        {"docker-server", true},
    }

Here is the method definition for the `StructuredGenererator` interface:

    func (s SecretForDockerRegistryGeneratorV1) StructuredGenerate() (runtime.Object, error) {
      if err := s.validate(); err != nil {
      ...
    }

The `SecretForDockerRegistryGeneratorV1` type also implements the `validate` method,
which is not in either of the interfaces:

    func (s SecretForDockerRegistryGeneratorV1) validate() error {
      if len(s.Name) == 0 {
        return fmt.Errorf("name must be specified")
      }
      ...
    }

There's one more function in `/pkg/kubectl/secret_for_docker_registry.go`.

    func handleDockercfgContent(username, password, email, server string) ([]byte, error) {
      dockercfgAuth := credentialprovider.DockerConfigEntry{
        Username: username,
        Password: password,
        Email:    email,
      }

      dockerCfg := map[string]credentialprovider.DockerConfigEntry{server: dockercfgAuth}
      return json.Marshal(dockerCfg)
    }




























## /pkg/kubectl/secret_for_docker_registry.go


