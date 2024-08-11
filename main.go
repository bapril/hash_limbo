package hash_limbo

import (
    "encoding/hex"
    "crypto/sha512"

)

func Hash_hex(bits int,input []byte) string {
	var nonce int = 0;
	var hex_out string;
	for {
		var sha512Hasher = sha512.New()
		sha512Hasher.Write(input)
		sha512Hasher.Write(make([]byte,nonce))
		output := sha512Hasher.Sum(nil)
		//if test
		if limbo(output,bits){
			hex_out = hex.EncodeToString(output)
			break
		}
		nonce++
	}
	return hex_out
}

func limbo(in []byte,bits int) bool{
	if bits >= 8 { // whole bytes
		if in[0] == 0 {
			if bits == 8 {
				return true
			} else {
				return limbo(in[1:],bits - 8)
			}
		} else {
			return false
		}
	} else { //sub byte bitmask
		var mask uint8 = 0
		for j := bits; j > 0; j-- {
			mask |= 1<<j
		}
		if uint8(in[0]) & mask == 0 {
			return true
		} else {
			return false
		}
		
	}
}
