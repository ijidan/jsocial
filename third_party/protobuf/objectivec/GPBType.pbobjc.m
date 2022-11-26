// Generated by the protocol buffer compiler.  DO NOT EDIT!
// clang-format off
// source: google/protobuf/type.proto

#import "GPBProtocolBuffers_RuntimeSupport.h"
#import "GPBType.pbobjc.h"

#import <stdatomic.h>

// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"
#pragma clang diagnostic ignored "-Wdollar-in-identifier-extension"

#pragma mark - Objective C Class declarations
// Forward declarations of Objective C classes that we can use as
// static values in struct initializers.
// We don't use [Foo class] because it is not a static value.
GPBObjCClassDeclaration(GPBAny);
GPBObjCClassDeclaration(GPBEnumValue);
GPBObjCClassDeclaration(GPBField);
GPBObjCClassDeclaration(GPBOption);
GPBObjCClassDeclaration(GPBSourceContext);

#pragma mark - GPBTypeRoot

@implementation GPBTypeRoot

// No extensions in the file and none of the imports (direct or indirect)
// defined extensions, so no need to generate +extensionRegistry.

@end

#pragma mark - GPBTypeRoot_FileDescriptor

static GPBFileDescriptor *GPBTypeRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"google.protobuf"
                                                 objcPrefix:@"GPB"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - Enum GPBSyntax

GPBEnumDescriptor *GPBSyntax_EnumDescriptor(void) {
  static _Atomic(GPBEnumDescriptor*) descriptor = nil;
  if (!descriptor) {
    static const char *valueNames =
        "SyntaxProto2\000SyntaxProto3\000";
    static const int32_t values[] = {
        GPBSyntax_SyntaxProto2,
        GPBSyntax_SyntaxProto3,
    };
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(GPBSyntax)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:GPBSyntax_IsValidValue
                                            flags:GPBEnumDescriptorInitializationFlag_None];
    GPBEnumDescriptor *expected = nil;
    if (!atomic_compare_exchange_strong(&descriptor, &expected, worker)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL GPBSyntax_IsValidValue(int32_t value__) {
  switch (value__) {
    case GPBSyntax_SyntaxProto2:
    case GPBSyntax_SyntaxProto3:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - GPBType

@implementation GPBType

@dynamic name;
@dynamic fieldsArray, fieldsArray_Count;
@dynamic oneofsArray, oneofsArray_Count;
@dynamic optionsArray, optionsArray_Count;
@dynamic hasSourceContext, sourceContext;
@dynamic syntax;

typedef struct GPBType__storage_ {
  uint32_t _has_storage_[1];
  GPBSyntax syntax;
  NSString *name;
  NSMutableArray *fieldsArray;
  NSMutableArray *oneofsArray;
  NSMutableArray *optionsArray;
  GPBSourceContext *sourceContext;
} GPBType__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "name",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBType_FieldNumber_Name,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(GPBType__storage_, name),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "fieldsArray",
        .dataTypeSpecific.clazz = GPBObjCClass(GPBField),
        .number = GPBType_FieldNumber_FieldsArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(GPBType__storage_, fieldsArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "oneofsArray",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBType_FieldNumber_OneofsArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(GPBType__storage_, oneofsArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "optionsArray",
        .dataTypeSpecific.clazz = GPBObjCClass(GPBOption),
        .number = GPBType_FieldNumber_OptionsArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(GPBType__storage_, optionsArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "sourceContext",
        .dataTypeSpecific.clazz = GPBObjCClass(GPBSourceContext),
        .number = GPBType_FieldNumber_SourceContext,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(GPBType__storage_, sourceContext),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "syntax",
        .dataTypeSpecific.enumDescFunc = GPBSyntax_EnumDescriptor,
        .number = GPBType_FieldNumber_Syntax,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(GPBType__storage_, syntax),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeEnum,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GPBType class]
                                     rootClass:[GPBTypeRoot class]
                                          file:GPBTypeRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GPBType__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown | GPBDescriptorInitializationFlag_ClosedEnumSupportKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

int32_t GPBType_Syntax_RawValue(GPBType *message) {
  GPBDescriptor *descriptor = [GPBType descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:GPBType_FieldNumber_Syntax];
  return GPBGetMessageRawEnumField(message, field);
}

void SetGPBType_Syntax_RawValue(GPBType *message, int32_t value) {
  GPBDescriptor *descriptor = [GPBType descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:GPBType_FieldNumber_Syntax];
  GPBSetMessageRawEnumField(message, field, value);
}

#pragma mark - GPBField

@implementation GPBField

@dynamic kind;
@dynamic cardinality;
@dynamic number;
@dynamic name;
@dynamic typeURL;
@dynamic oneofIndex;
@dynamic packed;
@dynamic optionsArray, optionsArray_Count;
@dynamic jsonName;
@dynamic defaultValue;

typedef struct GPBField__storage_ {
  uint32_t _has_storage_[1];
  GPBField_Kind kind;
  GPBField_Cardinality cardinality;
  int32_t number;
  int32_t oneofIndex;
  NSString *name;
  NSString *typeURL;
  NSMutableArray *optionsArray;
  NSString *jsonName;
  NSString *defaultValue;
} GPBField__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "kind",
        .dataTypeSpecific.enumDescFunc = GPBField_Kind_EnumDescriptor,
        .number = GPBField_FieldNumber_Kind,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(GPBField__storage_, kind),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeEnum,
      },
      {
        .name = "cardinality",
        .dataTypeSpecific.enumDescFunc = GPBField_Cardinality_EnumDescriptor,
        .number = GPBField_FieldNumber_Cardinality,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(GPBField__storage_, cardinality),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeEnum,
      },
      {
        .name = "number",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBField_FieldNumber_Number,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(GPBField__storage_, number),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeInt32,
      },
      {
        .name = "name",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBField_FieldNumber_Name,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(GPBField__storage_, name),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "typeURL",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBField_FieldNumber_TypeURL,
        .hasIndex = 4,
        .offset = (uint32_t)offsetof(GPBField__storage_, typeURL),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "oneofIndex",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBField_FieldNumber_OneofIndex,
        .hasIndex = 5,
        .offset = (uint32_t)offsetof(GPBField__storage_, oneofIndex),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeInt32,
      },
      {
        .name = "packed",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBField_FieldNumber_Packed,
        .hasIndex = 6,
        .offset = 7,  // Stored in _has_storage_ to save space.
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeBool,
      },
      {
        .name = "optionsArray",
        .dataTypeSpecific.clazz = GPBObjCClass(GPBOption),
        .number = GPBField_FieldNumber_OptionsArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(GPBField__storage_, optionsArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "jsonName",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBField_FieldNumber_JsonName,
        .hasIndex = 8,
        .offset = (uint32_t)offsetof(GPBField__storage_, jsonName),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "defaultValue",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBField_FieldNumber_DefaultValue,
        .hasIndex = 9,
        .offset = (uint32_t)offsetof(GPBField__storage_, defaultValue),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GPBField class]
                                     rootClass:[GPBTypeRoot class]
                                          file:GPBTypeRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GPBField__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown | GPBDescriptorInitializationFlag_ClosedEnumSupportKnown)];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\001\006\004\241!!\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

int32_t GPBField_Kind_RawValue(GPBField *message) {
  GPBDescriptor *descriptor = [GPBField descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:GPBField_FieldNumber_Kind];
  return GPBGetMessageRawEnumField(message, field);
}

void SetGPBField_Kind_RawValue(GPBField *message, int32_t value) {
  GPBDescriptor *descriptor = [GPBField descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:GPBField_FieldNumber_Kind];
  GPBSetMessageRawEnumField(message, field, value);
}

int32_t GPBField_Cardinality_RawValue(GPBField *message) {
  GPBDescriptor *descriptor = [GPBField descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:GPBField_FieldNumber_Cardinality];
  return GPBGetMessageRawEnumField(message, field);
}

void SetGPBField_Cardinality_RawValue(GPBField *message, int32_t value) {
  GPBDescriptor *descriptor = [GPBField descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:GPBField_FieldNumber_Cardinality];
  GPBSetMessageRawEnumField(message, field, value);
}

#pragma mark - Enum GPBField_Kind

GPBEnumDescriptor *GPBField_Kind_EnumDescriptor(void) {
  static _Atomic(GPBEnumDescriptor*) descriptor = nil;
  if (!descriptor) {
    static const char *valueNames =
        "TypeUnknown\000TypeDouble\000TypeFloat\000TypeInt"
        "64\000TypeUint64\000TypeInt32\000TypeFixed64\000Type"
        "Fixed32\000TypeBool\000TypeString\000TypeGroup\000Ty"
        "peMessage\000TypeBytes\000TypeUint32\000TypeEnum\000"
        "TypeSfixed32\000TypeSfixed64\000TypeSint32\000Typ"
        "eSint64\000";
    static const int32_t values[] = {
        GPBField_Kind_TypeUnknown,
        GPBField_Kind_TypeDouble,
        GPBField_Kind_TypeFloat,
        GPBField_Kind_TypeInt64,
        GPBField_Kind_TypeUint64,
        GPBField_Kind_TypeInt32,
        GPBField_Kind_TypeFixed64,
        GPBField_Kind_TypeFixed32,
        GPBField_Kind_TypeBool,
        GPBField_Kind_TypeString,
        GPBField_Kind_TypeGroup,
        GPBField_Kind_TypeMessage,
        GPBField_Kind_TypeBytes,
        GPBField_Kind_TypeUint32,
        GPBField_Kind_TypeEnum,
        GPBField_Kind_TypeSfixed32,
        GPBField_Kind_TypeSfixed64,
        GPBField_Kind_TypeSint32,
        GPBField_Kind_TypeSint64,
    };
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(GPBField_Kind)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:GPBField_Kind_IsValidValue
                                            flags:GPBEnumDescriptorInitializationFlag_None];
    GPBEnumDescriptor *expected = nil;
    if (!atomic_compare_exchange_strong(&descriptor, &expected, worker)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL GPBField_Kind_IsValidValue(int32_t value__) {
  switch (value__) {
    case GPBField_Kind_TypeUnknown:
    case GPBField_Kind_TypeDouble:
    case GPBField_Kind_TypeFloat:
    case GPBField_Kind_TypeInt64:
    case GPBField_Kind_TypeUint64:
    case GPBField_Kind_TypeInt32:
    case GPBField_Kind_TypeFixed64:
    case GPBField_Kind_TypeFixed32:
    case GPBField_Kind_TypeBool:
    case GPBField_Kind_TypeString:
    case GPBField_Kind_TypeGroup:
    case GPBField_Kind_TypeMessage:
    case GPBField_Kind_TypeBytes:
    case GPBField_Kind_TypeUint32:
    case GPBField_Kind_TypeEnum:
    case GPBField_Kind_TypeSfixed32:
    case GPBField_Kind_TypeSfixed64:
    case GPBField_Kind_TypeSint32:
    case GPBField_Kind_TypeSint64:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - Enum GPBField_Cardinality

GPBEnumDescriptor *GPBField_Cardinality_EnumDescriptor(void) {
  static _Atomic(GPBEnumDescriptor*) descriptor = nil;
  if (!descriptor) {
    static const char *valueNames =
        "CardinalityUnknown\000CardinalityOptional\000C"
        "ardinalityRequired\000CardinalityRepeated\000";
    static const int32_t values[] = {
        GPBField_Cardinality_CardinalityUnknown,
        GPBField_Cardinality_CardinalityOptional,
        GPBField_Cardinality_CardinalityRequired,
        GPBField_Cardinality_CardinalityRepeated,
    };
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(GPBField_Cardinality)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:GPBField_Cardinality_IsValidValue
                                            flags:GPBEnumDescriptorInitializationFlag_None];
    GPBEnumDescriptor *expected = nil;
    if (!atomic_compare_exchange_strong(&descriptor, &expected, worker)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL GPBField_Cardinality_IsValidValue(int32_t value__) {
  switch (value__) {
    case GPBField_Cardinality_CardinalityUnknown:
    case GPBField_Cardinality_CardinalityOptional:
    case GPBField_Cardinality_CardinalityRequired:
    case GPBField_Cardinality_CardinalityRepeated:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - GPBEnum

@implementation GPBEnum

@dynamic name;
@dynamic enumvalueArray, enumvalueArray_Count;
@dynamic optionsArray, optionsArray_Count;
@dynamic hasSourceContext, sourceContext;
@dynamic syntax;

typedef struct GPBEnum__storage_ {
  uint32_t _has_storage_[1];
  GPBSyntax syntax;
  NSString *name;
  NSMutableArray *enumvalueArray;
  NSMutableArray *optionsArray;
  GPBSourceContext *sourceContext;
} GPBEnum__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "name",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBEnum_FieldNumber_Name,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(GPBEnum__storage_, name),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "enumvalueArray",
        .dataTypeSpecific.clazz = GPBObjCClass(GPBEnumValue),
        .number = GPBEnum_FieldNumber_EnumvalueArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(GPBEnum__storage_, enumvalueArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "optionsArray",
        .dataTypeSpecific.clazz = GPBObjCClass(GPBOption),
        .number = GPBEnum_FieldNumber_OptionsArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(GPBEnum__storage_, optionsArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "sourceContext",
        .dataTypeSpecific.clazz = GPBObjCClass(GPBSourceContext),
        .number = GPBEnum_FieldNumber_SourceContext,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(GPBEnum__storage_, sourceContext),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "syntax",
        .dataTypeSpecific.enumDescFunc = GPBSyntax_EnumDescriptor,
        .number = GPBEnum_FieldNumber_Syntax,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(GPBEnum__storage_, syntax),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeEnum,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GPBEnum class]
                                     rootClass:[GPBTypeRoot class]
                                          file:GPBTypeRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GPBEnum__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown | GPBDescriptorInitializationFlag_ClosedEnumSupportKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

int32_t GPBEnum_Syntax_RawValue(GPBEnum *message) {
  GPBDescriptor *descriptor = [GPBEnum descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:GPBEnum_FieldNumber_Syntax];
  return GPBGetMessageRawEnumField(message, field);
}

void SetGPBEnum_Syntax_RawValue(GPBEnum *message, int32_t value) {
  GPBDescriptor *descriptor = [GPBEnum descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:GPBEnum_FieldNumber_Syntax];
  GPBSetMessageRawEnumField(message, field, value);
}

#pragma mark - GPBEnumValue

@implementation GPBEnumValue

@dynamic name;
@dynamic number;
@dynamic optionsArray, optionsArray_Count;

typedef struct GPBEnumValue__storage_ {
  uint32_t _has_storage_[1];
  int32_t number;
  NSString *name;
  NSMutableArray *optionsArray;
} GPBEnumValue__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "name",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBEnumValue_FieldNumber_Name,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(GPBEnumValue__storage_, name),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "number",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBEnumValue_FieldNumber_Number,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(GPBEnumValue__storage_, number),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeInt32,
      },
      {
        .name = "optionsArray",
        .dataTypeSpecific.clazz = GPBObjCClass(GPBOption),
        .number = GPBEnumValue_FieldNumber_OptionsArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(GPBEnumValue__storage_, optionsArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GPBEnumValue class]
                                     rootClass:[GPBTypeRoot class]
                                          file:GPBTypeRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GPBEnumValue__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown | GPBDescriptorInitializationFlag_ClosedEnumSupportKnown)];
    #if defined(DEBUG) && DEBUG
      NSAssert(descriptor == nil, @"Startup recursed!");
    #endif  // DEBUG
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - GPBOption

@implementation GPBOption

@dynamic name;
@dynamic hasValue, value;

typedef struct GPBOption__storage_ {
  uint32_t _has_storage_[1];
  NSString *name;
  GPBAny *value;
} GPBOption__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "name",
        .dataTypeSpecific.clazz = Nil,
        .number = GPBOption_FieldNumber_Name,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(GPBOption__storage_, name),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldClearHasIvarOnZero),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "value",
        .dataTypeSpecific.clazz = GPBObjCClass(GPBAny),
        .number = GPBOption_FieldNumber_Value,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(GPBOption__storage_, value),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GPBOption class]
                                     rootClass:[GPBTypeRoot class]
                                          file:GPBTypeRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GPBOption__storage_)
                                         flags:(GPBDescriptorInitializationFlags)(GPBDescriptorInitializationFlag_UsesClassRefs | GPBDescriptorInitializationFlag_Proto3OptionalKnown | GPBDescriptorInitializationFlag_ClosedEnumSupportKnown)];
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

// clang-format on