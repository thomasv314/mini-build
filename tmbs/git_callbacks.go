package tmbs

import (
	"fmt"
	git "github.com/libgit2/git2go"
	"net/url"
)

/*
func credentialsCallback(url, usernameFromURL string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
	i, cred := git.NewCredDefault()

	fmt.Println("WTF", i, cred)
	if allowedTypes&git.CredTypeUserpassPlaintext != 0 {
		i, cred = git.NewCredUserpassPlaintext(getUserPass())
		return git.ErrorCode(i), &cred
	}
	//	if allowedTypes&git.CredTypeSshKey != 0 {
	//		i, cred = git.NewCredSshKeyFromAgent(u.User.Username())
	//		fmt.Println("Can we use sshkey")
	//		fmt.Println(cred.Type())
	//		return git.ErrorCode(i), &cred
	//	}
	if allowedTypes&git.CredTypeSshCustom != 0 {
		fmt.Println(cred.Type())

		exitWithMessage("custom ssh not implemented")
	}
	if allowedTypes&git.CredTypeDefault == 0 {
		fmt.Println(cred.Type())
		exitWithMessage("invalid cred type")
	}

	fmt.Println("ERROR CODE", i, "CRED", cred)
	return git.ErrorCode(i), &cred
}
*/
func credentialsCallback(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
	ret, cred := git.NewCredSshKeyFromAgent(username)
	return git.ErrorCode(ret), &cred
}

// TODO find out what this does?
var i = &info{}

func calcPercent(partial, total uint) uint {
	percent := (float64(partial) / float64(total)) * 100
	return uint(percent)
}

func sidebandProgressCallback(str string) git.ErrorCode {
	fmt.Printf("\rremote: %v", str)
	return git.ErrorCode(0)
}

func transferProgressCallback(stats git.TransferProgress) git.ErrorCode {
	i.stats = stats
	i.update()
	fmt.Println("stat", stats)
	fmt.Println(i)
	return git.ErrorCode(0)
}

var uname string
var upass string

func getUserPass() (un string, pw string) {

	if uname == "" {
		uname = getInput("username")
	}

	if upass == "" {
		upass = getMaskedInput("password")
	}

	return uname, upass
}

func buildCertificateCheckCallback(u *url.URL) git.CertificateCheckCallback {
	return func(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
		if u.Scheme != "ssh" && valid == false {
			exitWithMessage("host key check failed for:", hostname)
		}
		return git.ErrorCode(0)
	}
}
