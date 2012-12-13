package goxep

import (
	"encoding/xml"
)

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

type xmlText struct {
	Lang string `xml:"xml:lang,attr,omitempty"`
	Body string `xml:",chardata"`
}

// TODO add ",omitempty" to optional elements/attributes
// TODO add lowercase tag names on all elements

// RFC 6120  A.1  Stream namespace

type streamFeatures struct {
	XMLName    xml.Name        `xml:"http://etherx.jabber.org/streams features"`
	StartTLS   *tlsStartTLS    `xml:""` //TODO
	Mechanisms *saslMechanisms `xml:""`
	Bind       *bindBind       `xml:""`
	//?? Session    bool
	// TODO Compression
}

type streamError struct {
	XMLName xml.Name  `xml:"http://etherx.jabber.org/streams error"`
	Info    *xml.Name `xml:",any"`
	Text    *xmlText  `xml:"text"`
}

// RFC 6120  A.3  STARTTLS namespace

type tlsStartTLS struct {
	XMLName  xml.Name `xml:":ietf:params:xml:ns:xmpp-tls starttls"`
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
	Info    *xml.Name `xml:",any"`
	Text    *xmlText  `xml:"text"`
}

// RFC 6120  A.7  Resource binding namespace

type bindBind struct {
	XMLName  xml.Name `xml:"urn:ietf:params:xml:ns:xmpp-bind bind"`
	Resource string   `xml:"resource,omitempty"`
	Jid      string   `xml:"jid,omitempty"`
}

// RFC 6120  A.5  Client namespace

type clientMessage struct {
	XMLName xml.Name       `xml:"jabber:client message"`
	Subject []xmlText      `xml:"subject"`
	Body    []xmlText      `xml:"body"`
	Thread  []clientThread `xml:"thread"`
	Error   *clientError   `xml:""`
	From    string         `xml:"from,attr,omitempty"`
	Id      string         `xml:"id,attr,omitempty"`
	To      string         `xml:"to,attr,omitempty"`
	Type    string         `xml:"type,attr,omitempty"` // chat, error, groupchat, headline, or normal
	Lang    string         `xml:"xml:lang,attr,omitempty"`
}

type clientThread struct {
	Parent string `xml:"parent,attr,omitempty"`
	Thread string `xml:",chardata"`
}

type clientPresence struct {
	XMLName  xml.Name     `xml:"jabber:client presence"`
	Show     []string     `xml:"show"` // away, chat, dnd, xa
	Status   []xmlText    `xml:"status"`
	Priority []byte       `xml:"priority"`
	Error    *clientError `xml:""`
	From     string       `xml:"from,attr,omitempty"`
	Id       string       `xml:"id,attr,omitempty"`
	To       string       `xml:"to,attr,omitempty"`
	Type     string       `xml:"type,attr,omitempty"` // error, probe, subscribe, subscribed, unavailable, unsubscribe, unsubscribed
	Lang     string       `xml:"xml:lang,attr,omitempty"`
}

type clientIQ struct {
	XMLName xml.Name `xml:"jabber:client iq"`
	//Any     xml.Name    `xml:",any"`
	Error *clientError `xml:""`
	From  string       `xml:"from,attr,omitempty"`
	Id    string       `xml:"id,attr"`
	To    string       `xml:"to,attr,omitempty"`
	Type  string       `xml:"type,attr"` // error, get, result, set
	Lang  string       `xml:"xml:lang,attr,omitempty"`
}

type clientError struct {
	XMLName xml.Name `xml:"jabber:client error"`
	Info    xml.Name `xml:",any"`
	Text    *xmlText `xml:"text"`
	By      string   `xml:"by,attr,omitempty"`
	Type    string   `xml:"type,attr"` // auth, cancel, continue, modify, wait
}
