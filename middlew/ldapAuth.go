package middlew

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap"
)

var ldapURL string = "ldap://localhost:389"
var baseDN = "ou=sounds,dc=sounds,dc=unal,dc=edu,dc=co"

func auth(conn *ldap.Conn, email string, password string) error {
	result, err := conn.Search(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(cn="+email+")",
		[]string{"dn"},
		nil,
	))

	if err != nil {
		return fmt.Errorf("Failed to find user. %s", err)
	}

	if len(result.Entries) < 1 {
		return fmt.Errorf("User does not exist")
	}

	if len(result.Entries) > 1 {
		return fmt.Errorf("Too many entries returned")
	}

	if err := conn.Bind(result.Entries[0].DN, password); err != nil {
		fmt.Printf("Failed to auth. %s", err)
	} else {
		fmt.Printf("Authenticated successfuly!")
	}

	return err
}

func AuthLDAP(Email string, password string) (*ldap.Conn, error) {
	l, err := ConnectLDAP()
	err = auth(l, Email, password)
	defer l.Close()
	return l, err
}

func ConnectLDAP() (*ldap.Conn, error) {
	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
	}
	err = l.Bind("cn=admin,dc=sounds,dc=unal, dc=edu, dc=co", "admin")
	if err != nil {
		log.Fatal(err)
	}
	return l, err
}

func AddRegisterUser(conn *ldap.Conn, email string, password string, nombre string, apellido string) {
	l, _ := ConnectLDAP()
	addReq := ldap.NewAddRequest("cn="+email+",ou=sounds,dc=sounds,dc=unal, dc=edu, dc=co", []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "user", "person"})
	addReq.Attribute("name", []string{nombre, apellido})
	addReq.Attribute("sAMAccountName", []string{email})
	addReq.Attribute("userPrincipalName", []string{email})
	addReq.Attribute("accountExpires", []string{fmt.Sprintf("%d", 0x00)})

	if err := l.Add(addReq); err != nil {
		log.Fatal("error adding service:", addReq, err)
	}
}
