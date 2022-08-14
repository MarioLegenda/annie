package annie

import (
	"fmt"
	"testing"
)

func BenchmarkDockerComposeConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		annie, err := NewAnnie("docker-compose-config.yml")

		if err != nil {
			b.Fail()
		}

		annie.CannotBeEmpty("version").
			IsString("version").
			IsMap("services").
			StepInto("services").
			IsMap("api").
			StepInto("api").
			IsMap("build").
			StepInto("build").
			IsString("context").
			IsString("dockerfile").
			StepOut().
			IsString("env_file").
			IsArray("ports").
			IsArray("volumes").
			IsArray("depends_on")

		errs := annie.Errors()

		if len(errs) != 0 {
			fmt.Println(errs)
			b.Fail()
		}

		annie.Close()
	}
}
