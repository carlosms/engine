// +build integration

package cmd

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	cmdtest "github.com/src-d/engine/cmd/test-utils"
	"github.com/src-d/engine/docker"

	"github.com/stretchr/testify/suite"
)

type ComponentsTestSuite struct {
	cmdtest.IntegrationSuite
	testDir string
}

func TestComponentsTestSuite(t *testing.T) {
	s := ComponentsTestSuite{}
	suite.Run(t, &s)
}

func (s *ComponentsTestSuite) SetupTest() {
	var err error
	s.testDir, err = ioutil.TempDir("", "components-test")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *ComponentsTestSuite) TearDownTest() {
	s.RunStop(context.Background())
	os.RemoveAll(s.testDir)
}

func (s *ComponentsTestSuite) TestListStopped() {
	require := s.Require()

	out, err := s.RunCommand(context.TODO(), "components", "list")
	require.NoError(err)

	expected := regexp.MustCompile(
		`^IMAGE                                INSTALLED    RUNNING    CONTAINER NAME
bblfsh/bblfshd:\S+ +(yes|no) +no         srcd-cli-bblfshd
bblfsh/web:\S+ +(yes|no) +no         srcd-cli-bblfsh-web
srcd/cli-daemon:\S+ +(yes|no) +no         srcd-cli-daemon
srcd/gitbase-web:\S+ +(yes|no) +no         srcd-cli-gitbase-web
srcd/gitbase:\S+ +(yes|no) +no         srcd-cli-gitbase
$`)

	s.Regexp(expected, out.String())
}

func (s *ComponentsTestSuite) TestListInit() {
	require := s.Require()

	_, err := s.RunInit(context.TODO(), s.testDir)
	require.NoError(err)

	out, err := s.RunCommand(context.TODO(), "components", "list")
	require.NoError(err)

	expected := regexp.MustCompile(`srcd/cli-daemon:\S+ +yes +yes +srcd-cli-daemon`)
	s.Regexp(expected, out.String())
}

func (s *ComponentsTestSuite) TestInstall() {
	require := s.Require()

	out, err := s.RunCommand(context.TODO(), "components", "list")
	require.NoError(err)

	// Get the exact image:version of gitbase
	exp := regexp.MustCompile(`(srcd/gitbase:\S+) +(yes|no)`)
	matches := exp.FindStringSubmatch(out.String())

	require.NotNil(matches)
	require.Len(matches, 3)

	imgVersion := matches[1]
	installed := matches[2]

	// If installed, remove it
	if installed == "yes" {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()
		err = docker.RemoveImage(ctx, imgVersion)
		require.NoError(err)
	}

	// Check it's not installed
	out, err = s.RunCommand(context.TODO(), "components", "list")
	require.NoError(err)

	expected := regexp.MustCompile(`srcd/gitbase:\S+ +no +no +srcd-cli-gitbase`)
	require.Regexp(expected, out.String())

	// Install
	_, err = s.RunCommand(context.TODO(), "components", "install", "srcd/gitbase")
	require.NoError(err)

	// Check it's installed
	out, err = s.RunCommand(context.TODO(), "components", "list")
	require.NoError(err)

	expected = regexp.MustCompile(`srcd/gitbase:\S+ +yes +no +srcd-cli-gitbase`)
	require.Regexp(expected, out.String())

	// Call install again, should be an exit 0
	_, err = s.RunCommand(context.TODO(), "components", "install", "srcd/gitbase")
	require.NoError(err)
}

func (s *ComponentsTestSuite) TestInstallAlias() {
	require := s.Require()

	// Install with image name
	_, err := s.RunCommand(context.TODO(), "components", "install", "srcd/cli-daemon")
	require.NoError(err)

	// Install with container name
	_, err = s.RunCommand(context.TODO(), "components", "install", "srcd-cli-daemon")
	require.NoError(err)
}

func (s *ComponentsTestSuite) TestInstallUnknown() {
	require := s.Require()

	// Call install with a srcd image not managed by engine
	out, err := s.RunCommand(context.TODO(), "components", "install", "srcd/lookout")
	require.Error(err)
	require.Contains(out.String(), "srcd/lookout is not valid. Component must be one of")
}

func (s *ComponentsTestSuite) TestInstallVersion() {
	require := s.Require()

	out, err := s.RunCommand(context.TODO(), "components", "list")
	require.NoError(err)

	// Get the exact image:version of gitbase
	exp := regexp.MustCompile(`(srcd/gitbase:\S+)`)
	matches := exp.FindStringSubmatch(out.String())

	require.NotNil(matches)
	require.Len(matches, 2)

	imgVersion := matches[1]

	// Call install with image:version
	out, err = s.RunCommand(context.TODO(), "components", "install", imgVersion)
	require.Error(err)
	require.Contains(out.String(), imgVersion+" is not valid. Component must be one of")
}
