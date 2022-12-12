package gitScanner

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"time"
	"reflect"
)

type Commits struct {
	Hash         string
	Author       string
	Date         time.Time
	Message      string
	ParentHashes []string
}

func GitScannerHandler() {
	url := "https://github.com/SkySecCoder/testing"
	path := "/tmp/foo"

	fmt.Println("\n\n[+] Cleaning clone directory " + path)
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Clone the given repository to the given directory
	fmt.Println("\n\n[+] Cloning: " + url + "\n`-- In path: " + path)
	repo, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	fmt.Println("\n\n[+] Get hash pointed by HEAD")
	// ... retrieving the branch being pointed by HEAD
	ref, err := repo.Head()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n\n[+] Get commit object")
	// ... retrieve commit object
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println("\n\n[+] List tree from head")
	// // ... List the tree from HEAD
	// tree, err := commit.Tree()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println("\n\n[+] View history of repo")
	// ... View history of repo
	commitIter, _ := repo.Log(&git.LogOptions{From: commit.Hash})

	allHashes := []Commits{}
	tempHash := Commits{}
	err = commitIter.ForEach(func(c *object.Commit) error {
		hash := c.Hash.String()
		tempHash.Hash = hash
		tempHash.Author = c.Author.Name
		tempHash.Date = c.Author.When
		tempHash.Message = c.Message

		tempParentHashes := []string{}
		for _, h := range c.ParentHashes {
			tempParentHashes = append(tempParentHashes, h.String())
		}
		tempHash.ParentHashes = tempParentHashes
		allHashes = append(allHashes, tempHash)
		return nil
	})

	fmt.Println("\n\n[+] Get diff between commits")
	// ... View history of repo
	// commit, _ = repo.CommitObject(plumbing.NewHash("4ffd9bc8144c2e31d9eb57007a24578d3a6ce17c"))
	// blob, _ := repo.GetBlob(tree.s, plumbing.NewHash("4ffd9bc8144c2e31d9eb57007a24578d3a6ce17c"))
	// fmt.Println(blob)
	

	for _, comm := range allHashes {
		fmt.Println("\n[+] Hash: "+comm.Hash)
		fmt.Println("[+] Author: "+comm.Author)
		fmt.Println("[+] Date: "+comm.Date.String())
		fmt.Println("[+] Message: "+comm.Message)
		fmt.Println(comm.ParentHashes)
		commit1, _ := repo.CommitObject(plumbing.NewHash(comm.Hash))
		for _, parentComm := range comm.ParentHashes {
			commit2, _ := repo.CommitObject(plumbing.NewHash(parentComm))
			data, _ := commit1.Patch(commit2)
			fmt.Println(data.String())
			fmt.Println(reflect.TypeOf(data.String()))
		}
	}
}
