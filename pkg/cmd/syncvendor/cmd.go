package syncvendor

import (
	"fmt"
	"io/ioutil"

	"io"

	"os"
	"path"

	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-billy.v3/osfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
)

type SyncBranch struct {
	TempDir string

	FromRepo      string
	FromBranch    string
	FromDirectory string
	ToRepo        string
	ToBranch      string

	Out    io.Writer
	ErrOut io.Writer
}

func NewCmdSyncBranch(out io.Writer, errOut io.Writer) *cobra.Command {
	o := &SyncBranch{
		Out:    out,
		ErrOut: errOut,
	}
	var err error
	o.TempDir, err = ioutil.TempDir("", "sync-branch-")
	if err != nil {
		panic(err)
	}

	cmd := &cobra.Command{
		Use:   "sync-vendor --from-repo repo --from-branch branch  --from-directory dir-to-sync --to-repo repo --to-branch branch",
		Short: "filters the original repo by directory and pushes to target",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run()
		},
	}

	cmd.Flags().StringVar(&o.TempDir, "working-dir", o.TempDir, "working directory")
	cmd.Flags().StringVar(&o.FromRepo, "from-repo", o.FromRepo, "source repository")
	cmd.Flags().StringVar(&o.FromBranch, "from-branch", o.FromBranch, "source branch")
	cmd.Flags().StringVar(&o.FromDirectory, "from-dir", o.FromDirectory, "source directory")
	cmd.Flags().StringVar(&o.ToRepo, "to-repo", o.ToRepo, "destination repository")
	cmd.Flags().StringVar(&o.ToBranch, "to-branch", o.ToBranch, "destination branch")

	return cmd
}

func (s *SyncBranch) Run() error {
	fmt.Printf("Starting\n")

	if err := os.MkdirAll(s.FromRepoDir(), 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(s.ToRepoDir(), 0755); err != nil {
		return err
	}

	fromFS := osfs.New(path.Join(s.FromRepoDir(), ".git"))
	fromStorage, err := filesystem.NewStorage(fromFS)
	if err != nil {
		return err
	}
	fromClone := &git.CloneOptions{
		URL:           s.FromRepo,
		ReferenceName: plumbing.ReferenceName("refs/heads/" + s.FromBranch),
		SingleBranch:  true,
		Tags:          git.NoTags,
	}
	fromRepo, err := git.Clone(fromStorage, fromFS, fromClone)
	fmt.Printf("#### 1a %v\n", err)
	if err == git.ErrRepositoryAlreadyExists {
		fmt.Printf("#### 1b\n")
		fromRepo, err = git.Open(fromStorage, fromFS)
		fmt.Printf("#### 1c %v\n", err)
		if err != nil {
			return nil
		}
		err = fromRepo.Fetch(&git.FetchOptions{
			RemoteName: "origin",
			RefSpecs:   []config.RefSpec{config.RefSpec("refs/heads/" + s.FromBranch + ":refs/remotes/origin/" + s.FromBranch)},
			Tags:       git.NoTags,
		})
		if err == git.NoErrAlreadyUpToDate {
			fmt.Printf("#### 1cc %v\n", err)
			err = nil
		}
		fmt.Printf("#### 1d %v\n", err)
	}
	fmt.Printf("#### 1e %v\n", err)
	if err != nil {
		return err
	}

	return nil
}

func (s *SyncBranch) FromRepoDir() string {
	return path.Join(s.TempDir, s.FromRepo)
}

func (s *SyncBranch) ToRepoDir() string {
	return path.Join(s.TempDir, s.ToRepo)
}
