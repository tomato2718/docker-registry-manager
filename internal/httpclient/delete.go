package httpclient

import (
	"errors"
	"net/http"
	"strings"
)

var manifestMediaTypes = []string{
	"application/vnd.docker.distribution.manifest.v2+json",
	"application/vnd.docker.distribution.manifest.list.v2+json",
	"application/vnd.oci.image.manifest.v1+json",
	"application/vnd.oci.image.index.v1+json",
}

func Delete(url string, name string, tag string) error {
	digest, err := resolveDigest(url, name, tag)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, url+"/v2/"+name+"/manifests/"+digest, nil)
	if err != nil {
		return errors.New("Unable to build request.")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New("Unable to reach registry.")
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusAccepted:
		return nil
	case http.StatusNotFound:
		return errors.New("Image not found.")
	case http.StatusMethodNotAllowed:
		return errors.New("Deletion is not enabled on this registry.")
	default:
		return errors.New("Unexpected response: " + res.Status)
	}
}

func resolveDigest(url string, name string, tag string) (string, error) {
	req, err := http.NewRequest(http.MethodHead, url+"/v2/"+name+"/manifests/"+tag, nil)
	if err != nil {
		return "", errors.New("Unable to build request.")
	}
	req.Header.Set("Accept", strings.Join(manifestMediaTypes, ", "))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.New("Unable to reach registry.")
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return "", errors.New("Image not found.")
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New("Unexpected response: " + res.Status)
	}

	digest := res.Header.Get("Docker-Content-Digest")
	if digest == "" {
		return "", errors.New("Registry did not return a manifest digest.")
	}

	return digest, nil
}
