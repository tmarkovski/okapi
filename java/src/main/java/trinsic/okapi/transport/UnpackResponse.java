// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: transport.proto

package trinsic.okapi.transport;

/**
 * Protobuf type {@code okapi.transport.UnpackResponse}
 */
public final class UnpackResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:okapi.transport.UnpackResponse)
    UnpackResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use UnpackResponse.newBuilder() to construct.
  private UnpackResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private UnpackResponse() {
    plaintext_ = com.google.protobuf.ByteString.EMPTY;
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new UnpackResponse();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private UnpackResponse(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10: {

            plaintext_ = input.readBytes();
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return trinsic.okapi.transport.Transport.internal_static_okapi_transport_UnpackResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return trinsic.okapi.transport.Transport.internal_static_okapi_transport_UnpackResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            trinsic.okapi.transport.UnpackResponse.class, trinsic.okapi.transport.UnpackResponse.Builder.class);
  }

  public static final int PLAINTEXT_FIELD_NUMBER = 1;
  private com.google.protobuf.ByteString plaintext_;
  /**
   * <code>bytes plaintext = 1;</code>
   * @return The plaintext.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString getPlaintext() {
    return plaintext_;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!plaintext_.isEmpty()) {
      output.writeBytes(1, plaintext_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!plaintext_.isEmpty()) {
      size += com.google.protobuf.CodedOutputStream
        .computeBytesSize(1, plaintext_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof trinsic.okapi.transport.UnpackResponse)) {
      return super.equals(obj);
    }
    trinsic.okapi.transport.UnpackResponse other = (trinsic.okapi.transport.UnpackResponse) obj;

    if (!getPlaintext()
        .equals(other.getPlaintext())) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + PLAINTEXT_FIELD_NUMBER;
    hash = (53 * hash) + getPlaintext().hashCode();
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static trinsic.okapi.transport.UnpackResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static trinsic.okapi.transport.UnpackResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static trinsic.okapi.transport.UnpackResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static trinsic.okapi.transport.UnpackResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static trinsic.okapi.transport.UnpackResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static trinsic.okapi.transport.UnpackResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static trinsic.okapi.transport.UnpackResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static trinsic.okapi.transport.UnpackResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static trinsic.okapi.transport.UnpackResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static trinsic.okapi.transport.UnpackResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static trinsic.okapi.transport.UnpackResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static trinsic.okapi.transport.UnpackResponse parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(trinsic.okapi.transport.UnpackResponse prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code okapi.transport.UnpackResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:okapi.transport.UnpackResponse)
      trinsic.okapi.transport.UnpackResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return trinsic.okapi.transport.Transport.internal_static_okapi_transport_UnpackResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return trinsic.okapi.transport.Transport.internal_static_okapi_transport_UnpackResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              trinsic.okapi.transport.UnpackResponse.class, trinsic.okapi.transport.UnpackResponse.Builder.class);
    }

    // Construct using trinsic.okapi.transport.UnpackResponse.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      plaintext_ = com.google.protobuf.ByteString.EMPTY;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return trinsic.okapi.transport.Transport.internal_static_okapi_transport_UnpackResponse_descriptor;
    }

    @java.lang.Override
    public trinsic.okapi.transport.UnpackResponse getDefaultInstanceForType() {
      return trinsic.okapi.transport.UnpackResponse.getDefaultInstance();
    }

    @java.lang.Override
    public trinsic.okapi.transport.UnpackResponse build() {
      trinsic.okapi.transport.UnpackResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public trinsic.okapi.transport.UnpackResponse buildPartial() {
      trinsic.okapi.transport.UnpackResponse result = new trinsic.okapi.transport.UnpackResponse(this);
      result.plaintext_ = plaintext_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof trinsic.okapi.transport.UnpackResponse) {
        return mergeFrom((trinsic.okapi.transport.UnpackResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(trinsic.okapi.transport.UnpackResponse other) {
      if (other == trinsic.okapi.transport.UnpackResponse.getDefaultInstance()) return this;
      if (other.getPlaintext() != com.google.protobuf.ByteString.EMPTY) {
        setPlaintext(other.getPlaintext());
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      trinsic.okapi.transport.UnpackResponse parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (trinsic.okapi.transport.UnpackResponse) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private com.google.protobuf.ByteString plaintext_ = com.google.protobuf.ByteString.EMPTY;
    /**
     * <code>bytes plaintext = 1;</code>
     * @return The plaintext.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString getPlaintext() {
      return plaintext_;
    }
    /**
     * <code>bytes plaintext = 1;</code>
     * @param value The plaintext to set.
     * @return This builder for chaining.
     */
    public Builder setPlaintext(com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      plaintext_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>bytes plaintext = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearPlaintext() {
      
      plaintext_ = getDefaultInstance().getPlaintext();
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:okapi.transport.UnpackResponse)
  }

  // @@protoc_insertion_point(class_scope:okapi.transport.UnpackResponse)
  private static final trinsic.okapi.transport.UnpackResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new trinsic.okapi.transport.UnpackResponse();
  }

  public static trinsic.okapi.transport.UnpackResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<UnpackResponse>
      PARSER = new com.google.protobuf.AbstractParser<UnpackResponse>() {
    @java.lang.Override
    public UnpackResponse parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new UnpackResponse(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<UnpackResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<UnpackResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public trinsic.okapi.transport.UnpackResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

