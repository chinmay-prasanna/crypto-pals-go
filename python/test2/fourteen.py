from Crypto.Cipher import AES
import base64
from Crypto.Random import get_random_bytes
from Crypto.Util.Padding import pad, unpad
import random

key = get_random_bytes(16)

def oracle(bstr = b''):
    r_prefix = get_random_bytes(random.randint(2, 30))
    text = """Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK"""
    b_text = base64.b64decode(text)
    e_text = r_prefix + bstr + b_text
    block_size = 16
    e_text = pad(e_text, block_size)
    o_text = b''
    cipher = AES.new(key, AES.MODE_ECB)
    for i in range(0, len(e_text), block_size):
        end = i + block_size
        encoded = cipher.encrypt(e_text[i:end])
        o_text += encoded

    return o_text

# Cant find key byte by byte method because random bytes are added, we 
# dont know block size
# len of unknown bytes is fixed. len of attack bytes are fixed.
ol = len(oracle())

while True:
    size = 0
    for i in range(1, 64):
        i_str = b'A'*i
        enc = oracle(i_str)
        if len(enc) > ol:
            size=len(enc)-ol
            ol = len(enc)

