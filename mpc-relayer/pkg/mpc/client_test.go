/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package mpc

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/logger"
)

var (
	pubkeySignature = []struct {
		name              string
		keyType           CryptoSystem
		key, address, sig string
	}{
		{"ec1", EcDSA, "03094781d60ed612edeaa122cfe01418c6c0f2c9fc9856f78daf4bf0ad06e3b0c2", "0x297543c8b698602d0b3bc029820a68dcccf1773b", "75c6da035c1d75b8c88eecde4418cb0cd5a75644ee452323ad97cc14698d1b8d319d90323257bac133190863e8586f77bd6142fc68ea56added20412b25173f2"},
		{"ec2", EcDSA, "034187b4d3b4f42436fecb3e926e18109fd3778afa28814d1b13626de829510950", "0x6054d343bc3d1aa800e9b7adfec35439fdc5bb90", "f4bcd188a16bd52d21322c5686c87fdde58e71523773518a6d0a97303e9b70e015a93c0918ea10c7bbe2633ebb94cf3832bc278b25f2c48743834b8c0f359b49"},
		{"ed1", EdDSA, "a50c3439c745cc1f79739df62b127edb3df18774bfc9973a598ae9d7f0879741", "addr1vxv3rjfr2d8tgwz0gt88r3y4sytk9wz8uhdxa288y2ktkdc38vev4", "3ff8d347f41dcbd76c5908432e278bd6ea711e664e16275e23b730ceb7895410816ba4407860c16c354ff8bea3e2344574211342092fd7c8e8cf72a612766809"},
		{"ed2", EdDSA, "005f114779027ae05344bbeea3633b8fa675f427dbbbdaa48c57da68d91c354d", "addr_test1vretyterdeqk68jgmsclm9y72f509na0nq7yy93x76ek43cgjzhl0", "39baac696af7bfcda507c7f49968ad0d34e0f37f917b8fdf46027f646f4bbbc555d8cca759327fc50b7c6f9e0afdd10b82166ea1c32f388e8d41be9367919e0d"},
	}

	signature = []struct {
		name           string
		keyType        CryptoSystem
		key, EcR, EcS  string
		EdPk, EdR, EdS string
	}{
		{name: "ec1", keyType: EcDSA, key: "037c6c6e7e1035d223847911054e00952bfe24ed0f91e45b3d2413310d6b8cb99b", EcR: "1ee4e814a43ed54d9ffe51558cde1a40baf095d97e98776a6c9f7155849ed713", EcS: "7dc08b61fe58dc74520d3834e904bf9b7cd55846951cc451408d03579cf18260"},
		{name: "ec2", keyType: EcDSA, key: "0325e37b80989a264959ec6fb9d3492b00efb1144d6651fb270c727f7a13c44d7f", EcR: "7686990ef6b21c6bbdd0a2b2782d00909f451d365a61ceafa707f1b4de706ff9", EcS: "1ae3487d02475ee9f44027870cad54fd28648634f322b02de346a1605277606f"},
		{name: "ec3", keyType: EcDSA, key: "03094781d60ed612edeaa122cfe01418c6c0f2c9fc9856f78daf4bf0ad06e3b0c2", EcR: "e4899cc47ab30e84aed02fe91cac71fa4e89110e81158b612c3e21b02ca37ea6", EcS: "56a5dcfd089c9b8e65c386c79433f4f2e66f0f7b4804c5dd6a6c002ab29c9d9c"},
		{name: "ec4", keyType: EcDSA, key: "034187b4d3b4f42436fecb3e926e18109fd3778afa28814d1b13626de829510950", EcR: "f877504f11b1171e422e4d9fba33628a9dbecd65bd12396fc19f698d828d36f1", EcS: "7dc3b1d4a3f64d06cbafde8c8ee4b55209a9614d9a9c4e02bd74ccdd61ab734d"},
		{name: "ec5", keyType: EcDSA, key: "0371f1426f330adbfeba3229e2279ad7569f636a292a6469fc5494f810f0930b7b", EcR: "cf743b33592351058527530eaa98f8392562c6cda86782905760b0cd110339d2", EcS: "1b36c97d2ac976776e4f7b37d1c5ae717506f33576a33f70718556fa76514783"},
		{name: "ec6", keyType: EcDSA, key: "03334d3b63e8779edf3a61c99af6e055fd25d21163dadc8608756d4df42bd4a0fc", EcR: "e24adb7a7624d5f28fb674ee8892f305cc590c18e40e09541853adf5eed70c06", EcS: "7fb5c87e1bed7d10515b1b63264e490cfcc799ed8353eac421d6784a9010b585"},
		{name: "ec7", keyType: EcDSA, key: "03b8d89460ccee97c9d4fcc8c2bfe06348e33f159d880efcd75dea82cd82a5991a", EcR: "f1af4026396fc2fcc01cca85682396bf36e6e5b59492da80b4ef1657ee572136", EcS: "27f2f162292aae7d7323a6ae783e0263ab74cdd34b5ae841eb83ed739180fe84"},
		{name: "ec8", keyType: EcDSA, key: "02fb062ae96f23f9fce28473a3f5723b09a18130f839b05b13fe55d00820ffab72", EcR: "830f6206890625f5fb3ecd76ae4953de08231325b33ed9b2ab6c1d336ab141a1", EcS: "4bffc5b41f1b98cbd9b5f7198421d8328df3cd3689ce1172658e9dd0e7ca7348"},
		{name: "ed1", keyType: EdDSA, EdPk: "c2224fafb6282fae4d49204afefec2de919fea1b1cfea8de188fc74328ea12be", EdR: "fccbd0f8907d069d93395c1e2107f380ab7d6566f74cb383a70c6105ef3df246", EdS: "ed08e2abc50edeeea1087a6c9fa342d36c551e950108e42d53911e14482d3900"},
		{name: "ed2", keyType: EdDSA, EdPk: "8f9eebdf1e0499bb5dfd1076fea50f72d37a19dcb3a4e91c0fc19023ac9ba926", EdR: "cbe83bface18971f62c3e530df92b2729833f3ed36973ecfeef755e417ec3493", EdS: "6db6899521d9cb26a0750a5c4b844f76c65fa9d351f09319256c708bead97103"},
		{name: "ed3", keyType: EdDSA, EdPk: "a50c3439c745cc1f79739df62b127edb3df18774bfc9973a598ae9d7f0879741", EdR: "ddc8382c2aaaf43b7a08f55a4c55af9b88f041ec9beedd51220e087b77b0e520", EdS: "e3d141c5cf2d513e3acac4cdeea6eb939ae3ab0e280bd12c4e2a7e3f814b6e01"},
		{name: "ed4", keyType: EdDSA, EdPk: "005f114779027ae05344bbeea3633b8fa675f427dbbbdaa48c57da68d91c354d", EdR: "d931895802e232bad99b8f5b3caab6741e8334ab9982261f23588056d031932c", EdS: "51b2c22f68cfc963f98fe9d1d9b35821b2c707ed1e0236d2f21a8daba865ec05"},
	}
)

func TestPubkeySignature(t *testing.T) {
	var key string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := &SigRequest{}
		_ = json.NewDecoder(r.Body).Decode(body)

		w.WriteHeader(200)
		if r.URL.String() == "/ecdsa/keys" {
			_ = json.NewEncoder(w).Encode(SigResponse{
				KeyID: body.KeyID,
				Pk:    key,
			})
		}
		if r.URL.String() == "/eddsa/keys" {
			_ = json.NewEncoder(w).Encode(SigResponse{
				KeyID: body.KeyID,
				EdPk:  key,
			})
		}
		if r.URL.String() == "/ecdsa/sign" {
			resp, _ := localMPCSign(body, 0, EcDSA)
			_ = json.NewEncoder(w).Encode(resp)
		}
		if r.URL.String() == "/eddsa/sign" {
			resp, _ := localMPCSign(body, 0, EdDSA)
			_ = json.NewEncoder(w).Encode(resp)
		}
	}))

	log, err := logger.NewLogger("error", "plain", false, "test")
	if err != nil {
		t.Fatal(err)
	}

	serverURL, _ := url.Parse(server.URL)
	classic := NewClient(Config{Node: []Node{{Host: serverURL.Hostname(), Port: serverURL.Port()}}}, log)
	local := NewClient(Config{Mock: true, Salt: 0}, log)

	tt := []struct {
		name   string
		client Client
	}{
		{"mock,MockMPCClient", local},
		{"MPC,live", classic},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, signatureData := range pubkeySignature {

				key = signatureData.key

				src := fmt.Sprintf("%0*v", 64, signatureData.name)
				keyID, err := hex.DecodeString(src)
				if err != nil {
					t.Fatalf("src: %v, error %v", src, err)
				}
				seed := sha256.Sum256(keyID)
				pubKey, _, err := tc.client.PublicKey(seed[:], signatureData.keyType)
				if err != nil {
					t.Fatalf("Client :%s keyType: %v error: %v", tc.name, signatureData, err)
				}

				resp, _, err := tc.client.PubkeySignature(pubKey, seed[:], signatureData.keyType)
				if err != nil {
					t.Fatal(err)
				}

				expectedSig, _ := hex.DecodeString(signatureData.sig)
				if !bytes.Equal(resp, expectedSig) {
					t.Errorf("Client :%s keyType:%v Unexpected signature, got %x,\n want %x", tc.name, signatureData.keyType, resp, expectedSig)
				}
			}

		})
	}
}

func TestSignature(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := &SigRequest{}
		_ = json.NewDecoder(r.Body).Decode(body)
		w.WriteHeader(200)
		if r.URL.String() == "/ecdsa/sign" {
			resp, _ := localMPCSign(body, 0, EcDSA)
			_ = json.NewEncoder(w).Encode(resp)
		}
		if r.URL.String() == "/eddsa/sign" {
			resp, _ := localMPCSign(body, 0, EdDSA)
			_ = json.NewEncoder(w).Encode(resp)
		}
	}))

	log, err := logger.NewLogger("error", "plain", false, "test")
	if err != nil {
		t.Fatal(err)
	}

	serverURL, _ := url.Parse(server.URL)

	classic := NewClient(Config{Node: []Node{{Host: serverURL.Hostname(), Port: serverURL.Port()}}}, log)
	local := NewClient(Config{Mock: true, Salt: 0}, log)

	var clients [2]Client
	clients[0] = classic
	clients[1] = local

	for cIndex, client := range clients {
		for _, c := range signature {
			src := fmt.Sprintf("%0*v", 64, c.name)

			keyID, err := hex.DecodeString(src)
			if err != nil {
				t.Fatalf("cIndex %d keyType %v : %s", cIndex, c.keyType, err)
			}
			message := sha256.Sum256([]byte("toto"))

			id := rand.Int63n(10)

			response, _, err := client.Signature(&SigRequestData{
				KeyID:   keyID,
				ID:      id,
				SigHash: message[:],
			}, c.keyType)
			if err != nil {
				t.Fatalf("cIndex %d keyType %v : %s", cIndex, c.keyType, err)
			}
			if got, want := response.Pk, c.key; got != want {
				t.Errorf("cIndex %d keyType %v got %v want %v", cIndex, c.keyType, got, want)
			}
			if got, want := response.EcR, c.EcR; got != want {
				t.Errorf("cIndex %d keyType %v got %v want %v", cIndex, c.keyType, got, want)
			}
			if got, want := response.EcS, c.EcS; got != want {
				t.Errorf("cIndex %d keyType %v got %v want %v", cIndex, c.keyType, got, want)
			}

			if got, want := response.EdPk, c.EdPk; got != want {
				t.Errorf("cIndex %d keyType %v got %v want %v", cIndex, c.keyType, got, want)
			}
			if got, want := response.EdR, c.EdR; got != want {
				t.Errorf("cIndex %d keyType %v got %v want %v", cIndex, c.keyType, got, want)
			}
			if got, want := response.EdS, c.EdS; got != want {
				t.Errorf("cIndex %d keyType %v got %v want %v", cIndex, c.keyType, got, want)
			}
		}
	}
}
