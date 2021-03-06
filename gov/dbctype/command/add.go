package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/frankbraun/codechain/secpkg"
	"github.com/frankbraun/codechain/util/log"
	"github.com/scritcash/scrit/netconf"
)

func add(net *netconf.Network, currency string, amount uint64) error {
	// make sure network has a future epoch
	if err := net.HasFuture(); err != nil {
		return err
	}
	// make sure DBCType has not been defined yet
	dbcTypes := net.DBCTypes()
	dbcType := netconf.DBCType{
		Currency: currency,
		Amount:   amount,
	}
	if dbcTypes[dbcType] {
		return fmt.Errorf("DBC type already defined: %v", dbcType)
	}
	// add DBCType
	net.DBCTypeAdd(dbcType)
	return nil
}

// Add implements the scrit-gov 'dbctype add' command.
func Add(argv0 string, args ...string) error {
	fs := flag.NewFlagSet(argv0, flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", argv0)
		fmt.Fprintf(os.Stderr, "Add new DBC type to future epoch of %s.\n", netconf.DefNetConfFile)
		fs.PrintDefaults()
	}
	currency := fs.String("currency", "", "Currency of DBC type to add")
	amount := fs.Uint64("amount", 0, "Amount of DBC type to add")
	verbose := fs.Bool("v", false, "Be verbose")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *verbose {
		log.Std = log.NewStd(os.Stdout)
	}
	if *currency == "" {
		fmt.Fprintf(os.Stderr, "%s: option -currency is mandatory\n", argv0)
		return flag.ErrHelp
	}
	if *amount == 0 {
		fmt.Fprintf(os.Stderr, "%s: option -amount is mandatory\n", argv0)
		return flag.ErrHelp
	}
	if fs.NArg() != 0 {
		fs.Usage()
		return flag.ErrHelp
	}
	if err := secpkg.UpToDate("scrit"); err != nil {
		return err
	}
	// load
	net, err := netconf.LoadNetwork(netconf.DefNetConfFile)
	if err != nil {
		return err
	}
	// validate
	if err := net.Validate(); err != nil {
		return err
	}
	// edit
	if err := add(net, *currency, *amount); err != nil {
		return err
	}
	// validate again
	if err := net.Validate(); err != nil {
		return err
	}
	// save
	if err := net.Save(netconf.DefNetConfFile); err != nil {
		return err
	}
	return nil
}
