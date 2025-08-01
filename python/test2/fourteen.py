from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad
from Crypto.Random import get_random_bytes
import random
from base64 import b64decode

KEY = get_random_bytes(16)
PREFIX = get_random_bytes(random.randint(1, 3*AES.block_size))

UNKNOWN_STRING = b"""
Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK"""

def encryption_oracle(string=None):
    plaintext = PREFIX + string + b64decode(UNKNOWN_STRING)
    block = AES.new(KEY, AES.MODE_ECB)
    return block.encrypt(pad(plaintext, AES.block_size))

def get_block_size():
    feed = b'A'
    length = 0
    while True:
        cipher = encryption_oracle(feed)
        feed += feed
        if length != 0 and len(cipher) - length > 1:
            return len(cipher) - length
        length = len(cipher)

def get_mode(cipher):
    size = 16
    chunks = []
    for i in range(0, len(cipher), size):
        end = i+size
        chunks.append(cipher[i:end])
    return len(chunks) > len(set(chunks))

def get_prefix_length():
    in1 = b'a'
    in2 = b'b'

    test1 = encryption_oracle(in1)
    test2 = encryption_oracle(in2)
    min_len = min(len(test1), len(test2))

    blocks = 0
    block_size = get_block_size()
    for i in range(0, min_len, block_size):
        if test1[i:i+block_size] != test2[i:i+block_size]:
            break
        blocks += 1

    # Perfectly aligned prefix blocks = blocks
    # length of bytes to start next test from is blocks * block_size
    length = blocks * block_size
    # start test from this length, provide incremental input. 
    # When two blocks become equal, input has aligned with start of new block. 
    # So the len of (block size - len of input) gives total padding required to align input with new block
    # extra_bytes_need = block_size - len(input)
    # total_prefix_length = length + extra_bytes_needed
    input = b''
    for _ in range(block_size):
        input += b'?'
        c1 = encryption_oracle(input)[length:block_size]
        c2 = encryption_oracle(input+b'?')[length:block_size]
        if c1 == c2:
            break
    
    return (length + (block_size-len(input))), blocks

def decrypt_ecb(block_size):
    # cleanly aligned prefix blocks (start check from) == (get_prefix_length // block_size) * block_size
    # extra prefix blocks == get_prefix_length % block_size
    # while
    # str_to_prepend = b'A'*(block_size-1-(len(recovered_bytes)+extra_prefix_blocks)%block_size)
    # actual_output is encrypted(str_to_prepend) -> aligns unkown text at start of new block
    # for i till 256
    # str = str_to_prepend + recovered_bytes + i
    # encrypt str, retrieve str[check_begin:length]. if this str == actuak_out[check_begin:length], recovered += i
    plaintext = b''
    prefix_len, n_blocks = get_prefix_length()
    begin_byte = (prefix_len//block_size) * block_size
    size = block_size
    extra_blocks = prefix_len % block_size
    while True:
        prepend_string = b'A'*(block_size-1-(len(plaintext)+extra_blocks)%block_size)
        actual_output = encryption_oracle(prepend_string)[begin_byte:begin_byte+size]
        found = False
        for i in range(256):
            val = bytes([i])
            input_string = prepend_string + plaintext + val
            controlled_output = encryption_oracle(input_string)[begin_byte:begin_byte+size]
            if actual_output == controlled_output:
                plaintext += val
                found = True
                break
        if not found:
            print(plaintext.decode())
            return
        if (len(plaintext)+extra_blocks)%block_size==0:
            size += block_size

print(decrypt_ecb(get_block_size()))
