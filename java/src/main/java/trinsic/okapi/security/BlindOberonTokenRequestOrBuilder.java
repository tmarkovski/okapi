// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: security.proto

package trinsic.okapi.security;

public interface BlindOberonTokenRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:okapi.security.BlindOberonTokenRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * raw token bytes
   * </pre>
   *
   * <code>bytes token = 1;</code>
   * @return The token.
   */
  com.google.protobuf.ByteString getToken();

  /**
   * <pre>
   * blinding to apply to the token
   * </pre>
   *
   * <code>repeated bytes blinding = 2;</code>
   * @return A list containing the blinding.
   */
  java.util.List<com.google.protobuf.ByteString> getBlindingList();
  /**
   * <pre>
   * blinding to apply to the token
   * </pre>
   *
   * <code>repeated bytes blinding = 2;</code>
   * @return The count of blinding.
   */
  int getBlindingCount();
  /**
   * <pre>
   * blinding to apply to the token
   * </pre>
   *
   * <code>repeated bytes blinding = 2;</code>
   * @param index The index of the element to return.
   * @return The blinding at the given index.
   */
  com.google.protobuf.ByteString getBlinding(int index);
}
