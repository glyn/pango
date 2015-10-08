package parse_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/glyn/pango/packages"
	"github.com/glyn/pango/parse"
)

func TestParseAll(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	par := parse.New("github.com/cloudfoundry-incubator/guardian",
		filepath.Join(pwd, "../fixtures/guardian"))
	graph := par.Parse()

	// Check the packages are correct.
	s := graph.Packages()
	expected := guardianPackages
	if !s.Equals(expected) {
		t.Errorf("Expected %v to equal %v", s, expected)
	}

	// Check a representative sample of imports are correct.
	guardianGraphSample.Walk(func(p Pkg, i PSet) {
		j, ok := graph.Imports(p)
		if !ok {
			t.Fail()
		}
		if !i.Equals(j) {
			t.Errorf("The imports of %s were %v rather than the expected %v", p, j, i)
		}
	})
}

var guardianPackages PSet = NewPSet("github.com/cloudfoundry-incubator/guardian/cmd/guardian",
	"github.com/cloudfoundry-incubator/guardian/rundmc",
	"github.com/cloudfoundry-incubator/guardian/rundmc/runrunc",
	"github.com/cloudfoundry-incubator/guardian/rundmc/iodaemon",
	"github.com/cloudfoundry-incubator/guardian/gardener",
	"github.com/cloudfoundry-incubator/guardian/rundmc/process_tracker",
	"github.com/cloudfoundry-incubator/guardian/rundmc/process_tracker/fake_msg_sender",
	"github.com/cloudfoundry-incubator/guardian/kawasaki",
	"github.com/cloudfoundry-incubator/guardian/rundmc/iodaemon/cmd/iodaemon",
	"github.com/cloudfoundry-incubator/guardian/rundmc/depot",
	"github.com/cloudfoundry-incubator/guardian/rundmc/runrunc/fakes",
	"github.com/cloudfoundry-incubator/guardian/rundmc/process_tracker/fake_signaller",
	"github.com/cloudfoundry-incubator/guardian/rundmc/depot/fakes",
	"github.com/cloudfoundry-incubator/guardian/rundmc/fakes",
	"github.com/cloudfoundry-incubator/guardian/gardener/fakes",
	"github.com/cloudfoundry-incubator/guardian/rundmc/process_tracker/writer",
	"github.com/cloudfoundry-incubator/guardian/log",
	"github.com/cloudfoundry-incubator/guardian/gqt",
	"github.com/cloudfoundry-incubator/guardian/script",
	"github.com/cloudfoundry-incubator/guardian/rundmc/iodaemon/cmd",
	"github.com/cloudfoundry-incubator/guardian/rundmc/iodaemon/test_print_signals",
	"github.com/cloudfoundry-incubator/guardian/cmd",
	"github.com/cloudfoundry-incubator/guardian/rundmc/process_tracker/fake_process_tracker",
	"github.com/cloudfoundry-incubator/guardian/gqt/runner",
	"github.com/cloudfoundry-incubator/guardian",
	"github.com/cloudfoundry-incubator/guardian/rundmc/iodaemon/link")

var guardianGraphSample PGraph = NewPGraph().
	AddImports("github.com/cloudfoundry-incubator/guardian/cmd/guardian",
	NewPSet("github.com/cloudfoundry-incubator/cf-debug-server",
		"github.com/cloudfoundry-incubator/cf-lager",
		"github.com/cloudfoundry-incubator/guardian/log",
		"github.com/cloudfoundry-incubator/guardian/rundmc/runrunc",
		"github.com/cloudfoundry-incubator/guardian/rundmc",
		"github.com/cloudfoundry-incubator/goci/specs",
		"github.com/nu7hatch/gouuid",
		"github.com/pivotal-golang/lager",
		"github.com/cloudfoundry-incubator/guardian/rundmc/depot",
		"github.com/cloudfoundry-incubator/guardian/gardener",
		"github.com/cloudfoundry-incubator/guardian/kawasaki",
		"github.com/cloudfoundry/gunk/command_runner/linux_command_runner",
		"github.com/cloudfoundry-incubator/goci",
		"github.com/cloudfoundry-incubator/garden/server",
		"github.com/cloudfoundry-incubator/guardian/rundmc/process_tracker")).
	AddImports("github.com/cloudfoundry-incubator/guardian/rundmc",
	NewPSet("github.com/cloudfoundry/gunk/command_runner/fake_command_runner",
		"github.com/cloudfoundry-incubator/goci",
		"github.com/cloudfoundry-incubator/goci/specs",
		"github.com/cloudfoundry/gunk/command_runner/fake_command_runner/matchers",
		"github.com/cloudfoundry-incubator/guardian/gardener",
		"github.com/cloudfoundry-incubator/garden",
		"github.com/pivotal-golang/lager",
		"github.com/cloudfoundry-incubator/guardian/rundmc/depot",
		"github.com/cloudfoundry-incubator/guardian/log",
		"github.com/cloudfoundry-incubator/guardian/rundmc/fakes",
		"github.com/onsi/ginkgo",
		"github.com/cloudfoundry-incubator/guardian/rundmc/runrunc",
		"github.com/onsi/gomega",
		"github.com/cloudfoundry/gunk/command_runner")).
	AddImports("github.com/cloudfoundry-incubator/guardian/kawasaki",
	NewPSet()).
	AddImports("github.com/cloudfoundry-incubator/guardian/rundmc/fakes",
	NewPSet("github.com/cloudfoundry-incubator/guardian/rundmc",
		"github.com/cloudfoundry-incubator/guardian/rundmc/depot",
		"github.com/cloudfoundry-incubator/garden",
		"github.com/cloudfoundry-incubator/goci",
		"github.com/cloudfoundry-incubator/guardian/gardener"))
