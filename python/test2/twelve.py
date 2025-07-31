from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad
from Crypto.Random import get_random_bytes
import base64

key = get_random_bytes(16)
iv = get_random_bytes(16)

def oracle(bstr = b''):
    text = """Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK"""
    b_text = base64.b64decode(text)
    e_text = bstr + b_text
    block_size = 16
    e_text = pad(e_text, block_size)
    o_text = b''
    cipher = AES.new(key, AES.MODE_ECB)
    for i in range(0, len(e_text), block_size):
        end = i + block_size
        encoded = cipher.encrypt(e_text[i:end])
        o_text += encoded

    return o_text

size = 0
last_l = len(oracle(b''))
for i in range(1, 64):
    input = b'A' * i
    new_l = len(oracle(input))
    if new_l > last_l:
        size = new_l - last_l
        break

block_size = size

recovered = []
for n in range(1, block_size):
    cipher_map = {}
    input = b'A'*(block_size-n)
    for j in range(256):
        m_input = bytes(input) + bytes(recovered) + bytes([j])
        m_output = oracle(m_input)
        cipher_map.update({
            m_output[:block_size]:j
        })
    wo_input = oracle(input)
    if wo_input[:block_size] in cipher_map:
        recovered.append(cipher_map[wo_input[:block_size]])
