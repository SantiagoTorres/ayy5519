#!/usr/bin/env python

import securesystemslib.keys as k
import securesystemslib.ed25519_keys as ed
import binascii as b

with open('whatup', 'rb')  as fp:
    key = k.import_ed25519key_from_private_json(fp.read())
pubkey  = b.unhexlify(key['keyval']['public'])

for i in range(10):
    print(b.hexlify(ed.create_signature(pubkey, b.unhexlify(key['keyval']['private']), b'ayylmao', 'ed25519')[0]))
