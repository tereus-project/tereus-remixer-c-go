package remixer

import (
	"os"
	"strings"
	"testing"
)

func testRemix(t *testing.T, source string, target string) {
	source = strings.TrimSpace(source)
	target = strings.TrimSpace(target)

	dir := t.TempDir()
	sourceFile := dir + "/test.go"
	err := os.WriteFile(sourceFile, []byte(source), 0644)
	if err != nil {
		t.Error(err)
	}

	output, err := Remix(sourceFile)
	if err != nil {
		t.Error(err)
	}

	output = strings.TrimSpace(output)
	if output != target {
		t.Errorf("Expected %s, got %s", target, output)
	}
}

func TestEmptyFunction(t *testing.T) {
	source := `
int main() {

}
`

	target := `
package main

func main() {

}
`

	testRemix(t, source, target)
}

func TestMainWithArgcAndArgv(t *testing.T) {
	source := `
int main(int argc, char **argv) {

}
`

	target := `
package main

import (
	"os"
)

func main() {
	argc := len(os.Args)
	argv := os.Args
}
`

	testRemix(t, source, target)
}

func TestFunctionCall(t *testing.T) {
	source := `
int add(int a, int b) {
	return a + b;
}

int main() {
	add(1, 2);
}
`

	target := `
package main

func add(a int, b int) int {
	return a + b
}

func main() {
	add(1, 2)
}
`

	testRemix(t, source, target)
}

func TestIfCondition(t *testing.T) {
	source := `
int main() {
	int a = 1;

	if (a >= 1) {
		a = 2;
	}

	return 0;
}
`

	target := `
package main

import (
	"os"
)

func main() {
	a := 1
	if a >= 1 {
		a = 2
	}
	os.Exit(0)
}
`

	testRemix(t, source, target)
}

func TestUnaryPreExpression(t *testing.T) {
	source := `
int main() {
	int a = -1;
	return 0;
}
`

	target := `
package main

import (
	"os"
)

func main() {
	a := -1
	os.Exit(0)
}
`

	testRemix(t, source, target)
}

func TestUnaryPostExpression(t *testing.T) {
	source := `
int main() {
	int a = 1;
	a++;
	return 0;
}
`

	target := `
package main

import (
	"os"
)

func main() {
	a := 1
	a++
	os.Exit(0)
}
`

	testRemix(t, source, target)
}
