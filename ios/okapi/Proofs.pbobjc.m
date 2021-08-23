// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proofs.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers_RuntimeSupport.h>
#else
 #import "GPBProtocolBuffers_RuntimeSupport.h"
#endif

#import <stdatomic.h>

#import "Proofs.pbobjc.h"
#import "Keys.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"
#pragma clang diagnostic ignored "-Wdollar-in-identifier-extension"

#pragma mark - Objective C Class declarations
// Forward declarations of Objective C classes that we can use as
// static values in struct initializers.
// We don't use [Foo class] because it is not a static value.
GPBObjCClassDeclaration(GPBStruct);
GPBObjCClassDeclaration(JsonWebKey);

#pragma mark - ProofsRoot

@implementation ProofsRoot

// No extensions in the file and none of the imports (direct or indirect)
// defined extensions, so no need to generate +extensionRegistry.

@end

#pragma mark - ProofsRoot_FileDescriptor

static GPBFileDescriptor *ProofsRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"okapi.proofs"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - Enum LdSuite

GPBEnumDescriptor *LdSuite_EnumDescriptor(void) {
  static _Atomic(GPBEnumDescriptor*) descriptor = nil;
  if (!descriptor) {
    static const char *valueNames =
        "JcsEd25519Signature2020\000";
    static const int32_t values[] = {
        LdSuite_JcsEd25519Signature2020,
    };
    static const char *extraTextFormatInfo = "\001\000\027\000";
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(LdSuite)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:LdSuite_IsValidValue
                              extraTextFormatInfo:extraTextFormatInfo];
    GPBEnumDescriptor *expected = nil;
    if (!atomic_compare_exchange_strong(&descriptor, &expected, worker)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL LdSuite_IsValidValue(int32_t value__) {
  switch (value__) {
    case LdSuite_JcsEd25519Signature2020:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - CreateProofRequest

@implementation CreateProofRequest

@dynamic hasDocument, document;
@dynamic hasKey, key;
@dynamic suite;

typedef struct CreateProofRequest__storage_ {
  uint32_t _has_storage_[1];
  LdSuite suite;
  GPBStruct *document;
  JsonWebKey *key;
} CreateProofRequest__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "document",
        .dataTypeSpecific.clazz = GPBObjCClass(GPBStruct),
        .number = CreateProofRequest_FieldNumber_Document,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(CreateProofRequest__storage_, document),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "key",
        .dataTypeSpecific.clazz = GPBObjCClass(JsonWebKey),
        .number = CreateProofRequest_FieldNumber_Key,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(CreateProofRequest__storage_, key),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "suite",
        .dataTypeSpecific.enumDescFunc = LdSuite_EnumDescriptor,
        .number = CreateProofRequest_FieldNumber_Suite,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(CreateProofRequest__storage_, suite),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeEnum,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[CreateProofRequest class]
                                     rootClass:[ProofsRoot class]
                                          file:ProofsRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(CreateProofRequest__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

int32_t CreateProofRequest_Suite_RawValue(CreateProofRequest *message) {
  GPBDescriptor *descriptor = [CreateProofRequest descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:CreateProofRequest_FieldNumber_Suite];
  return GPBGetMessageRawEnumField(message, field);
}

void SetCreateProofRequest_Suite_RawValue(CreateProofRequest *message, int32_t value) {
  GPBDescriptor *descriptor = [CreateProofRequest descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:CreateProofRequest_FieldNumber_Suite];
  GPBSetMessageRawEnumField(message, field, value);
}

#pragma mark - CreateProofResponse

@implementation CreateProofResponse

@dynamic hasSignedDocument, signedDocument;

typedef struct CreateProofResponse__storage_ {
  uint32_t _has_storage_[1];
  GPBStruct *signedDocument;
} CreateProofResponse__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "signedDocument",
        .dataTypeSpecific.clazz = GPBObjCClass(GPBStruct),
        .number = CreateProofResponse_FieldNumber_SignedDocument,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(CreateProofResponse__storage_, signedDocument),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[CreateProofResponse class]
                                     rootClass:[ProofsRoot class]
                                          file:ProofsRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(CreateProofResponse__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - VerifyProofRequest

@implementation VerifyProofRequest


typedef struct VerifyProofRequest__storage_ {
  uint32_t _has_storage_[1];
} VerifyProofRequest__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[VerifyProofRequest class]
                                     rootClass:[ProofsRoot class]
                                          file:ProofsRoot_FileDescriptor()
                                        fields:NULL
                                    fieldCount:0
                                   storageSize:sizeof(VerifyProofRequest__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - VerifyProofResponse

@implementation VerifyProofResponse


typedef struct VerifyProofResponse__storage_ {
  uint32_t _has_storage_[1];
} VerifyProofResponse__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[VerifyProofResponse class]
                                     rootClass:[ProofsRoot class]
                                          file:ProofsRoot_FileDescriptor()
                                        fields:NULL
                                    fieldCount:0
                                   storageSize:sizeof(VerifyProofResponse__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
