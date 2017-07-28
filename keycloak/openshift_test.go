package keycloak_test

import (
	"testing"

	"github.com/fabric8-services/fabric8-notification/keycloak"
	"github.com/stretchr/testify/assert"
)

// ignore for now, require vcr recording
func xTestOpenshiftToken(t *testing.T) {

	c := keycloak.Config{
		BaseURL: "https://sso.prod-preview.openshift.io",
		Realm:   "fabric8",
		Broker:  "openshift-v3",
	}

	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ6RC01N29CRklNVVpzQVdxVW5Jc1Z1X3g3MVZJamQxaXJHa0dVT2lUc0w4In0.eyJqdGkiOiI3NzA1YjU0NS1iZjcwLTQ2NTMtOTFmNS01YzZhY2FkMzdkMzAiLCJleHAiOjE0ODk3OTgzNjUsIm5iZiI6MCwiaWF0IjoxNDg5Nzk2NTY1LCJpc3MiOiJodHRwOi8vc3NvLnByb2QtcHJldmlldy5vcGVuc2hpZnQuaW8vYXV0aC9yZWFsbXMvZmFicmljOCIsImF1ZCI6ImZhYnJpYzgtb25saW5lLXBsYXRmb3JtIiwic3ViIjoiZTE1OThjMTgtMTg0Ny00ZDg5LWE3OWMtNmYyNjY4MWNhYjM5IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiZmFicmljOC1vbmxpbmUtcGxhdGZvcm0iLCJhdXRoX3RpbWUiOjE0ODk3OTY1NjUsInNlc3Npb25fc3RhdGUiOiJiNzc4M2QyYy1kZThjLTRmOWMtYmJlNC1iOGY4MDQ2MWE1ODMiLCJhY3IiOiIxIiwiY2xpZW50X3Nlc3Npb24iOiJmMTBkNjdmNy01NzMwLTQwMTktYmMwZS04ZGUzNDI3NmQzYjciLCJhbGxvd2VkLW9yaWdpbnMiOlsiKiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJicm9rZXIiOnsicm9sZXMiOlsicmVhZC10b2tlbiJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsInZpZXctcHJvZmlsZSJdfX0sIm5hbWUiOiJBc2xhayBLbnV0c2VuIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiYXNsYWtANGZzLm5vIiwiZ2l2ZW5fbmFtZSI6IkFzbGFrIiwiZmFtaWx5X25hbWUiOiJLbnV0c2VuIiwiZW1haWwiOiJhc2xha0A0ZnMubm8ifQ.DSrcUCFWOemNjrmtLvdMEWoCtP4etATl-Baoj94nXoCC785fEusOBMl2h_cQtiv_rR3DJgnImms-T3h4fWi3AsKrYxogslijlh1cMP1fgUOwDCV7U_zziNF2sAL9-r6WoVpv9caq2S7VXpFSAYLQC60Adlc8asqaXuWnmpPhkradEkRr_e-XUULNtXhG-klDSOYetgMGO5oqfft37thsc6n3YE7GXGX8tD_ve39jvdlT-XwxicQesG6KjhE9lkzf1AVD3Bhc6__1CXlrTqfEDUpwlXONWtjsBA36YG99BL1RwlYRvVrVNXDBbf03R0Vv8GWjYZpopJckRXL0vh1VCw"

	u, err := keycloak.OpenshiftToken(c, token)
	assert.NoError(t, err)
	assert.NotEqual(t, "", u)
}
