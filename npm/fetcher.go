package npm

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DownloadAndExtract(name, version, dest string) error {
	url := fmt.Sprintf("https://registry.npmjs.org/%s/%s", name, version)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := new(strings.Builder)
	io.Copy(buf, resp.Body)
	data := buf.String()
	start := strings.Index(data, `"tarball":"`) + len(`"tarball":"`)
	end := strings.Index(data[start:], `"`) + start
	tarballURL := strings.ReplaceAll(data[start:end], "\\u0026", "&")

	tarResp, err := http.Get(tarballURL)
	if err != nil {
		return err
	}
	defer tarResp.Body.Close()

	os.MkdirAll(dest, 0755)
	return extractTarGz(tarResp.Body, dest)
}

func extractTarGz(gzipStream io.Reader, target string) error {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		return err
	}
	defer uncompressedStream.Close()

	tarReader := tar.NewReader(uncompressedStream)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(path, 0755)
		case tar.TypeReg:
			os.MkdirAll(filepath.Dir(path), 0755)
			outFile, err := os.Create(path)
			if err != nil {
				return err
			}
			io.Copy(outFile, tarReader)
			outFile.Close()
		}
	}
	return nil
}