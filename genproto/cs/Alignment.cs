// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: Alignment.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Deviant {

  /// <summary>Holder for reflection information generated from Alignment.proto</summary>
  public static partial class AlignmentReflection {

    #region Descriptor
    /// <summary>File descriptor for Alignment.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static AlignmentReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cg9BbGlnbm1lbnQucHJvdG8SB0RldmlhbnQqNgoJQWxpZ25tZW50EgwKCEZS",
            "SUVORExZEAASDgoKVU5GUklFTkRMWRABEgsKB05FVVRSQUwQAkIJWgdkZXZp",
            "YW50YgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Deviant.Alignment), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum Alignment {
    [pbr::OriginalName("FRIENDLY")] Friendly = 0,
    [pbr::OriginalName("UNFRIENDLY")] Unfriendly = 1,
    [pbr::OriginalName("NEUTRAL")] Neutral = 2,
  }

  #endregion

}

#endregion Designer generated code
