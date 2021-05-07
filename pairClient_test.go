package ptcpayclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	"github.com/btcsuite/btcd/btcec"
)

func TestBTCPayClient_PairClient(t *testing.T) {
	var pkh = "31eb31ecf1a640cd91e0a1105501f36235f8c7d51d67dcf74ccc968d74cb6b25"

	var cryt Cryptography
	cc := cryt.New()

	kp := cc.LoadKeyPair(pkh, btcec.S256())

	var ptc BTCPayClient
	var head Headers
	ptc.SetHeader(head)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	ptc.OverrideProxy(&gp)

	c := ptc.New("http://127.0.0.1:49392", kp)
	ptc.SetLogLevel(lg.AllLevel)

	ps := c.PairClient("123")

	fmt.Println("ps: ", ps)

	//t.Fail()
}
