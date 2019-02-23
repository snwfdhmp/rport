package rportserver

import (
	"encoding/json"

	"fmt"

	"github.com/spf13/afero"

	"github.com/sirupsen/logrus"
)

type report struct {
	Name    string
	Payload interface{}
}

type storage struct {
	collections map[string][]report
}

var (
	fs = afero.NewOsFs()
)

func (s *storage) Store(r *report) error {
	if _, ok := s.collections[r.Name]; !ok {
		s.collections[r.Name] = make([]report, 0)
		logrus.Infof("store: collection '%s': collection created.", r.Name)
	}

	s.collections[r.Name] = append(s.collections[r.Name], *r)
	logrus.Infof("store: collection '%s': item added.", r.Name)
	return nil
}

func (s *storage) Get(reportName string) ([]report, error) {
	if _, ok := s.collections[reportName]; !ok {
		logrus.Warnf("store: collection '%s' does not exist.", reportName)
		return nil, fmt.Errorf("collection '%s' does not exist", reportName)
	}

	logrus.Warnf("store: collection '%s': displaying.", reportName)
	return s.collections[reportName], nil
}

func (s *storage) SaveTo(filepath string) error {
	b, err := json.Marshal(s.collections)
	if err != nil {
		return err
	}

	if err := afero.WriteFile(fs, filepath, b, 0700); err != nil {
		return err
	}

	logrus.Infof("store: saved to '%s'.", filepath)
	return nil

}

//newStorage inits a new empty storage object
func newStorage() *storage {
	c := make(map[string][]report)

	return &storage{
		collections: c,
	}
}

//newStorageFrom tries to create a new storage object by parsing a yaml file
func newStorageFrom(filepath string) (*storage, error) {
	b, err := afero.ReadFile(fs, filepath)
	if err != nil {
		return nil, err
	}

	var collections map[string][]report

	if err := json.Unmarshal(b, &collections); err != nil {
		return nil, err
	}
	logrus.Infof("store: loaded from '%s'.", filepath)

	if collections == nil {
		logrus.Warnf("doctor: collections were nil. initing to avoid nil map issues.")
		collections = make(map[string][]report)
	}

	return &storage{
		collections: collections,
	}, nil
}
