// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: examples.proto

package trinsic.okapi.examples;

public final class Examples {
  private Examples() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_okapi_examples_BasicMessage_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_okapi_examples_BasicMessage_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\016examples.proto\022\016okapi.examples\032\021pbmse/" +
      "pbmse.proto\"\034\n\014BasicMessage\022\014\n\004text\030\001 \001(" +
      "\t2\221\001\n\rSecureExample\0229\n\005Unary\022\027.pbmse.Enc" +
      "ryptedMessage\032\027.pbmse.EncryptedMessage\022E" +
      "\n\017ServerStreaming\022\027.pbmse.EncryptedMessa" +
      "ge\032\027.pbmse.EncryptedMessage0\001BH\n\026trinsic" +
      ".okapi.examplesP\001Z\033github.com/trinsic-id" +
      "/okapi\252\002\016Okapi.Examplesb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          trinsic.okapi.pbmse.Pbmse.getDescriptor(),
        });
    internal_static_okapi_examples_BasicMessage_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_okapi_examples_BasicMessage_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_okapi_examples_BasicMessage_descriptor,
        new java.lang.String[] { "Text", });
    trinsic.okapi.pbmse.Pbmse.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
