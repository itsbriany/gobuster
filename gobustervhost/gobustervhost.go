package gobustervhost

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/OJ/gobuster/libgobuster"
)

// GobusterVhost is the main type to implement the interface
type GobusterVhost struct {
	options    *OptionsVhost
	globalopts *libgobuster.Options
	http       *libgobuster.HTTPClient
}

// NewGobusterVhost creates a new initialized GobusterDir
func NewGobusterVhost(cont context.Context, globalopts *libgobuster.Options, opts *OptionsVhost) (*GobusterVhost, error) {
	return nil, nil
}

// PreRun is the pre run implementation of gobusterdir
func (d *GobusterVhost) PreRun() error {
	return nil
}

// Run is the process implementation of gobusterdir
func (d *GobusterVhost) Run(word string) ([]libgobuster.Result, error) {
	return nil, nil
}

// ResultToString is the to string implementation of gobusterdir
func (d *GobusterVhost) ResultToString(r *libgobuster.Result) (*string, error) {
	return nil, nil
}

// GetConfigString returns the string representation of the current config
func (d *GobusterVhost) GetConfigString() (string, error) {
	var buffer bytes.Buffer
	bw := bufio.NewWriter(&buffer)
	tw := tabwriter.NewWriter(bw, 0, 5, 3, ' ', 0)
	o := d.options
	if _, err := fmt.Fprintf(tw, "[+] Url:\t%s\n", o.URL); err != nil {
		return "", err
	}
	if _, err := fmt.Fprintf(tw, "[+] Threads:\t%d\n", d.globalopts.Threads); err != nil {
		return "", err
	}

	wordlist := "stdin (pipe)"
	if d.globalopts.Wordlist != "-" {
		wordlist = d.globalopts.Wordlist
	}
	if _, err := fmt.Fprintf(tw, "[+] Wordlist:\t%s\n", wordlist); err != nil {
		return "", err
	}

	if o.Proxy != "" {
		if _, err := fmt.Fprintf(tw, "[+] Proxy:\t%s\n", o.Proxy); err != nil {
			return "", err
		}
	}

	if o.Cookies != "" {
		if _, err := fmt.Fprintf(tw, "[+] Cookies:\t%s\n", o.Cookies); err != nil {
			return "", err
		}
	}

	if o.UserAgent != "" {
		if _, err := fmt.Fprintf(tw, "[+] User Agent:\t%s\n", o.UserAgent); err != nil {
			return "", err
		}
	}

	if o.Username != "" {
		if _, err := fmt.Fprintf(tw, "[+] Auth User:\t%s\n", o.Username); err != nil {
			return "", err
		}
	}

	if d.globalopts.Verbose {
		if _, err := fmt.Fprintf(tw, "[+] Verbose:\ttrue\n"); err != nil {
			return "", err
		}
	}

	if _, err := fmt.Fprintf(tw, "[+] Timeout:\t%s\n", o.Timeout.String()); err != nil {
		return "", err
	}

	if err := tw.Flush(); err != nil {
		return "", fmt.Errorf("error on tostring: %v", err)
	}

	if err := bw.Flush(); err != nil {
		return "", fmt.Errorf("error on tostring: %v", err)
	}

	return strings.TrimSpace(buffer.String()), nil
}
