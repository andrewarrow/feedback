package email

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	dkim "github.com/emersion/go-dkim"
	"github.com/emersion/go-message"
	"io"
	"io/ioutil"
	"net"
	"net/smtp"
	"strings"
)

type smtpClient interface {
	Extension(string) (bool, string)
	StartTLS(*tls.Config) error
	Auth(smtp.Auth) error
	Hello(localName string) error
	Mail(string) error
	Rcpt(string) error
	Data() (io.WriteCloser, error)
	Quit() error
	Close() error
}

func MakeEmailHTML(s string) string {
	var b bytes.Buffer
	var h message.Header
	h.SetContentType("text/html", nil)
	w, err := message.CreateWriter(&b, h)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	io.WriteString(w, s)

	w.Close()
	return b.String()
}

func HelloSend(addr, from string, to []string, msg string) bool {
	var c smtpClient

	c, err := smtp.Dial(addr)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer c.Close()
	c.Hello("mail.many.pw")

	if err = c.Mail(from); err != nil {
		fmt.Println(err)
		return false
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			fmt.Println(err)
			return false
		}
	}

	w, err := c.Data()
	if err != nil {
		fmt.Println(err)
		return false
	}

	private, _ := ioutil.ReadFile("private.key")
	block, _ := pem.Decode(private)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return false
	}

	r := strings.NewReader(msg)

	options := &dkim.SignOptions{
		Domain:   "socialdistance.app",
		Selector: "manypw",
		Signer:   privateKey,
	}

	var b bytes.Buffer
	err = dkim.Sign(&b, r, options)
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = w.Write(b.Bytes())
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}

	c.Quit()
	return true
}

func Send(to, from, subj, content string) {
	tokens := strings.Split(to, "@")
	domain := tokens[1]
	fmt.Println("|", domain, "|")
	mxrecords, err := net.LookupMX(domain)
	fmt.Println(err)
	for _, mx := range mxrecords {
		fmt.Println(mx.Host, mx.Pref)
	}
	if len(mxrecords) == 0 {
		return
	}
	server := mxrecords[0].Host
	recipients := []string{to}
	headers := []string{"From: " + from,
		"To: " + to,
		"Subject: " + subj}
	body := []string{content}

	ports := []int{25, 465, 587, 2525}

	for _, port := range ports {

		if HelloSend(fmt.Sprintf("%s:%d", server, port), from, recipients,
			strings.Join(append(headers, body...), "\r\n")) {
			break
		}
	}
}
