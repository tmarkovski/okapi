require_relative 'test_helper'
require_relative '../lib/proto/okapi/keys/v1/keys_pb'
require_relative '../lib/proto/okapi/security/v1/security_pb'
require 'google/protobuf/well_known_types'
require_relative '../lib/okapi'
require_relative '../lib/version'
require 'base64'
require 'base58'

class OberonTest < Minitest::Test
  def before_setup
    Okapi::set_library_path('./libs')
    Okapi::load_native_library
  end

  def test_oberon_demo
    key = Okapi::Oberon::create_key(Security_V1::CreateOberonKeyRequest.new)
    data = "alice"
    nonce = "1234"

    token = Okapi::Oberon.create_token(Security_V1::CreateOberonTokenRequest.new(:data=>data, :sk=>key.sk))
    proof = Okapi::Oberon.create_proof(Security_V1::CreateOberonProofRequest.new(:data=>data, :nonce=>nonce, :token=>token.token))
    result = Okapi::Oberon.verify_proof(Security_V1::VerifyOberonProofRequest.new(:data=>data, :nonce=>nonce, :pk=>key.pk, :proof=>proof.proof))

    assert result.valid == true
  end

  def test_demo_with_blinding
    key = Okapi::Oberon::create_key(Security_V1::CreateOberonKeyRequest.new)
    data = "alice"
    nonce = "1234"
    # Blinding code to be used by issuer and given to holder to transfer the token securely
    issuer_2fa = "issuer code"

    token_request = Security_V1::CreateOberonTokenRequest.new(:data=>data, :sk=>key.sk)
    token_request.blinding += [issuer_2fa]
    blinded_token = Okapi::Oberon.create_token(token_request)

    # Holder unblinds the token
    unblind_request = Security_V1::UnBlindOberonTokenRequest.new(:token=>blinded_token.token)
    unblind_request.blinding += [issuer_2fa]
    token = Okapi::Oberon.unblind_token(unblind_request)

    # Holder prepares a proof with blinding
    proof = Okapi::Oberon.create_proof(Security_V1::CreateOberonProofRequest.new(:data=>data, :nonce=>nonce, :token=>token.token))
    # Verifier verifies the proof
    result = Okapi::Oberon.verify_proof(Security_V1::VerifyOberonProofRequest.new(:data=>data, :nonce=>nonce, :pk=>key.pk, :proof=>proof.proof))
    assert result.valid == true

    # holder blinds the token with a personal pin
    user_pin = "00042"
    blind_request = Security_V1::BlindOberonTokenRequest.new(:token=>token.token)
    blind_request.blinding += [user_pin]
    user_blinded_token = Okapi::Oberon.blind_token(blind_request)

    # Holder prepares a proof using the pin binding
    proof_request = Security_V1::CreateOberonProofRequest.new(:data=>data, :nonce=>nonce, :token=>user_blinded_token.token)
    proof_request.blinding += [user_pin]
    proof = Okapi::Oberon.create_proof(proof_request)

    # Verifier verifies the proof
    result = Okapi::Oberon.verify_proof(Security_V1::VerifyOberonProofRequest.new(:data=>data, :nonce=>nonce, :pk=>key.pk, :proof=>proof.proof))
    assert result.valid == true

    # Bad actor creates a proof with incorrect blinding pin
    proof_request = Security_V1::CreateOberonProofRequest.new(:data=>data, :nonce=>nonce, :token=>user_blinded_token.token)
    proof_request.blinding += ["invalid pin"]
    proof = Okapi::Oberon.create_proof(proof_request)

    # verifier tries to verify proof, fails
    result = Okapi::Oberon.verify_proof(Security_V1::VerifyOberonProofRequest.new(:data=>data, :nonce=>nonce, :pk=>key.pk, :proof=>proof.proof))
    assert result.valid == false
  end
end

