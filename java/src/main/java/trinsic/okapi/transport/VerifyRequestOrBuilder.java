// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: transport.proto

package trinsic.okapi.transport;

public interface VerifyRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:okapi.transport.VerifyRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.pbmse.SignedMessage message = 1;</code>
   * @return Whether the message field is set.
   */
  boolean hasMessage();
  /**
   * <code>.pbmse.SignedMessage message = 1;</code>
   * @return The message.
   */
  trinsic.okapi.pbmse.SignedMessage getMessage();
  /**
   * <code>.pbmse.SignedMessage message = 1;</code>
   */
  trinsic.okapi.pbmse.SignedMessageOrBuilder getMessageOrBuilder();

  /**
   * <code>.okapi.keys.JsonWebKey key = 2;</code>
   * @return Whether the key field is set.
   */
  boolean hasKey();
  /**
   * <code>.okapi.keys.JsonWebKey key = 2;</code>
   * @return The key.
   */
  trinsic.okapi.keys.JsonWebKey getKey();
  /**
   * <code>.okapi.keys.JsonWebKey key = 2;</code>
   */
  trinsic.okapi.keys.JsonWebKeyOrBuilder getKeyOrBuilder();
}
