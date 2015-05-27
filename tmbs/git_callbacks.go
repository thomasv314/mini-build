package tmbs

import (
	"fmt"
	git "github.com/libgit2/git2go"
	"net/url"
	"os"
)

func buildCertCallback(url *url.URL, writer TestWriter) git.CertificateCheckCallback {
	return func(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
		if url.Host != hostname {
			writer.Write("Hostname mismatch: " + url.Host + " does not match " + hostname)
			os.Exit(22)
		}
		return git.ErrorCode(0)
	}
}

func buildCredCallback(u *url.URL, writer TestWriter) git.CredentialsCallback {

	return func(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {

		i, cred := git.NewCredDefault()

		if allowedTypes&git.CredTypeSshKey != 0 {
			writer.Write("Cloning repository with SSH Authentication")
			i, cred = git.NewCredSshKey("git", GetTmbsDirectory()+"/id_rsa.pub", GetTmbsDirectory()+"/id_rsa", "")
		}

		if allowedTypes&git.CredTypeUserpassPlaintext != 0 {
			return git.ErrorCode(1), &cred
		} else {
			return git.ErrorCode(i), &cred
		}
	}
}

// TODO find out what this does?
var i = &info{}

func calcPercent(partial, total uint) uint {
	percent := (float64(partial) / float64(total)) * 100
	return uint(percent)
}

func buildSidebandProgressCallback(writer TestWriter) git.TransportMessageCallback {

	return func(str string) git.ErrorCode {
		writer.Write(" remote: " + str)
		return git.ErrorCode(0)
	}

}

func transferProgressCallback(stats git.TransferProgress) git.ErrorCode {
	i.stats = stats
	i.update()
	return git.ErrorCode(0)
}
