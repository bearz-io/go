package keepass

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/bearz-io/go/errors"
	"github.com/bearz-io/go/fs"
	"github.com/bearz-io/go/vaults"
	"github.com/tobischo/gokeepasslib/v3"
)

const TEST = "TEST"

func init() {
}

type KdbxOptions struct {
	Path           string
	Secret         *string
	SecretFileData []byte
	Create         bool
}

type Kdbx struct {
	db        *gokeepasslib.Database
	features  *vaults.Features
	options   KdbxOptions
	open      bool
	delimiter string
}

func New(options KdbxOptions) *Kdbx {

	features := vaults.NewFeatures(map[string]bool{
		"password": true,
	})
	return &Kdbx{

		features:  features,
		options:   options,
		open:      false,
		delimiter: "/",
	}
}

func Open(options KdbxOptions) (*Kdbx, error) {
	kdbx := New(options)
	return kdbx, kdbx.Open()
}

func (kdbx *Kdbx) Open() error {
	if kdbx == nil {
		return errors.NewArgumentError("kdbx", "kdbx is nil")
	}

	if kdbx.open {
		return nil
	}

	var creds *gokeepasslib.DBCredentials
	db := gokeepasslib.NewDatabase()
	if kdbx.options.Secret != nil && kdbx.options.SecretFileData != nil {
		c, err := gokeepasslib.NewPasswordAndKeyDataCredentials(*kdbx.options.Secret, kdbx.options.SecretFileData)
		if err != nil {
			return err
		}

		creds = c
	} else if kdbx.options.Secret != nil {
		c := gokeepasslib.NewPasswordCredentials(*kdbx.options.Secret)

		creds = c
	} else if kdbx.options.SecretFileData != nil {
		c, err := gokeepasslib.NewKeyDataCredentials(kdbx.options.SecretFileData)
		if err != nil {
			return err
		}

		creds = c
	} else {
		return errors.NewArgumentError("kdbx", "no secret provided")
	}

	exists := fs.Exists(kdbx.options.Path)
	if !exists {

		if !kdbx.options.Create {
			return errors.NewArgumentError("kdbx", "file does not exist and create option is false")
		}

		db = gokeepasslib.NewDatabase(
			gokeepasslib.WithDatabaseKDBXVersion4(),
		)

		rg := rootGroupFromPath(kdbx.options.Path)
		db.Content.Meta.DatabaseName = rg
		db.Content.Root.Groups[0].Name = rg
		db.Credentials = creds

		dir := filepath.Dir(kdbx.options.Path)
		err := fs.EnsureDirDefault(dir)
		if err != nil {
			return err
		}

		slog.Debug("creating new keepass database", "path", kdbx.options.Path)
		file, err := os.Create(kdbx.options.Path)
		if err != nil {
			return err
		}

		err = db.LockProtectedEntries()
		if err != nil {
			return err
		}
		enc := gokeepasslib.NewEncoder(file)
		if err := enc.Encode(db); err != nil {
			file.Close()
			return err
		}

		err = file.Close()
		if err != nil {
			return err
		}

		err = db.UnlockProtectedEntries()
		if err != nil {
			return err
		}

		kdbx.open = true
		kdbx.db = db
		return nil
	}

	file, err := fs.OpenReadDefault(kdbx.options.Path)
	if err != nil {
		return err
	}

	defer file.Close()

	dec := gokeepasslib.NewDecoder(file)
	if err := dec.Decode(db); err != nil {
		return err
	}

	db.Credentials = creds
	err = db.UnlockProtectedEntries()
	if err != nil {
		return err
	}

	kdbx.open = true
	kdbx.db = db
	return nil
}

func rootGroupFromPath(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(path)
	if ext != "" {
		base = base[:len(base)-len(ext)]
	}
	return base
}
