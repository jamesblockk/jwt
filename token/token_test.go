package token

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	in := map[string]string{}
	in["id"] = "1"
	token , err := Build(in)
	if !assert.NoError(t, err) {
		fmt.Println(err)
		return
	}
	if !assert.NotEmpty(t,token){
		fmt.Println(token)
	}
}

func TestVerify(t *testing.T) {
	in := map[string]string{}
	in["id"] = "1"
	token , err := Build(in)
	if !assert.NoError(t, err) {
		fmt.Println(err)
		return
	}
	result , err := Verify(token)
	if !assert.NoError(t, err) {
		fmt.Println(err)
		return
	}
	if !assert.NotEmpty(t,result){
		fmt.Println(result)
	}

	if !assert.Equal(t,"1",result["id"]){
		fmt.Println(result)
	}
}
