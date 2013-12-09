package goxep

import (
	"encoding/xml"
)

// See also http://play.golang.org/p/eiX7aFD14S

const (
	nsStream    = "http://etherx.jabber.org/streams"
	nsTLS       = "urn:ietf:params:xml:ns:xmpp-tls"
	nsSASL      = "urn:ietf:params:xml:ns:xmpp-sasl"
	nsBind      = "urn:ietf:params:xml:ns:xmpp-bind"
	nsClient    = "jabber:client"
	streamStart = xml.Header + `<stream:stream
	from='%s'
	to='%s'
	version='1.0'
	xml:lang='en'
	xmlns='` + nsClient + `'
	xmlns:stream='` + nsStream + `'>`
	streamEnd = "</stream:stream>"
)

type xmppText struct {
	Lang string `xml:"xml:lang,attr,omitempty"`
	Body string `xml:",chardata"`
}

type xmppThread struct {
	Parent string `xml:"parent,attr,omitempty"`
	Thread string `xml:",chardata"`
}

type xmppError struct {
	Info xml.Name  `xml:",any"` // possible element names: see stanzaErrorGroup
	Text *xmppText `xml:"text"`
	By   string    `xml:"by,attr,omitempty"`
	Type string    `xml:"type,attr"` // auth, cancel, continue, modify, wait
}

type xmppStanza struct {
	From string `xml:"from,attr,omitempty"`
	Id   string `xml:"id,attr,omitempty"`
	To   string `xml:"to,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`
	Lang string `xml:"xml:lang,attr,omitempty"`
}

// RFC 6120  A.1  Stream namespace

type streamFeatures struct {
	XMLName    xml.Name        `xml:"http://etherx.jabber.org/streams features"`
	StartTLS   *tlsStartTLS    `xml:""`
	Mechanisms *saslMechanisms `xml:""`
	Bind       *bindBind       `xml:""`
	Session    bool            `xml:"session,omitempty"`
	// TODO see http://xmpp.org/registrar/stream-features.html
}

type streamError struct {
	XMLName xml.Name  `xml:"http://etherx.jabber.org/streams error"`
	Info    *xml.Name `xml:",any"` // possible element names: see streamErrorGroup
	Text    *xmppText `xml:"text"`
}

// RFC 6120  A.3  STARTTLS namespace

type tlsStartTLS struct {
	XMLName  xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-tls starttls"`
	Required bool     `xml:"required,omitempty"`
}

type tlsProceed struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-tls proceed"`
}

type tlsFailure struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-tls failure"`
}

// RFC 6120  A.4  SASL namespace

type saslMechanisms struct {
	XMLName   xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl mechanisms"`
	Mechanism []string `xml:"mechanism"`
}

type saslAbort struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl abort"`
}

type saslAuth struct {
	XMLName   xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl auth"`
	Mechanism string   `xml:"mechanism,attr"`
	Data      string   `xml:",chardata"`
}

type saslChallenge struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl challenge"`
	Data    string   `xml:",chardata"`
}

type saslResponse struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl response"`
	Data    string   `xml:",chardata"`
}

type saslSuccess struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-sasl success"`
	Data    string   `xml:",chardata"`
}

type saslFailure struct {
	XMLName xml.Name  `xml:"urn:ietf:params:xml:ns:xmpp-sasl failure"`
	Info    *xml.Name `xml:",any"` // aborted, account-disabled, credentials-expired, encryption-required, incorrect-encoding, invalid-authzid, invalid-mechanism, malformed-request, mechanism-too-weak, not-authorized, temporary-auth-failure
	Text    *xmppText `xml:"text"`
}

// RFC 6120  A.5  Client namespace

type clientMessage struct {
	XMLName    xml.Name     `xml:"jabber:client message"`
	Subject    []xmppText   `xml:"subject"`
	Body       []xmppText   `xml:"body"`
	Thread     []xmppThread `xml:"thread"`
	Error      *xmppError   `xml:"error"`
	xmppStanza              // Type: chat, error, groupchat, headline, normal
}

type clientPresence struct {
	XMLName    xml.Name   `xml:"jabber:client presence"`
	Show       []string   `xml:"show"` // away, chat, dnd, xa
	Status     []xmppText `xml:"status"`
	Priority   []byte     `xml:"priority"`
	Error      *xmppError `xml:"error"`
	xmppStanza            // Type: error, probe, subscribe, subscribed, unavailable, unsubscribe, unsubscribed
}

type clientIQ struct {
	XMLName    xml.Name   `xml:"jabber:client iq"`
	IQ         string     `xml:",innerxml"`
	Error      *xmppError `xml:"error"`
	xmppStanza            // Type: error, get, result, set
}

// RFC 6120  A.6  Server namespace

type serverMessage struct {
	XMLName    xml.Name     `xml:"jabber:server message"`
	Subject    []xmppText   `xml:"subject"`
	Body       []xmppText   `xml:"body"`
	Thread     []xmppThread `xml:"thread"`
	Error      *xmppError   `xml:"error"`
	xmppStanza              // Type: chat, error, groupchat, headline, normal
}

type serverPresence struct {
	XMLName    xml.Name   `xml:"jabber:server presence"`
	Show       []string   `xml:"show"` // away, chat, dnd, xa
	Status     []xmppText `xml:"status"`
	Priority   []byte     `xml:"priority"`
	Error      *xmppError `xml:"error"`
	xmppStanza            // Type: error, probe, subscribe, subscribed, unavailable, unsubscribe, unsubscribed
}

type serverIQ struct {
	XMLName    xml.Name   `xml:"jabber:server iq"`
	IQ         string     `xml:",innerxml"`
	Error      *xmppError `xml:"error"`
	xmppStanza            // Type: error, get, result, set
}

// RFC 6120  A.7  Resource binding namespace

type bindBind struct {
	XMLName  xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-bind bind"`
	Resource string   `xml:"resource,omitempty"`
	Jid      string   `xml:"jid,omitempty"`
}
