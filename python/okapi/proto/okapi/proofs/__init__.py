# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: proofs.proto
# plugin: python-betterproto
from dataclasses import dataclass

import betterproto


class LdSuite(betterproto.Enum):
    JcsEd25519Signature2020 = 0


@dataclass(eq=False, repr=False)
class CreateProofRequest(betterproto.Message):
    # The input JSON document that will be used to create the LD Proof. This
    # document must also contain a "proof" object, with the desired values filled
    # in.
    document: "betterproto_lib_google_protobuf.Struct" = betterproto.message_field(1)
    # The signer of the proof. This field must include the 'kid' in full URI
    # format. Example:  did:example:alice#key-1
    key: "_keys__.JsonWebKey" = betterproto.message_field(3)
    # The LD Suite to use to produce this proof
    suite: "LdSuite" = betterproto.enum_field(4)


@dataclass(eq=False, repr=False)
class CreateProofResponse(betterproto.Message):
    signed_document: "betterproto_lib_google_protobuf.Struct" = (
        betterproto.message_field(1)
    )


@dataclass(eq=False, repr=False)
class VerifyProofRequest(betterproto.Message):
    pass


@dataclass(eq=False, repr=False)
class VerifyProofResponse(betterproto.Message):
    pass


from .. import keys as _keys__
import betterproto.lib.google.protobuf as betterproto_lib_google_protobuf
