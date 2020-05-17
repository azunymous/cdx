package watch

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/azunymous/cdx/watch/cryptopasta"
	"github.com/azunymous/cdx/watch/diff"
	"golang.org/x/crypto/hkdf"
	"google.golang.org/grpc"
	"io"
	"os/exec"
)

func NewShareClient(target string, insecure bool) (diff.DiffClient, func(), error) {
	opts := createDialOptions(insecure)
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, nil, err
	}

	return diff.NewDiffClient(conn),
		func() {
			defer conn.Close()
		}, nil
}

func Upload(ctx context.Context, client diff.DiffClient, name, password string) error {
	commits, err := getDiffCommits(ctx)
	if err != nil {
		return err
	}

	var salt = new([32]byte)
	encrypted := false
	if password != "" {
		encrypted = true
		commits, salt, err = encrypt(commits, password)
		if err != nil {
			return err
		}
	}

	reply := &diff.DiffCommits{
		Name:      name,
		Commits:   commits,
		Salt:      salt[:],
		Encrypted: encrypted,
	}
	_, err = client.UploadDiff(ctx, reply)
	return err
}

func encrypt(commits []byte, password string) ([]byte, *[32]byte, error) {
	// derive an encryption key from the master key and the nonce
	var key [32]byte
	var salt = cryptopasta.NewEncryptionKey()
	kdf := hkdf.New(sha256.New, []byte(password), salt[:], nil)
	if _, err := io.ReadFull(kdf, key[:]); err != nil {
		return nil, nil, fmt.Errorf("failed to derive encryption key: %v", err)
	}

	ciphertext, err := cryptopasta.Encrypt(commits, &key)
	if err != nil {
		return nil, nil, err
	}
	return ciphertext, salt, nil
}

func decrypt(secret []byte, password string, salt []byte) ([]byte, error) {
	// derive an encryption key from the master key and the nonce
	var key [32]byte
	kdf := hkdf.New(sha256.New, []byte(password), salt[:], nil)
	if _, err := io.ReadFull(kdf, key[:]); err != nil {
		return nil, fmt.Errorf("failed to derive encryption key: %v", err)
	}

	plaintext, err := cryptopasta.Decrypt(secret, &key)
	if err != nil {
		return nil, err
	}
	return plaintext, err
}

func getDiffCommits(ctx context.Context) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "git", "format-patch", "origin/master", "--stdout")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		return nil, errors.New(stdErr.String())
	}
	return stdOut.Bytes(), nil
}
