// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

package com.google.crypto.tink.mac;

import static com.google.crypto.tink.internal.TinkBugException.exceptionIsBug;

import com.google.crypto.tink.AccessesPartialKey;
import com.google.crypto.tink.KeyTemplate;
import com.google.crypto.tink.Mac;
import com.google.crypto.tink.Parameters;
import com.google.crypto.tink.Registry;
import com.google.crypto.tink.SecretKeyAccess;
import com.google.crypto.tink.config.internal.TinkFipsUtil;
import com.google.crypto.tink.internal.KeyTypeManager;
import com.google.crypto.tink.internal.MutableParametersRegistry;
import com.google.crypto.tink.internal.MutablePrimitiveRegistry;
import com.google.crypto.tink.internal.PrimitiveConstructor;
import com.google.crypto.tink.internal.PrimitiveFactory;
import com.google.crypto.tink.internal.Util;
import com.google.crypto.tink.mac.internal.ChunkedHmacImpl;
import com.google.crypto.tink.mac.internal.HmacProtoSerialization;
import com.google.crypto.tink.proto.HashType;
import com.google.crypto.tink.proto.HmacKey;
import com.google.crypto.tink.proto.HmacKeyFormat;
import com.google.crypto.tink.proto.HmacParams;
import com.google.crypto.tink.proto.KeyData.KeyMaterialType;
import com.google.crypto.tink.subtle.PrfHmacJce;
import com.google.crypto.tink.subtle.PrfMac;
import com.google.crypto.tink.subtle.Random;
import com.google.crypto.tink.subtle.Validators;
import com.google.protobuf.ByteString;
import com.google.protobuf.ExtensionRegistryLite;
import com.google.protobuf.InvalidProtocolBufferException;
import java.io.InputStream;
import java.security.GeneralSecurityException;
import java.util.Collections;
import java.util.HashMap;
import java.util.Map;
import javax.annotation.Nullable;
import javax.crypto.spec.SecretKeySpec;

/**
 * This key manager generates new {@code HmacKey} keys and produces new instances of {@code
 * PrfHmacJce}.
 */
public final class HmacKeyManager extends KeyTypeManager<HmacKey> {
  public HmacKeyManager() {
    super(
        HmacKey.class,
        new PrimitiveFactory<Mac, HmacKey>(Mac.class) {
          @Override
          public Mac getPrimitive(HmacKey key) throws GeneralSecurityException {
            HashType hash = key.getParams().getHash();
            byte[] keyValue = key.getKeyValue().toByteArray();
            SecretKeySpec keySpec = new SecretKeySpec(keyValue, "HMAC");
            int tagSize = key.getParams().getTagSize();
            switch (hash) {
              case SHA1:
                return new PrfMac(new PrfHmacJce("HMACSHA1", keySpec), tagSize);
              case SHA224:
                return new PrfMac(new PrfHmacJce("HMACSHA224", keySpec), tagSize);
              case SHA256:
                return new PrfMac(new PrfHmacJce("HMACSHA256", keySpec), tagSize);
              case SHA384:
                return new PrfMac(new PrfHmacJce("HMACSHA384", keySpec), tagSize);
              case SHA512:
                return new PrfMac(new PrfHmacJce("HMACSHA512", keySpec), tagSize);
              default:
                throw new GeneralSecurityException("unknown hash");
            }
          }
        });
  }

  /** Minimum key size in bytes. */
  private static final int MIN_KEY_SIZE_IN_BYTES = 16;

  /** Minimum tag size in bytes. This provides minimum 80-bit security strength. */
  private static final int MIN_TAG_SIZE_IN_BYTES = 10;

  private static final PrimitiveConstructor<com.google.crypto.tink.mac.HmacKey, ChunkedMac>
      CHUNKED_MAC_PRIMITIVE_CONSTRUCTOR =
          PrimitiveConstructor.create(
              ChunkedHmacImpl::new, com.google.crypto.tink.mac.HmacKey.class, ChunkedMac.class);
  private static final PrimitiveConstructor<com.google.crypto.tink.mac.HmacKey, Mac>
      MAC_PRIMITIVE_CONSTRUCTOR =
          PrimitiveConstructor.create(
              PrfMac::create, com.google.crypto.tink.mac.HmacKey.class, Mac.class);

  @Override
  public String getKeyType() {
    return "type.googleapis.com/google.crypto.tink.HmacKey";
  }

  @Override
  public int getVersion() {
    return 0;
  }

  @Override
  public KeyMaterialType keyMaterialType() {
    return KeyMaterialType.SYMMETRIC;
  }

  @Override
  public void validateKey(HmacKey key) throws GeneralSecurityException {
    Validators.validateVersion(key.getVersion(), getVersion());
    if (key.getKeyValue().size() < MIN_KEY_SIZE_IN_BYTES) {
      throw new GeneralSecurityException("key too short");
    }
    validateParams(key.getParams());
  }

  @Override
  public HmacKey parseKey(ByteString byteString) throws InvalidProtocolBufferException {
    return HmacKey.parseFrom(byteString, ExtensionRegistryLite.getEmptyRegistry());
  }

  private static void validateParams(HmacParams params) throws GeneralSecurityException {
    if (params.getTagSize() < MIN_TAG_SIZE_IN_BYTES) {
      throw new GeneralSecurityException("tag size too small");
    }
    switch (params.getHash()) {
      case SHA1:
        if (params.getTagSize() > 20) {
          throw new GeneralSecurityException("tag size too big");
        }
        break;
      case SHA224:
        if (params.getTagSize() > 28) {
          throw new GeneralSecurityException("tag size too big");
        }
        break;
      case SHA256:
        if (params.getTagSize() > 32) {
          throw new GeneralSecurityException("tag size too big");
        }
        break;
      case SHA384:
        if (params.getTagSize() > 48) {
          throw new GeneralSecurityException("tag size too big");
        }
        break;
      case SHA512:
        if (params.getTagSize() > 64) {
          throw new GeneralSecurityException("tag size too big");
        }
        break;
      default:
        throw new GeneralSecurityException("unknown hash type");
    }
  }

  @Override
  public KeyFactory<HmacKeyFormat, HmacKey> keyFactory() {
    return new KeyFactory<HmacKeyFormat, HmacKey>(HmacKeyFormat.class) {
      @Override
      public void validateKeyFormat(HmacKeyFormat format) throws GeneralSecurityException {
        if (format.getKeySize() < MIN_KEY_SIZE_IN_BYTES) {
          throw new GeneralSecurityException("key too short");
        }
        validateParams(format.getParams());
      }

      @Override
      public HmacKeyFormat parseKeyFormat(ByteString byteString)
          throws InvalidProtocolBufferException {
        return HmacKeyFormat.parseFrom(byteString, ExtensionRegistryLite.getEmptyRegistry());
      }

      @Override
      public HmacKey createKey(HmacKeyFormat format) throws GeneralSecurityException {
        return HmacKey.newBuilder()
            .setVersion(getVersion())
            .setParams(format.getParams())
            .setKeyValue(ByteString.copyFrom(Random.randBytes(format.getKeySize())))
            .build();
      }

      @Override
      public com.google.crypto.tink.mac.HmacKey createKeyFromRandomness(
          Parameters parameters,
          InputStream stream,
          @Nullable Integer idRequirement,
          SecretKeyAccess access)
          throws GeneralSecurityException {
        if (parameters instanceof HmacParameters) {
          return createHmacKeyFromRandomness(
              (HmacParameters) parameters, stream, idRequirement, access);
        }
        throw new GeneralSecurityException(
            "Unexpected parameters: expected HmacParameters, but got: " + parameters);
      }
    };
  }

  @AccessesPartialKey
  static com.google.crypto.tink.mac.HmacKey createHmacKeyFromRandomness(
      HmacParameters parameters,
      InputStream stream,
      @Nullable Integer idRequirement,
      SecretKeyAccess access)
      throws GeneralSecurityException {
    return com.google.crypto.tink.mac.HmacKey.builder()
        .setParameters(parameters)
        .setKeyBytes(Util.readIntoSecretBytes(stream, parameters.getKeySizeBytes(), access))
        .setIdRequirement(idRequirement)
        .build();
  }

  private static Map<String, Parameters> namedParameters() throws GeneralSecurityException {
        Map<String, Parameters> result = new HashMap<>();
        result.put("HMAC_SHA256_128BITTAG", PredefinedMacParameters.HMAC_SHA256_128BITTAG);
        result.put(
            "HMAC_SHA256_128BITTAG_RAW",
            HmacParameters.builder()
                .setKeySizeBytes(32)
                .setTagSizeBytes(16)
                .setVariant(HmacParameters.Variant.NO_PREFIX)
                .setHashType(HmacParameters.HashType.SHA256)
                .build());
        result.put(
            "HMAC_SHA256_256BITTAG",
            HmacParameters.builder()
                .setKeySizeBytes(32)
                .setTagSizeBytes(32)
                .setVariant(HmacParameters.Variant.TINK)
                .setHashType(HmacParameters.HashType.SHA256)
                .build());
        result.put(
            "HMAC_SHA256_256BITTAG_RAW",
            HmacParameters.builder()
                .setKeySizeBytes(32)
                .setTagSizeBytes(32)
                .setVariant(HmacParameters.Variant.NO_PREFIX)
                .setHashType(HmacParameters.HashType.SHA256)
                .build());
        result.put(
            "HMAC_SHA512_128BITTAG",
            HmacParameters.builder()
                .setKeySizeBytes(64)
                .setTagSizeBytes(16)
                .setVariant(HmacParameters.Variant.TINK)
                .setHashType(HmacParameters.HashType.SHA512)
                .build());
        result.put(
            "HMAC_SHA512_128BITTAG_RAW",
            HmacParameters.builder()
                .setKeySizeBytes(64)
                .setTagSizeBytes(16)
                .setVariant(HmacParameters.Variant.NO_PREFIX)
                .setHashType(HmacParameters.HashType.SHA512)
                .build());
        result.put(
            "HMAC_SHA512_256BITTAG",
            HmacParameters.builder()
                .setKeySizeBytes(64)
                .setTagSizeBytes(32)
                .setVariant(HmacParameters.Variant.TINK)
                .setHashType(HmacParameters.HashType.SHA512)
                .build());
        result.put(
            "HMAC_SHA512_256BITTAG_RAW",
            HmacParameters.builder()
                .setKeySizeBytes(64)
                .setTagSizeBytes(32)
                .setVariant(HmacParameters.Variant.NO_PREFIX)
                .setHashType(HmacParameters.HashType.SHA512)
                .build());
        result.put("HMAC_SHA512_512BITTAG", PredefinedMacParameters.HMAC_SHA512_512BITTAG);
        result.put(
            "HMAC_SHA512_512BITTAG_RAW",
            HmacParameters.builder()
                .setKeySizeBytes(64)
                .setTagSizeBytes(64)
                .setVariant(HmacParameters.Variant.NO_PREFIX)
                .setHashType(HmacParameters.HashType.SHA512)
                .build());
        return Collections.unmodifiableMap(result);
  }

  public static void register(boolean newKeyAllowed) throws GeneralSecurityException {
    Registry.registerKeyManager(new HmacKeyManager(), newKeyAllowed);
    HmacProtoSerialization.register();
    MutablePrimitiveRegistry.globalInstance()
        .registerPrimitiveConstructor(CHUNKED_MAC_PRIMITIVE_CONSTRUCTOR);
    MutablePrimitiveRegistry.globalInstance()
        .registerPrimitiveConstructor(MAC_PRIMITIVE_CONSTRUCTOR);
    MutableParametersRegistry.globalInstance().putAll(namedParameters());
  }

  /**
   * @return A {@link KeyTemplate} that generates new instances of HMAC keys with the following
   *     parameters:
   *     <ul>
   *       <li>Key size: 32 bytes
   *       <li>Tag size: 16 bytes
   *       <li>Hash function: SHA256
   *       <li>Prefix type: {@link KeyTemplate.OutputPrefixType#TINK}
   *     </ul>
   */
  public static final KeyTemplate hmacSha256HalfDigestTemplate() {
    return exceptionIsBug(
        () ->
            KeyTemplate.createFrom(
                HmacParameters.builder()
                    .setKeySizeBytes(32)
                    .setTagSizeBytes(16)
                    .setHashType(HmacParameters.HashType.SHA256)
                    .setVariant(HmacParameters.Variant.TINK)
                    .build()));
  }

  /**
   * @return A {@link KeyTemplate} that generates new instances of HMAC keys with the following
   *     parameters:
   *     <ul>
   *       <li>Key size: 32 bytes
   *       <li>Tag size: 32 bytes
   *       <li>Hash function: SHA256
   *       <li>Prefix type: {@link KeyTemplate.OutputPrefixType#TINK}
   *     </ul>
   */
  public static final KeyTemplate hmacSha256Template() {
    return exceptionIsBug(
        () ->
            KeyTemplate.createFrom(
                HmacParameters.builder()
                    .setKeySizeBytes(32)
                    .setTagSizeBytes(32)
                    .setHashType(HmacParameters.HashType.SHA256)
                    .setVariant(HmacParameters.Variant.TINK)
                    .build()));
  }

  /**
   * @return A {@link KeyTemplate} that generates new instances of HMAC keys with the following
   *     parameters:
   *     <ul>
   *       <li>Key size: 64 bytes
   *       <li>Tag size: 32 bytes
   *       <li>Hash function: SHA512
   *       <li>Prefix type: {@link KeyTemplate.OutputPrefixType#TINK}
   *     </ul>
   */
  public static final KeyTemplate hmacSha512HalfDigestTemplate() {
    return exceptionIsBug(
        () ->
            KeyTemplate.createFrom(
                HmacParameters.builder()
                    .setKeySizeBytes(64)
                    .setTagSizeBytes(32)
                    .setHashType(HmacParameters.HashType.SHA512)
                    .setVariant(HmacParameters.Variant.TINK)
                    .build()));
  }

  /**
   * @return A {@link KeyTemplate} that generates new instances of HMAC keys with the following
   *     parameters:
   *     <ul>
   *       <li>Key size: 64 bytes
   *       <li>Tag size: 64 bytes
   *       <li>Hash function: SHA512
   *       <li>Prefix type: {@link KeyTemplate.OutputPrefixType#TINK}
   *     </ul>
   */
  public static final KeyTemplate hmacSha512Template() {
    return exceptionIsBug(
        () ->
            KeyTemplate.createFrom(
                HmacParameters.builder()
                    .setKeySizeBytes(64)
                    .setTagSizeBytes(64)
                    .setHashType(HmacParameters.HashType.SHA512)
                    .setVariant(HmacParameters.Variant.TINK)
                    .build()));
  }

  @Override
  public TinkFipsUtil.AlgorithmFipsCompatibility fipsStatus() {
    return TinkFipsUtil.AlgorithmFipsCompatibility.ALGORITHM_REQUIRES_BORINGCRYPTO;
  }
}
